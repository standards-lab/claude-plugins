# marathon plan

`plan` defines goals and makes them ready to run: it creates goals and their tasks, sets up their
repositories and records, stages and orders them, and writes their notes and next briefs. `start`
runs a ready task. `plan` is also the coordinator's planning and administrative channel: its
changes there are short, direct commits on the default branch, and it stays short.

`plan` runs LOCATE and START (`mechanics/pipeline.md`); the goal is optional.

## Rounds

Enter plan mode and run plan rounds with the architect (`references/briefs.md`, "Plan round").
The plan file holds the open round. The planner finds the facts. Look for the questions whose
answers would change what the workspace does next, not only the areas with the most unknowns.

Nothing changes until the architect approves the outcome.

## Edits

Apply what the rounds decided, in one commit per kind of change:

- **Roadmap**: stage a goal into `active`, pivot any of the three arrays, add or delete goals
  and tasks, and sharpen the next task's entry (`mechanics/goals.md`, `references/manifest.md`).
  Staging checks the goal's `root`, `repos`, and lock first.
- **Notes**: write a note a coming goal needs, sharpen one, or delete one the discussion ruled
  out (`references/context-engineering.md`).
- **Catalog and configuration**: the workspace's repository catalog, `order`, and
  `marathon.toml`.
- **Goal setup**: for an active goal whose repositories, `root`, record, or tasks aren't set up
  yet, create what is missing: run `init` in each new repository the rounds decided
  (`commands/init.md`), add it to `order`, the repository catalog, and the goal's `repos` (and
  `root`, when there is none), write the tasks into the manifest, and create or extend the goal
  record in the root, with the tasks the rounds settled checked. A goal with no `root` locks
  nothing, and `start` on it stops and asks for this `plan`.
- **A next brief**: when an active goal's root is on its default branch with no task in
  progress, write the approved brief for its next task into the goal record there, as a direct
  commit, firing `on-record` first. The next `start` presents it for approval without new rounds.

A change the rounds decide for a member repository is recorded as a task of the goal that locks
it, and lands with that goal's next `start`.

Pull before each commit, stage explicit paths, and push after it. Delete the plan file once its
edits are committed.
