# The architecture layer

The architecture layer holds the principles, definitions, and conventions that apply beyond one
repository, written for a general reader and relied on from outside the project. A validated note
moves here when the built work can't express it, because the knowledge isn't about the built
work.

## Where it lives

A standalone project keeps the layer in a top-level `architecture/` directory with a README as
its index. A workspace keeps one layer in one repository, named by `repo` in the coordinator's
`.claude/marathon-architecture.toml`. That repository is a `context` project whose directory tree
is the architecture, with a README in every directory as its index.

## What belongs in it

- Only knowledge that applies beyond one repository. Nothing a reader could learn from a
  repository's source. A page that restates an implementation is a defect: reduce it to its
  principle or remove it.
- A repository links the principles it follows from its README, and states next to each link any
  convention of its own that narrows the principle. It never restates the page. The layer may
  list the repositories that implement it, one line and a link each.
- Only what is true now, like every note: no changelog, no revision log, and no record of who
  promoted a page.
- No project documentation. `docs/` is one repository's guide, written by a documentation step.
  marathon's rule that deletes a note once the built work expresses it never applies to the
  layer, which changes only when a principle changes.

## How knowledge reaches it

Knowledge reaches the layer only by promotion. A note is proven once the work it describes is
built and validated, which a closeout's **Validated** entry records. If the built work or its
documentation then expresses the note, the note is deleted under marathon's rule. If the note
applies beyond its repository, it becomes a page. In a workspace, a member's `close` or `review`
lands the note in the architecture repository, recorded under **Cross-repo**, and that repository
writes the page in a session of its own.

## Starting index

A standalone project's `architecture/README.md`, created at `on-start`:

```markdown
# Architecture

The principles, definitions, and conventions that apply beyond this repository, written for a
general reader. Nothing a reader could learn from the source belongs here, and a page that
restates the implementation is a defect.

A page arrives by promotion: a note in `context/`, proven by validated work, becomes a page here
once the knowledge applies beyond this repository.
```
