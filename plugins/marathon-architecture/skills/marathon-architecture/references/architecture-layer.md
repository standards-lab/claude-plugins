# The architecture layer

The architecture layer holds the principles, definitions, and conventions that apply beyond one
repository, written for a general reader. Its readers are the standards-reviewer, through each
repository's `STANDARDS.md`, and the people who build on the architecture.

## Where it lives

A standalone project keeps the layer in a top-level `architecture/` directory with a README as
its index. A workspace keeps one layer in one repository, named by `repo` in the coordinator's
`.claude/marathon-architecture.toml`. That repository is a `context` project whose directory tree
is the architecture, with a README in every directory as its index.

## What belongs in it

- Only knowledge that applies beyond one repository, and that code expresses or a check enforces
  somewhere. Nothing a reader could learn from a repository's source: a page that restates an
  implementation is a defect, reduced to its principle or removed. A principle no code follows
  yet is speculation, and stays a note until code proves it.
- Only what is true now: no changelog, no revision log, and no record of who wrote a page.
- No project documentation. `docs/` is one repository's guide.

## How a repository points to it

A repository's `STANDARDS.md` holds its judgement calls, one line each, and points to the layer's
pages that apply, by path: `architecture/principles/<page>.md` in a workspace, relative to the
workspace, or relative to the repository when standalone. A pointer may narrow the page with a
convention of the repository's own, stated beside it; it never restates the page. A rule a check
can enforce goes in the check, not in `STANDARDS.md` and not in a page.

## How a page arrives

A page arrives only from validated code. When the same principle is expressed in more than one
repository's merged code, `retro` proposes it as a finding, and the architect's tick makes it a
task: in a workspace, a task of a goal that locks the architecture repository; when standalone, a
task in the project. That task writes the page and adds the pointer to each `STANDARDS.md` that
should follow it.

## Starting index

A standalone project's `architecture/README.md`, created at `on-start`:

```markdown
# Architecture

The principles, definitions, and conventions that apply beyond this repository, written for a
general reader. Nothing a reader could learn from the source belongs here, and a page that
restates the implementation is a defect.

A page arrives from validated code, once the same principle holds in more than one place.
`STANDARDS.md` points to the pages that apply.
```
