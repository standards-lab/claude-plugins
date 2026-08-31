# marathon-references: the references catalog as an extension

Captured 2026-08-31 during the workspace-sweep session, after establishing that the references
system is a workspace convention with no marathon footprint. This concept proposes evaluating
its promotion to a marathon extension, alongside `concepts/marathon-docs-extension.md` and
`concepts/marathon-sitrep.md`. Everything here is candidate direction for a future session on
this repository.

## The gap

The reference-architecture workspace maintains a references system at its coordinator:
`references.toml` (portable identity — one key per repository, mapping to its canonical remote),
`references.local.toml` (gitignored, the same keys mapping to local checkouts on this machine),
and `references.md` (the described catalog: effort repositories, external references, prior
R&D). The convention is defined solely by the workspace's own design note
(`standards-lab/context/design/repo-references.md`); the marathon skill never mentions it.

The shape is general. Any marathon workspace draws on repositories beyond its members — prior
R&D to salvage from, external illustrations, sibling efforts — and wants portable identity plus
a per-machine location map so the workspace reconstructs anywhere. That is the same
schema-driven-artifact shape the roadmap manifest already has.

## Proposal

`marathon-references`, an extension owning the three-file references artifact, resolved the way
the roadmap manifest is: at the coordinator for a workspace, at the project itself standalone.
The extension codifies the file contract — the key join, the never-duplicate-locations rule,
committed/gitignored split — and stays non-prescriptive about what a workspace catalogs.

## Open questions

- Facet: integration, enhancement, or hybrid (the taxonomy in `concepts/marathon-sitrep.md`).
  The catalog may need no hooks at all — sessions read it for orientation but never advance it —
  which would make it a pure artifact convention; whether the extension contract even has a
  shape for that is the interesting question.
- Whether `on-start` should layer a convention in (cite catalog keys in prose the way roadmap
  tasks are cited by dotted path), or nothing fires and the artifact is consulted on demand.
- Artifact location: the workspace keeps the files at the coordinator's repository root; an
  extension would have to either adopt that or standardize `context/`, and the migration story
  for the existing workspace follows from that choice.
- Whether `references.md` stays hand-authored or becomes a projection of the TOML plus described
  entries — the source-of-truth rule pulls toward one authored home.
