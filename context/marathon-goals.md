# marathon goals: the unit of parallel work

This note covers how marathon tracks work that runs at the same time. Goals and tasks are the
whole vocabulary. This replaces marathon 0.15's waves, lanes and folds, and moves
marathon-roadmap into marathon core. It is tracked as `factory.goals`. The pipeline each task
runs is `marathon-factory.md`.

## The problem it solves

Under 0.15, a wave was locked in once it started:

- A lane made of an array of tasks had no single reference point across its sessions.
- Coordinator edits were deferred until every lane finished, and then applied in one fold.
- Lanes reached the coordinator through worktrees.

Pivoting meant waiting for the fold. The v1-storage lane finished on 2026-09-24, but its
coordinator edits stayed unapplied for a week while the other lanes ran. That lane also had to
bend the rules to work at all: it named its record after its goal and kept that record for the
goal's whole life.

## One primitive

- **A goal is like a directory.** It holds sub-goals and tasks.
- **A task is like a file.** It is a leaf stored in a goal, and one session does it.
- **Any goal, at any depth, is a unit of parallel work.** It is the single reference point across
  its sessions.

## Three states

Every goal is in exactly one of three top-level arrays of dotted paths:

- **`active`:** the goals running now. It is a live snapshot that can be edited at any time.
- **`planned`:** the long-term order, top first.
- **`backlog`:** goals with no slot in the order yet. A backlog goal is an ordinary `[goals.x]`
  table, with no `backlog.` prefix.

Three verbs act on them:

- **stage:** move a goal from `planned` to `active`
- **sync:** close a goal out into the workspace
- **pivot:** edit any of the three arrays at any time

## The manifest

`context/roadmap.toml` is plain TOML: no inline tables and no timestamps.

```toml
active  = ["factory", "v1.messaging", "v1.ai.experiment"]
planned = ["v1.ai", "quality", "cli", "v1.data"]
backlog = ["second-providers", "docs-site"]

[goals.factory]
name = "marathon as a software factory"
home = "claude-plugins"     # where the goal record lives
repos = ["claude-plugins"]  # what the goal locks while active
summary = "..."

[goals.factory.tasks.pipeline]
name = "The pipeline and its profiles"
summary = "..."
```

`home` and `repos` default to the nearest ancestor's values. A goal with no `repos`, and none
to inherit, touches only the coordinator's `context/`.

## The repository lock

- **An active goal locks every repository in its `repos`.** No two active goals may name the same
  repository, so each goal works on the main checkout of its own repositories. There are no
  worktrees.
- **The coordinator is the one shared repository.** It takes only short, direct commits on its
  default branch, from three sources:
  - `plan` edits
  - stage
  - sync
- **A workspace can narrow the lock.** The policy lives in the coordinator's configuration, not
  in core. This workspace allows at most one active goal on the reference-architecture member
  repositories at a time; spikes and harness work run beside it. While `v1.messaging` holds
  that slot, `v1.ai` stays `planned` even after `v1.ai.experiment` finishes.

## The goal record

`context/goals/<goal>.md` lives in the goal's `home` repository. Every task session updates it.
It holds:

- the tasks and their progress
- the current task brief and its slice list
- handoff state, when a session stops mid-task
- decisions made along the way
- pending coordinator edits: notes, catalog rows, workspace `order`, roadmap changes

Each task merges its own small PR into the home repository's default branch, so no branch lives
for the whole goal. The record replaces 0.15's reset file and lane records, the Disposition
included.

## Sync

When a goal's last task ships, its session makes one coordinator commit:

1. Apply the goal record's pending coordinator edits.
2. Remove the goal from `active` and delete its table.
3. Delete the goal record in the home repository.
4. Release the lock on its repositories.

Concurrent syncs touch separate goal tables and one array line each, so they merge cleanly. The
next `planned` goal can be staged straight away.

## Spikes

A spike is a goal whose `home` and `repos` are the spike repository. It runs and syncs like any
other goal. Taking in its result is a task of the goal it serves, such as the intake in
`v1.messaging.experiment`.

## The roadmap moves into core

Goals now carry concurrency, locking and the session record, so marathon can't run without the
manifest. `factory.goals` folds marathon-roadmap's skill, hooks and `references/manifest.md`
into marathon core and retires the extension.

## Retired terms

| 0.15 | Now |
|---|---|
| wave | `active` |
| lane | goal |
| fold | sync |
| `next` | `active` and `planned` |
| `backlog.x` tables | `backlog` array, `[goals.x]` tables |
| lane record, reset file, Disposition | goal record |
| checkpoint, stage list | task brief and slices (`marathon-factory.md`) |
| coordinator worktree | not needed (repository lock) |

## Migration

`factory.goals` migrates the records 0.15 left open:

- standards-lab `context/reset/messaging-experiment.md` becomes the `v1.messaging` goal record.
  Its home is decided in that goal's next `plan` session.
- standards-lab `context/reset/ai-experiment.md` becomes the `v1.ai.experiment` goal record.
  The goal spans three spikes, so its home is decided at the migration.
- standards-lab `context/reset.md` is deleted once no session reads it.
