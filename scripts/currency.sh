#!/usr/bin/env bash
# currency reports every pin in .github/workflows that trails its latest
# release, one stdout line each, and exits non-zero when it reports any. The
# pins are GitHub Actions, checked against the action's latest release tag, and
# npm packages installed globally, checked against the package's stable
# dist-tag, the channel the pin follows. A pin is current only when it equals
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

# latest_stable sets latest to an npm package's stable dist-tag, looking each
# package up once. Like latest_tag, it runs in the current shell.
declare -A latest_stables=()
latest_stable() {
  local package=$1
  if [ -z "${latest_stables[$package]+set}" ]; then
    latest_stables[$package]=$(npm view "$package" dist-tags.stable)
  fi
  latest=${latest_stables[$package]}
  if [ -z "$latest" ]; then
    echo "currency: no stable dist-tag for $package" >&2
    exit 1
  fi
}

for file in .github/workflows/*.yml .github/workflows/*.yaml; do
  # grep exits 1 on no match, which is no pins; any other status fails.
  # A pin is the first package of an npm install -g, or its i and --global
  # spellings, with a version after a later @; quotes around a run line are not
  # part of it.
  pins=$(grep -oP 'npm (?:install|i) +(?:-g|--global) +\K[^\s"\x27]+@[^\s"\x27]+' "$file" || [ $? -eq 1 ])
  pins=$(sort -u <<<"$pins")
  while read -r spec; do
    [ -n "$spec" ] || continue
    # The version follows the last @, since a scoped package name begins with one.
    package=${spec%@*}
    pin=${spec##*@}
    latest_stable "$package"
    [ "$pin" = "$latest" ] || trailing+=("$file: $package $pin -> $latest")
  done <<<"$pins"
done

if [ ${#trailing[@]} -gt 0 ]; then
  printf '%s\n' "${trailing[@]}"
  exit 1
fi
