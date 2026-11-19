#!/usr/bin/env bash
# Repository consistency checks, run by .github/workflows/ci.yml and runnable locally.
set -uo pipefail

fail=0
err() {
  echo "FAIL: $*" >&2
  fail=1
}

# Per plugin: the manifest version, the top CHANGELOG heading, and each skill's
# Version: line agree.
for plugin in plugins/*/; do
  manifest="$plugin.claude-plugin/plugin.json"
  name=$(jq -r .name "$manifest")
  version=$(jq -r .version "$manifest")

  top=$(grep -m1 -oP '^## v\K[0-9]+\.[0-9]+\.[0-9]+$' "${plugin}CHANGELOG.md")
  [ "$version" = "$top" ] ||
    err "$name: plugin.json version $version != CHANGELOG top heading ${top:-<none>};" \
      "give ${plugin}CHANGELOG.md a top '## v$version' heading or correct plugin.json"

  for skill in "$plugin"skills/*/SKILL.md; do
    skill_version=$(grep -m1 -oP '^Version: \K[0-9]+\.[0-9]+\.[0-9]+$' "$skill")
    [ "$version" = "$skill_version" ] ||
      err "$name: plugin.json version $version != $skill Version: ${skill_version:-<none>};" \
        "set its Version: line to $version or correct plugin.json"
  done
done

# Every marketplace source resolves to a plugin directory. The sources are read
# into a variable first, so a marketplace jq can't parse fails the check rather
# than yielding no sources to test.
if sources=$(jq -r '.plugins[].source' .claude-plugin/marketplace.json); then
  while read -r source; do
    [ -f "$source/.claude-plugin/plugin.json" ] ||
      err "marketplace source $source does not resolve to a plugin;" \
        "point it at a directory holding .claude-plugin/plugin.json"
  done <<<"$sources"
else
  err ".claude-plugin/marketplace.json does not parse; fix the lines above, then rerun"
fi

# Claude Code validates each plugin and the marketplace. CI installs claude at
# the version pinned in ci.yml; locally it is the one on PATH. A missing claude
# fails the check rather than skipping it, so a passing check always means the
# manifests were validated. --strict fails on warnings too, since the runtime
# tolerates what it warns about and the check is where it gets caught.
if command -v claude >/dev/null; then
  for target in plugins/*/ .; do
    # The report is printed only on failure, so a passing run stays quiet.
    if ! report=$(claude plugin validate --strict "$target" 2>&1); then
      echo "$report" >&2
      err "claude plugin validate failed for $target; fix the lines above, then rerun"
    fi
  done
else
  err "claude is not on PATH, so the plugins and marketplace were not validated;" \
    "install claude or put it on PATH, then rerun"
fi

# Claude Code expands an @-include line only in CLAUDE.md, never in a skill's
# markdown, so a plugin loads its files through links instead. Every ./ and ../
# link in the plugins' markdown resolves to a real file.
while read -r file; do
  dir=$(dirname "$file")
  while IFS=: read -r line _; do
    err "$file:$line: @-include line; link the file instead"
  done < <(grep -nP '^\s*@\S+\s*$' "$file")
  while read -r target; do
    [ -e "$dir/$target" ] ||
      err "$file: link $target does not resolve; point the link at an existing file or remove it"
  done < <(grep -oP '\]\(\K\.\.?/[^)#]+' "$file")
done < <(find plugins -name '*.md')

# Prints "line<TAB>owner<TAB>span" for each backtick span in a markdown file
# that reads as a file path: it contains a / and ends in an extension, with no
# < or * placeholder and no URL scheme. The owner is the word before a
# possessive just ahead of the span, as in marathon's `references/briefs.md`,
# even across a line break, or - when there is none. Fenced code blocks are
# skipped. A CHANGELOG is read only up to its first release heading, because
# released entries cite the tree as it stood then.
citations() {
  awk '
    /^[ \t]*(```|~~~)/ { fenced = !fenced; next }
    fenced { next }
    FILENAME ~ /CHANGELOG\.md$/ && /^## v/ { exit }
    {
      rest = carry " " $0
      carry = match($0, /[^ \t`]+$/) ? substr($0, RSTART) : ""
      while (match(rest, /`[^`]+`/)) {
        before = substr(rest, 1, RSTART - 1)
        span = substr(rest, RSTART + 1, RLENGTH - 2)
        rest = substr(rest, RSTART + RLENGTH)
        owner = "-"
        if (match(before, /[A-Za-z0-9-]+\047s[ \t]+$/)) {
          owner = substr(before, RSTART)
          sub(/\047s[ \t]+$/, "", owner)
        }
        if (span ~ /^[A-Za-z0-9._${}\/-]+\/[A-Za-z0-9._-]+\.[a-z]+$/ && span !~ /:\/\//)
          print FNR "\t" owner "\t" span
      }
    }
  ' "$1"
}

# Prints "root<TAB>dir" for the first directory in the given roots, taken in
# order and sorted within each, whose name is exactly one edit (an insertion,
# deletion, or substitution) away from name, and fails when there is none.
near_dir() {
  local name=$1 root
  shift
  for root in "$@"; do
    find "$root" -mindepth 1 -maxdepth 1 -type d -printf '%H\t%f\n' | sort
  done | awk -F '\t' -v name="$name" '
    # Past the first mismatch the names agree once the longer drops its extra
    # character, or, with equal lengths, once both drop the substituted one.
    function near(a, b, t, i) {
      if (length(a) < length(b)) { t = a; a = b; b = t }
      if (length(a) - length(b) > 1 || a == b) return 0
      for (i = 1; i <= length(b) && substr(a, i, 1) == substr(b, i, 1); i++) {}
      return substr(a, i + 1) == substr(b, i + (length(a) == length(b)))
    }
    !found && near(name, $2) { print; found = 1 }
    END { exit !found }
  '
}

# Every backtick file citation in the plugins' markdown names a real file.
# ${CLAUDE_PLUGIN_ROOT}/ paths resolve against the plugin, and ./ and ../ paths
# against the citing file. A bare path resolves inside a skill against the
# skill root and then the plugin root, so a skill can cite the plugin's agents/
# profiles, and outside one (profiles, README, CHANGELOG) against the plugin
# root or any of its skill roots. A path another plugin owns, written as
# marathon's `references/briefs.md`, resolves against that plugin's roots. A
# bare path is a plugin citation only when its first segment is a directory in
# one of those roots; anything else, such as context/roadmap.toml or
# .claude/marathon.toml, is a path in the consumer's repository.
#
# A first segment one edit away from a directory in those roots is a typo,
# such as refrences/ for references/, rather than a consumer path, so it fails
# instead of being skipped.
cite_fix="cite an existing file or remove the citation"
while read -r file; do
  plugin=$(echo "$file" | cut -d/ -f1-2)
  if [[ $file == "$plugin"/skills/*/* ]]; then
    roots=("$(echo "$file" | cut -d/ -f1-4)" "$plugin")
  else
    roots=("$plugin" "$plugin"/skills/*)
  fi
  while IFS=$'\t' read -r line owner cite; do
    bases=("${roots[@]}")
    if [ -d "plugins/$owner" ]; then
      bases=("plugins/$owner" "plugins/$owner"/skills/*)
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
    cited=0
    for root in "${bases[@]}"; do
      [ -d "$root/${cite%%/*}" ] || continue
      cited=1
      [ -f "$root/$cite" ] && continue 2
    done
    if [ "$cited" -eq 1 ]; then
      err "$file:$line: citation $cite does not resolve; $cite_fix"
    elif near=$(near_dir "${cite%%/*}" "${bases[@]}"); then
      IFS=$'\t' read -r root dir <<<"$near"
      err "$file:$line: citation $cite: directory '${cite%%/*}' is not in $root; did you mean '$dir'?"
    fi
  done < <(citations "$file")
done < <(find plugins -name '*.md')

# Word budgets, recorded in scripts/budgets.txt. Each is the size measured when
# it was set, rounded up to the next 100 words, so routine edits fit and growth
# beyond that is a deliberate budget change.
declare -A budget
while read -r path words; do
  budget[$path]=$words
done < <(grep -v '^\s*\(#\|$\)' scripts/budgets.txt)

# Prints the files a markdown file's Read first list links, as written.
read_first() {
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

# Fails when the words in a file or files exceed the budget for key. The
# message gives the budget to record: the size rounded up to the next 100.
within_budget() {
  local key=$1 words rounded
  shift
  words=$(cat "$@" | wc -w)
  rounded=$(((words + 99) / 100 * 100))
  if [ -z "${budget[$key]:-}" ]; then
    err "$key: $words words, no budget in scripts/budgets.txt;" \
      "add the line '$key $rounded' to it"
  elif [ "$words" -gt "${budget[$key]}" ]; then
    err "$key: $words words, over its budget of ${budget[$key]} in scripts/budgets.txt;" \
      "trim it or raise the budget to $rounded"
  fi
}

# A command's read-first set: its skill's SKILL.md, the playbook, and every
# file reached through Read first lists from there, each counted once.
# Each playbook has exactly one Read first line, so a renamed or duplicated
# line can't silently shrink the set the budget measures.
for playbook in plugins/*/skills/*/commands/*.md; do
  count=$(grep -c '^Read first' "$playbook")
  [ "$count" -eq 1 ] ||
    err "$playbook: has $count \"Read first\" lines; give it exactly one \"Read first, before acting:\" list"
  declare -A seen=()
  queue=("$playbook")
  set_files=("$(dirname "$(dirname "$playbook")")/SKILL.md")
  while [ ${#queue[@]} -gt 0 ]; do
    file=${queue[0]}
    queue=("${queue[@]:1}")
    [ -z "${seen[$file]:-}" ] && [ -f "$file" ] || continue
    seen[$file]=1
    set_files+=("$file")
    while read -r target; do
      queue+=("$(realpath -m --relative-to=. "$(dirname "$file")/$target")")
    done < <(read_first "$file")
  done
  within_budget "$playbook" "${set_files[@]}"
done

for profile in plugins/*/agents/*.md; do
  within_budget "$profile" "$profile"
done

# A skill's description is capped at 1,024 characters. It is measured as YAML
# reads it: the indented lines of a block or multi-line scalar join with one
# character each, so the raw lines would overcount. The trailing newline YAML
# keeps is not counted.
for skill in plugins/*/skills/*/SKILL.md; do
  length=$(awk '
    NR == 1 { next }
    /^---$/ { exit }
    joining && /^[ \t]/ { sub(/^[ \t]+/, ""); text = text (text == "" ? "" : " ") $0; next }
    joining { exit }
    /^description:/ {
      text = $0
      sub(/^description:[ \t]*/, "", text)
      if (text ~ /^[>|][-+]?$/) text = ""
      joining = 1
    }
    END { gsub(/^["\047]|["\047]$/, "", text); printf "%s", text }
  ' "$skill" | LC_ALL=C.UTF-8 wc -m)
  [ "$length" -le 1024 ] ||
    err "$skill: description is $length characters, over the 1024 limit;" \
      "shorten it to 1024 characters or fewer"
done

# lint-context runs from the working tree, not the installed plugin, so the
# check tests the copy being changed. -B keeps the tests from writing bytecode
# into the tree.
plugins/marathon/bin/lint-context || err "lint-context failed; fix the lines above, then rerun"
if ! report=$(python3 -B -m unittest discover -s scripts/tests 2>&1); then
  echo "$report" >&2
  err "scripts/tests failed; fix the lines above, then rerun"
fi

if [ "$fail" -ne 0 ]; then
  exit 1
fi
echo "All checks passed."
