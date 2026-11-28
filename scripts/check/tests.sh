#!/usr/bin/env bash
# The check's tests step: scripts/tests pass, its node --test files given a build of marathon/ by
# scripts/build.sh as MARATHON_BUILD, and its *.test.sh scripts.
set -uo pipefail
source "$(dirname "$0")/lib.sh"

build=$(mktemp -d)
trap 'rm -rf "$build"' EXIT

run_tests() {
  local status=0 test
  scripts/build.sh "$build" || status=1
  MARATHON_BUILD=$build node --test 'scripts/tests/*.test.mjs' || status=1
  for test in scripts/tests/*.test.sh; do
    [ -f "$test" ] || continue
    bash "$test" || status=1
  done
  return "$status"
}

if ! failed_tests_report=$(run_tests 2>&1); then
  echo "$failed_tests_report" >&2
  err "scripts/tests failed; fix the lines above, then rerun"
fi

finish
