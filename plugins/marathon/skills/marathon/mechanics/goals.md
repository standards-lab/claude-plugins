# Goals

Goals and tasks are marathon's whole vocabulary for work, including work that runs at the same
time. The roadmap manifest, `context/roadmap.toml`, holds them (`references/manifest.md`); a
standalone project keeps its own, and a workspace keeps one at the coordinator.

## Goals and tasks

- **A goal** names an outcome and holds sub-goals and tasks, as a directory holds directories and
  files.
- **A task** is a leaf of a goal, sized to one session's work.
- **Any goal, at any depth, can run in parallel with other goals.** Its goal record, in its root
  repository, is the single reference point across its sessions (`mechanics/goal-record.md`).
- **An experiment** is a goal under `experiment` that settles what another goal is built from:
  `experiment.ai` investigates, and `v1.ai` builds. Its tasks are its spikes and a final intake
  (`commands/experiment.md`).

## Three states

Every goal the manifest lists is in exactly one of three root arrays of dotted paths:

- **`active`**: the goals running now. It is a live snapshot, edited at any time.
- **`planned`**: the long-term order, top first.
- **`backlog`**: goals with no place in the order yet.

A goal that isn't listed is a container whose sub-goals are listed, such as `v1` when `v1.data` is
listed. A goal and one of its ancestors are never both active.

Three verbs move goals:

- **stage**: put a goal in `active`, from `planned`, `backlog`, or as a new goal. Staging
  requires a `root` among the goal's `repos`, `repos` that cover every repository its tasks
  touch other than the coordinator, and a free lock. A goal whose repositories don't exist yet
  is staged without them, and its first session is a `plan` that sets it up
  (`commands/plan.md`).
- **sync**: carry a finished goal's context into the workspace, then delete the goal
- **pivot**: edit any of the three arrays, at any time, as the architect decides

`plan` stages and pivots. The session that ships a goal's last task syncs it.

## The repository lock

- **An active goal locks every repository in its `repos`.** No two active goals name the same
  repository, so each goal works on the main checkout of its own repositories, with no worktrees.
  Staging a goal whose `repos` overlap an active goal's waits, or pivots the other goal out.
- **The coordinator is never locked.** It is the one shared repository, and it changes only two
  ways: `plan`'s planning and administrative edits, and syncs. No goal's record lives there, and a
  task's changes to it wait in its goal record's pending edits. Each coordinator commit pulls
  first and stages explicit paths, never everything.
- **A workspace can narrow the lock.** `[workspace] exclusive` in the coordinator's
  `marathon.toml` lists groups of repositories that at most one active goal may touch at a time
  (`mechanics/configuration.md`).

## Sync

When a goal's last task merges, its session syncs the goal. Every piece of context lands before
anything is deleted:

1. Read the goal record's pending edits from the root's default branch.
2. Apply the pending edits for the coordinator in one commit there: the notes, the catalog, the
   workspace `order`, and the manifest, removing the goal from `active` and deleting its table
   with everything under it. Delete any ancestor goal whose criteria now hold, and any ancestor
   left empty.
3. Apply each pending edit for another repository on a branch there, published and merged as
   SHIP does, when no active goal locks that repository. When one does, add the edit as a task of
   the locking goal in step 2's commit instead.
4. Delete the goal record from the root, as a direct commit on its default branch. For an
   experiment, archive the spikes' remotes; step 2 has already moved each spike into the
   workspace's repository catalog as archived.
5. The lock is released. Tell the architect which `planned` goal could be staged next; staging it
   is a `plan` decision.

If a merge can't happen, because there is no `[remote] merge` command or a check fails, nothing is
synced and the record stays: set State to `handoff` with the next move "merge, then sync", and
stop.
