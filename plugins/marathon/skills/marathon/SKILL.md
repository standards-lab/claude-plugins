---
name: marathon
argument-hint: "[init | start <development|context|experiment> | reset | close | review | docs]"
description: >
  Sustainable long-haul development workflow built on context engineering. Use this skill whenever
  the developer is initializing a project from a planning concept; starting, pausing, resuming, or
  closing out a development session; handing off work because the context window is growing large; or
  reviewing whether design notes have drifted from the code. Natural triggers include "start a
  session", "begin working on this", "hand off / I'm running low on context", "resume where I left
  off", "close out this session", "wrap up and open a PR", "review drift", "initialize this project",
  and "set up marathon here". marathon keeps the repository itself the single source of truth: it
  manages a volatile-vs-stable top-level context/ tree, promotes and decays knowledge deliberately,
  and drives branch-based sessions. Prefer this skill for any structured, multi-session development
  on a marathon-managed repo, even when the developer doesn't name it explicitly.
---

# Marathon

marathon is a workflow for long-haul development that stays sustainable. Within a session you work
fast: plan the next step, build it, close it out. Across the project you go the distance, advancing one
finished step at a time toward a production-quality version of the original concept.

What keeps that sustainable is how marathon handles context. A model's working context is limited, and
if you let it fill with old plans, answered questions, and documentation the code has outgrown, you lose
room to work and the old notes start to contradict the code. marathon keeps the project's written
context in a top-level `context/` directory and maintains it as it goes — promoting notes that prove
out, deleting notes the code has caught up to. The repository is the source of truth, not an external
tracker and not the conversation.

## When to use this skill

Use marathon whenever you're working on a marathon-managed repository (one with a top-level `context/`
directory) and you're:

- initializing a new project from a planning concept,
- starting, pausing, resuming, or closing out a session,
- handing off because the context window is getting full but the work isn't done,
- checking whether the notes in `context/` still match the code, or
- authoring or curating the project's human-oriented reference documentation (the optional `docs/` tier).

If a repository doesn't have a `context/` directory yet and you want this workflow, start with `init`.

## Sub-commands

Route on the first argument. Each command has a detailed playbook under `commands/`.

| Command | When | Detail |
|---------|------|--------|
| `init` | One-time, to set up marathon on a repo from a planning concept | `commands/init.md` |
| `start [type]` | Begin or resume a session; `type` is `development` (default), `context`, or `experiment` | `commands/start.md` |
| `reset` | Hand off mid-session: context is filling but the work isn't done | `commands/reset.md` |
| `close` | The session's work is finished and validated | `commands/close.md` |
| `review` | On demand: check the notes against the code and clean them up | `commands/review.md` |
| `docs` | On demand: author and curate the optional human-oriented `docs/` tier | `commands/docs.md` |

## The session loop

```
  start ──► (plan mode: settle scope with developer) ──► implementation guide ──► STOP
    ▲                                                                              │
    │                                                              developer applies code
    │                                                                              │
    │                                                                              ▼
    └──◄── close (work done: publish) ─or── reset (handing off: resume later) ──◄── AI closeout
                                                                    (tests, docs, decay, reset file)
```

A session is one unit of work on one branch. It ends in one of two ways. If the work is finished,
`close` commits and publishes it. If the context window is filling up before the work is done, `reset`
checkpoints it and hands off, so a fresh context can pick up the same branch. Both write
`context/reset.md`; `close` is the version that also finishes the work. See `references/session-loop.md`
for the full walk-through.

## Session types

`start` takes a type that decides what closeout does:

- **development** (the default) — plan the next step, write an implementation guide, then stop and let
  the developer apply the production code. Closeout adds tests and documentation and cleans up the notes.
- **context** — planning and writing work. Edit the files in `context/` directly; there is no code
  handoff. Use it for bigger-picture planning between builds.
- **experiment** — a spike. Work happens in the top-level `experiments/` directory. Results don't
  quietly become settled design; you promote them deliberately at closeout, if at all.

## Iterative development

marathon plans one step at a time. You work out the immediate next step in detail and nothing past it,
and you don't let a session's focus spread beyond that one area. Start from the lowest-level requirement,
build it, and let the next step come into view once it's done. Over many sessions, these small finished
steps stack up into the complete solution.

This is why the notes in `context/` stay shallow until the work on them is close. Planning far ahead
commits you to decisions you haven't earned yet, and it's usually the planning-ahead, not the building,
that turns out wrong. Keep the focus narrow and let the design grow out of working code.

## Planning is half the work

Planning carries as much weight as building. In `init`, in a fresh `start`, and in `review`, the
planning phase is where the real architectural thinking happens — you work out what the step involves,
how deep it needs to go, and how it fits the larger design. The quality of the implementation is largely
set here, so plan with enough depth and clarity that you come out with a clear picture of what you're
about to build. Rushing planning to get to the code is how a session builds the wrong thing, or the
right thing at the wrong depth.

Because the discussion ranges wider than the single step, planning is also where you tend the context:
capture ideas that belong to later steps as candidate notes in `concepts/`, cull candidates the
discussion has ruled out, and start to spot what the next session's focus should be.

marathon plans in plan mode and settles things with the developer before any consequential action —
`init` before scaffolding, a fresh `start` before writing the guide, `review` before changing notes. (A
`start` that resumes a handoff skips this and picks up the existing plan.)

## How context is organized

The `context/` directory holds the project's written knowledge, grouped by how often each part changes:

- `context/README.md` — vision and a broad map of the capabilities the project will need.
- `context/design/` — design notes: settled intent the code can't express yet.
- `context/concepts/` — candidate notes: ideas that aren't settled. Most of the churn is here.
- `context/guide.md` — the current session's implementation guide; deleted at closeout.
- `context/reset.md` — the latest session's record and the pointer to the next step.
- `experiments/` — top-level, for isolated spike work, created on demand.
- `docs/` — top-level, optional, for human-oriented reference documentation; opted into via `docs`.
- the source code — the implementation, and the last word on what the project does.

A `design/` note gets deleted once the code expresses it; a candidate gets promoted to `design/` when it
proves out, or culled when it doesn't. The full rules, and how to decide where something belongs, are in
`references/context-engineering.md`.

`docs/` is a separate, optional tier: human-oriented reference documentation, a peer to the agent-oriented
`context/`. It shares the same maintenance discipline but not the decay rule — a `docs/` page describes
code that exists, so it doesn't decay toward the code the way a `design/` note does. The `docs` command
authors and curates it; see `commands/docs.md`.

## Role boundary

The developer owns the production code — they apply it, adjust it, and answer for it. The agent drafts
each change in the implementation guide for the developer to apply, and writes everything else outright:
tests, comments and API docs, prose documentation, the files in `context/`, the reset file, and the
guide itself. If you're unsure, ask who `git blame` should show on a production line in six months — the
developer, because they put it there. Keep the rule language-neutral. See `references/role-boundary.md`
and `references/implementation-guides.md`.

## Continuity

Continuity lives in the repository, not in a tracker or the chat. `context/reset.md` holds the last
session's record, and its Next-focus line is the handoff: written at the end of one session, read at the
start of the next. The Status line — `handoff` or `closeout` — tells `start` how to resume: pick up the
open branch after a handoff, or start a new branch after a closeout.

marathon's git workflow is branches and commits. A finished branch is published as the change proposal
the project's remote uses — a pull request on GitHub, a merge request on GitLab, or the equivalent — with
its description copied from the reset file. The remote platform is declared at `init`. Issues, boards,
and other platform project management aren't part of the core; they come from optional extensions.

## The reset file

`context/reset.md` is a single file, rewritten at each `reset` or `close`:

```markdown
# reset · wire-config-loader

- **Status:** closeout            # handoff | closeout
- **Session type:** development
- **Branch:** wire-config-loader

## Disposition
- **Integrated:** removed the "three-phase load" note from design/config.md — the code now expresses it (config/loader).
- **Promoted:** concepts/config-validation.md → design/ (validation rules settled this session).
- **Culled:** dropped the env-override idea — the loader implementation went another way.
- **Retained:** design/config.md "secret sourcing" — still unbuilt.

## Next-focus
Add secret sourcing on top of the validated loader. Start here next session.
```

On a handoff, Next-focus records the in-progress state and the exact next move, so a fresh context can
resume the same branch without working it out again.

## Extension hooks

The core's git workflow is platform-neutral; the remote platform, and the command that publishes a
finished branch (`gh pr create`, `glab mr create`, or another), is declared at `init` and stored in
`.claude/marathon.toml`. On top of that, the core calls named hook points — `on-init`,
`on-session-start`, `on-commit`, `on-closeout` — where an optional project-management extension can
mirror the repo's state onto a platform's tracker (issues, boards). The points do nothing when no
extension is set up. Information flows one way: an extension reads the repo and projects outward, never
back into the core. See `references/extension-hooks.md`.
