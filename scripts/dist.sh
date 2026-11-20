#!/usr/bin/env bash
# dist builds the assets of a marathon/v<version> release into a directory, as the release workflow
# publishes them and scripts/gate.sh cross-builds them before the tag: one archive per target,
# marathon-<version>-<os>-<arch>.tar.gz for linux and darwin on amd64 and arm64, each holding the
# four commands that scripts/build.sh builds, stamped with the version, at its top level, and
# SHA256SUMS, listing each archive's sha256 sum as sha256sum writes it. The names and the layout
# are what mise's github: backend reads with version_prefix = "marathon/v" alone: it picks the
# archive whose name carries the host's OS and architecture, and puts the directory it extracts to
# on PATH since that directory holds executables. A target that fails to build fails dist, naming
# the target.
set -euo pipefail

fail() {
  echo "dist: $*" >&2
  exit 1
}

[ $# -eq 2 ] || fail "usage: scripts/dist.sh <version> <directory>"
version=$1
[[ "$version" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] ||
  fail "malformed version '$version'; pass the tag's <major>.<minor>.<patch>, without marathon/v"
scripts=$(cd "$(dirname "$0")" && pwd)
mkdir -p "$2"
out=$(cd "$2" && pwd)
stage=$(mktemp -d)
trap 'rm -rf "$stage"' EXIT

archives=()
for os in linux darwin; do
  for arch in amd64 arm64; do
    target=$stage/$os-$arch
    GOOS=$os GOARCH=$arch "$scripts/build.sh" "$target" "$version" ||
      fail "marathon/ does not build for $os/$arch; fix the lines above, then rerun"
    archive=marathon-$version-$os-$arch.tar.gz
    commands=("$target"/*)
    tar -czf "$out/$archive" -C "$target" "${commands[@]##*/}"
    archives+=("$archive")
  done
done
(cd "$out" && sha256sum "${archives[@]}" >SHA256SUMS)
