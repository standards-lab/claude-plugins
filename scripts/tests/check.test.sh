#!/usr/bin/env bash
# Tests for scripts/check.sh, run against a seeded copy of the repository.
#
# The outer check runs the real test suite and validates the plugins, so the copy's scripts/tests
# holds one passing test in their place, which keeps the check from recursing into these tests, and
# a stub claude goes first on PATH. The copy is a git repository, so lint-references and the
# Python check read its tracked files. Each test_ function is one case; the script prints a FAIL
# line per failed assertion and exits 1 if any fails.
set -uo pipefail

repo=$(cd "$(dirname "$0")/../.." && pwd)
skill=plugins/marathon/skills/marathon
unbudgeted_seed=$skill/references/seeded.md
read_first_fix='give it exactly one "Read first, before acting:" list'
rerun="fix the lines above, then rerun"
extension=plugins/marathon-architecture
check_tools_except_claude=(
  as awk bash cat cut dirname find gcc git go gofmt golangci-lint grep head jq ld mkdir mktemp node
  realpath rm sort wc
)
failures=0

set_up() {
  tmp=$(mktemp -d)
  root=$tmp/repo
  mkdir "$root"
  tar -C "$repo" --exclude=.git -cf - . | tar -C "$root" -xf -
  rm -rf "$root/scripts/tests"
  mkdir "$root/scripts/tests"
  printf '#!/usr/bin/env bash\nexit 0\n' >"$root/scripts/tests/passes.test.sh"
  git -C "$root" init -q
  git -C "$root" add -A
  mkdir "$tmp/bin"
  write_executable "$tmp/bin/claude" $'#!/bin/sh\nexit 0'
  check_path=$tmp/bin:$PATH
}

write_executable() {
  printf '%s\n' "$2" >"$1"
  chmod +x "$1"
}

write() {
  printf '%s' "$2" >"$root/$1"
}

append() {
  printf '%s' "$2" >>"$root/$1"
}

replace() {
  local path=$root/$1 old=$2 new=$3 content rest
  content=$(
    cat "$path"
    printf x
  )
  content=${content%x}
  rest=${content//"$old"/}
  if [ $(((${#content} - ${#rest}) / ${#old})) -ne 1 ]; then
    echo "FAIL: ${FUNCNAME[1]}: $1 does not hold [$old] exactly once" >&2
    failures=1
  fi
  printf '%s' "${content/"$old"/"$new"}" >"$path"
}

run_check() {
  local stderr
  stderr=$(cd "$root" && PATH=$check_path scripts/check.sh 2>&1 >/dev/null)
  code=$?
  lines=$stderr
}

expect() {
  local expected_code=$1 expected_lines
  shift
  expected_lines=$(printf '%s\n' "$@")
  if [ "$code" != "$expected_code" ] || [ "$lines" != "${expected_lines%$'\n'}" ]; then
    printf 'FAIL: %s: got status %s and stderr:\n%s\nwant status %s and stderr:\n%s\n' \
      "${FUNCNAME[1]}" "$code" "$lines" "$expected_code" "${expected_lines%$'\n'}" >&2
    failures=1
  fi
}

test_playbook_needs_exactly_one_read_first_line() {
  append "$skill/commands/status.md" $'\nRead first, before acting:\n'
  replace "$skill/commands/reset.md" $'\nRead first' $'\nRead these first'
  run_check
  expect 1 \
    "FAIL: $skill/commands/reset.md: has 0 \"Read first\" lines; $read_first_fix" \
    "FAIL: $skill/commands/status.md: has 2 \"Read first\" lines; $read_first_fix"
}

test_playbook_citing_an_unlisted_format_fails() {
  replace "$skill/commands/retro.md" \
    $'- [`../references/briefs/retro.md`](../references/briefs/retro.md)\n' ""
  run_check
  expect 1 "FAIL: $skill/commands/retro.md: cites references/briefs/retro.md, which its read-first\
 set lacks; add it to its Read first list"
}

test_listed_format_citing_an_unlisted_format_fails() {
  replace "$skill/commands/intake.md" \
    $'- [`../references/briefs/plan-round.md`](../references/briefs/plan-round.md)\n' ""
  run_check
  expect 1 "FAIL: $skill/commands/intake.md: its references/briefs/intake-round.md cites\
 references/briefs/plan-round.md, which its read-first set lacks; add it to its Read first list"
}

test_citing_formats_the_read_first_set_reaches_passes() {
  replace "$skill/commands/experiment.md" "(\`commands/plan.md\`)" \
    "(\`commands/plan.md\`, \`references/briefs/plan-round.md\`)"
  run_check
  expect 0
}

test_typo_in_first_directory_fails() {
  write "$unbudgeted_seed" $'# Seeded\n\nSee `refrences/build.md`.\n'
  run_check
  expect 1 "FAIL: $unbudgeted_seed:3: citation refrences/build.md: directory 'refrences' is not in\
 $skill; did you mean 'references'?"
}

test_skill_cites_plugin_agents() {
  write "$unbudgeted_seed" $'# Seeded\n\nSee `agents/planner.md` and `agents/nobody.md`.\n'
  run_check
  expect 1 "FAIL: $unbudgeted_seed:3: citation agents/nobody.md does not resolve;\
 cite an existing file or remove the citation"
}

test_consumer_paths_pass() {
  write "$unbudgeted_seed" \
    $'# Seeded\n\nSee `roadmap/index.toml`, `scripts/check.sh`, and `src/app/main.py`.\n'
  run_check
  expect 0
}

test_version_mismatch_names_the_files_to_align() {
  local manifest=$root/$extension/.claude-plugin/plugin.json version
  version=$(jq -r .version "$manifest")
  jq '.version = "9.9.9"' "$manifest" >"$tmp/plugin.json" && mv "$tmp/plugin.json" "$manifest"
  run_check
  expect 1 \
    "FAIL: marathon-architecture: plugin.json version 9.9.9 != CHANGELOG top heading $version;\
 give $extension/CHANGELOG.md a top '## v9.9.9' heading or correct plugin.json" \
    "FAIL: marathon-architecture: plugin.json version 9.9.9 !=\
 $extension/skills/marathon-architecture/SKILL.md Version: $version; set its Version: line to\
 9.9.9 or correct plugin.json"
}

test_shim_version_mismatch_names_the_fix() {
  local version
  version=$(jq -r .version "$root/plugins/marathon/.claude-plugin/plugin.json")
  replace plugins/marathon/bin/wrap-formats $'\nversion='"$version"$'\n' $'\nversion=0.0.1\n'
  run_check
  expect 1 "FAIL: marathon: plugin.json version $version != plugins/marathon/bin/wrap-formats\
 version=0.0.1; set its version= line to $version or correct plugin.json"
}

test_tracked_python_fails() {
  write scripts/seeded.py $'print("seeded")\n'
  write scripts/seeded-tool $'#!/usr/bin/env python3\nprint("seeded")\n'
  write scripts/untracked.py $'print("untracked")\n'
  git -C "$root" add scripts/seeded.py scripts/seeded-tool
  run_check
  expect 1 \
    "FAIL: scripts/seeded-tool: a tracked Python file; port it to Go, bash, or a node --test file,\
 then remove it" \
    "FAIL: scripts/seeded.py: a tracked Python file; port it to Go, bash, or a node --test file,\
 then remove it"
}

test_marketplace_source_must_be_a_plugin() {
  replace .claude-plugin/marketplace.json "./$extension" ./plugins/gone
  run_check
  expect 1 "FAIL: marketplace source ./plugins/gone does not resolve to a plugin;\
 point it at a directory holding .claude-plugin/plugin.json"
}

test_unresolved_link_fails() {
  write "$unbudgeted_seed" $'# Seeded\n\nSee [gone](./gone.md).\n'
  run_check
  expect 1 "FAIL: $unbudgeted_seed: link ./gone.md does not resolve;\
 point the link at an existing file or remove it"
}

test_unresolved_reference_fails() {
  append scripts/gate.sh $'# See scripts/gone.md\n'
  run_check
  expect 1 "FAIL: lint-references failed; $rerun"
}

test_eval_results_markdown_is_not_checked() {
  local results=$root/plugins/marathon/evals/results/run
  mkdir -p "$results"
  printf 'See [gone](./gone.md).\n' >"$results/case-with-r1.last-message.md"
  run_check
  expect 0
}

test_budget_messages_give_the_budget_to_record() {
  local implementer=plugins/marathon/agents/implementer.md planner=plugins/marathon/agents/planner.md
  replace scripts/budgets.txt "$implementer 300"$'\n' "$implementer 100"$'\n'
  replace scripts/budgets.txt "$planner 300"$'\n' ""
  local implementer_words planner_words
  implementer_words=$(wc -w <"$root/$implementer")
  planner_words=$(wc -w <"$root/$planner")
  run_check
  expect 1 \
    "FAIL: $implementer: $implementer_words words, over its budget of 100 in scripts/budgets.txt;\
 trim it or raise the budget to $(((implementer_words + 99) / 100 * 100))" \
    "FAIL: $planner: $planner_words words, no budget in scripts/budgets.txt; add the line\
 '$planner $(((planner_words + 99) / 100 * 100))' to it"
}

test_long_description_fails() {
  local extension_skill=$extension/skills/marathon-architecture/SKILL.md long
  long=$(printf 'x%.0s' {1..1024})
  replace "$extension_skill" $'description: >\n' $'description: >\n  '"$long"$'\n'
  run_check
  local pattern="^FAIL: $extension_skill: description is [0-9]+ characters, over the 1024 limit;\
 shorten it to 1024 characters or fewer$"
  if [ "$code" != 1 ] || ! [[ $lines =~ $pattern ]]; then
    echo "FAIL: ${FUNCNAME[0]}: got status $code and stderr: $lines" >&2
    failures=1
  fi
}

test_extension_skill_must_not_be_user_invocable() {
  local extension_skill=$extension/skills/marathon-architecture/SKILL.md
  run_check
  expect 0
  replace "$extension_skill" $'user-invocable: false\n' ""
  run_check
  expect 1 "FAIL: $extension_skill: a marathon extension that users can invoke;\
 add 'user-invocable: false' to its front matter"
}

test_release_heading_carries_only_its_version() {
  local changelog=$extension/CHANGELOG.md
  replace "$changelog" $'## v0.2.2\n' $'## Unreleased (next)\n\n## v0.2.2 - 2026-01-02\n'
  replace "$changelog" $'## v0.2.1\n' $'## [v0.2.1](https://example.com/v0.2.1)\n'
  local fix="carries more than its version; write it as '## vX.Y.Z' or '## Unreleased', leaving\
 dates and links to the GitHub release"
  local heading expected=()
  for heading in "## Unreleased (next)" "## v0.2.2 - 2026-01-02" \
    "## [v0.2.1](https://example.com/v0.2.1)"; do
    expected+=("FAIL: $changelog:$(grep -nxF -m1 "$heading" "$root/$changelog" | cut -d: -f1):\
 release heading '$heading' $fix")
  done
  run_check
  expect 1 "${expected[@]}"
}

test_changelog_defines_no_link_references() {
  local changelog=$extension/CHANGELOG.md line
  line=$(($(wc -l <"$root/$changelog") + 2))
  append "$changelog" $'\n[v0.1.0]: https://example.com/v0.1.0\n'
  run_check
  expect 1 "FAIL: $changelog:$line: link reference definition; write the link inline instead"
}

test_missing_claude_is_the_only_failure() {
  local tools=$tmp/tools tool
  mkdir "$tools"
  for tool in "${check_tools_except_claude[@]}"; do
    ln -s "$(command -v "$tool")" "$tools/$tool"
  done
  check_path=$tools
  run_check
  expect 1 "FAIL: claude is not on PATH, so the plugins and marketplace were not validated;\
 install claude or put it on PATH, then rerun"
}

test_wrapped_tools_point_at_their_output() {
  write_executable "$tmp/bin/claude" $'#!/bin/sh\necho invalid\nexit 1'
  write .claude-plugin/marketplace.json "{"
  write context/seeded.md $'[gone](gone.md)\n'
  write scripts/tests/seeded.test.sh $'#!/usr/bin/env bash\nexit 1\n'
  run_check
  local actual expected
  actual=$(grep '^FAIL: ' <<<"$lines" | sort)
  expected=$(sort <<EOF
FAIL: .claude-plugin/marketplace.json does not parse; $rerun
FAIL: claude plugin validate failed for plugins/marathon/; $rerun
FAIL: claude plugin validate failed for $extension/; $rerun
FAIL: claude plugin validate failed for .; $rerun
FAIL: lint-context failed; $rerun
FAIL: scripts/tests failed; $rerun
EOF
  )
  if [ "$code" != 1 ] || [ "$actual" != "$expected" ]; then
    printf 'FAIL: %s: got status %s and FAIL lines:\n%s\nwant status 1 and FAIL lines:\n%s\n' \
      "${FUNCNAME[0]}" "$code" "$actual" "$expected" >&2
    failures=1
  fi
}

for test in $(declare -F | awk '$3 ~ /^test_/ { print $3 }'); do
  set_up
  "$test"
  rm -rf "$tmp"
done
exit "$failures"
