#!/usr/bin/env bash
# build compiles marathon's Go commands from the working tree's marathon/ module into one
# directory: its first argument, or by default plugins/marathon/.build/, the gitignored local build
# that the plugin's bin/ shims run in place of the release. Its second argument, when given, is the
# version each command's --version reports, stamped at link time; without it they report dev. The
# commands build for GOOS and GOARCH from the environment, the host's by default, with cgo off and
# no local paths, as scripts/dist.sh builds the release.
set -euo pipefail

root=$(cd "$(dirname "$0")/.." && pwd)
out=${1:-$root/plugins/marathon/.build}
version=${2-}
ldflags=
if [ -n "$version" ]; then
  ldflags="-X github.com/standards-lab/claude-plugins/marathon/internal/version.version=$version"
fi
mkdir -p "$out"
GOWORK=off CGO_ENABLED=0 go build -C "$root/marathon" -trimpath -ldflags "$ldflags" \
  -o "$(cd "$out" && pwd)/" ./cmd/...
