#!/usr/bin/env bash
# build compiles marathon's Go commands from the working tree's marathon/ module into one
# directory: its one argument, or by default plugins/marathon/.build/, the gitignored local build
# that the plugin's bin/ shims run in place of the release.
set -euo pipefail

root=$(cd "$(dirname "$0")/.." && pwd)
out=${1:-$root/plugins/marathon/.build}
mkdir -p "$out"
GOWORK=off go build -C "$root/marathon" -o "$(cd "$out" && pwd)/" ./cmd/...
