# Staged execution

Every working session carries out its step in stages. Each stage commits once its check passes,
and the architect confirms the results at checkpoints.

## Stages

A stage is the smallest set of changes that leaves one unit consistent on its own, including its
tests and in-source comments.

- A stage that produces **source** has a compilation unit as its unit: a package in Go, a crate in
  Rust, a module in Python. Its check runs the repository's own tooling on that unit: the build,
  vet, and tests (in Go, `go build`, `go vet`, and `go test` on the package path), plus any lint
  or conventions tool the repository runs in its CI or task runner. Run the tool and fix what it
  finds; never restate what it checks. The module may fail to build or pass its tests between
  stages, because each broken caller is fixed by a later stage.
- A stage that produces **prose or configuration** has as its unit the smallest set of files that
  must change together: a playbook and the reference it cites, one note, or one documentation
  page. Its check is the repository's consistency script, where one exists, and a read for
  coherence.

`plan` and `review` sessions work in stages of context edits under the second rule.

## The stage list

The stage list is what SETTLE produces, and the architect approves it before anything changes.
The planner profile may draft it (`behavior/delegation.md`). The list orders stages by
dependency, lowest first, and gives each stage its unit, its files, and its reason. A step that
spans member repositories has one list, grouped by repository in the coordinator's `order`. The
list lives in the conversation and, on a handoff, in the reset file's Next-focus.

## Checkpoints

The list groups its stages under checkpoints. A checkpoint is a behavior the architect can run
or watch, never "the tests pass". The session runs every stage up to a checkpoint without
stopping, then stops and shows the checkpoint.

- The final validation is always a checkpoint. Add an earlier checkpoint wherever a behavior that
  later stages build on first becomes visible. A checkpoint after every stage stops the session at
  every stage.
- For a library with nothing to run, write a program that exercises it, such as an example
  program or a conformance run, and use that as the checkpoint.
- On a context project, a checkpoint is a walkthrough: one concrete scenario traced through the
  changed prose, citing where each changed rule applies.

## Running a stage

1. **State the delegation decision**: the executor profile or the session itself, and why. For
   the executor, state the model too.
2. **Do the work** and run the check until it passes. Read a delegate's work yourself.
3. **Commit.** Fire `on-commit`, then commit, with the stage's decisions in the message.
4. **Log** one entry in the conversation: the `diff --stat`, the check's result, the delegation
   decision. Add prose only for a decision the plan didn't cover.
5. **Continue**, or at a checkpoint, stop and report: the stages it covers, how to see the
   behavior, what the session observed, and what the session is least sure of.

## Checkpoint outcomes

- **Confirmed**: continue with the next group of stages.
- **Adjust**: the behavior is wrong within the checkpoint's stages. Fix it in a new commit that
  names the checkpoint, run the checks again, and report again.
- **Re-plan**: the finding reaches beyond the checkpoint. The architect enters plan mode, and the
  session returns to SETTLE for the remaining stages. Committed stages stay unless a finding
  invalidates them, in which case the re-plan reverts them first.
- **Interrupt**: the architect interrupts partway through. If the stage in progress can pass its
  check quickly, the session finishes it. Otherwise, the session abandons it and restores the
  working tree. Either way, it reports as it would at a checkpoint.

## Validation

After the last stage commits, validate the whole step as the final checkpoint:

- **Code**: build the whole module and run the full test suite (in Go, `./...`), then run the
  run-and-verify check, with its commands and what to look for.
- **Context**: read the whole change for coherence, and run the consistency script.
- **An experiment's final step**: the answer to its question, with the evidence.

When the branch changed prose, the editor pass comes first. Engage the editor profile
(`behavior/delegation.md`), or edit the prose yourself when briefing the editor would cost more,
and commit the edits as one stage before the rest of validation runs.

Don't close on a failure. Once the architect confirms, `close` begins with the branch review.

## Stay within the step

Make no opportunistic refactors or unrelated cleanups. When a change outside the step looks
worthwhile, record it as a note and leave it.
