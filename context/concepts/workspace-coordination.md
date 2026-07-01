# Workspace coordination

marathon manages one repository. As the reference architecture spans several — a harness, libraries, a
service template, a service — a single change often crosses repositories in dependency order, as this
effort's consolidation session did (the org context first, then the library level). marathon should support
running from a workspace: a directory that contains several marathon projects and coordinates a change
across them.

## The workspace holds no context

The estate-level knowledge already has a home: the organization's context repository (its references
catalog and capability organization) is the one place the whole is described. A workspace `context/` would
be a second place describing the estate — a duplication that violates marathon's single-source-of-truth
rule. So the workspace layer is contextless by design, and the coordination knowledge splits in two,
neither part a workspace file:

- **Mechanism (general, in the skill).** How to detect a workspace, enumerate the nested marathon projects,
  and run a coordination session that fans out to per-repo sessions in dependency order — each a normal
  marathon session with its own branch, `reset.md`, and PR. The coordination session itself produces no
  branch, no reset, and no context; it is an orchestration wrapper.
- **Map (effort-specific, in the org context repo).** Which repositories exist, how they depend, and which
  one coordinates. marathon does not hardcode this; the agent reads the effort's org context to learn the
  order.

## Detection and degradation

A directory that contains marathon projects but is not itself one (no `context/`) is a workspace. The
coordinator repository self-declares (for example, a field in its `marathon.toml`), and its references
catalog is the map. With no coordinator present, marathon degrades gracefully: it enumerates the projects
and asks the developer for the order rather than assuming one.

## Continuity stays per-repo

A coordinated change checkpoints in each repository's `reset.md` independently. An interrupted fan-out
resumes per repository via each `reset.md` Status; the cross-repository order on resume comes from the map,
not a workspace tracker. This is why no workspace state tier is needed.

## Open questions

- How the coordinator declares itself, and whether a workspace with no coordinator needs any marker at all.
- Whether a coordination session leaves any record, and if so where — each repo's reset, or nowhere.
- How this interacts with the extension hooks when a coordinated change touches several trackers.
- Awareness direction: the mechanism must respect that a lower layer never names its consumers, so
  coordination reads the map from the org context, never by teaching lower repos about higher ones.
