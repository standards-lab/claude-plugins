#!/usr/bin/env bash
# Tests for scripts/gate.sh's builds, run against a seeded copy of the repository whose
# scripts/check.sh passes at once, so the gate reaches its builds without the whole check. A stub
# claude goes first on PATH and stands for the eval: it runs the plugin's marathon shim and records
# what that prints, and exits with the status the test sets. Each test_ function is one case; the
# script prints a FAIL line per failed assertion and exits 1 if any fails.
set -uo pipefail

repo=$(cd "$(dirname "$0")/../.." && pwd)
tag=marathon/v0.21.0
local_build=plugins/marathon/.build
failures=0

fail() {
  echo "FAIL: ${FUNCNAME[1]}: $*" >&2
  failures=1
}

set_up() {
  tmp=$(mktemp -d)
  root=$tmp/repo
  mkdir "$root"
  tar -C "$repo" --exclude=.git --exclude="$local_build" -cf - . | tar -C "$root" -xf -
  printf '#!/usr/bin/env bash\nexit 0\n' >"$root/scripts/check.sh"
  git -C "$root" init -q
  mkdir "$tmp/bin"
  stub_claude 0
}

stub_claude() {
  printf '#!/bin/sh\n"%s/plugins/marathon/bin/marathon" --version >"%s/eval-saw"\nexit %s\n' \
    "$root" "$tmp" "$1" >"$tmp/bin/claude"
  chmod +x "$tmp/bin/claude"
}

without_eval_cases() {
  rm -rf "$root/plugins/marathon/evals"
}

break_the_build() {
  echo 'func broken() {' >>"$root/marathon/cmd/marathon/main.go"
}

run_gate() {
  stderr=$(cd "$root" && PATH=$tmp/bin:$PATH scripts/gate.sh "$1" 2>&1 >/dev/null)
  code=$?
}

assert_status() {
  [ "$code" = "$1" ] || fail "got status $code, want $1; stderr:"$'\n'"$stderr"
}

assert_stderr_has() {
  grep -qxF -- "$1" <<<"$stderr" || fail "stderr lacks [$1]; stderr:"$'\n'"$stderr"
}

assert_no_local_build() {
  [ ! -e "$root/$local_build" ] || fail "$local_build is left behind"
}

test_a_good_build_passes_a_marathon_tag_without_eval_cases() {
  without_eval_cases
  run_gate "$tag"
  assert_status 0
  assert_stderr_has "gate: marathon has no eval cases, so $tag passes on scripts/check.sh and the release build"
  assert_no_local_build
}

test_a_broken_build_holds_a_marathon_tag_naming_the_fix() {
  without_eval_cases
  break_the_build
  run_gate "$tag"
  assert_status 1
  assert_stderr_has "dist: marathon/ does not build for linux/amd64; fix the lines above, then rerun"
  assert_stderr_has "gate: the release build failed, so $tag is held; fix the lines above, then rerun"
  assert_no_local_build
}

test_the_eval_runs_the_working_trees_tools_through_the_shims() {
  run_gate "$tag"
  assert_status 0
  [ "$(cat "$tmp/eval-saw" 2>/dev/null)" = "marathon 0.21.0" ] ||
    fail "the eval's marathon shim printed [$(cat "$tmp/eval-saw" 2>/dev/null)], want [marathon 0.21.0]"
  assert_no_local_build
}

test_a_failed_eval_still_removes_the_local_build() {
  stub_claude 1
  run_gate "$tag"
  assert_status 1
  [ -s "$tmp/eval-saw" ] || fail "the stub eval did not run"
  assert_no_local_build
}

test_the_gate_keeps_a_local_build_it_did_not_replace() {
  without_eval_cases
  mkdir -p "$root/$local_build"
  touch "$root/$local_build/kept"
  run_gate "$tag"
  assert_status 0
  [ -e "$root/$local_build/kept" ] || fail "the gate removed a local build it did not make"
}

test_another_plugins_tag_builds_no_marathon_binaries() {
  break_the_build
  run_gate marathon-architecture/v0.1.0
  assert_status 0
  assert_stderr_has "gate: marathon-architecture has no eval cases, so marathon-architecture/v0.1.0 passes on scripts/check.sh alone"
  assert_no_local_build
}

for test in $(declare -F | awk '$3 ~ /^test_/ { print $3 }'); do
  set_up
  "$test"
  rm -rf "$tmp"
done
exit "$failures"
