#!/usr/bin/env bash
# currency reports every pin that trails its latest release, one stdout line
# each, and exits non-zero when it reports any. The pins are the tools in
# mise.toml and CI's mise.ci.toml, checked by mise outdated, and the GitHub
# Actions in .github/workflows, checked against the action's latest release tag.
# A pin is current only when it equals that latest exactly, and a pin with no
# latest version is not reported. It reads the network through mise and gh and
# writes nothing; it is a development-time command and CI never runs it.
#
# Each lookup captures its command's output in a variable before reading it, so
# under set -e a failed lookup fails currency instead of reporting nothing. The
# report is printed only once every lookup has succeeded, so a failure partway
# through exits nonzero with empty stdout, which the contract reads as a failed
# command rather than a partial list of what trails.
set -euo pipefail
shopt -s nullglob
cd "$(git -C "$(dirname "$0")" rev-parse --show-toplevel)"

trailing=()

grep_matches_or_none() {
  grep "$@" || [ $? -eq 1 ]
}

outdated_mise_tools=$(MISE_ENV=ci mise outdated --bump --local --json)
outdated_mise_tools=$(jq -r 'to_entries[]
  | (.value.bump // .value.latest) as $latest
  | select($latest != null and $latest != .value.requested)
  | "mise: \(.key) \(.value.requested) -> \($latest)"' <<<"$outdated_mise_tools")
while read -r line; do
  [ -n "$line" ] || continue
  trailing+=("$line")
done <<<"$outdated_mise_tools"

declare -A latest_release_tag_cache=()
set_latest_to_release_tag() {
  local repo=$1
  if [ -z "${latest_release_tag_cache[$repo]+set}" ]; then
    latest_release_tag_cache[$repo]=$(gh api "repos/$repo/releases/latest" --jq .tag_name)
  fi
  latest=${latest_release_tag_cache[$repo]}
  if [ -z "$latest" ]; then
    echo "currency: no latest release tag for $repo" >&2
    exit 1
  fi
}

action_pin='uses: *\K[^ ]*@[^ ]*'
for file in .github/workflows/*.yml .github/workflows/*.yaml; do
  pins=$(grep_matches_or_none -oP "$action_pin" "$file")
  pins=$(sort -u <<<"$pins")
  while read -r uses; do
    [ -n "$uses" ] || continue
    action=${uses%@*}
    pin=${uses#*@}
    set_latest_to_release_tag "$(cut -d/ -f1,2 <<<"$action")"
    [ "$pin" = "$latest" ] || trailing+=("$file: $action $pin -> $latest")
  done <<<"$pins"
done

if [ ${#trailing[@]} -gt 0 ]; then
  printf '%s\n' "${trailing[@]}"
  exit 1
fi
