# The architecture layer

The principles, the definitions, and the conventions that have generalized past one repository,
written for a general reader and relied on from outside the project. It is the top of marathon's
context lifecycle: where a settled note goes when the built work cannot express it, because the
knowledge is not about the built work.

## Where it lives

A standalone project keeps it as a top-level `architecture/` directory with a README index. A
workspace keeps one layer as one repository, named by `repo` in the coordinator's
`.claude/marathon-architecture.toml`: a `context` project whose tree is the architecture, with a
README in every directory as its index.

## What belongs in it

- Only what generalizes. Nothing a reader could infer from a repository's source; a page that
  restates an implementation is a defect, reduced to its principle or removed.
- A repository links the principles it follows from its README and states beside the link any
  convention of its own that narrows one; it never restates the page. The layer may catalog the
  repositories that implement it, one line and a link each.
- Current truth only, like every note: no changelog, revision log, or record of who promoted it.
- Not project documentation: `docs/` is one repository's guide, written by a documentation step.
  marathon's decay rule never applies to the layer, which changes only when a principle does.

## How knowledge reaches it

By promotion only. A note proves out once the work it describes is built and validated, which
a closeout's **Validated** records. If the built work or its documentation then expresses it, it decays under marathon's rule; if it has generalized past the
repository, it becomes a page. In a workspace, a member's `close` or `review` lands it as a
note in the architecture repository, recorded under **Cross-repo**, and that repository
authors the page in a session of its own.

## Bootstrap index

A standalone project's `architecture/README.md`, created at `on-start`:

```markdown
# Architecture

The principles, definitions, and conventions that have generalized past this repository, written
for a general reader. Nothing a reader could infer from the source belongs here, and a page that
restates the implementation is a defect.

A page arrives by promotion: a note in `context/`, proven by validated work, becomes a page here
once the knowledge has generalized past this repository.
```
