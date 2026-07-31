# marathon

A sustainable long-haul development workflow built on context engineering. marathon keeps the
repository itself the single source of truth: it maintains a top-level `context/` directory —
promoting notes that prove out, deleting notes the code has caught up to — and drives work as small,
branch-based sessions, one finished step at a time.

This README is a quick reference. The skill itself is the source of truth for how marathon behaves;
its files under [`skills/marathon/`](./skills/marathon/) — `SKILL.md`, the `commands/` playbooks, and
the `references/` deep-dives — describe the full detail.

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
| `coordinate` | Run one change across several marathon projects in a workspace. |
| `reset` | Hand off mid-session so a fresh context can resume the same branch. |
| `close` | Finish and publish a completed session. |
| `review` | Audit the notes for drift from the code and clean them up. |
| `docs` | Author and curate the optional human-oriented `docs/` tier. |

## The session loop

A session is one step, on one branch. A working session — `plan`, `start`, or `experiment` — begins by
reading `context/reset.md`: a `closeout` status means plan a new step in plan mode; a `handoff` status
means resume the open branch. The session ends with `close` (finished → publish) or `reset` (handing
off → resume later). Both rewrite `context/reset.md`, whose Next-focus line is the handoff to the next
session.

## Project kinds

A project is declared `code` or `context` at `init`, in `.claude/marathon.toml`:

- **code** — the repository holds production source the developer owns. `start` drafts an implementation
  guide, the developer applies it, and closeout adds tests and documentation.
- **context** — the repository *is* context (skills, prose, configuration). The agent authors it
  directly under the developer's review; there is no implementation guide and no tests.

## Workspace coordination

When several marathon projects live as siblings under one directory — a workspace — `coordinate` runs a
single change across them in dependency order, fanning out a session to each and honoring its kind. The
workspace holds no context of its own; continuity stays per repository. See
[`references/workspace-coordination.md`](./skills/marathon/references/workspace-coordination.md).

## Releases

Releases are tag-driven, cut from [`CHANGELOG.md`](./CHANGELOG.md). Pushing a tag `marathon/v<version>`
triggers the host's release workflow.
