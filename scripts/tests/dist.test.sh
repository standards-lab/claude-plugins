#!/usr/bin/env bash
# Tests for scripts/dist.sh, the build the release workflow publishes, run on the working tree into
# a temporary directory. Each test_ function is one case; the script prints a FAIL line per failed
# assertion and exits 1 if any fails.
set -uo pipefail

repo=$(cd "$(dirname "$0")/../.." && pwd)
commands=(lint-context lint-references marathon wrap-formats)
targets=(darwin-amd64 darwin-arm64 linux-amd64 linux-arm64)
failures=0

fail() {
  echo "FAIL: ${FUNCNAME[1]}: $*" >&2
  failures=1
}

set_up() {
  tmp=$(mktemp -d)
}

host_target() {
  local arch
  arch=$(uname -m)
  case $arch in
  x86_64) arch=amd64 ;;
  aarch64) arch=arm64 ;;
  esac
  echo "$(uname -s | tr '[:upper:]' '[:lower:]')-$arch"
}

test_each_target_gets_an_archive_of_the_four_commands_and_a_checksum() {
  "$repo/scripts/dist.sh" 0.21.0 "$tmp/out" || fail "dist failed"
  local want_assets=SHA256SUMS target
  for target in "${targets[@]}"; do
    want_assets+=$'\n'"marathon-0.21.0-$target.tar.gz"
  done
  [ "$(cd "$tmp/out" && ls | LC_ALL=C sort)" = "$want_assets" ] ||
    fail "assets: got [$(ls "$tmp/out")], want [$want_assets]"
  for target in "${targets[@]}"; do
    local listing
    listing=$(tar -tvzf "$tmp/out/marathon-0.21.0-$target.tar.gz" | awk '{ print $1, $NF }' | sort)
    [ "$listing" = "$(printf -- '-rwxr-xr-x %s\n' "${commands[@]}")" ] ||
      fail "$target archive holds [$listing], want the four commands executable at its top level"
  done
  (cd "$tmp/out" && sha256sum --check --strict --quiet SHA256SUMS) || fail "SHA256SUMS does not match"
  [ "$(wc -l <"$tmp/out/SHA256SUMS")" = 4 ] || fail "SHA256SUMS lists other than the four archives"
}

test_the_hosts_commands_report_the_version() {
  "$repo/scripts/dist.sh" 0.21.0 "$tmp/out" || fail "dist failed"
  mkdir "$tmp/host"
  tar -xzf "$tmp/out/marathon-0.21.0-$(host_target).tar.gz" -C "$tmp/host"
  local command
  for command in "${commands[@]}"; do
    [ "$("$tmp/host/$command" --version)" = "$command 0.21.0" ] ||
      fail "$command --version: got [$("$tmp/host/$command" --version)], want [$command 0.21.0]"
  done
}

test_a_malformed_version_fails_naming_the_fix() {
  local stderr
  stderr=$("$repo/scripts/dist.sh" marathon/v0.21.0 "$tmp/out" 2>&1)
  [ $? = 1 ] || fail "dist passed a malformed version"
  [ "$stderr" = "dist: malformed version 'marathon/v0.21.0'; pass the tag's <major>.<minor>.<patch>, without marathon/v" ] ||
    fail "stderr: [$stderr]"
}

for test in $(declare -F | awk '$3 ~ /^test_/ { print $3 }'); do
  set_up
  "$test"
  rm -rf "$tmp"
done
exit "$failures"
