---
name: marathon
argument-hint: "[init | plan | start | experiment | reset | close | review]"
description: >
  A workflow for long-running development built on context engineering. Use this skill whenever
  the architect initializes a project from a planning concept; plans, advances, pauses, resumes,
  or closes a session; spikes an idea; runs one step across several projects in a workspace; hands
  off work because the context window is filling; or checks whether the notes have drifted from
  the code. Typical requests include "start a session", "plan the next step", "begin working on
  this", "hand off", "I'm running low on context", "resume where I left off", "close out this
  session", "wrap up and open a PR", "coordinate this across the repos", "review drift",
  "initialize this project", and "set up marathon here". marathon treats the repository as the
  only source of truth: it keeps a flat, top-level context/ directory of notes, deletes each note
  once the built work expresses it, and runs each session on its own branch. Prefer this skill
  for any structured, multi-session work on a marathon repository, even when the architect doesn't
  name it.
---

# Marathon

Version: 0.15.0

marathon is a workflow for long-running development. Each session plans, builds, and closes one
finished step, and the steps add up to a production-quality version of the original concept. The
project's written context lives in a top-level `context/` directory, and every session keeps it
current, so the context window holds the task instead of stale notes. The repository is the
source of truth, not an issue tracker and not the conversation.

Use this skill whenever a session begins, advances, pauses, resumes, or ends on a marathon
repository: one with a top-level `context/`, or a workspace of such repositories. Set up a new
repository with `init`.

## Behavior

These rules apply in every session: how a session plans, and how it hands work to another agent.

@behavior/planning.md

@behavior/delegation.md

## Mechanics

The session pipeline applies to every command:

@mechanics/pipeline.md

The pipeline refers to these files where it needs them:

- [`mechanics/reset-file.md`](./mechanics/reset-file.md): the session record, where it lives, its
  schema, and its Status values.
- [`mechanics/waves.md`](./mechanics/waves.md): lanes that run at the same time, their records,
  their checkouts and worktrees, and folding.
- [`mechanics/configuration.md`](./mechanics/configuration.md): the layout of
  `.claude/marathon.toml`.
- [`mechanics/hooks.md`](./mechanics/hooks.md): how extension hooks resolve and fire.

## Commands

The first argument names the command. Each command's playbook supplies the content of the
pipeline's stages.

| Command | When to use it | Playbook |
|---------|----------------|----------|
| `init` | Once, to set up marathon on a repository from a planning concept | [`commands/init.md`](./commands/init.md) |
| `plan` | To refine notes and decide the next step; changes only `context/` | [`commands/plan.md`](./commands/plan.md) |
| `start` | To advance the product one concrete step | [`commands/start.md`](./commands/start.md) |
| `experiment` | To set up a spike as a standalone project | [`commands/experiment.md`](./commands/experiment.md) |
| `reset` | To hand off when the context is filling but the step isn't done | [`commands/reset.md`](./commands/reset.md) |
| `close` | To finish a session whose work is done and validated | [`commands/close.md`](./commands/close.md) |
| `review` | On demand, to check the notes against the code and clean them up | [`commands/review.md`](./commands/review.md) |

Every session ends with `close` when its work is finished and validated, or with `reset` when it
hands off.

## References

Read each reference when its subject comes up:

- [`references/context-engineering.md`](./references/context-engineering.md): writing and
  maintaining the notes in `context/`.
- [`references/staged-execution.md`](./references/staged-execution.md): stages, checkpoints,
  and validation.
- [`references/workspace-coordination.md`](./references/workspace-coordination.md): the
  coordinator, the dependency order, and steps that span repositories.
- [`references/extensions.md`](./references/extensions.md): the extension system.
