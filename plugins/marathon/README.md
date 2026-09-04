# marathon

A sustainable long-haul development workflow built on context engineering. marathon keeps the
repository itself the single source of truth: it maintains a top-level `context/` directory —
promoting notes that prove out, deleting notes the code has caught up to — and drives work as small,
branch-based sessions, one finished step at a time.

This README is a quick reference. The skill itself is the source of truth for how marathon behaves;
its files under [`skills/marathon/`](./skills/marathon/) — `SKILL.md`, the `commands/` playbooks, the
`mechanics/` execution specs, the `behavior/` always-active conduct, and the `references/` deep-dives —
describe the full detail.

## Install

```bash
claude plugin marketplace add standards-lab/claude-plugins
claude plugin install marathon@standards-lab
```

Then run `marathon init` in a repository to set it up, or `marathon start` on a repository that already
has a top-level `context/` directory.

## Commands

Invoke as `marathon <command>` (or `/marathon:marathon <command>`).

| Command | What it does |
|---------|--------------|
| `init` | One-time setup of marathon on a repository, from a planning concept. |
| `plan` | Planning/curation session that touches only `context/` — refine concepts, decide the next step. |
| `start` | Advance the product one concrete step. |
| `experiment` | Spike an idea in the isolated `experiments/` directory. |
| `reset` | Hand off mid-session so a fresh context can resume the same branch. |
| `close` | Finish and publish a completed session. |
| `review` | Audit the notes for drift from the code and clean them up. |
| `docs` | Author and curate the optional human-oriented `docs/` tier. |

## The session loop

A session is one step, on one branch. A working session — `plan`, `start`, or `experiment` — begins by
reading the reset file — a standalone project's `context/reset.md`, or in a workspace the single
reset kept at the coordinator: a `closeout` status means plan a new step in plan mode; a `handoff`
status means resume the open branch. The session ends with `close` (finished → publish) or `reset`
(handing off → resume later). Both rewrite the reset file, whose Next-focus line is the handoff to
the next session.

## Project kinds

A project is declared `code` or `context` at `init`, in `.claude/marathon.toml`:

Both kinds execute in stages, and the architect reviews each stage, uncommitted, before it commits.
The kind sets what a stage is:

- **code** — the repository contains production source. A stage is a compilation unit brought to
  green with its tests, and the step ends with a whole-module validation.
- **context** — the repository *is* context (skills, prose, configuration). A stage is the smallest
  set of files that must change together, checked by the repository's own consistency script and
  a coherent read; there are no tests.

## Workspaces

When several marathon projects live as siblings under one directory — a workspace — a session's step
may span them: one session works the touched repos in the coordinator's declared dependency order,
each on its own branch under the step's shared slug. The workspace has no context of its own; its
continuity lives in the single reset file at the coordinator, and every experiment lives in the
coordinator's `experiments/`. See
[`references/workspace-coordination.md`](./skills/marathon/references/workspace-coordination.md).

## Extensions

marathon is extensible: a separately installed skill can act at the five universal hooks every
session fires, and a repository enables it in `.claude/marathon.toml`, for itself under `[project]`
or for a whole workspace under the coordinator's `[workspace]`. The repository stays the source of
truth; anything an extension projects outward is a read-only mirror. See
[`references/extensions.md`](./skills/marathon/references/extensions.md) for the system and
[`mechanics/hooks.md`](./skills/marathon/mechanics/hooks.md) for the firing spec.

## Releases

Releases are tag-driven, cut from [`CHANGELOG.md`](./CHANGELOG.md). Pushing a tag `marathon/v<version>`
triggers the host's release workflow.
