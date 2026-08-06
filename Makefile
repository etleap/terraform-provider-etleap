default: generate

get-latest-schema:
	@curl -s https://stoplight.io/api/v1/projects/etleap/api-v2/nodes/api-specs/external-api-v2/reference/etleap-api.v2.json?deref=optimizedBundle -o schemas/schema.json

preprocess: get-latest-schema
	@python3 preprocessor/process.py schemas/schema.json schemas/schema-adapted.json

generate: preprocess
	@speakeasy generate sdk --lang terraform -o . -s schemas/schema-adapted.json

apply-patch:
	@git apply etleap.patch

generate-patch:
	@git reset
	@git diff ':!etleap.patch' > etleap.patch
	@git checkout -- . ":!etleap.patch"

release-build:
	@goreleaser build --skip=validate --clean

release-publish:
	@goreleaser release --skip=validate --clean
