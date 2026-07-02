# claude-plugins

The standards-lab organization's Claude Code plugin marketplace. This is the harness level of the
organization's reference architecture: the agentic infrastructure that codifies the organization's
development processes, so the same processes are applied consistently across every project. `marathon`
ships first as v0; the set grows over time. It is managed with marathon itself, so the workflow is
battle-tested against its own source.

## Capability map

Broad and shallow; detail is added when a piece is about to be built.

- **marathon** — the sustainable long-haul development workflow, built on context engineering. The
  substance of the repository today; it distinguishes code from context projects, runs `plan`/`start`/
  `experiment` sessions, and coordinates a change across several projects in a workspace.
- **marketplace host** — the marketplace manifest and the per-plugin, independent versioning and release of
  each hosted plugin.
- **further plugins** (candidate) — the set grows as processes prove worth codifying: diagram generation
  (`tau-diagrams`), the dev-blog infrastructure (`tau-blog`), and skills that make AI-assisted development
  on the organization's architecture more effective.

## How this repository works

- **Plugin design lives in the skill files**, not in `context/`. See
  `design/skill-is-the-source-of-truth.md`.
- Candidates in flight are in `concepts/`; the latest session record and the next step are in `reset.md`.
