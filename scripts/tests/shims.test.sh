#!/usr/bin/env bash
# Tests for marathon's bin/ shims and the MessageDisplay hook that runs one, on a copy of the
# plugin's bin/ and hooks/ beside stub commands. A stub in .build/ stands for the local build, and
# stubs in release/ for the release's binaries. A stub mise first on PATH records each call with its
# working directory and the MISE_OVERRIDE_ settings the shim hands it; its where prints release/
# once the release is installed, which a call to its install, or an installed file, marks. Each
# test_ function is one case; the script prints a FAIL line per failed assertion and exits 1 if
# any fails.
set -uo pipefail

repo=$(cd "$(dirname "$0")/../.." && pwd)
plugin_version=$(jq -r .version "$repo/plugins/marathon/.claude-plugin/plugin.json")
release="github:standards-lab/claude-plugins[version_prefix=marathon/v]@$plugin_version"
commands=(lint-context lint-references marathon wrap-formats)
failures=0

fail() {
  echo "FAIL: ${FUNCNAME[1]}: $*" >&2
  failures=1
}

assert_equal() {
  local what=$1 actual=$2 expected=$3
  [ "$actual" = "$expected" ] || {
    echo "FAIL: ${FUNCNAME[1]}: $what: got [$actual], want [$expected]" >&2
    failures=1
  }
}

write_stub() {
  local path=$1 body=$2
  mkdir -p "$(dirname "$path")"
  printf '#!/usr/bin/env bash\n%s\n' "$body" >"$path"
  chmod +x "$path"
}

set_up() {
  tmp=$(mktemp -d)
  plugin=$tmp/plugin
  mkdir -p "$plugin" "$tmp/caller" "$tmp/stubs"
  cp -R "$repo/plugins/marathon/bin" "$repo/plugins/marathon/hooks" "$plugin/"
  write_stub "$tmp/stubs/mise" "overrides=\$(env | grep '^MISE_OVERRIDE_' | sort)
echo \"\$PWD \$* |\" \$overrides >>'$tmp/mise-calls'
echo \"\$overrides\" >'$tmp/mise-env'
case \$1 in
where)
  [ -e '$tmp/installed' ] || { echo 'mise ERROR not installed' >&2; exit 1; }
  echo '$tmp/release' ;;
install) touch '$tmp/installed' ;;
esac"
  for command in "${commands[@]}"; do
    write_stub "$tmp/release/$command" \
      "echo \"$command \$PWD \$# \$*\"; env | grep '^MISE_OVERRIDE_' | sort"
  done
  stubbed_path=$tmp/stubs:$PATH
}

mise_calls() {
  cat "$tmp/mise-calls" 2>/dev/null
}

overrides_on_mise_calls="| MISE_OVERRIDE_CONFIG_FILENAMES=no-project-config MISE_OVERRIDE_TOOL_VERSIONS_FILENAMES=none"

mise_overrides() {
  cat "$tmp/mise-env" 2>/dev/null
}

with_the_shims_overrides() {
  local -a overrides
  mapfile -t overrides < <(mise_overrides)
  (cd "$tmp/caller" && env "${overrides[@]}" MISE_TRUSTED_CONFIG_PATHS="$tmp/caller" "$@")
}

hook_command() {
  jq -r '.hooks.MessageDisplay[0].hooks[0].command' "$plugin/hooks/hooks.json"
}

run_hook() {
  (cd "$tmp/caller" && CLAUDE_PLUGIN_ROOT=$plugin PATH=$1 sh -c "$(hook_command)")
}

test_each_command_runs_its_local_build_in_the_callers_directory() {
  for command in "${commands[@]}"; do
    write_stub "$plugin/.build/$command" "echo \"$command \$PWD \$# \$*\""
    local output
    output=$(cd "$tmp/caller" && PATH=$stubbed_path "$plugin/bin/$command" a 'b c')
    assert_equal "$command output" "$output" "$command $tmp/caller 2 a b c"
  done
  assert_equal "mise calls" "$(mise_calls)" ""
}

test_each_command_without_a_local_build_execs_the_installed_release_at_the_plugins_version() {
  touch "$tmp/installed"
  for command in "${commands[@]}"; do
    rm -f "$tmp/mise-calls"
    local output errors
    output=$(cd "$tmp/caller" && PATH=$stubbed_path "$plugin/bin/$command" a 'b c' 2>"$tmp/errors")
    errors=$(cat "$tmp/errors")
    assert_equal "$command output" "$output" "$command $tmp/caller 2 a b c"
    assert_equal "$command errors" "$errors" ""
    assert_equal "$command: mise calls" "$(mise_calls)" \
      "$tmp/caller where $release $overrides_on_mise_calls"
  done
}

test_each_command_installs_a_release_not_yet_installed_quietly() {
  for command in "${commands[@]}"; do
    rm -f "$tmp/mise-calls" "$tmp/installed"
    local output errors
    output=$(cd "$tmp/caller" && PATH=$stubbed_path "$plugin/bin/$command" a 'b c' 2>"$tmp/errors")
    errors=$(cat "$tmp/errors")
    assert_equal "$command output" "$output" "$command $tmp/caller 2 a b c"
    assert_equal "$command errors" "$errors" ""
    assert_equal "$command: mise calls" "$(mise_calls)" \
      "$tmp/caller where $release $overrides_on_mise_calls
$tmp/caller install $release $overrides_on_mise_calls
$tmp/caller where $release $overrides_on_mise_calls"
  done
}

test_a_failed_install_surfaces_and_fails_the_command() {
  write_stub "$tmp/stubs/mise" \
    "[ \"\$1\" = install ] && echo 'mise ERROR download failed' >&2; exit 1"
  for command in "${commands[@]}"; do
    local output errors status
    output=$(cd "$tmp/caller" && PATH=$stubbed_path "$plugin/bin/$command" 2>"$tmp/errors")
    status=$?
    errors=$(cat "$tmp/errors")
    [ "$status" -ne 0 ] || fail "$command: status 0, want nonzero"
    assert_equal "$command output" "$output" ""
    assert_equal "$command errors" "$errors" "mise ERROR download failed"
  done
}

test_the_release_sees_the_callers_environment_unchanged() {
  touch "$tmp/installed"
  for command in "${commands[@]}"; do
    local output
    output=$(cd "$tmp/caller" && PATH=$stubbed_path "$plugin/bin/$command" a)
    assert_equal "$command: the release's output without caller overrides" "$output" \
      "$command $tmp/caller 1 a"
    rm -f "$tmp/mise-calls"
    output=$(cd "$tmp/caller" && PATH=$stubbed_path MISE_OVERRIDE_CONFIG_FILENAMES=mine.toml \
      "$plugin/bin/$command" a)
    assert_equal "$command: the release's output with the caller's override" "$output" \
      "$command $tmp/caller 1 a
MISE_OVERRIDE_CONFIG_FILENAMES=mine.toml"
    assert_equal "$command: mise calls under the caller's override" "$(mise_calls)" \
      "$tmp/caller where $release $overrides_on_mise_calls"
  done
}

test_the_release_ignores_the_callers_mise_config() {
  command -v mise >/dev/null || {
    fail "mise is not on PATH; install mise, then rerun"
    return
  }
  printf '[env]\nFROM_CALLERS_CONFIG = "read"\n' >"$tmp/caller/mise.toml"
  local config_value='echo "${FROM_CALLERS_CONFIG:-unset}"'
  local read_by_default
  read_by_default=$(cd "$tmp/caller" && MISE_TRUSTED_CONFIG_PATHS=$tmp/caller \
    mise exec -- sh -c "$config_value")
  assert_equal "the config read without an override" "$read_by_default" "read"
  for command in "${commands[@]}"; do
    (cd "$tmp/caller" && PATH=$stubbed_path "$plugin/bin/$command" >/dev/null)
    local ignored
    ignored=$(with_the_shims_overrides mise exec -- sh -c "$config_value")
    assert_equal "$command: the config read with the shim's overrides" "$ignored" "unset"
  done
}

test_the_release_ignores_the_callers_tool_versions() {
  command -v mise >/dev/null || {
    fail "mise is not on PATH; install mise, then rerun"
    return
  }
  echo 'nonexistent-tool-xyz 9.9.9' >"$tmp/caller/.tool-versions"
  local read_by_default
  read_by_default=$(cd "$tmp/caller" && mise exec -- true 2>&1)
  [[ $read_by_default == *nonexistent-tool-xyz* ]] ||
    fail "the .tool-versions read without an override: got [$read_by_default], want a warning naming nonexistent-tool-xyz"
  for command in "${commands[@]}"; do
    (cd "$tmp/caller" && PATH=$stubbed_path "$plugin/bin/$command" >/dev/null)
    local ignored
    ignored=$(with_the_shims_overrides mise exec -- true 2>&1)
    assert_equal "$command: mise's output with the shim's overrides" "$ignored" ""
  done
}

test_hook_passes_the_local_builds_output_and_input_through() {
  write_stub "$plugin/.build/wrap-formats" 'cat'
  local output
  output=$(echo '{"hookSpecificOutput": {}}' | run_hook "$stubbed_path")
  assert_equal "hook status" "$?" 0
  assert_equal "hook output" "$output" '{"hookSpecificOutput": {}}'
}

test_hook_without_mise_passes() {
  local bare=$tmp/bare
  mkdir "$bare"
  local tool
  for tool in bash sh basename dirname; do
    ln -s "$(command -v "$tool")" "$bare/$tool"
  done
  local output
  output=$(echo '{}' | run_hook "$bare" 2>/dev/null)
  assert_equal "hook status" "$?" 0
  assert_equal "hook output" "$output" ""
}

test_hook_passes_when_mise_fails() {
  write_stub "$tmp/stubs/mise" 'echo "{\"hookSpecificOutput\": {}}"; exit 2'
  local output
  output=$(echo '{}' | run_hook "$stubbed_path" 2>/dev/null)
  assert_equal "hook status" "$?" 0
  assert_equal "hook output" "$output" ""
}

test_hook_passes_when_the_release_fails() {
  touch "$tmp/installed"
  write_stub "$tmp/release/wrap-formats" 'echo "{\"hookSpecificOutput\": {}}"; exit 1'
  local output
  output=$(echo '{}' | run_hook "$stubbed_path")
  assert_equal "hook status" "$?" 0
  assert_equal "hook output" "$output" ""
}

test_hook_passes_when_the_local_build_fails() {
  write_stub "$plugin/.build/wrap-formats" 'echo "{\"hookSpecificOutput\": {}}"; exit 1'
  local output
  output=$(echo '{}' | run_hook "$stubbed_path")
  assert_equal "hook status" "$?" 0
  assert_equal "hook output" "$output" ""
}

for test in $(declare -F | awk '$3 ~ /^test_/ { print $3 }'); do
  set_up
  "$test"
  rm -rf "$tmp"
done
exit "$failures"
