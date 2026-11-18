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
    err "$name: plugin.json version $version != CHANGELOG top heading ${top:-<none>}"

  for skill in "$plugin"skills/*/SKILL.md; do
    skill_version=$(grep -m1 -oP '^Version: \K[0-9]+\.[0-9]+\.[0-9]+$' "$skill")
    [ "$version" = "$skill_version" ] ||
      err "$name: plugin.json version $version != $skill Version: ${skill_version:-<none>}"
  done
done

# Every marketplace source resolves to a plugin directory. The sources are read
# into a variable first, so a marketplace jq can't parse fails the check rather
# than yielding no sources to test.
if sources=$(jq -r '.plugins[].source' .claude-plugin/marketplace.json); then
  while read -r source; do
    [ -f "$source/.claude-plugin/plugin.json" ] ||
      err "marketplace source $source does not resolve to a plugin"
  done <<<"$sources"
else
  err ".claude-plugin/marketplace.json does not parse"
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
      err "claude plugin validate failed for $target"
    fi
  done
else
  err "claude is not on PATH, so the plugins and marketplace were not validated"
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
    [ -e "$dir/$target" ] || err "$file: link $target does not resolve"
  done < <(grep -oP '\]\(\K\.\.?/[^)#]+' "$file")
done < <(find plugins -name '*.md')

# Prints "line<TAB>owner<TAB>span" for each backtick span in a markdown file
# that reads as a file path: it contains a / and ends in an extension, with no
# < or * placeholder and no URL scheme. The owner is the word before a
# possessive just ahead of the span, as in marathon's `references/briefs.md`,
# even across a line break, and - when there is none. Fenced code blocks are skipped. A CHANGELOG is
# read only up to its first release heading, because released entries cite the
# tree as it stood then.
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

# Every backtick file citation in the plugins' markdown names a real file.
# ${CLAUDE_PLUGIN_ROOT}/ paths resolve against the plugin, and ./ and ../ paths
# against the citing file. A bare path resolves against the skill root inside a
# skill, and outside one (profiles, README, CHANGELOG) against the plugin root
# or any of its skill roots. A path another plugin owns, written as
# marathon's `references/briefs.md`, resolves against that plugin's roots. A
# bare path is a plugin citation only when its first segment is a directory in
# one of those roots; anything else, such as context/roadmap.toml or
# .claude/marathon.toml, is a path in the consumer's repository.
while read -r file; do
  plugin=$(echo "$file" | cut -d/ -f1-2)
  if [[ $file == "$plugin"/skills/*/* ]]; then
    roots=("$(echo "$file" | cut -d/ -f1-4)")
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
        [ -f "$plugin/${cite#*/}" ] || err "$file:$line: citation $cite does not resolve"
        continue ;;
      ./* | ../*)
        [ -f "$(dirname "$file")/$cite" ] || err "$file:$line: citation $cite does not resolve"
        continue ;;
    esac
    cited=0
    for root in "${bases[@]}"; do
      [ -d "$root/${cite%%/*}" ] || continue
      cited=1
      [ -f "$root/$cite" ] && continue 2
    done
    [ "$cited" -eq 0 ] || err "$file:$line: citation $cite does not resolve"
  done < <(citations "$file")
done < <(find plugins -name '*.md')

if [ "$fail" -ne 0 ]; then
  exit 1
fi
echo "All checks passed."
