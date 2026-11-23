# Goals

Goals and tasks are marathon's whole vocabulary for work, including work that runs at the same time.
The roadmap manifest holds them: an index, `roadmap/index.toml`, of the three arrays, and one file
per goal, `roadmap/<goal>.toml` (`references/manifest.md`). A standalone project keeps its own,
and a workspace keeps one at the coordinator.

## Goals and tasks

- **A goal** names an outcome and holds sub-goals and tasks, as a directory holds directories and
  files.
- **A task** is a leaf of a goal, sized to one session's work.
- **Any goal, at any depth, can run in parallel with other goals.** Its goal record, in its root
  repository, is the single reference point across its sessions (`mechanics/goal-record.md`).
- **An experiment** is a goal under `experiment` that settles what another goal is built from:
  `experiment.ai` investigates, and `v1.ai` builds. Each spike is a sub-goal,
  `experiment.<topic>.<spike>`, with its own repository and its path steps as tasks. The
  experiment itself holds only its `intake` task (`commands/experiment.md`).

## Three states

Every goal the manifest lists is in exactly one of the index's three arrays of dotted paths:

- **`active`**: the goals running now. It is a live snapshot, edited at any time.
- **`planned`**: the long-term order, top first.
- **`backlog`**: goals with no place in the order yet.

A goal that isn't listed is a container whose sub-goals are listed, such as `v1` when `v1.data` is
listed. A goal and one of its ancestors are never both active.

An experiment goal is never listed. Its spike sub-goals are listed and staged on their own, and
several spikes of one experiment can be active at once, each rooted in its own repository.

Three verbs move goals:

- **stage**: put a goal in `active`, from `planned`, `backlog`, or as a new goal. Staging
  requires a `root` among the goal's `repos`, `repos` that cover every repository its tasks
  touch other than the coordinator, and the repository lock's staging rule. A goal whose
  repositories don't exist yet is staged without them, and its first session is a `plan` that
  sets it up (`commands/plan.md`).
- **sync**: carry a finished goal's context into the workspace, then delete the goal
- **pivot**: edit any of the three arrays, at any time, as the architect decides

Staging and pivoting edit only the index. Adding a goal creates its file, and adding a task adds
its table to its goal's file and appends its slug to `order`.

`plan` stages and pivots, through its pull request. The session that ships a goal's last task
syncs it, and stages the goal its session brief proposes when staging's checks pass.

## The repository lock

- **Staging reserves only the root.** No two active goals declare the same `root`; a goal's
  `repos` may name any other repository, another active goal's root included. Staging fails only
  on a shared root or a broken `exclusive` group, and the goal waits or a pivot takes the other
  out.
- **A repository is held only while a session works in it.** Any repository, root included, is
  held from every other session while it is off its default branch, on a task, `<slug>-fix`,
  `sync-<goal>`, or session branch (`mechanics/pipeline.md`, "Branches and pull requests"), and
  while an active goal's Handoff reads "tag <names>" with it short of released
  (`mechanics/releasing.md`). Otherwise it is free, and goals work on its main checkout, with no
  worktrees.
- **The coordinator is in no goal's `repos`.** It changes only through sessions: `plan`'s
  planning and administrative edits, `experiment`, `intake`, `retro`'s findings, and syncs, each
  on its own session branch and pull request. No goal's record lives there, and a task's changes
  to it wait in its goal record's pending edits.
- **`[workspace] exclusive` is the one reservation besides the root.** At most one active goal's
  `repos` may touch each of its groups, for the goal's whole life (`mechanics/configuration.md`).

## Sync

When a goal's last task merges, and every tag its task brief's Release line names is released
(`mechanics/releasing.md`), its session syncs the goal. Every piece of context lands before
anything is deleted:

1. Read the goal record's pending edits from the root's default branch.
2. Apply the pending edits for the coordinator in one commit on a `sync-<goal>` branch there: the
   notes, the catalog, the workspace `order`, the index, and the goal files. Remove the goal from
   `active`, and delete its file and its sub-goals' files. Then delete the file of each ancestor
   whose criteria now hold, or that is left empty: one with no tasks of its own and no file whose
   name extends its own. When the session brief's Sync section reads "staged on accept", stage
   that goal in the same commit (`references/briefs/session-brief.md`), checking staging's
   requirements again; if one no longer holds, leave the goal where it is and give the architect
   the reason.
3. Apply each pending edit for another repository on a `sync-<goal>` branch there; one for an
   active goal's root becomes its task in step 2's commit, and one for a held repository waits
   ("The repository lock").
4. Publish and merge each `sync-<goal>` branch, lowest layer first, with the Sync section of the
   last task's session brief as its body (`mechanics/pipeline.md`, "Branches and pull requests").
5. Once every sync pull request merges, delete the goal record from the root, as a direct commit
   on its default branch.
6. The goal's root is released, and the staged goal is ready for its next session. A goal the
   Sync section only proposed waits for a `plan` to stage it.

A spike sub-goal syncs the same way. Its last task's session brief gives the spike's answer, and
its pending edits carry the spike's answer section for the note its experiment cites
(`references/briefs/answer-section.md`), which step 2 lands. The goal its Sync section
proposes is the next planned spike of the same experiment, or, with none left, the experiment's
intake (`references/briefs/session-brief.md`). The experiment stays until its intake,
which takes in the spikes' answers and archives their remotes (`commands/intake.md`).

If the task's merge can't happen, because `[remote] merge` is unset or a failing check needs the
architect, nothing is synced and the record stays:
set State to `handoff` with the next move "merge, then sync", and stop. If the context fills while the release is unfinished, the next move is "tag
<names>, then sync". If a sync pull request can't merge, or an edit waits, the record stays in
`handoff` with the next move "merge `sync-<goal>`, then delete the record" or "sync <repo>, then
delete the record", as goal-record bookkeeping, and stop.
