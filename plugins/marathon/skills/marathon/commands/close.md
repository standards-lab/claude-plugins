# marathon close

Close out a session whose work is finished and validated: every stage committed and every
checkpoint confirmed (`references/staged-execution.md`). A `plan` session has nothing to build, a
`review` session's cleanup is its deliverable, and a documentation step's pages are checked against
the current code, where the page is what's wrong when they disagree. Don't close on a failure.

## 1. Review the branch

A `start` session's branch gets one holistic review; other sessions skip this step.

1. The report is `.claude/report.md`: the project's own, or the coordinator's in a workspace.
   Before anything writes it, confirm that repository's `.gitignore` lists it, adding the line if
   needed.
2. Engage the reviewer profile (`behavior/delegation.md`), or review directly when briefing it
   would cost more. The brief carries the stage list, the branch in each touched repository, the
   confirmed checkpoints, the report's path, and any architecture or practice the project
   declares, cited by path. Without one, the review checks against the ecosystem's idiom.
3. Read the report firsthand, correct it, and tell the architect it's ready.
4. Settle each finding with the architect: fix it in a new commit and re-run the checks, or record
   it for a later step. A finding past the step is a re-plan.

## 2. Tend the context

Bring the written context the step touched in line with what now exists, under
`references/context-engineering.md`, including the edits noted at SETTLE. The scope is the
project's `context/`, or in a workspace each touched repo's `context/`, the coordinator's notes on
the changed capability, `docs/` pages the change moved out from under, and claims about it in
other members' context; a stale claim there is fixed and recorded under **Cross-repo**. Confirm
with the architect before applying.

## 3. Agree on the next step

Settle the single next step with the architect; don't set the direction alone.

## 4. Record what happened

Rewrite the reset file (`mechanics/reset-file.md`) with `Status: closeout`: the coordinator's in a
workspace, or the session's own record in a wave. The Disposition carries the tending ledger and,
under **Validated**, each checkpoint with its evidence. Next-focus is the step from step 3.

## 5. Commit and publish

Delete `.claude/report.md`. Commit in each touched repo, with the coordinator's record as its own
commit there, then publish each branch with its repo's `[remote] publish` command, using the
change description from the record. A project with no remote stops after its commit.
