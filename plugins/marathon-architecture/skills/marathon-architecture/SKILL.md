---
name: marathon-architecture
user-invocable: false
description: >
  A marathon extension that adds the architecture layer: the principles, definitions, and
  conventions that have generalized past one repository, written for a general reader and
  relied on from outside the project. A standalone project keeps the layer as a top-level
  architecture/ directory; a workspace keeps it as the one member repository this extension's
  own .claude/marathon-architecture.toml names, at the coordinator. marathon's sessions fill it
  by promotion from settled context/ notes through the extension hooks — the conventions layered
  in at session start, a generalized note landed and recorded before the session record is
  written, the landing committed at closeout. Load this skill when a marathon session resolves it as an
  enabled extension, or when the architect asks about the architecture layer, promoting a note
  past the repository that owns it, or what belongs in the architecture repository.
---

# marathon-architecture

Version: 0.2.0

marathon deletes a note once the built work expresses it. A principle, definition, or convention
that holds across repositories never reaches that point, because no one repository's code
expresses it. This extension gives that knowledge a home, the architecture layer, filled by
promotion through marathon's session hooks. Enable it when a settled note has outgrown the
repository it lives in.

## Declaration

- **Artifact:** the architecture layer. Enabled under `[project]` on a standalone project, it is
  a top-level `architecture/` directory with a README as its index. Enabled under `[workspace]`
  at a coordinator, it is the one member repository named by `repo` in this extension's own
  `.claude/marathon-architecture.toml` at the coordinator, serving every member project's
  sessions.
- **Hooks:** `on-start`, `on-reset`, `on-close`.
- **Targets:** marathon 0.13.

## Mechanics

Loaded with this skill: the hook map and where the layer resolves.

@mechanics/pipeline.md

The hook instructions the pipeline acts from:

- [`mechanics/on-start.md`](./mechanics/on-start.md) — the conventions; bootstrap or adopt.
- [`mechanics/on-reset.md`](./mechanics/on-reset.md) — land a generalized note and record it.
- [`mechanics/on-close.md`](./mechanics/on-close.md) — commit the landing.

## References

- [`references/architecture-layer.md`](./references/architecture-layer.md) — what the layer
  holds and how it fills. Load it before promoting a note or authoring a page.

How a project organizes the layer, its hierarchy and page format, is its own.
