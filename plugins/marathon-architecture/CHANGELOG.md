# Changelog

All notable changes to the marathon-architecture plugin are documented here. Versions follow
[Semantic Versioning](https://semver.org/spec/v2.0.0.html); dates and release links live on the
GitHub releases the tags cut.

## v0.2.1

### Changed

- **Targets marathon 0.14.** A note is promoted once two things hold: the work it describes is
  built and validated, as a closeout's **Validated** entry records, and its knowledge applies
  beyond its repository. Notes no longer state that they are settled.
- **The prose is rewritten for clarity**, with the rules unchanged.

## v0.2.0

### Changed

- **Targets marathon 0.13** and its flat `context/`: the promotion path reads a concept note, the
  same note once it states itself settled, and a page once the knowledge generalizes. A
  workspace landing is a concept note in the architecture repository's `context/`.
- **Trimmed**: `on-start` layers in `references/architecture-layer.md` instead of restating it,
  and the skill's files are about 40% shorter.

## v0.1.0

### Added

- **The architecture layer as a marathon extension**, on the marathon 0.12 extension contract.
  The layer marathon core carried as a mandatory tier through 0.11 is now a convention a project
  opts into: a top-level `architecture/` directory on a standalone project, or the one member
  repository this extension's own `.claude/marathon-architecture.toml` names, at the coordinator.
  A project whose knowledge never generalizes past itself enables nothing and scaffolds nothing.
- **Hook connections** — `on-start` layers the conventions in, bootstraps a standalone project's
  directory, adopts a layer that already exists, and adds the two promotion checks to a `review`
  session's drift scan; `on-reset` lands a design note that has generalized and records the
  landing in the session record, holding off on a handoff; `on-close` commits the landing where it
  landed, riding the closeout commit in the same repository or taking its own **Cross-repo**
  commit in the architecture repository.
- **Skill tiers mirroring marathon's** — `mechanics/pipeline.md` maps marathon's firing points to
  the extension's hook instructions (`on-start.md`, `on-reset.md`, `on-close.md`) and states how
  the layer resolves in each shape; `references/architecture-layer.md` is what the layer holds,
  what it refuses, how a repository cites it, the promotion sequence that fills it, the rule that
  a page states only current truth, and the bootstrap index.
