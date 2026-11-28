#!/usr/bin/env bash
# The check's Go step: the Go module in marathon/ passes the Go check (build, vet, gofmt, go fix,
# go mod tidy, go test -race, and golangci-lint, with GOWORK=off), and lint-context and
# lint-references, built from marathon/ by scripts/build.sh, pass on the repository.
set -uo pipefail
source "$(dirname "$0")/lib.sh"

go_module=marathon
tools=$(mktemp -d)
trap 'rm -rf "$tools"' EXIT
if command -v go >/dev/null && command -v gofmt >/dev/null && command -v golangci-lint >/dev/null; then
  export GOWORK=off
  rerun_after_output() {
    local step=$1
    shift
    if ! failed_step_report=$("$@" 2>&1); then
      echo "$failed_step_report" >&2
      err "$step failed in $go_module/; fix the lines above, then rerun"
    fi
  }
  rerun_after_output "go build" go build -C "$go_module" ./...
  rerun_after_output "go vet" go vet -C "$go_module" ./...
  while read -r file; do
    [ -n "$file" ] && err "$go_module/$file is not gofmt-formatted; run gofmt -w $go_module/$file"
  done < <(cd "$go_module" && gofmt -l .)
  if ! go_fix_diff=$(go fix -C "$go_module" -diff ./... 2>&1) || [ -n "$go_fix_diff" ]; then
    echo "$go_fix_diff" >&2
    err "go fix would change $go_module/; run 'go fix ./...' in $go_module/"
  fi
  if ! go_mod_tidy_diff=$(go mod tidy -C "$go_module" -diff 2>&1); then
    echo "$go_mod_tidy_diff" >&2
    err "go.mod or go.sum in $go_module/ is not tidy; run 'go mod tidy' in $go_module/"
  fi
  rerun_after_output "go test -race" go test -C "$go_module" -race ./...
  rerun_after_output "golangci-lint" bash -c 'cd "$1" && golangci-lint run ./...' _ "$go_module"

  if scripts/build.sh "$tools"; then
    "$tools/lint-context" || err "lint-context failed; fix the lines above, then rerun"
    "$tools/lint-references" || err "lint-references failed; fix the lines above, then rerun"
  else
    err "the $go_module/ commands do not build, so lint-context and lint-references did not" \
      "run; fix the lines above, then rerun"
  fi
else
  err "go, gofmt, or golangci-lint is not on PATH, so $go_module/ was not checked and" \
    "lint-context and lint-references did not run; run 'mise install' and activate mise," \
    "then rerun"
fi

finish
