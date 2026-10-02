# marathon plan

`plan` defines goals and makes them ready to run: it creates goals and their tasks, sets up their
repositories and records, stages and orders them, and writes their notes and next briefs. `start`
runs a ready task. `plan` is also the coordinator's planning and administrative channel: its
changes land through a `plan-<goal>` branch and pull request, and it stays short.

`plan` runs LOCATE and START (`mechanics/pipeline.md`); the goal is optional.

## Rounds

Enter plan mode and run plan rounds with the architect (`references/briefs.md`, "Plan round").
The plan file holds the open round. The planner finds the facts. Look for the questions whose
answers would change what the workspace does next, not only the areas with the most unknowns.

Nothing changes until the architect approves the outcome.

## Edits

Apply what the rounds decided on a `plan-<goal>` branch in each repository they change, in one
commit per kind of change:

- **Roadmap**: stage a goal into `active`, pivot any of the three arrays, add or delete goals
  and tasks, and sharpen the next task's entry (`mechanics/goals.md`, `references/manifest.md`).
  Staging checks the goal's `root`, `repos`, and lock first. A spike added to an existing
  experiment is a sub-goal shaped as `experiment` shapes one (`commands/experiment.md`).
- **Notes**: write a note a coming goal needs, sharpen one, or delete one the discussion ruled
  out (`references/context-engineering.md`).
- **Catalog and configuration**: the workspace's repository catalog, `order`, and
  `marathon.toml`.
- **Goal setup**: for an active goal whose repositories, `root`, record, or tasks aren't set up
  yet, create what is missing: run `init` in each new repository the rounds decided
  (`commands/init.md`), add it to `order`, the repository catalog, and the goal's `repos` (and
  `root`, when there is none), write the tasks into the manifest, and create or extend the goal
  record in the root, with the tasks the rounds settled checked. A goal with no `root` locks
  nothing, and `start` on it stops and asks for this `plan`. A spike sub-goal adds its own steps
  ("Spike setup").
- **A next brief**: when an active goal's root is on its default branch with no task in
  progress, write the approved brief for its next task into the goal record there, as a direct
  commit on that default branch, never on the plan branch, firing `on-record` first. It is
  goal-record bookkeeping, the one edit here that skips the pull request. The next `start`
  presents it for approval without new rounds.

## Spike setup

`plan experiment.<topic>.<spike>` stages the spike, unless the previous spike's sync already
staged it, and sets it up through Goal setup. Its repository doesn't exist yet, so it is staged
without it (`mechanics/goals.md`). A spike always starts in a new repository: an existing project
it builds on goes in its references, read-only, and never becomes the spike's root. The rounds
also settle:

- **The evidence**: a numbered list of what the spike must show to answer its question.
- **`init`'s founding decisions** (`commands/init.md`).
- **The path**: the steps that answer the question, written as the sub-goal's tasks.
- **The references**: the repositories the spike reads.

Goal setup then runs with these steps:

1. Run `init` at the sub-goal's `path` with the settled decisions. Its `context/README.md` states
   the question, the evidence list, and the sub-goal; its goals live in the coordinator's
   roadmap. It writes no `[experiment]` table.
2. Record the repositories the spike reads, using the served project's own convention. In a
   workspace, that is a committed list of remotes plus a gitignored map to local checkouts, both
   in the coordinator. The spike reads those repositories and never writes to them. Its code
   dependencies are published versions, never a replace directive.
3. Create the repository at the sub-goal's `remote` and push `init`'s commit, which creates its
   default branch. It is the one commit that can't land through a pull request, since there is
   no branch yet to merge into (`mechanics/pipeline.md`, "Branches and pull requests").
4. Write the path as the sub-goal's tasks and create the goal record in the spike's repository.
   The spike stays out of `order` and the repository catalog: LOCATE finds it through its `path`
   (`mechanics/pipeline.md`).

A change the rounds decide for a member repository is recorded as a task of the goal that locks
it, and lands with that goal's next `start`.

Publish and merge each `plan-<goal>` branch with the approved round outcome as its body
(`mechanics/pipeline.md`, "Branches and pull requests"). Delete the plan file once its edits are
published.
