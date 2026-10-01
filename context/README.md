# claude-plugins

claude-plugins is the Standards Lab organization's Claude Code plugin marketplace. It is the
harness level of the organization's reference architecture: plugins that codify the
organization's development processes, so every project applies them the same way. The
repository is developed with marathon, so the workflow runs against its own source.

## Capability map

- **marathon**: the long-haul development workflow, being reshaped into a software factory
  (`factory`). The architect is involved twice, at the plan and at the session brief. Agents
  implement, review against standards, check against the spec, and maintain the context in
  between. See `marathon-factory.md` for the pipeline and profiles, `marathon-goals.md` for goals
  as the unit of parallel work, `marathon-briefs.md` for what the architect reads, and
  `marathon-commands.md` for the command set.
- **marathon-roadmap**: the extension that keeps `context/roadmap.toml` current. It folds into
  marathon core in `factory.goals` (`marathon-goals.md`).
- **marathon-architecture**: the extension that keeps the architecture layer. It shrinks to the
  pointers each repository's `STANDARDS.md` follows, which only the standards-reviewer reads
  (`marathon-factory.md`).
- **marathon-sitrep** (planned): situation reports over a date range for a chosen audience. See
  `marathon-sitrep.md`.
- **marathon-extraction** (planned): carries patterns a consumer proves back to the blueprint.
  See `marathon-extraction.md`.
- **marathon-references** (proposed): the references catalog as an extension. See
  `marathon-references-extension.md`.
- **Harness testing**: CI checks the repository's consistency on every push. Each release
  runs `claude plugin eval` suites, whose cases come from observed failures. See
  `harness-testing.md`.
- **Marketplace host**: the marketplace manifest, with each plugin versioned and released on its
  own.
- **Further plugins** (planned): the set grows as processes prove worth codifying, such as
  diagram generation, dev-blog infrastructure, and skills for the organization's architecture.

## How this repository works

- Each plugin's design lives in its own files under `plugins/<name>/`, not in `context/`.
  `scripts/check.sh` verifies that every pointer inside them resolves.
- Notes for planned work sit directly under `context/`. The session record is the workspace's
  reset file at the coordinator (standards-lab), so this repository keeps no `reset.md`.
