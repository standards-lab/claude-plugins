---
name: marathon-architecture
user-invocable: false
description: >
  A marathon extension that adds the architecture layer: the principles, definitions, and
  conventions that apply beyond one repository, written for a general reader and relied on from
  outside the project. A standalone project keeps the layer in a top-level architecture/ directory.
  A workspace keeps it in one member repository, which this extension's own
  .claude/marathon-architecture.toml at the coordinator names. marathon's sessions fill the layer by
  promoting validated context/ notes through the extension hooks: the session loads the layer's
  rules at start, lands and records a note that applies beyond its repository before writing the
  session record, and commits the landed note at closeout. Load this skill when a marathon session
  finds it enabled, or when the architect asks about the architecture layer, promoting a note beyond
  the repository that owns it, or what belongs in the architecture repository.
---

# marathon-architecture

Version: 0.2.1

marathon deletes a note once the built work expresses it. A principle, definition, or convention
that holds across repositories never reaches that point, because no single repository's code
expresses it. This extension gives that knowledge a home, the architecture layer, which sessions
fill by promotion through marathon's hooks. Enable it when a validated note applies beyond the
repository it lives in.

## Declaration

- **Artifact:** the architecture layer. Enabled under `[project]` on a standalone project, the
  layer is a top-level `architecture/` directory with a README as its index. Enabled under
  `[workspace]` at a coordinator, the layer is the one member repository named by `repo` in this
  extension's `.claude/marathon-architecture.toml` at the coordinator, and it serves every member
  project's sessions.
- **Hooks:** `on-start`, `on-reset`, `on-close`.
- **Targets:** marathon 0.15.

## Mechanics

The hook map, and where the layer is found, load with this skill:

@mechanics/pipeline.md

Each hook's instructions:

- [`mechanics/on-start.md`](./mechanics/on-start.md): load the rules, and create or adopt the
  layer.
- [`mechanics/on-reset.md`](./mechanics/on-reset.md): land a note that applies beyond its
  repository, and record it.
- [`mechanics/on-close.md`](./mechanics/on-close.md): commit the landed note.

## References

- [`references/architecture-layer.md`](./references/architecture-layer.md): what the layer holds
  and how knowledge reaches it. Read it before promoting a note or writing a page.

Each project decides how to organize its layer, including the hierarchy and the page format.
