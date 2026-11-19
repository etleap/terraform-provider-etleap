"""Checks that the public Terraform provider keeps up with the API spec, and that the Terraform Registry serves
the latest release.

Prints one line per check, writes the lines to $BASH_ENV as HEALTH_SUMMARY for the Slack notification, and exits
with status 1 if any check fails. GITHUB_TOKEN must be able to read the etleap repository.
"""
import json
import os
import re
import shlex
import sys
import urllib.error
import urllib.request
from datetime import datetime, timedelta, timezone

PROVIDER_REPO = 'etleap/terraform-provider-etleap'
SPEC_REPO = 'etleap/etleap'
SPEC_PATH = 'api-specs/external-api-v2/reference/etleap-api.v2.json'
REGISTRY_URL = 'https://registry.terraform.io/v1/providers/etleap/etleap'

# How long each stage may lag behind the previous one before it counts as a failure. The pipeline normally
# finishes each stage within an hour; the extra time covers reruns.
SPEC_GRACE = timedelta(hours=24)
RELEASE_GRACE = timedelta(hours=2)
REGISTRY_GRACE = timedelta(hours=2)

NOW = datetime.now(timezone.utc)


def fetch(url, raw=False):
    headers = {'Accept': 'application/vnd.github.raw' if raw else 'application/json'}
    if url.startswith('https://api.github.com/') and os.environ.get('GITHUB_TOKEN'):
        headers['Authorization'] = 'Bearer ' + os.environ['GITHUB_TOKEN']
    with urllib.request.urlopen(urllib.request.Request(url, headers=headers), timeout=60) as response:
        body = response.read().decode()
    return body if raw else json.loads(body)


def github(path, raw=False):
    return fetch('https://api.github.com/' + path, raw)


def github_file(repo, path, ref):
    return github(f'repos/{repo}/contents/{path}?ref={ref}', raw=True)


def parse_time(value):
    return datetime.fromisoformat(value.replace('Z', '+00:00'))


def age(since):
    hours = int((NOW - since).total_seconds() // 3600)
    return f'{hours // 24}d {hours % 24}h' if hours >= 24 else f'{hours}h'


def check_spec():
    spec_version = json.loads(github_file(SPEC_REPO, SPEC_PATH, 'master'))['info']['version']
    gen_lock = github_file(PROVIDER_REPO, '.speakeasy/gen.lock', 'main')
    provider_spec_version = re.search(r'^\s*docVersion:\s*(\S+)', gen_lock, re.MULTILINE).group(1)
    if spec_version == provider_spec_version:
        return True, f'API spec: the provider is built from {spec_version}, the version on master'

    # Every regeneration changes gen.lock, so the spec changes after its last commit are the ones that the
    # provider is missing. The age of the oldest of them shows how long the provider has been behind.
    regenerations = github(f'repos/{PROVIDER_REPO}/commits?sha=main&path=.speakeasy/gen.lock&per_page=1')
    regenerated_at = regenerations[0]['commit']['committer']['date']
    missed_changes = github(f'repos/{SPEC_REPO}/commits?sha=master&path={SPEC_PATH}&since={regenerated_at}&per_page=100')
    behind_since = parse_time(missed_changes[-1]['commit']['committer']['date'] if missed_changes else regenerated_at)
    return NOW - behind_since < SPEC_GRACE, (
        f'API spec: master is at {spec_version}, but the provider is built from {provider_spec_version}. '
        f'The provider has been behind for {age(behind_since)}')


def check_release():
    gen_yaml = github_file(PROVIDER_REPO, '.speakeasy/gen.yaml', 'main')
    tag = 'v' + re.search(r'^terraform:\s*\n\s+version:\s*(\S+)', gen_yaml, re.MULTILINE).group(1)
    # The publish job pushes the tag before it creates the release, so only a published release shows that the
    # version shipped. GitHub returns 404 for a tag without a release and for a draft release.
    try:
        github(f'repos/{PROVIDER_REPO}/releases/tags/{tag}')
        return True, f'Release: main is released as {tag}'
    except urllib.error.HTTPError as e:
        if e.code != 404:
            raise

    # Every regeneration changes the version in gen.yaml, so its last commit is when this version first needed a
    # release. Later commits on main must not restart the grace period.
    version_commits = github(f'repos/{PROVIDER_REPO}/commits?sha=main&path=.speakeasy/gen.yaml&per_page=1')
    versioned_at = parse_time(version_commits[0]['commit']['committer']['date'])
    statuses = github(f'repos/{PROVIDER_REPO}/commits/main/status')['statuses']
    test_state = next((s['state'] for s in statuses if s['context'] == 'ci/circleci: test'), 'not run')
    return NOW - versioned_at < RELEASE_GRACE, (
        f'Release: {tag} on main is not released. It has been waiting for {age(versioned_at)}, '
        f'and the test job status of the last commit on main is {test_state}')


def check_registry():
    release = github(f'repos/{PROVIDER_REPO}/releases/latest')
    github_version = release['tag_name'].lstrip('v')
    registry_version = fetch(REGISTRY_URL)['version']
    if registry_version == github_version:
        return True, f'Registry: serves {registry_version}, the latest GitHub release'

    published_at = parse_time(release['published_at'])
    return NOW - published_at < REGISTRY_GRACE, (
        f'Registry: serves {registry_version}, but the latest GitHub release is {github_version}. '
        f'It was published {age(published_at)} ago')


def main():
    results = []
    for name, check in [('API spec', check_spec), ('Release', check_release), ('Registry', check_registry)]:
        try:
            results.append(check())
        except Exception as e:
            results.append((False, f'{name}: the check could not run: {e}'))

    summary = '\n'.join((':white_check_mark: ' if ok else ':x: ') + line for ok, line in results)
    print(summary)
    if os.environ.get('BASH_ENV'):
        with open(os.environ['BASH_ENV'], 'a') as bash_env:
            bash_env.write(f'export HEALTH_SUMMARY={shlex.quote(summary)}\n')
    sys.exit(0 if all(ok for ok, _ in results) else 1)


if __name__ == '__main__':
    main()
