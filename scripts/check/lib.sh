# Helpers each step of scripts/check.sh sources: it moves to the repository root, err prints a FAIL
# line to stderr and marks the step failed, and finish exits 1 if any err ran.
cd "$(dirname "${BASH_SOURCE[0]}")/../.." || exit 1

fail=0
err() {
  echo "FAIL: $*" >&2
  fail=1
}

finish() {
  [ "$fail" -eq 0 ]
  exit
}
