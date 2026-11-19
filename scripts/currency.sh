#!/usr/bin/env bash
# currency reports every pin in .github/workflows that trails its latest
# release, one stdout line each, and exits non-zero when it reports any. The
# pins are GitHub Actions, checked against the action's latest release tag, and
# npm packages installed globally, checked against the package's latest
# dist-tag, its latest release. A pin is current only when it equals
# that latest exactly. It reads the network through gh and npm and writes
# nothing; it is a development-time command and CI never runs it.
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

declare -A latest_dist_tag_cache=()
set_latest_to_dist_tag() {
  local package=$1
  if [ -z "${latest_dist_tag_cache[$package]+set}" ]; then
    latest_dist_tag_cache[$package]=$(npm view "$package" dist-tags.latest)
  fi
  latest=${latest_dist_tag_cache[$package]}
  if [ -z "$latest" ]; then
    echo "currency: no latest dist-tag for $package" >&2
    exit 1
  fi
}

first_global_npm_install_pin='npm (?:install|i) +(?:-g|--global) +\K[^\s"\x27]+@[^\s"\x27]+'
for file in .github/workflows/*.yml .github/workflows/*.yaml; do
  pins=$(grep_matches_or_none -oP "$first_global_npm_install_pin" "$file")
  pins=$(sort -u <<<"$pins")
  while read -r spec; do
    [ -n "$spec" ] || continue
    package_with_any_scope=${spec%@*}
    pin_after_last_at=${spec##*@}
    set_latest_to_dist_tag "$package_with_any_scope"
    [ "$pin_after_last_at" = "$latest" ] ||
      trailing+=("$file: $package_with_any_scope $pin_after_last_at -> $latest")
  done <<<"$pins"
done

if [ ${#trailing[@]} -gt 0 ]; then
  printf '%s\n' "${trailing[@]}"
  exit 1
fi
