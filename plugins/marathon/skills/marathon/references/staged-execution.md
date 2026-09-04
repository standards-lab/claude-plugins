# Staged execution

How every working session executes the step it settled. The agent implements the step in
stages, and the architect reviews each stage before the next begins. Execution starts only after
the planning phase: the stage list below is settled and approved at SETTLE, like any other plan,
before anything changes.

## What a stage is

A stage is the smallest change set that leaves one unit of the project consistent on its own. A
stage includes its tests and its in-source comments, so nothing is left invalidated for later.
What the unit is, and what its check is, follows the project kind:

- On a **code** project the unit is a compilation unit, the smallest unit the language builds on
  its own: a package in Go, a crate in Rust, a module in Python. The check is scoped to that unit
  (in Go: `go build`, `go vet`, and `go test` on the package path). The module as a whole may be
  red between stages, because the stage sequence is in dependency order and every broken caller
  is a later stage.
- On a **context** project the unit is the smallest set of files that must change together to
  stay consistent: a command playbook and the reference it cites, or one design note. The check
  is the repository's own consistency script where one exists, and a read of the touched files
  for coherence.

An `experiment` runs in stages the same way. Its unit is whatever the spike builds next, and its
check is the spike's own. A `plan`, `review`, or `docs` session runs in stages of context edits
under the context-project rule.

## The stage list

The stage list is the SETTLE artifact: written during planning, approved by the architect before
execution begins. It orders the stages by dependency, lowest first, so a stage that changes an
exported interface is followed by the stages that consume it. Each entry names the stage's unit,
the files it touches, and one line on why.

In a workspace, a step that spans member repos has one list: the stages group by repository and
order across repositories by the coordinator's `order` map, lowest layer first.

The list is the plan's artifact and is not committed anywhere. It lives in the conversation and,
on a handoff, in the reset file's Next-focus.

## Executing a stage

A stage is not complete until the architect has reviewed it. Per-stage review is the point of
stages: a misstep never spreads across the code base, and the architect never reviews a whole
session at once. Every stage runs in this order:

1. **Execute.** Implement the stage and run its check; fix until the check passes.
2. **Report, uncommitted.** Stop and report with the working tree uncommitted, so the diff reads
   cleanly in the architect's tools. Iterate on adjustments until the architect approves. Never
   run ahead into the next stage.
3. **Commit on approval.** Fire `on-commit`, then commit with the stage's decision lines in the
   message. The architect states whether a `reset` follows, to keep the context small.
4. **Move to the next stage.**

## The stage report

The report is conversational, not a file: the `diff --stat`, the check result, then prose only on
the decisions the plan did not spell out and the parts you are least confident of. Code carries
the what; the report carries the why. Do not restate what the diff shows.

## Review outcomes

- **Approve.** Commit, then move to the next stage.
- **Adjust.** Edit in place, re-run the check, and re-report; the stage stays uncommitted until
  the architect approves it. A finding against a stage already committed is fixed in a new commit
  on the same branch.
- **Re-plan.** The outcome for findings that reach beyond the current stage: the architect enters
  plan mode and sends the findings, and the session re-enters SETTLE for stages k..N. Committed
  stages stay committed unless a finding invalidates them, in which case reverting those commits
  is the first act of the re-plan. A revised stage list from k onward is approved like the
  original, and execution resumes at k.

## Validation

After the last stage is approved, validate the whole step before `close`:

- On a **code** project: the whole-module build and full test run (in Go: `./...`), then the
  run-and-verify behavior check, with the concrete commands and what to look for.
- On a **context** project: a read of the whole change for coherence and consistency, and the
  repository's consistency script where one exists.
- For an `experiment`: the answer to the question the spike was settled to answer, with the
  evidence that supports it.

Do not close on a failure; fix it and validate again.

## Stay within the step

A session covers its settled step and nothing else. No opportunistic refactors, no unrelated
cleanups; note the temptation as a concept if it is worth keeping, and do not take the detour.
