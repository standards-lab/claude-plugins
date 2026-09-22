# Flattening `context/design` and `context/concepts` into one tier

Captured 2026-09-22 from a workflow-refinement planning session at the coordinator
(`standards-lab`). The architect questioned whether the two-tier context model still earns its
complexity and called for collapsing it: no `concepts/` or `design/` subdirectory at all, a note
is a file directly under `context/`. This concept is settled direction for the `start` session
that implements it.

## The gap

`references/context-engineering.md` frames `design/` as validated intent the code can't express
yet, promoted from `concepts/` once something real exercises the shape. In practice the two
directories don't separate settled from unsettled — the notes already carry their own settledness
in their own prose regardless of which directory holds them
(`go-web-service/context/concepts/data-layer.md`: "candidate until a task's session settles it";
`design/domain-architecture.md`: "settled in the `organization-reads` session") — and `design/`
has drifted into holding restatements of already-built interfaces that the decay rule's own
qualifier ("duplication in API documentation alone is not decay") was protecting from cleanup.
`design/storage-strategy.md`'s section 2, for instance, reproduces `go-storage`'s interface
verbatim for a library already released at v0.1.0. `context-architecture.md` and
`architecture-layer.md` already point the other way: a repository's design reasoning belongs in
its own README and documentation, each repository documents itself. A permanent `design/` tier at
the coordinator for material about member libraries contradicts that.

## Design

- `context-engineering.md`'s "How `context/` is organized" collapses to: Stable —
  `context/README.md`. Volatile — `context/*.md` (notes, flat, no subdirectory) and
  `context/reset.md`. "Deciding where something goes" collapses to: orientation → README; anything
  the built work and its own documentation don't yet express → a note; already expressed → delete
  it.
- **Decay inverts to match.** A note is deleted once the built work *or its own documentation*
  expresses what it described. The reasoning that used to justify a permanent `design/` note — a
  cross-repo strategy record's rationale, for example — moves into the *owning* repository's
  documentation (`doc.go`, README, `docs/`) once that capability is built, written by the step
  that builds it, and the note is deleted in the same close.
- **The ledger simplifies.** Core keeps **Integrated** (absorbed by built work or its docs),
  **Culled**, **Retained** (kept with a stated reason — legitimate when the owning repository's
  `docs/` doesn't exist yet, since `v1.repository-docs` is sequenced last), and **Cross-repo**.
  **Promote** leaves core entirely; the one remaining relocation — a note that has generalized
  past its repository — becomes `marathon-architecture`'s own vocabulary, unchanged in trigger
  (something real exercised the shape and it holds beyond this repository), sourced from
  `context/` directly instead of from `design/`.
- A note states its own settledness in its opening lines in place of the directory signal it
  loses — this is already how the strongest existing notes read; the change makes it a rule
  instead of a habit.

## Files this touches

Core: `references/context-engineering.md` (rewritten as above), `mechanics/reset-file.md` (schema
example, ledger vocabulary), `mechanics/pipeline.md` (START step 2's orientation read),
`commands/init.md` (scaffold `context/README.md` + `context/reset.md` only; drop the
settled-vs-concept founding decision), `commands/plan.md` and `commands/close.md` (Execute list:
Promote/Cull/Add-or-sharpen → Cull/Add-or-sharpen/Integrate), `commands/review.md` ("concepts that
have quietly proven out and belong in `design/`" → "notes the built work now expresses").

`marathon-architecture` (version bump to 0.2.0, targeting marathon 0.14): `SKILL.md`'s three-state
description, `mechanics/on-reset.md`, `references/architecture-layer.md` ("How knowledge reaches
it"), its bootstrap index text — the promotion source moves from `design/` to `context/`.

## Migration, once this lands

Only two repositories have a `design/` directory: `standards-lab` (15 notes) and `go-web-service`
(4 notes). Each is a `review` session on that repository, dispositioning every note into one of
four outcomes: moved to flat `context/` unchanged with a status line added; split (the reasoning
that describes built work moves into the owning repository's own docs as a Cross-repo edit, the
rest stays a note); promoted to the architecture repository, then deleted; or collapsed entirely
into the repository's own `CLAUDE.md`/README where it's already restated there. `standards-lab`'s
15 notes likely need two sessions given the Cross-repo edits several require into `go-storage` and
`go-observability`.

## Open questions for the implementing session

- Whether `marathon-architecture`'s promotion trigger needs any wording change beyond the source
  directory, now that the note it reads carries its own settledness statement rather than living
  in a directory reserved for settled material.
- Whether `init`'s dropped founding decision leaves a gap worth replacing with something lighter
  (e.g., nothing — a fresh project just starts writing notes and orientation).
