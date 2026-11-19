#!/usr/bin/env bash
# Repository consistency checks, run by .github/workflows/ci.yml and runnable locally: each
# plugin's versions agree, the marketplace sources resolve, claude validates every plugin and
# the marketplace (a missing claude fails the check), the plugins' markdown has no @-include
# lines and its links and backtick file citations resolve, the word budgets in
# scripts/budgets.txt hold, every skill description fits in 1,024 characters, and lint-context
# and scripts/tests pass.
set -uo pipefail

fail=0
err() {
  echo "FAIL: $*" >&2
  fail=1
}

for plugin in plugins/*/; do
  manifest="$plugin.claude-plugin/plugin.json"
  name=$(jq -r .name "$manifest")
  version=$(jq -r .version "$manifest")

  changelog_top_version=$(grep -m1 -oP '^## v\K[0-9]+\.[0-9]+\.[0-9]+$' "${plugin}CHANGELOG.md")
  [ "$version" = "$changelog_top_version" ] ||
    err "$name: plugin.json version $version != CHANGELOG top heading" \
      "${changelog_top_version:-<none>};" \
      "give ${plugin}CHANGELOG.md a top '## v$version' heading or correct plugin.json"

  for skill in "$plugin"skills/*/SKILL.md; do
    skill_version=$(grep -m1 -oP '^Version: \K[0-9]+\.[0-9]+\.[0-9]+$' "$skill")
    [ "$version" = "$skill_version" ] ||
      err "$name: plugin.json version $version != $skill Version: ${skill_version:-<none>};" \
        "set its Version: line to $version or correct plugin.json"
  done
done

if marketplace_sources=$(jq -r '.plugins[].source' .claude-plugin/marketplace.json); then
  while read -r source; do
    [ -f "$source/.claude-plugin/plugin.json" ] ||
      err "marketplace source $source does not resolve to a plugin;" \
        "point it at a directory holding .claude-plugin/plugin.json"
  done <<<"$marketplace_sources"
else
  err ".claude-plugin/marketplace.json does not parse; fix the lines above, then rerun"
fi

if command -v claude >/dev/null; then
  for target in plugins/*/ .; do
    if ! failed_validation_report=$(claude plugin validate --strict "$target" 2>&1); then
      echo "$failed_validation_report" >&2
      err "claude plugin validate failed for $target; fix the lines above, then rerun"
    fi
  done
else
  err "claude is not on PATH, so the plugins and marketplace were not validated;" \
    "install claude or put it on PATH, then rerun"
fi

plugin_markdown() {
  find plugins -name '*.md'
}

at_include_line='^\s*@\S+\s*$'
relative_link_target='\]\(\K\.\.?/[^)#]+'

while read -r file; do
  dir=$(dirname "$file")
  while IFS=: read -r line _; do
    err "$file:$line: @-include line; link the file instead"
  done < <(grep -nP "$at_include_line" "$file")
  while read -r target; do
    [ -e "$dir/$target" ] ||
      err "$file: link $target does not resolve; point the link at an existing file or remove it"
  done < <(grep -oP "$relative_link_target" "$file")
done < <(plugin_markdown)

unreleased_path_citations_with_owner() {
  awk '
    /^[ \t]*(```|~~~)/ { in_fenced_code = !in_fenced_code; next }
    in_fenced_code { next }
    FILENAME ~ /CHANGELOG\.md$/ && /^## v/ { exit }
    {
      rest = previous_line_last_word " " $0
      previous_line_last_word = match($0, /[^ \t`]+$/) ? substr($0, RSTART) : ""
      while (match(rest, /`[^`]+`/)) {
        before_span = substr(rest, 1, RSTART - 1)
        span = substr(rest, RSTART + 1, RLENGTH - 2)
        rest = substr(rest, RSTART + RLENGTH)
        possessive_owner = "-"
        if (match(before_span, /[A-Za-z0-9-]+\047s[ \t]+$/)) {
          possessive_owner = substr(before_span, RSTART)
          sub(/\047s[ \t]+$/, "", possessive_owner)
        }
        is_path_with_extension = span ~ /^[A-Za-z0-9._${}\/-]+\/[A-Za-z0-9._-]+\.[a-z]+$/
        if (is_path_with_extension && span !~ /:\/\//)
          print FNR "\t" possessive_owner "\t" span
      }
    }
  ' "$1"
}

first_dir_one_edit_away() {
  local name=$1 root
  shift
  for root in "$@"; do
    find "$root" -mindepth 1 -maxdepth 1 -type d -printf '%H\t%f\n' | sort
  done | awk -F '\t' -v name="$name" '
    function one_edit_apart(longer, shorter, swap, first_mismatch, same_length) {
      if (length(longer) < length(shorter)) { swap = longer; longer = shorter; shorter = swap }
      if (length(longer) - length(shorter) > 1 || longer == shorter) return 0
      first_mismatch = 1
      while (first_mismatch <= length(shorter) &&
        substr(longer, first_mismatch, 1) == substr(shorter, first_mismatch, 1))
        first_mismatch++
      same_length = length(longer) == length(shorter)
      return substr(longer, first_mismatch + 1) == substr(shorter, first_mismatch + same_length)
    }
    !found && one_edit_apart(name, $2) { print; found = 1 }
    END { exit !found }
  '
}

cite_fix="cite an existing file or remove the citation"
while read -r file; do
  plugin=$(echo "$file" | cut -d/ -f1-2)
  if [[ $file == "$plugin"/skills/*/* ]]; then
    citing_file_roots=("$(echo "$file" | cut -d/ -f1-4)" "$plugin")
  else
    citing_file_roots=("$plugin" "$plugin"/skills/*)
  fi
  while IFS=$'\t' read -r line owner cite; do
    citation_roots=("${citing_file_roots[@]}")
    if [ -d "plugins/$owner" ]; then
      citation_roots=("plugins/$owner" "plugins/$owner"/skills/*)
    fi
    case $cite in
      '${CLAUDE_PLUGIN_ROOT}/'*)
        [ -f "$plugin/${cite#*/}" ] ||
          err "$file:$line: citation $cite does not resolve; $cite_fix"
        continue ;;
      ./* | ../*)
        [ -f "$(dirname "$file")/$cite" ] ||
          err "$file:$line: citation $cite does not resolve; $cite_fix"
        continue ;;
    esac
    first_segment=${cite%%/*}
    is_plugin_citation=0
    for root in "${citation_roots[@]}"; do
      [ -d "$root/$first_segment" ] || continue
      is_plugin_citation=1
      [ -f "$root/$cite" ] && continue 2
    done
    if [ "$is_plugin_citation" -eq 1 ]; then
      err "$file:$line: citation $cite does not resolve; $cite_fix"
    elif typo_of=$(first_dir_one_edit_away "$first_segment" "${citation_roots[@]}"); then
      IFS=$'\t' read -r root dir <<<"$typo_of"
      err "$file:$line: citation $cite: directory '$first_segment' is not in $root; did you mean '$dir'?"
    fi
  done < <(unreleased_path_citations_with_owner "$file")
done < <(plugin_markdown)

declare -A word_budget
while read -r path words; do
  word_budget[$path]=$words
done < <(grep -v '^\s*\(#\|$\)' scripts/budgets.txt)

read_first_list_links() {
  awk '
    /^Read first/ { listed = 1; next }
    !listed { next }
    /^[^ \t-]/ { exit }
    {
      line = $0
      while (match(line, /\]\(\.\.?\/[^)#]+/)) {
        print substr(line, RSTART + 2, RLENGTH - 2)
        line = substr(line, RSTART + RLENGTH)
      }
    }
  ' "$1"
}

within_budget() {
  local key=$1 words budget_to_record
  shift
  words=$(cat "$@" | wc -w)
  budget_to_record=$(((words + 99) / 100 * 100))
  if [ -z "${word_budget[$key]:-}" ]; then
    err "$key: $words words, no budget in scripts/budgets.txt;" \
      "add the line '$key $budget_to_record' to it"
  elif [ "$words" -gt "${word_budget[$key]}" ]; then
    err "$key: $words words, over its budget of ${word_budget[$key]} in scripts/budgets.txt;" \
      "trim it or raise the budget to $budget_to_record"
  fi
}

for playbook in plugins/*/skills/*/commands/*.md; do
  read_first_lines=$(grep -c '^Read first' "$playbook")
  [ "$read_first_lines" -eq 1 ] ||
    err "$playbook: has $read_first_lines \"Read first\" lines;" \
      "give it exactly one \"Read first, before acting:\" list"
  declare -A seen=()
  queue=("$playbook")
  read_first_set=("$(dirname "$(dirname "$playbook")")/SKILL.md")
  while [ ${#queue[@]} -gt 0 ]; do
    file=${queue[0]}
    queue=("${queue[@]:1}")
    [ -z "${seen[$file]:-}" ] && [ -f "$file" ] || continue
    seen[$file]=1
    read_first_set+=("$file")
    while read -r target; do
      queue+=("$(realpath -m --relative-to=. "$(dirname "$file")/$target")")
    done < <(read_first_list_links "$file")
  done
  within_budget "$playbook" "${read_first_set[@]}"
done

for profile in plugins/*/agents/*.md; do
  within_budget "$profile" "$profile"
done

description_as_yaml_reads_it() {
  awk '
    NR == 1 { next }
    /^---$/ { exit }
    in_description && /^[ \t]/ {
      sub(/^[ \t]+/, "")
      description = description (description == "" ? "" : " ") $0
      next
    }
    in_description { exit }
    /^description:/ {
      description = $0
      sub(/^description:[ \t]*/, "", description)
      if (description ~ /^[>|][-+]?$/) description = ""
      in_description = 1
    }
    END { gsub(/^["\047]|["\047]$/, "", description); printf "%s", description }
  ' "$1"
}

for skill in plugins/*/skills/*/SKILL.md; do
  description_length=$(description_as_yaml_reads_it "$skill" | LC_ALL=C.UTF-8 wc -m)
  [ "$description_length" -le 1024 ] ||
    err "$skill: description is $description_length characters, over the 1024 limit;" \
      "shorten it to 1024 characters or fewer"
done

working_tree_lint_context=plugins/marathon/bin/lint-context
"$working_tree_lint_context" || err "lint-context failed; fix the lines above, then rerun"
if ! failed_tests_report=$(python3 -B -m unittest discover -s scripts/tests 2>&1); then
  echo "$failed_tests_report" >&2
  err "scripts/tests failed; fix the lines above, then rerun"
fi

if [ "$fail" -ne 0 ]; then
  exit 1
fi
echo "All checks passed."
