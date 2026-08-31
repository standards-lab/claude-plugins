# Changelog

All notable changes to the marathon-roadmap plugin are documented here. Versions follow
[Semantic Versioning](https://semver.org/spec/v2.0.0.html); dates and release links live on the
GitHub releases the tags cut.

## v0.1.3

### Changed

- **Targets marathon 0.9** — marathon's staged execution and cross-repo steps leave the hook
  contract unchanged: a cross-repo step is one session with one close, so `on-close` fires once
  per finished task as before. The SKILL.md gains the `Version:` line marathon 0.9 makes
  checkable, and the prose follows marathon 0.9's terminology: the human is the architect.

## v0.1.2

### Changed

- **SKILL.md is an index** — the entry file follows marathon 0.8's index layout: front matter, the
  description, and an index into the sub-layers, with the hook map (`mechanics/pipeline.md`)
  `@`-loaded and the hook instructions and manifest reference linked.
- **Targets marathon 0.8** — the reset file the `on-reset` hook shapes now resolves per marathon's
  `mechanics/reset-file.md`: a standalone project's own, or the workspace's single reset at the
  coordinator.

## v0.1.1

### Changed

- **The skill is passive** — `user-invocable: false` in the SKILL.md frontmatter keeps it out of
  the slash-command menu. An extension skill is resolved by marathon's hook firing, or loaded by
  the model on a roadmap question; it is never typed as a command.

## v0.1.0

### Added

- **The roadmap manifest convention as a marathon extension**, on the marathon 0.7 extension
  contract. One `context/roadmap.toml` holds goals, tasks, and a backlog — ephemeral, unordered
  past its `next` list, detailed only in front — at the workspace coordinator or on a standalone
  project, per where `marathon.toml` enables the extension.
- **Hook connections** — `on-start` layers the conventions into the session and bootstraps a
  missing manifest; `on-reset` cites tasks by dotted slug path in the session record; `on-close`
  deletes the finished task, closes emptied goals, and advances `next`, landing inside the
  closeout commit (or as its own coordinator commit from a member project's session).
- **Skill tiers mirroring marathon's** — `mechanics/pipeline.md` maps marathon's firing points
  to the extension's hook instructions (`on-start.md`, `on-reset.md`, `on-close.md`);
  `references/manifest.md` is the full manifest format: structure, fields, dotted citations,
  lifecycle, bootstrap header, and a worked example.
