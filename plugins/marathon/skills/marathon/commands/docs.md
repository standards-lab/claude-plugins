# marathon docs

Author and curate human-oriented documentation in a top-level `docs/` directory. `docs/` is
optional: a project opts into it the first time you run this command, and many projects never
need it. Use `docs` when the project has reached a point where documentation a person reads is
worth writing and keeping.

Documentation is high-effort work that has to stay coherent as the project evolves, so it gets
its own deliberate pass rather than being folded into the build loop. `docs` runs the session
pipeline (`mechanics/pipeline.md`), recorded under `Session: docs`, and writes its pages in
stages under `references/staged-execution.md`: each stage is reported with the working tree
uncommitted and commits on approval. Branch slug: the pass.

## The two kinds of docs directory

A `docs/` directory holds one of two kinds of documentation, and which kind decides what the
command writes there. The rule and the lifecycle are `references/context-engineering.md`.

- **A repository's own documentation** is the accessibility layer over that repository's
  README, API documentation, and source: the guide a person reads to use or contribute to the
  repository, indexed by the README. It describes the repository's implementation, and the
  code is its source of truth.
- **A landing zone's documentation** is the architecture, its principles, and a catalog of the
  repositories that implement it. It holds only what generalizes past one repository: a
  principle, a convention, a definition. Nothing that a reader could infer from a repository's
  source belongs in it, and a page that restates a repository is a defect.

A standalone project's `docs/` is the first kind, and it may carry the second when the project
defines conventions of its own. In a workspace, the coordinator's `[workspace] docs` field
names the one project whose `docs/` is the landing zone, and every member project's `docs/` is
the first kind.

## Find which kind this run writes

Before bootstrapping or curating anything, work out which kind of directory this project holds.
Check whether this project is a standalone project or sits in a workspace (the LOCATE stage of
`mechanics/pipeline.md` makes the same check).

- **Standalone.** `docs/` is the repository's own documentation; continue below.
- **In a workspace, and `[workspace] docs` names this project.** `docs/` is the landing zone.
  Continue below under the landing-zone rule, and route anything implementation-shaped back to
  the repository it describes.
- **In a workspace, and `[workspace] docs` names a different project.** `docs/` is this
  repository's own documentation; continue below. Route anything that generalizes past this
  repository, a principle or a convention, to the landing zone through the promotion sequence:
  a concept here, a design note once it settles, and a landing-zone page in a `docs` session
  there.
- **In a workspace, and the coordinator declares no `[workspace] docs`.** No landing zone is
  settled. Ask the architect whether this project should be the one, and add the field to the
  coordinator's `marathon.toml` if so, or whether another project holds it.

## Bootstrapping (first run, no docs/)

If the project has no `docs/` yet, this run establishes it:

1. Settle with the architect what the documentation needs to cover and how to structure it:
   the topics a reader moves through, and where to start.
2. Create `docs/`, index it from the README, and write the first pages. A repository's own
   documentation is written from the code and the README as they stand; a landing zone's pages
   are written from the design notes that have settled and generalized.
3. Keep it shallow where the project is still moving; document what is settled enough to
   explain.

## Curating (later runs, docs/ exists)

If `docs/` already exists, this run extends and maintains it:

1. Settle what to add or rework: a topic the project grew, a page the code outpaced, structure
   that no longer reads well.
2. Write or revise those pages. In a repository's own documentation the code wins: when the
   code and a page disagree, the page is what is wrong. In a landing zone the rule wins: a page
   that has come to restate a repository's implementation is reduced to the principle it
   states or removed.
3. Don't pad it. Add the documentation a reader needs, not coverage for its own sake.

## Conclude

Close with `close`, which records the pass in the reset file under `Session: docs` and publishes
the branch. The `review` command keeps `docs/` honest between `docs` passes: in a repository's
own documentation it flags pages the code has moved out from under, and in a landing zone it
flags pages that restate a repository.
