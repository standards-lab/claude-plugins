# marathon close

Finish a session whose work is done and validated: every stage committed and every checkpoint
confirmed (`references/staged-execution.md`). A `plan` session has nothing to build. A `review`
session's cleanup is its deliverable. A documentation step's pages are checked against the
current code, and when they disagree, the page is wrong. Don't close on a failure.

## 1. Review the branch

A `start` session's branch gets one review of the whole branch. Other sessions skip this step.

1. The report is `.claude/report.md`, in the project's own repository or, in a workspace, the
   coordinator's. Before anything writes it, confirm that repository's `.gitignore` lists it, and
   add the line if it doesn't. A wave's lane also confirms `.claude/worktrees/` in each touched
   repository.
2. Engage the reviewer profile (`behavior/delegation.md`), or review the branch yourself when
   briefing the reviewer would cost more. The brief gives the stage list, the branch in each
   touched repository, the confirmed checkpoints, and any architecture or practice the project
   declares, cited by path. Without a declared architecture or practice, the review checks
   the branch against the ecosystem's usual practice. The reviewer returns the report as text.
3. Verify the findings yourself, write the report to its path, and tell the architect it's ready.
4. Settle each finding with the architect: fix it, or record it for a later step. A fix lands as a
   checkpoint Adjust: a new commit, the checks run again, and the fix reported for the architect
   to confirm and cited under **Validated**. A finding that reaches beyond the step needs a
   re-plan.

## 2. Tend the context

Bring the written context the step touched in line with what now exists, following
`references/context-engineering.md` and including the edits noted at SETTLE. The scope is:

- the project's `context/`, or in a workspace each touched repository's `context/`
- the coordinator's notes on the changed capability
- `docs/` pages the change made wrong
- claims about the change in other members' context

Fix a stale claim in another member and record it under **Cross-repo**. Confirm the changes with
the architect before applying them.

## 3. Agree on the next step

Settle the one next step with the architect. Don't set the direction alone.

## 4. Record the session

Rewrite the reset file (`mechanics/reset-file.md`) with `Status: closeout`. In a workspace, write
the coordinator's record, or the lane's record during a wave; when this session finishes the last
lane, fold the wave (`mechanics/reset-file.md`). The Disposition lists each note operation and,
under **Validated**, each checkpoint with its evidence. Next-focus is the step agreed in step 3.

## 5. Commit and publish

Delete `.claude/report.md`. Commit in each touched repository, with the coordinator's record as a
commit of its own there. Then publish each branch with its repository's `[remote] publish`
command, using the change description from the record. A project with no remote stops after the
commit. A wave's lane then leaves its worktrees and removes them with `git worktree remove`
(`references/workspace-coordination.md`).
