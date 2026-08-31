---
name: marathon
argument-hint: "[init | plan | start | experiment | reset | close | review | docs]"
description: >
  Sustainable long-haul development workflow built on context engineering. Use this skill whenever
  the architect is initializing a project from a planning concept; planning, advancing, pausing,
  resuming, or closing out a session; spiking an idea; running one step across several projects
  in a workspace; handing off work because the context window is growing large; or reviewing whether
  design notes have drifted from the code. Natural triggers include "start a session", "plan the next
  step", "begin working on this", "hand off / I'm running low on context", "resume where I left off",
  "close out this session", "wrap up and open a PR", "coordinate this across the repos", "review
  drift", "initialize this project", and "set up marathon here". marathon keeps the repository itself
  the single source of truth: it manages a volatile-vs-stable top-level context/ tree, promotes and
  decays knowledge deliberately, and drives branch-based sessions. Prefer this skill for any
  structured, multi-session work on a marathon-managed repo, even when the architect doesn't name it
  explicitly.
---

# Marathon

Version: 0.9.0

marathon is a workflow for long-haul development that stays sustainable. Within a session you work
fast: plan the next step, build it, close it out. Across the project you go the distance, advancing one
finished step at a time toward a production-quality version of the original concept.

What keeps that sustainable is how marathon handles context. A model's working context is limited, and
if you let it fill with old plans, answered questions, and documentation the code has outgrown, you lose
room to work and the old notes start to contradict the code. marathon keeps the project's written
context in a top-level `context/` directory and maintains it as it goes — promoting notes that prove
out, deleting notes the code has caught up to. The repository is the source of truth, not an external
tracker and not the conversation.

Use this skill whenever you're working on a marathon-managed repository (one with a top-level
`context/` directory, or a workspace of them) and a session is beginning, advancing, pausing,
resuming, or ending. If a repository doesn't have marathon yet and you want this workflow, start
with `init`.

## Behavior

Always active, loaded with this skill: the planning conduct every session settles scope by.

@behavior/planning.md

## Mechanics

The execution layer, loaded with this skill: the five-stage session pipeline every command runs —
locate, start, settle, execute, conclude — with its LOCATE routing and invariants.

@mechanics/pipeline.md

Consulted where the pipeline points at them:

- [`mechanics/reset-file.md`](./mechanics/reset-file.md) — the reset file: where it lives
  (standalone project, or the workspace's single reset at the coordinator), the schema, and the
  Status semantics.
- [`mechanics/configuration.md`](./mechanics/configuration.md) — the canonical
  `.claude/marathon.toml` layout: project kind, remote, workspace declaration, extensions.
- [`mechanics/hooks.md`](./mechanics/hooks.md) — the extension hook firing spec: resolution of
  the enabled set, the firing table, per-command ordering constraints.

## Commands

Route on the first argument; each command's playbook supplies its stages' content.

| Command | When | Playbook |
|---------|------|----------|
| `init` | One-time, to set up marathon on a repo from a planning concept | [`commands/init.md`](./commands/init.md) |
| `plan` | Refine concepts and settle what the next session should focus on; touches only `context/` | [`commands/plan.md`](./commands/plan.md) |
| `start` | Advance the product one concrete step | [`commands/start.md`](./commands/start.md) |
| `experiment` | Spike an idea in the isolated `experiments/` directory | [`commands/experiment.md`](./commands/experiment.md) |
| `reset` | Hand off mid-session: context is filling but the work isn't done | [`commands/reset.md`](./commands/reset.md) |
| `close` | The session's work is finished and validated | [`commands/close.md`](./commands/close.md) |
| `review` | On demand: check the notes against the code and clean them up | [`commands/review.md`](./commands/review.md) |
| `docs` | On demand: author and curate the optional human-oriented `docs/` tier | [`commands/docs.md`](./commands/docs.md) |

`plan`, `start`, and `experiment` are the working sessions; `review` and `docs` are on-demand
maintenance passes. All five end through `close` (finished and validated) or `reset` (handing
off), recorded under their own Session values.

## References

Consulted when their subject is in play:

- [`references/context-engineering.md`](./references/context-engineering.md) — how `context/` is
  organized and maintained: the tiers, promote/decay/cull, assumption annotations, the `docs/`
  tier's place in the lifecycle.
- [`references/staged-execution.md`](./references/staged-execution.md) — how a code project's
  `start` executes: stages, the stage list, the report, review outcomes, re-plan, validation.
- [`references/workspace-coordination.md`](./references/workspace-coordination.md) — the design
  behind workspaces: the coordinator, the order map, cross-repo steps, continuity.
- [`references/extensions.md`](./references/extensions.md) — the extension system: installed vs.
  enabled, declarations, the hook points, the source-of-truth rule.
