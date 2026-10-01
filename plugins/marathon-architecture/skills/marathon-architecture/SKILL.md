---
name: marathon-architecture
user-invocable: false
description: >
  A marathon extension that adds the architecture layer: the principles, definitions, and
  conventions that apply beyond one repository, written for a general reader. A standalone project
  keeps the layer in a top-level architecture/ directory. A workspace keeps it in one member
  repository, which this extension's own .claude/marathon-architecture.toml at the coordinator
  names. Each repository's STANDARDS.md points to the pages that apply to it, and marathon's
  standards-reviewer reads those pages; this extension finds the layer at the start of every
  session so the pointers resolve. Load this skill when a marathon session finds it enabled, or
  when the architect asks about the architecture layer, what a STANDARDS.md should point to, or
  what belongs in the architecture repository.
---

# marathon-architecture

Version: 0.3.0

marathon deletes a note once the built work expresses it. A principle, definition, or convention
that holds across repositories never reaches that point, because no single repository's code
expresses it. This extension gives that knowledge a home, the architecture layer, and makes it the
reference each repository's `STANDARDS.md` points to, so the standards-reviewer applies it.

## Declaration

- **Artifact:** the architecture layer. Enabled under `[project]` on a standalone project, the
  layer is a top-level `architecture/` directory with a README as its index. Enabled under
  `[workspace]` at a coordinator, the layer is the one member repository named by `repo` in this
  extension's `.claude/marathon-architecture.toml` at the coordinator, and it serves every member.
- **Hooks:** `on-start`.
- **Targets:** marathon 0.16.

## Mechanics

The hook map, and where the layer is found, load with this skill:

@mechanics/pipeline.md

- [`mechanics/on-start.md`](./mechanics/on-start.md): find the layer, create or adopt it, and
  give the session its location.

## References

- [`references/architecture-layer.md`](./references/architecture-layer.md): what the layer holds,
  how a `STANDARDS.md` points to it, and how a page arrives. Read it before writing a page or a
  pointer.

Each project decides how to organize its layer, including the hierarchy and the page format.
