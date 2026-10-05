#!/usr/bin/env bash
# currency reports every GitHub Actions pin in .github/workflows that trails the
# action's latest release, one stdout line each, and exits non-zero when it
# reports any. A pin is current only when it equals the latest release tag
# exactly. It reads the network through gh and writes nothing; it is a
# development-time command and CI never runs it.
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

# latest_tag sets latest to the tag of a repository's latest release, looking
# each repository up once. It runs in the current shell, not a command
# substitution, so its cache persists and set -e fails currency when gh fails.
declare -A latest_tags=()
latest_tag() {
  local repo=$1
  if [ -z "${latest_tags[$repo]+set}" ]; then
    latest_tags[$repo]=$(gh api "repos/$repo/releases/latest" --jq .tag_name)
  fi
  latest=${latest_tags[$repo]}
  if [ -z "$latest" ]; then
    echo "currency: no latest release tag for $repo" >&2
    exit 1
  fi
}

for file in .github/workflows/*.yml .github/workflows/*.yaml; do
  # grep exits 1 on no match, which is no pins; any other status fails.
  pins=$(grep -oP 'uses: *\K[^ ]*@[^ ]*' "$file" || [ $? -eq 1 ])
  pins=$(sort -u <<<"$pins")
  while read -r uses; do
    [ -n "$uses" ] || continue
    action=${uses%@*}
    pin=${uses#*@}
    latest_tag "$(cut -d/ -f1,2 <<<"$action")"
    [ "$pin" = "$latest" ] || trailing+=("$file: $action $pin -> $latest")
  done <<<"$pins"
done

if [ ${#trailing[@]} -gt 0 ]; then
  printf '%s\n' "${trailing[@]}"
  exit 1
fi
