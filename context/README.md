# claude-plugins

claude-plugins is the Standards Lab organization's Claude Code plugin marketplace. It is the
harness level of the organization's reference architecture: plugins that codify the
organization's development processes, so every project applies them the same way. The
repository is developed with marathon, so the workflow runs against its own source.

## Capability map

- **marathon**: the long-haul development workflow, run as a software factory. The architect is
  involved twice per task, at the task brief and at the session brief, and subagents implement,
  review against standards and the spec, and keep the context current in between. Goals and
  tasks in the roadmap manifest are the units of work. PLAN runs each repository's optional
  `currency` command first, so a task begins by upgrading what trails, and SHIP runs an optional
  release gate before each tag. Its design lives in `plugins/marathon/`.
- **marathon-architecture**: the extension that finds the architecture layer each repository's
  `STANDARDS.md` points to, so the standards-reviewer can apply it.
- **marathon-sitrep** (planned): situation reports over a date range for a chosen audience. See
  `marathon-sitrep.md`.
- **marathon-extraction** (planned): carries patterns a consumer proves back to the blueprint.
  See `marathon-extraction.md`.
- **marathon-references** (proposed): the references catalog as an extension. See
  `marathon-references-extension.md`.
- **Harness testing**: `scripts/check.sh` checks consistency and runs `claude plugin validate` on
  every push. The release gate, `scripts/gate.sh`, adds a plugin's `claude plugin eval` suite
  before each tag; marathon has one case, from observed failures. See `harness-testing.md`.
- **Marketplace host**: the marketplace manifest, with each plugin versioned and released on its
  own.
- **Further plugins** (planned): the set grows as processes prove worth codifying, such as
  diagram generation, dev-blog infrastructure, and skills for the organization's architecture.

## How this repository works

- Each plugin's design lives in its own files under `plugins/<name>/`, not in `context/`.
  `scripts/check.sh` verifies that every pointer inside them resolves.
- Notes for planned work sit directly under `context/`. The `factory` goal is rooted here, so its
  record is `context/goals/factory.md`; the workspace roadmap is at the coordinator
  (standards-lab).
