---
name: marathon-architecture
user-invocable: false
description: >
  A marathon extension that adds the architecture layer: the principles, definitions, and
  conventions that have generalized past one repository, written for a general reader and
  relied on from outside the project. A standalone project keeps the layer as a top-level
  architecture/ directory; a workspace keeps it as the one member repository this extension's
  own .claude/marathon-architecture.toml names, at the coordinator. marathon's sessions fill it
  by promotion from design/ through the extension hooks — the conventions layered in at session
  start, a generalized design note landed and recorded before the session record is written, the
  landing committed at closeout. Load this skill when a marathon session resolves it as an
  enabled extension, or when the architect asks about the architecture layer, promoting a note
  past the repository that owns it, or what belongs in the architecture repository.
---

# marathon-architecture

Version: 0.1.0

marathon keeps a project's written context decaying toward the built work: a concept proves out
and becomes a design note, and the design note goes once the code expresses it. Some knowledge
never reaches that end, because it is not about this repository. A principle, a definition, or a
convention that holds across repositories has no code to decay into. This extension gives that
knowledge a home, the architecture layer, and binds the promotion into it to marathon's session
hooks, so the layer fills deliberately instead of by accumulation.

The layer is the top of the context lifecycle and the last stop on the promotion path:

- **concept** — an idea in the repository that owns it.
- **design note** — the same idea once it settles, still about that repository.
- **architecture page** — the idea once it has generalized past that repository, written for a
  general reader.

A project that never produces the third does not need this extension. Enable it when a design
note has outgrown the repository it lives in and needs somewhere more general to go.

## Declaration

- **Artifact:** the architecture layer. Enabled under `[project]` on a standalone project, it is
  a top-level `architecture/` directory with a README as its index. Enabled under `[workspace]`
  at a coordinator, it is the one member repository named by `repo` in this extension's own
  `.claude/marathon-architecture.toml` at the coordinator, serving every member project's
  sessions.
- **Hooks:** `on-start`, `on-reset`, `on-close`.
- **Targets:** marathon 0.12.

## Mechanics

The execution layer, loaded with this skill: the map from marathon's firing points to this
extension's hooks, and where the layer resolves.

@mechanics/pipeline.md

The hook instructions the pipeline acts from:

- [`mechanics/on-start.md`](./mechanics/on-start.md) — layer the conventions in; bootstrap a
  missing layer, or adopt the one that already exists.
- [`mechanics/on-reset.md`](./mechanics/on-reset.md) — land a design note that has generalized,
  and record the landing in the session record.
- [`mechanics/on-close.md`](./mechanics/on-close.md) — commit the landing where it landed.

## References

- [`references/architecture-layer.md`](./references/architecture-layer.md) — what the layer
  holds, what it refuses, how a repository cites it, the promotion sequence that fills it, and
  the bootstrap index. Load it before promoting a note or authoring a page.

## Scope

The extension codifies where the layer lives, what belongs in it, and how knowledge reaches it,
and stays non-prescriptive past them. How a project organizes the layer — its hierarchy, its
page format, its front matter — is its own.
