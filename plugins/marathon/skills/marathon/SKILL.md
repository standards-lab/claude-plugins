---
name: marathon
argument-hint: "[init | plan | start | experiment | intake | status | retro | reset] [goal]"
description: >
  A workflow for long-running development run as a software factory. The architect engages twice
  per task, approving the task brief and accepting the session brief, while subagents implement,
  review against standards and against the spec, and keep the context current in between. Goals
  and tasks in a roadmap manifest are the units of work, and active goals run side by side, each
  locking its own repositories. Use this skill whenever the architect sets up a project from a
  planning concept; plans, stages, or reorders goals; starts, resumes, or ships a task; spikes an
  idea; asks what is running; reviews past work; or hands off because the context window is
  filling. Typical requests include "start factory", "what's the status", "plan the next goal",
  "take in the experiment", "resume where I left off", "run a retro", "hand off", "initialize this
  project", and "set up marathon here". marathon treats the repository as the only source of
  truth: a flat context/ of notes, a roadmap manifest, and one goal record per active goal. Prefer
  this skill for any structured, multi-session work on a marathon repository, even when the
  architect doesn't name it.
---

# Marathon

Version: 0.18.0

marathon is a workflow for long-running development. Each session takes one task of an active goal
from an approved plan to merged work, and the tasks add up to a production-quality version of the
original concept. The architect approves the task brief and accepts the session brief; agents do
the rest. The project's written context lives in a top-level `context/` directory, which every
task keeps current, so the context window holds the task instead of stale notes. The repository
is the source of truth, not an issue tracker and not the conversation.

Use this skill whenever a session begins, advances, pauses, resumes, or ends on a marathon
repository: one with a top-level `context/`, or a workspace of such repositories. Set up a new
repository with `init`.

## Behavior

These rules apply in every session: how a session plans, and how it hands work to subagents.

@behavior/planning.md

@behavior/delegation.md

## Mechanics

The session pipeline applies to every command:

@mechanics/pipeline.md

The pipeline refers to these files where it needs them:

- [`mechanics/goals.md`](./mechanics/goals.md): goals and tasks, the active, planned, and backlog
  states, the repository lock, and sync.
- [`mechanics/goal-record.md`](./mechanics/goal-record.md): the goal record, where it lives, its
  schema, and its State values.
- [`mechanics/configuration.md`](./mechanics/configuration.md): the layout of
  `.claude/marathon.toml`.
- [`mechanics/hooks.md`](./mechanics/hooks.md): how extension hooks resolve and fire.

## Commands

The first argument names the command. Each command's playbook supplies the content of the
pipeline's stages.

| Command | When to use it | Playbook |
|---------|----------------|----------|
| `init` | Once, to set up marathon on a repository from a planning concept | [`commands/init.md`](./commands/init.md) |
| `plan` | To define goals and make them ready: create goals and tasks, set up their repositories and records, stage, order, and write notes and next briefs | [`commands/plan.md`](./commands/plan.md) |
| `start` | To run one task of an active goal: plan, build, brief, ship, and sync on the last task | [`commands/start.md`](./commands/start.md) |
| `experiment` | Once per topic, to create an `experiment.<topic>` goal, its `intake` task, and its planned spikes; `plan` sets up each spike | [`commands/experiment.md`](./commands/experiment.md) |
| `intake` | Once per experiment, after every spike has synced: to decide what the served goals build from the answers, close the experiment, and archive its spikes | [`commands/intake.md`](./commands/intake.md) |
| `status` | To print one line per active goal, with what needs the architect first | [`commands/status.md`](./commands/status.md) |
| `retro` | To turn comments on past work into checks, standards, skill changes, or notes | [`commands/retro.md`](./commands/retro.md) |
| `reset` | To hand off a task partway; runs on its own when the context fills | [`commands/reset.md`](./commands/reset.md) |

A `start` session ends by shipping its task, or with `reset` when it hands off.

## References

Read each reference when its subject comes up:

- [`references/context-engineering.md`](./references/context-engineering.md): writing and
  maintaining the notes in `context/`.
- [`references/build.md`](./references/build.md): checks, slices, the build loop, escalation,
  and handoff.
- [`references/briefs.md`](./references/briefs.md): every format the architect reads.
- [`references/workspace-coordination.md`](./references/workspace-coordination.md): the
  coordinator, the dependency order, and steps that span repositories.
- [`references/manifest.md`](./references/manifest.md): the `roadmap.toml` format. Read it before
  editing or creating the manifest.
- [`references/extensions.md`](./references/extensions.md): the extension system.
