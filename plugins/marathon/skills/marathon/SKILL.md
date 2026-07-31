---
name: marathon
argument-hint: "[init | plan | start | experiment | coordinate | reset | close | review | docs]"
description: >
  Sustainable long-haul development workflow built on context engineering. Use this skill whenever
  the developer is initializing a project from a planning concept; planning, advancing, pausing,
  resuming, or closing out a session; spiking an idea; coordinating one change across several projects
  in a workspace; handing off work because the context window is growing large; or reviewing whether
  design notes have drifted from the code. Natural triggers include "start a session", "plan the next
  step", "begin working on this", "hand off / I'm running low on context", "resume where I left off",
  "close out this session", "wrap up and open a PR", "coordinate this across the repos", "review
  drift", "initialize this project", and "set up marathon here". marathon keeps the repository itself
  the single source of truth: it manages a volatile-vs-stable top-level context/ tree, promotes and
  decays knowledge deliberately, and drives branch-based sessions. Prefer this skill for any
  structured, multi-session work on a marathon-managed repo, even when the developer doesn't name it
  explicitly.
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
- planning the next step, advancing it, pausing, resuming, or closing out a session,
- spiking an idea before committing to it,
- coordinating a single change across several marathon projects in a workspace,
- handing off because the context window is getting full but the work isn't done,
- checking whether the notes in `context/` still match the code, or
- authoring or curating the project's human-oriented reference documentation (the optional `docs/` tier).

If a repository doesn't have a `context/` directory yet and you want this workflow, start with `init`.

## Sub-commands

Route on the first argument. Each command has a detailed playbook under `commands/`.

| Command | When | Detail |
|---------|------|--------|
| `init` | One-time, to set up marathon on a repo from a planning concept | `commands/init.md` |
| `plan` | Refine concepts and settle what the next session should focus on; touches only `context/` | `commands/plan.md` |
| `start` | Advance the product one concrete step | `commands/start.md` |
| `experiment` | Spike an idea in the isolated `experiments/` directory | `commands/experiment.md` |
| `coordinate` | Run one change across several marathon projects in a workspace | `commands/coordinate.md` |
| `reset` | Hand off mid-session: context is filling but the work isn't done | `commands/reset.md` |
| `close` | The session's work is finished and validated | `commands/close.md` |
| `review` | On demand: check the notes against the code and clean them up | `commands/review.md` |
| `docs` | On demand: author and curate the optional human-oriented `docs/` tier | `commands/docs.md` |

## The session loop

```
  start ─► (plan mode: settle scope with developer) ─► do the step ─┬─ finished ─► close
    ▲                                                               │
    │      code project:    the step is an implementation guide;    └─ context filling,
    │                       the developer applies it                   work unfinished ─► reset
    │      context project: the step is the change itself,
    │                       authored directly (no guide, no handoff)
    │
    ├──◄── close — work done: (code: add tests/docs) · decay notes · write reset file · publish
    └──◄── reset — handing off: write reset file · leave the branch open · resume later
```

A session is one unit of work on one branch. It ends in one of two ways. If the work is finished,
`close` commits and publishes it. If the context window is filling up before the work is done, `reset`
checkpoints it and hands off, so a fresh context can pick up the same branch. Both write
`context/reset.md`; `close` is the version that also finishes the work. See `references/session-loop.md`
for the full walk-through.

## Project kind

A marathon project is one of two kinds, declared once at `init` in `.claude/marathon.toml` and read by
every session afterward:

```toml
[project]
kind = "code"      # production source the developer authors and answers for
# kind = "context" # the whole repository is context — the agent authors it directly
```

- **code** — the repository contains production source code: the implementation logic that makes a
  program behave. The role boundary applies. The agent drafts each change as an implementation guide,
  the developer applies it, and closeout adds tests and documentation.
- **context** — the repository *is* context: prose, configuration, and skills (which are advanced
  context, not source). There is no production code to hand off, so the agent authors the whole
  repository directly — no implementation guide, no tests. The developer sets direction and reviews and
  approves each change; the pull request is the ownership seam. A context project can still version and
  release what it ships (a plugin, a document set); it just has no code layer.

Project kind decides how `start` and `close` behave, and how `coordinate` treats each project in a
fan-out. When in doubt, ask whether `git blame` on the repository's real deliverable should point at
the developer (code) or at the agent under the developer's review (context). See
`references/role-boundary.md`.

## Working sessions: plan, start, experiment

Three commands do the actual work of a session; which one you run says what the session is for.

- **plan** — a planning and curation session that touches only `context/`. Create and refine concepts,
  settle a design, and decide what the next `start` should focus on. It changes no product — no code,
  no skill, no prose deliverable — only the agent's written context. Use it for the bigger-picture
  thinking between builds, or to get a concept ready before building it.
- **start** — advance the product one concrete step. On a **code** project you plan in plan mode, write
  the implementation guide, and stop for the developer to apply it. On a **context** project you plan
  in plan mode, then author the change directly. Either way the session covers one step and nothing
  wider.
- **experiment** — a spike in the top-level `experiments/` directory. Results don't quietly become
  settled design; you promote them deliberately at closeout, if at all. On a context project this is
  where a new skill or agent idea gets tried before it becomes real.

`plan`, `start`, and `experiment` all end the same way: `close` when the work is finished, or `reset`
to hand off. `review` and `docs` are on-demand maintenance passes rather than working sessions, but
they end the same way too — their changes land through `close`, recorded under their own Session
values.

## Iterative development

marathon plans one step at a time. You work out the immediate next step in detail and nothing past it,
and you don't let a session's focus spread beyond that one area. Start from the lowest-level requirement,
build it, and let the next step come into view once it's done. Over many sessions, these small finished
steps stack up into the complete solution.

This is why the notes in `context/` stay shallow until the work on them is close. Planning far ahead
commits you to decisions you haven't earned yet, and it's usually the planning-ahead, not the building,
that turns out wrong. Keep the focus narrow and let the design grow out of working code.

## Planning is half the work

Planning carries as much weight as building. In `init`, in a fresh `start`, in a `plan` session, and in
`review`, the planning phase is where the real architectural thinking happens — you work out what the
step involves, how deep it needs to go, and how it fits the larger design. The quality of the
implementation is largely set here, so plan with enough depth and clarity that you come out with a clear
picture of what you're about to build. Rushing planning to get to the code is how a session builds the
wrong thing, or the right thing at the wrong depth.

Because the discussion ranges wider than the single step, planning is also where you tend the context:
capture ideas that belong to later steps as concepts in `concepts/`, cull the ones the discussion has
ruled out, and start to spot what the next session's focus should be.

marathon plans in plan mode and settles things with the developer before any consequential action —
`init` before scaffolding, a fresh `start` before writing the guide or authoring the change, `plan` and
`review` before changing notes. (A `start` that resumes a handoff skips this and picks up the existing
plan.)

## How context is organized

The `context/` directory holds the project's written knowledge, grouped by how often each part changes:

- `context/README.md` — vision and a broad map of the capabilities the project will need.
- `context/design/` — design notes: settled intent the code can't express yet.
- `context/concepts/` — concepts: ideas that aren't settled. Most of the churn is here.
- `context/guide.md` — the current session's implementation guide, on a **code** project; deleted at
  closeout. A context project has no guide.
- `context/reset.md` — the latest session's record and the pointer to the next step.
- `experiments/` — top-level, for isolated spike work, created on demand.
- `docs/` — top-level, optional, for human-oriented reference documentation; opted into via `docs`.
- the source code — the implementation, and the last word on what the project does.

A `design/` note gets deleted once the code expresses it; a concept gets promoted to `design/` when it
proves out, or culled when it doesn't. The full rules, and how to decide where something belongs, are in
`references/context-engineering.md`.

`docs/` is a separate, optional tier: human-oriented reference documentation, a peer to the agent-oriented
`context/`. It shares the same maintenance discipline but not the decay rule — a `docs/` page describes
code that exists, so it doesn't decay toward the code the way a `design/` note does. The `docs` command
authors and curates it; see `commands/docs.md`.

## Role boundary

marathon divides the work between the developer and the agent, and where the line falls depends on the
project kind.

On a **code** project, the developer owns the production code — they apply it, adjust it, and answer for
it. The agent drafts each change in the implementation guide for the developer to apply, and writes
everything else outright: tests, comments and API docs, prose documentation, the files in `context/`,
the reset file, and the guide itself. If you're unsure who owns a line, ask who `git blame` should show
on it in six months — the developer, because they put it there.

On a **context** project there is no production code, so there is nothing to hand off. The agent authors
the repository directly — its skills, prose, configuration, and everything in `context/`. The developer
owns direction and reviews and approves the result; the pull request is where ownership is exercised.

Keep the rule language-neutral. See `references/role-boundary.md`. Everything the agent writes in
prose, on either kind of project, follows the voice standard in `references/writing-voice.md`.

## Continuity

Continuity lives in the repository, not in a tracker or the chat. `context/reset.md` holds the last
session's record, and its Next-focus line is the handoff: written at the end of one session, read at the
start of the next. The Status line — `handoff` or `closeout` — tells the next session how to resume:
pick up the open branch after a handoff, or start a new branch after a closeout.

A repository in a workspace can also reach a resting point: its deliverable is released and it has no
next step of its own. `close` then deletes the reset file instead of rewriting it, and continuity
moves to the workspace coordinator. A session that finds no `context/reset.md` reads that as the
resting state and defers to the coordinator's reset file.

marathon's git workflow is branches and commits. A finished branch is published as the change proposal
the project's remote uses — a pull request on GitHub, a merge request on GitLab, or the equivalent — with
its description copied from the reset file. The remote platform is declared at `init`. Issues, boards,
and other platform project management aren't part of the core; they come from optional extensions.

## The reset file

`context/reset.md` is a single file, rewritten at each `reset` or `close`:

```markdown
# reset · wire-config-loader

- **Status:** closeout            # handoff | closeout
- **Session:** start              # init | plan | start | experiment | review | docs
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
resume the same branch without working it out again. (Older reset files may carry a `Session type:` line
naming `development`/`context`/`experiment`; read it as the equivalent working session.)

## Workspace coordination

A single marathon project manages one repository. When several marathon projects live together as
siblings under one directory — a workspace — a change sometimes crosses more than one of them in
dependency order. `coordinate` runs that change: it detects the workspace, enumerates the participating
projects in order, and fans out a session to each, honoring each project's kind. The workspace itself
holds no context, and the coordination run leaves no branch, reset, or committed context of its own;
continuity stays per repository. See `commands/coordinate.md` and `references/workspace-coordination.md`.

Any marathon command run from the workspace directory routes by the same map: detect the workspace,
resolve the coordinator through its `[workspace] role`, and read the coordinator's `context/reset.md`
as the continuity anchor. The reset file's Branch line names the repository the session continues in,
so a session started at the workspace root finds its way without being told.

## Extension hooks

The core's git workflow is platform-neutral; the remote platform, and the command that publishes a
finished branch (`gh pr create`, `glab mr create`, or another), is declared at `init` and stored in
`.claude/marathon.toml`. On top of that, the core calls named hook points — `on-init`,
`on-session-start`, `on-commit`, `on-closeout` — where an optional project-management extension can
mirror the repo's state onto a platform's tracker (issues, boards). The points do nothing when no
extension is set up. Information flows one way: an extension reads the repo and projects outward, never
back into the core. See `references/extension-hooks.md`.
