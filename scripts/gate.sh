#!/usr/bin/env bash
# gate decides whether a Release tag may be pushed: marathon's [project] gate,
# which SHIP runs on the merge commit before each tag, with the tag as its one
# argument. The tag names a plugin and its version, as <plugin>/v<version>. The
# gate runs scripts/check.sh, then that plugin's eval suite, and exits 0 only
# when both pass. It is a development-time command and CI never runs it, since
# the eval suite runs paid claude sessions on the caller's credential.
#
# For a marathon tag, the gate also builds the release's assets with
# scripts/dist.sh, the build the release workflow publishes, into a temporary
# directory, stamped with the tag's version, after the check, and fails when a
# target doesn't build. Before the eval it builds the host's commands into
# plugins/marathon/.build/, the local build the plugin's bin/ shims run, so the
# eval's sessions run the working tree's tools; it removes that directory, and
# any local build it replaced, when it exits, pass or fail. Another plugin's
# tag ships no binaries, so its gate builds neither.
#
# The suite passes when every case scores 1.0 over its three runs, against the
# plugin alone (--ablation none); it builds its fixtures without a prompt
# (--scaffold --trust-plugin), and keeps its results local (--no-publish). It
# grants Bash, Edit, and Write (--allow-tools), which eval withholds without an
# operator's grant; each case still gets only the tools its allowed_tools lists.
# The gate's pass rests on deterministic graders alone, and a gated suite holds no
# llm grader: eval keeps no judge reasoning and can't re-grade a stored reply,
# so no one can reproduce a judge's verdict. A plugin keeps its judges in a
# suite the gate doesn't run, as marathon does in evals-judged/.
#
# eval's results hold scores but not what a run produced, so the gate keeps
# each run's scaffold (--keep-temp), copies the run's trace and final message
# into the results directory beside aggregate-result.json, as
# <case>-<arm>-r<run>.trace.jsonl and <case>-<arm>-r<run>.last-message.md, and
# then removes the scaffold, but only a trace path of the shape eval writes,
# <absolute dir>/claude-eval-<id>/out/trace.jsonl with no . or .. component. Every failed run can be read back from there.
#
# A plugin with no eval directory, or one holding no case, releases on the
# check alone, a marathon tag on the check and the release build, and the gate
# says so on stderr. A malformed tag, or one naming no plugin under plugins/,
# fails the gate rather than passing it unchecked.
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

local_build=plugins/marathon/.build
if [ "$plugin" = marathon ]; then
  version=${tag#marathon/v}
  dist=$(mktemp -d) || fail "could not create a directory for the release build"
  trap 'rm -rf "$dist"' EXIT
  scripts/dist.sh "$version" "$dist" ||
    fail "the release build failed, so $tag is held; fix the lines above, then rerun"
fi

eval_dir="$dir/$(jq -r '.experimental.evals // "evals"' "$dir/.claude-plugin/plugin.json")" ||
  fail "$dir/.claude-plugin/plugin.json does not parse"
cases_outside_results=
if [ -d "$eval_dir" ]; then
  cases_outside_results=$(find "$eval_dir" -path "$eval_dir/results" -prune -o \
    \( -name case.yaml -o -name prompt.md \) -print) ||
    fail "could not list the cases in $eval_dir"
fi
if [ -z "$cases_outside_results" ]; then
  passed_on="scripts/check.sh alone"
  [ "$plugin" = marathon ] && passed_on="scripts/check.sh and the release build"
  echo "gate: $plugin has no eval cases, so $tag passes on $passed_on" >&2
  exit 0
fi

keep_run_evidence() {
  local results=$1 case arm run trace scaffold
  [ -f "$results/aggregate-result.json" ] || return
  jq -r '.cases[] | .name as $case | .arms | to_entries[] | .key as $arm | .value
    | to_entries[] | select((.value.tracePath // "") != "")
    | [$case, $arm, .key + 1, .value.tracePath] | @tsv' "$results/aggregate-result.json" |
    while IFS=$'\t' read -r case arm run trace; do
      cp "$trace" "$results/$case-$arm-r$run.trace.jsonl"
      jq -rn '[inputs | select(.type == "result")] | last | .result // empty' "$trace" \
        >"$results/$case-$arm-r$run.last-message.md"
      [[ $trace =~ ^(/([^./][^/]*/)*claude-eval-[^/]+)/out/trace\.jsonl$ ]] || continue
      scaffold=${BASH_REMATCH[1]}
      chmod -R u+rwx "$scaffold" && rm -rf "$scaffold"
    done
}

if [ "$plugin" = marathon ]; then
  trap 'rm -rf "$dist" "$local_build"' EXIT
  rm -rf "$local_build"
  scripts/build.sh "$local_build" "$version" ||
    fail "the local build for the eval failed, so $tag is held; fix the lines above, then rerun"
fi

results="$eval_dir/results/$(date -u +%Y-%m-%dT%H-%M-%S-%3NZ)"
eval_status=0
claude plugin eval "$dir" --threshold 1.0 --runs 3 --ablation none --scaffold \
  --trust-plugin --allow-tools Bash Edit Write --no-publish --keep-temp \
  --output-dir "$results" || eval_status=$?
keep_run_evidence "$results"
[ "$eval_status" -eq 0 ] ||
  fail "the $plugin eval suite failed, so $tag is held; each run's reply is in $results"
echo "gate: $tag passes."
