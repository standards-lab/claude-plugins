#!/usr/bin/env bash
# Repository consistency checks, run by .github/workflows/check.yml and runnable locally.
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

# Every marketplace source resolves to a plugin directory.
while read -r source; do
  [ -f "$source/.claude-plugin/plugin.json" ] ||
    err "marketplace source $source does not resolve to a plugin"
done < <(jq -r '.plugins[].source' .claude-plugin/marketplace.json)

# Every @-pointer and ./-link in the plugins' markdown resolves to a real file.
while read -r file; do
  dir=$(dirname "$file")
  while read -r target; do
    [ -e "$dir/$target" ] || err "$file: @$target does not resolve"
  done < <(grep -oP '^@\K\S+' "$file")
  while read -r target; do
    [ -e "$dir/$target" ] || err "$file: link $target does not resolve"
  done < <(grep -oP '\]\(\K\./[^)#]+' "$file")
done < <(find plugins -name '*.md')

if [ "$fail" -ne 0 ]; then
  exit 1
fi
echo "All checks passed."
