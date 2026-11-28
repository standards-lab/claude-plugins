#!/usr/bin/env bash
# Repository consistency checks, run by .github/workflows/ci.yml and runnable locally. The check
# runs its three steps in order, each runnable alone: scripts/check/plugins.sh, the plugins'
# consistency; scripts/check/go.sh, the Go check over marathon/ and the lints it builds; and
# scripts/check/tests.sh, the tests in scripts/tests. Every step runs even when an earlier one
# fails, so the check reports every failure, and it exits 1 if any step fails.
set -uo pipefail

fail=0
for step in plugins go tests; do
  bash "$(dirname "$0")/check/$step.sh" || fail=1
done

if [ "$fail" -ne 0 ]; then
  exit 1
fi
echo "All checks passed."
