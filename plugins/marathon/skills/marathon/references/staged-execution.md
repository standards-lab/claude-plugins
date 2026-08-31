# Staged execution

How a `code` project's `start` executes: the agent implements the settled step in stages, and the
architect reviews each stage before the next begins. Execution starts only after the planning
phase — the stage list below is settled and approved at SETTLE, like any other plan, before any
code changes.

## What a stage is

A stage is the smallest change set that leaves one compilation unit green — the smallest unit the
language builds on its own: a package in Go, a crate in Rust, a module in Python. A stage includes
its tests and its in-source comments, so nothing is left invalidated for later. The green check is
scoped to that unit only (in Go: `go build`, `go vet`, `go test` on the package path); the module
as a whole may be red between stages, because the stage sequence is in dependency order and every
broken caller is a later stage.

## The stage list

The stage list is the SETTLE artifact: written during planning, approved by the architect before
execution begins. It orders the stages by dependency, lowest first, so a stage that changes an
exported interface is followed by the stages that consume it. Each entry names the stage's
compilation unit, the files it touches, and one line on why.

In a workspace, a step that spans member repos has one list: the stages group by repository and
order across repositories by the coordinator's `order` map, lowest layer first.

The list is the plan's artifact and is not committed anywhere — it lives in the conversation and,
on a handoff, in the reset file's Next-focus.

## Executing a stage

Implement the stage, run the unit-scoped build, lint, and tests, and fix until green. Commit, with
the stage's decision lines in the message (`on-commit` fires first). Then stop and report. Never
run ahead into the next stage.

## The stage report

The report is conversational, not a file: the `diff --stat`, the check result, then prose only on
the decisions the plan didn't spell out and the parts you are least confident of. Code carries the
what; the report carries the why. Don't restate what the diff shows.

## Review outcomes

- **Approve** — move to the next stage.
- **Adjust** — edit in place, re-run to green, amend or add a commit, re-report.
- **Re-plan** — the convention for findings that reach beyond the current stage: the architect
  enters plan mode and sends the findings, and the session re-enters SETTLE for stages k..N.
  Landed stages stay landed unless a finding invalidates them, in which case reverting those
  commits is the first act of the re-plan. A revised stage list from k onward is approved like
  the original; execution resumes at k.

## Validation

After the last stage is approved: the whole-module build and full test run (in Go: `./...`), then
the run-and-verify behavior check — the concrete commands and what to look for. Don't close on a
failure; fix it and re-validate.

## Stay within the step

A session covers its settled step and nothing else. No opportunistic refactors, no unrelated
cleanups; note the temptation as a concept if it's worth keeping, and don't take the detour.
