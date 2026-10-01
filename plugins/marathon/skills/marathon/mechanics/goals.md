# Goals

Goals and tasks are marathon's whole vocabulary for work, including work that runs at the same
time. The roadmap manifest, `context/roadmap.toml`, holds them (`references/manifest.md`); a
standalone project keeps its own, and a workspace keeps one at the coordinator.

## Goals and tasks

- **A goal** names an outcome and holds sub-goals and tasks, as a directory holds directories and
  files.
- **A task** is a leaf of a goal, sized to one session's work.
- **Any goal, at any depth, can run in parallel with other goals.** Its goal record is the single
  reference point across its sessions (`mechanics/goal-record.md`).

## Three states

Every goal the manifest lists is in exactly one of three root arrays of dotted paths:

- **`active`**: the goals running now. It is a live snapshot, edited at any time.
- **`planned`**: the long-term order, top first.
- **`backlog`**: goals with no place in the order yet.

A goal that isn't listed is a container whose sub-goals are listed, such as `v1` when
`v1.ai.experiment` is listed. A goal and one of its ancestors are never both listed.

Three verbs move goals:

- **stage**: put a goal in `active`, from `planned`, `backlog`, or as a new goal, once its lock
  is free
- **sync**: apply a finished goal's pending coordinator edits and delete it from the manifest
- **pivot**: edit any of the three arrays, at any time, as the architect decides

`plan` stages and pivots. The session that ships a goal's last task syncs it.

## The repository lock

- **An active goal locks every repository in its `repos`.** No two active goals name the same
  repository, so each goal works on the main checkout of its own repositories, with no worktrees.
  Staging a goal whose `repos` overlap an active goal's waits, or pivots the other goal out.
- **A goal with no `repos`** touches only the coordinator's `context/`, and locks nothing.
- **The coordinator is never locked.** It is the one shared repository, and it takes only short,
  direct commits on its default branch, from `plan` edits, staging, sync, and the records of
  goals homed there. Each of these touches separate goal tables, one array line, or one record
  file, so concurrent commits rebase cleanly. Pull before each one.
- **A workspace can narrow the lock.** `[workspace] exclusive` in the coordinator's
  `marathon.toml` lists groups of repositories that at most one active goal may touch at a time
  (`mechanics/configuration.md`).

## Sync

When a goal's last task ships, its session syncs the goal:

1. In the last task's branch, delete the goal record from the home repository. When the home is
   the coordinator, the deletion goes in step 2's commit instead.
2. After the task merges, make one direct commit at the coordinator that applies the record's
   pending coordinator edits, removes the goal from `active`, and deletes the goal's table with
   everything under it. Delete any ancestor goal whose criteria now hold, and any ancestor left
   empty.
3. The lock on the goal's repositories is released. Tell the architect which `planned` goal could
   be staged next; staging it is a `plan` decision.

## Spikes

A spike is a goal whose `home` and `repos` are the spike's own repository
(`commands/experiment.md`). It runs and syncs like any other goal. Taking in its result is a task
of the goal it serves, such as `v1.messaging.experiment`'s intake.
