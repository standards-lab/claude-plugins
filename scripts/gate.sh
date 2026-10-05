#!/usr/bin/env bash
# gate decides whether a Release tag may be pushed: marathon's [project] gate,
# which SHIP runs on the merge commit before each tag, with the tag as its one
# argument. The tag names a plugin and its version, as <plugin>/v<version>. The
# gate runs scripts/check.sh, then that plugin's eval suite, and exits 0 only
# when both pass. It is a development-time command and CI never runs it, since
# the eval suite runs paid claude sessions on the caller's credential.
#
# A plugin with no evals/ directory, or one holding no case, releases on the
# check alone, and the gate says so on stderr. A malformed tag, or one naming no
# plugin under plugins/, fails the gate rather than passing it unchecked.
set -uo pipefail
cd "$(git -C "$(dirname "$0")" rev-parse --show-toplevel)" || exit 1

fail() {
  echo "gate: $*" >&2
  exit 1
}

[ $# -eq 1 ] || fail "usage: scripts/gate.sh <plugin>/v<version>"
tag=$1
[[ "$tag" =~ ^([a-z0-9][a-z0-9-]*)/v[0-9]+\.[0-9]+\.[0-9]+$ ]] ||
  fail "malformed tag '$tag', expected <plugin>/v<major>.<minor>.<patch>"
plugin=${BASH_REMATCH[1]}
dir="plugins/$plugin"
[ -f "$dir/.claude-plugin/plugin.json" ] ||
  fail "unknown plugin '$plugin': no $dir/.claude-plugin/plugin.json"

scripts/check.sh || fail "scripts/check.sh failed, so $tag is held"

# A case is what claude plugin eval runs: a case.yaml, or a prompt.md with its
# graders, anywhere below evals/. The results/ directory eval writes holds no
# case, so a suite with only past results still counts as none.
cases=$(find "$dir/evals" -path "$dir/evals/results" -prune -o \
  \( -name case.yaml -o -name prompt.md \) -print 2>/dev/null)
if [ -z "$cases" ]; then
  echo "gate: $plugin has no eval cases, so $tag passes on scripts/check.sh alone" >&2
  exit 0
fi

# Every case must score 1.0 over its three runs (--threshold 1.0 --runs 3),
# against the plugin alone (--ablation none). --scaffold and --trust-plugin let the
# suite build its fixtures without a prompt, --no-publish keeps the results
# local, and --allow-tools EnterPlanMode grants the tool marathon's
# plan-round-in-reply case lists.
claude plugin eval "$dir" --threshold 1.0 --runs 3 --ablation none --scaffold \
  --trust-plugin --no-publish --allow-tools EnterPlanMode ||
  fail "the $plugin eval suite failed, so $tag is held"
echo "gate: $tag passes."
