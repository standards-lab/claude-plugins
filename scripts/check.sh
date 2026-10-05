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
