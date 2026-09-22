# Staged execution

How every working session executes its settled step: in stages, each committed once its check
passes, with the architect confirming the result at checkpoints.

## Stages

A stage is the smallest change set that leaves one unit consistent on its own, including its tests
and in-source comments.

- A stage that produces **source** has a compilation unit as its unit: a package in Go, a crate in
  Rust, a module in Python. Its check runs the repository's own tooling on that unit: the build,
  vet, and test (in Go, `go build`, `go vet`, and `go test` on the package path), and any lint or
  conventions tool the repository wires into its CI or task runner. Run the tool and resolve its
  findings; never restate what it checks. The module may be red between stages, since every
  broken caller is a later stage.
- A stage that produces **prose or configuration** has as its unit the smallest set of files that
  must change together: a playbook and the reference it cites, one note, one documentation page.
  Its check is the repository's consistency script where one exists, and a read for coherence.

`plan` and `review` sessions run in stages of context edits under the second rule.

## The stage list

The stage list is the SETTLE artifact, approved by the architect before anything changes. The
planner profile may draft it (`behavior/delegation.md`). It orders stages by dependency, lowest
first, and names each stage's unit, files, and reason. A step spanning member repos has one list,
grouped by repository in the coordinator's `order`. The list lives in the conversation and, on a
handoff, in the reset file's Next-focus.

## Checkpoints

The list groups its stages under checkpoints. A checkpoint is an observable behavior the architect
can run or watch run, never "the tests are green". The session runs every stage up to a checkpoint
without stopping, then stops and shows it.

- The final validation is always a checkpoint. Place an earlier one wherever a behavior first
  becomes observable that later stages build on. A checkpoint after every stage stops at every
  stage.
- A library with no runnable surface earns its checkpoint through a consumer exercise written for
  the purpose, such as an example program or a conformance run.
- A context project's checkpoint is a walkthrough: one concrete scenario traced through the
  changed prose, citing where each changed rule applies.

## Executing a stage

1. **State the delegation call**: the executor profile or the session itself, and why; for the
   executor, the model too.
2. **Execute** and run the check until it passes. Read a delegate's work firsthand.
3. **Commit.** Fire `on-commit`, then commit with the stage's decisions in the message.
4. **Log** one entry in the conversation: the `diff --stat`, the check result, the delegation
   call, and prose only for a decision the plan didn't spell out.
5. **Continue**, or at a checkpoint, stop and report: the stages it covers, how to see the
   behavior, what the session observed, and what it is least confident of.

## Checkpoint outcomes

- **Confirmed** — continue to the next group of stages.
- **Adjust** — the behavior is wrong within the checkpoint's stages. Fix it forward in a commit
  naming the checkpoint, re-run the checks, and report again.
- **Re-plan** — the finding reaches past the checkpoint. The architect enters plan mode, and the
  session re-enters SETTLE for stages k..N. Committed stages stay unless a finding invalidates
  them, in which case reverting them is the re-plan's first act.
- **Interrupt** — the architect breaks in mid-run. The session finishes the stage in flight if its
  check can pass quickly, or abandons it and restores the tree, then reports as at a checkpoint.

## Validation

After the last stage commits, validate the whole step as the final checkpoint:

- **Code**: the whole-module build and full test run (in Go, `./...`), then the run-and-verify
  behavior check with its commands and what to look for.
- **Context**: a read of the whole change for coherence, and the consistency script.
- **An experiment's final step**: the answer to its question, with the evidence.

Don't close on a failure. Once the architect confirms, `close` opens with the branch review.

## Stay within the step

No opportunistic refactors or unrelated cleanups; note a worthwhile temptation as an open note
and don't take the detour.
