# claude-plugins

The Standards Lab organization's Claude Code plugin marketplace. This is the harness level of the
organization's reference architecture: the agentic infrastructure that codifies the organization's
development processes, so the same processes are applied consistently across every project. `marathon`
ships first as v0; the set grows over time. It is managed with marathon itself, so the workflow is
exercised against its own source.

## Capability map

Broad and shallow; detail is added when a capability is about to be built.

- **marathon** — the sustainable long-haul development workflow, built on context engineering. The
  substance of the repository today; it distinguishes code from context projects and runs `plan`/
  `start`/`experiment` sessions on a standalone repository or in a workspace, where one session can
  carry a step across several member projects.
- **marathon-roadmap** — the first marathon extension: the roadmap manifest convention
  (`context/roadmap.toml` — goals, tasks, backlog) kept current through marathon's session hooks.
- **marathon-sitrep** (candidate) — audience-calibrated situation reports over a date range, read
  from the workspace's git history, session record, and roadmap deltas; the first enhancement-facet
  extension. Concept: `concepts/marathon-sitrep.md`.
- **marketplace host** — the marketplace manifest and the per-plugin, independent versioning and release of
  each hosted plugin.
- **further plugins** (candidate) — the set grows as processes prove worth codifying: diagram
  generation, dev-blog infrastructure, and skills aligned to the organization's architecture.

## How this repository works

- **Plugin design lives in the skill files**, not in `context/`. See
  `design/skill-is-the-source-of-truth.md`.
- Candidates in flight go to `concepts/`. The session record is the workspace's single reset file
  at the coordinator (standards-lab); this repository keeps no `reset.md` of its own.
