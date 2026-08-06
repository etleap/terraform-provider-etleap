#!/usr/bin/env bash
set -euo pipefail

PORT="${BRIDGE_PORT:-5555}"

pkill -f "socat TCP-LISTEN:${PORT}" 2>/dev/null || true

echo "bridge-ready"
echo "watching for plugin socket"
start=$(date +%s)
sock=""
while :; do
  sock=$(find /tmp -maxdepth 1 -name 'plugin*' -type s -newermt "@${start}" 2>/dev/null | head -n1)
  [ -n "$sock" ] && break
  sleep 0.2
done

echo "bridging $sock -> :${PORT}"

reattach=$(printf '{"registry.terraform.io/etleap/etleap":{"Protocol":"grpc","ProtocolVersion":6,"Pid":%d,"Test":true,"Addr":{"Network":"tcp","String":"127.0.0.1:%s"}}}' "$$" "$PORT")
printf 'export TF_REATTACH_PROVIDERS=%q\n' "$reattach" > /tmp/tf-reattach.host.env
echo "wrote host reattach env to /tmp/tf-reattach.host.env"
echo "on host: source \$(devcontainer path)/tmp/tf-reattach.host.env  # or copy the line"
echo "---"
cat /tmp/tf-reattach.host.env
echo "---"

exec socat "TCP-LISTEN:${PORT},reuseaddr,fork" "UNIX-CONNECT:${sock}"
