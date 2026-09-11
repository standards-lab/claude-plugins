# The architecture layer

The principles, the definitions, and the conventions that have generalized past one repository,
written for a general reader and relied on from outside the project. It is the top of marathon's
context lifecycle: where a design note goes when the built work cannot express it, because the
knowledge is not about the built work.

## Where it lives

- A **standalone project** keeps the layer as a top-level `architecture/` directory, with a
  README as its index.
- A **workspace** keeps the layer as one repository, named by the `repo` key in this extension's
  own `.claude/marathon-architecture.toml` at the coordinator — this extension's config, never
  marathon core's. That repository is a `context` project whose whole tree is the architecture,
  with a README at its root and in every directory as the index GitHub renders, and it runs
  sessions like any other project.

One workspace keeps one layer. The extension is enabled at the coordinator, under
`[workspace] extensions`, so every member session promotes into the same repository.

## What belongs in it

The layer holds only what generalizes. Nothing a reader could infer from a repository's source
belongs in it, so it stays stable while the repositories change beneath it, and a page that
restates a repository's implementation is a defect, reduced to the principle it states or
removed.

A repository links the layer's principles from its README and states beside the link any
convention of its own that narrows a principle; it never restates the page. The layer may
catalog the repositories that implement it, with a description and a link for each, and goes no
deeper.

## States current truth only

A page holds what is true now, and nothing about how it got that way: no changelog entry, no
dated revision log, no note of the session or repository that promoted it. That history belongs
to the promoting session's own reset file, never the page — the same rule marathon's core applies
to `design/` and `concepts/` notes (`references/context-engineering.md`).

## How knowledge reaches it

By promotion, and only by promotion. A concept proves out and is promoted into `design/` in the
repository that owns it. A design note is then expressed either by the built work, at which point
it decays under marathon's rule, or by an architecture page, once the knowledge has generalized
past that one repository.

In a workspace the second case is a cross-repository step: a member's `close` or `review` finds
that a design note has generalized and lands it as a concept in the architecture repository,
recorded under **Cross-repo**, and the architecture repository authors the page in a session of
its own. A design note that describes one repository's implementation is expressed by the code
and the repository's own documentation, never by the architecture.

## What the layer is not

The layer is not the project's own documentation. `docs/` is the guide a person reads to use or
contribute to one repository; the code is its source of truth, and a documentation step of a
`start` session writes it. What generalizes past the repository reaches the layer by promotion
from `design/`, never by a documentation step.

marathon's decay rule applies to `context/` and never to the layer. `context/` is agent-oriented
and decays toward the built work; the architecture states what generalizes and changes only when
a principle does.

## Bootstrap index

A standalone project's `architecture/README.md`, created at `on-start` when the directory does
not exist yet. It states what the layer holds and how it fills, and nothing else; pages arrive by
promotion.

```markdown
# Architecture

The principles, definitions, and conventions that have generalized past this repository, written
for a general reader. Nothing a reader could infer from the source belongs here, and a page that
restates the implementation is a defect.

A page arrives by promotion: a concept in `context/concepts/`, a design note in `context/design/`
once it settles, and a page here once the design has generalized past this repository.
```
