# claude-plugins

claude-plugins is the Standards Lab organization's Claude Code plugin marketplace. It is the
harness level of the organization's reference architecture: plugins that codify the
organization's development processes, so every project applies them the same way. The
repository is developed with marathon, so the workflow runs against its own source.

## Capability map

- **marathon**: the long-haul development workflow. Each session plans, builds, and closes one
  step, in stages that commit once their check passes and stop at checkpoints the architect
  confirms. It runs on a standalone repository or across a workspace of repositories, and it
  ships planner, executor, editor, and reviewer subagent profiles.
- **marathon-roadmap**: an extension that keeps `context/roadmap.toml` current through marathon's
  session hooks. The manifest holds the goals, tasks, backlog, and `next` sequence.
- **marathon-architecture**: an extension that keeps the architecture layer, the principles and
  conventions that apply beyond one repository, filled by promotion from validated notes.
- **marathon-sitrep** (planned): situation reports over a date range for a chosen audience. See
  `marathon-sitrep.md`.
- **marathon-extraction** (planned): carries patterns a consumer proves back to the blueprint.
  See `marathon-extraction.md`.
- **marathon-references** (proposed): the references catalog as an extension. See
  `marathon-references-extension.md`.
- **Harness testing**: CI checks the repository's consistency, and behavioral tests wait for real
  failure modes. See `harness-testing.md`.
- **Marketplace host**: the marketplace manifest, with each plugin versioned and released on its
  own.
- **Further plugins** (planned): the set grows as processes prove worth codifying, such as
  diagram generation, dev-blog infrastructure, and skills for the organization's architecture.

## How this repository works

- Each plugin's design lives in its own files under `plugins/<name>/`, not in `context/`.
  `scripts/check.sh` verifies that every pointer inside them resolves.
- Notes for planned work sit directly under `context/`. The session record is the workspace's
  reset file at the coordinator (standards-lab), so this repository keeps no `reset.md`.
