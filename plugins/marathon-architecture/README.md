# marathon-architecture

A [marathon](../marathon/) extension that adds the architecture layer: the principles,
definitions, and conventions that have generalized past one repository. A standalone project
keeps the layer as a top-level `architecture/` directory; a workspace keeps it as the one member
repository the coordinator names. marathon's sessions fill it through the extension hooks — the
conventions layered in at session start, a generalized settled note landed and recorded before the
session record is written, the landing committed at closeout.

This README is a quick reference. The skill itself is the source of truth for how the extension
behaves; its files under [`skills/marathon-architecture/`](./skills/marathon-architecture/) —
`SKILL.md`, the `mechanics/` execution specs, and the `references/` description of the layer —
describe the full detail.

## Install

```bash
claude plugin marketplace add standards-lab/claude-plugins
claude plugin install marathon-architecture@standards-lab
```

## Enable

Installation makes the extension available; a repository enables it in `.claude/marathon.toml`:

```toml
[project]
extensions = ["marathon-architecture"]   # a top-level architecture/ directory here
```

or, at a workspace coordinator, for every member project against one shared layer:

```toml
[workspace]
extensions = ["marathon-architecture"]
```

and a second file beside it, owned entirely by this extension — never read or written by
marathon core:

```toml
# .claude/marathon-architecture.toml
repo = "architecture"   # the order key of the member repository holding the layer
```

A workspace keeps one layer, so the extension belongs at the coordinator; a member that enables
it alone is reported as a mismatch rather than given a layer of its own.

The next marathon session creates a standalone project's `architecture/` directory if it does not
exist, and adopts one that does, unchanged. In a workspace, if
`.claude/marathon-architecture.toml` doesn't exist yet or names no `repo`, the session settles it
with the architect and writes the file; the architecture repository is itself a marathon project,
initialized by `marathon init` as a `context` project — the hook never creates it. Requires
marathon 0.13 or later.

## What the layer holds

Only what generalizes: a principle, a definition, a convention. Nothing a reader could infer from
a repository's source, so a page that restates an implementation is a defect. Pages arrive by
promotion and only by promotion — a note in the repository that owns the knowledge, proven by
validated work, becomes a page once the knowledge has generalized past that repository.
