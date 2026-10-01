# marathon plan

Shape the path rather than build on it: stage, reorder, or drop goals; write or sharpen the notes a
coming goal needs; take in a finished spike; or work out a design the roadmap doesn't hold yet.
`plan` changes only the coordinator (or a standalone project's) `context/`, in short, direct
commits on its default branch, and stays short. Planning a single task happens inside `start`.

`plan` runs LOCATE and START (`mechanics/pipeline.md`); the goal is optional.

## Rounds

Enter plan mode and run plan rounds with the architect (`references/briefs.md`, "Plan round").
The planner finds the facts. Look for the questions whose answers would change what the
workspace does next, not only the areas with the most unknowns. A spike intake leads with the
spike's result (`references/briefs.md`, "Spike intake").

Nothing changes until the architect approves the outcome.

## Edits

Apply what the rounds decided, in one commit per kind of change:

- **Roadmap**: stage a goal into `active` once its lock is free, pivot any of the three arrays,
  add or delete goals and tasks, and sharpen the next task's entry (`mechanics/goals.md`,
  `references/manifest.md`).
- **Notes**: write a note a coming goal needs, sharpen one, or delete one the discussion ruled
  out (`references/context-engineering.md`).
- **Goal records**: for a goal homed at the coordinator, create or update its record, such as
  the brief for its next task, or check a task that `plan` itself completes, such as an intake.
  When that task is the goal's last, sync it (`mechanics/goals.md`). Fire `on-record` first.
- **Intake**: update the experiment's catalog entry and archive its remote
  (`commands/experiment.md`).

A change the rounds decide for a member repository is recorded in that goal's record, or as a task,
and lands with that goal's next `start`.

Pull before each commit, and push after it.
