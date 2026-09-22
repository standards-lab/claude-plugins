---
name: marathon
argument-hint: "[init | plan | start | experiment | reset | close | review]"
description: >
  Sustainable long-haul development workflow built on context engineering. Use this skill whenever
  the architect is initializing a project from a planning concept; planning, advancing, pausing,
  resuming, or closing out a session; spiking an idea; running one step across several projects
  in a workspace; handing off work because the context window is growing large; or reviewing whether
  the notes have drifted from the code. Natural triggers include "start a session", "plan the next
  step", "begin working on this", "hand off / I'm running low on context", "resume where I left off",
  "close out this session", "wrap up and open a PR", "coordinate this across the repos", "review
  drift", "initialize this project", and "set up marathon here". marathon keeps the repository itself
  the single source of truth: it manages a flat top-level context/ directory of notes, settles and
  deletes them deliberately, and drives branch-based sessions. Prefer this skill for any
  structured, multi-session work on a marathon-managed repo, even when the architect doesn't name it
  explicitly.
---

# Marathon

Version: 0.13.0

marathon is a workflow for long-haul development: each session plans, builds, and closes one
finished step, and the steps add up to a production-quality version of the original concept. It
keeps the project's written context in a top-level `context/` directory and maintains it as it
goes, so a limited context window holds the task rather than stale notes. The repository is the
source of truth, not an external tracker and not the conversation.

Use this skill on a marathon-managed repository (one with a top-level `context/`, or a workspace
of them) whenever a session begins, advances, pauses, resumes, or ends. Start a new repository with
`init`.

## Behavior

Always active: how sessions plan, and how they hand a unit of work to another agent.

@behavior/planning.md

@behavior/delegation.md

## Mechanics

Always active: the session pipeline every command runs.

@mechanics/pipeline.md

Consulted where the pipeline points at them:

- [`mechanics/reset-file.md`](./mechanics/reset-file.md) — the session record: location, waves,
  schema, and Status.
- [`mechanics/configuration.md`](./mechanics/configuration.md) — the `.claude/marathon.toml`
  layout.
- [`mechanics/hooks.md`](./mechanics/hooks.md) — extension hook resolution and firing.

## Commands

Route on the first argument; each command's playbook supplies its stages' content.

| Command | When | Playbook |
|---------|------|----------|
| `init` | One-time, to set up marathon on a repo from a planning concept | [`commands/init.md`](./commands/init.md) |
| `plan` | Refine concepts and settle what the next session should focus on; touches only `context/` | [`commands/plan.md`](./commands/plan.md) |
| `start` | Advance the product one concrete step | [`commands/start.md`](./commands/start.md) |
| `experiment` | Set up a spike as a standalone project of its own | [`commands/experiment.md`](./commands/experiment.md) |
| `reset` | Hand off mid-session: context is filling but the work isn't done | [`commands/reset.md`](./commands/reset.md) |
| `close` | The session's work is finished and validated | [`commands/close.md`](./commands/close.md) |
| `review` | On demand: check the notes against the code and clean them up | [`commands/review.md`](./commands/review.md) |

Every session ends through `close` (finished and validated) or `reset` (handing off).

## References

Consulted when their subject is in play:

- [`references/context-engineering.md`](./references/context-engineering.md) — writing and
  tending the notes in `context/`.
- [`references/staged-execution.md`](./references/staged-execution.md) — stages, checkpoints,
  and validation.
- [`references/workspace-coordination.md`](./references/workspace-coordination.md) — the
  coordinator, the order map, and cross-repo steps.
- [`references/extensions.md`](./references/extensions.md) — the extension system.
