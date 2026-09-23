# marathon-architecture

A [marathon](../marathon/) extension that adds the architecture layer: the principles,
definitions, and conventions that apply beyond one repository. A standalone project keeps the
layer in a top-level `architecture/` directory. A workspace keeps it in one member repository,
named at the coordinator. marathon's sessions fill the layer through the extension hooks: the
session loads the layer's rules at start, lands and records a note that applies beyond its
repository before writing the session record, and commits the landed note at closeout.

This README is a quick reference. The skill's files under
[`skills/marathon-architecture/`](./skills/marathon-architecture/) define how the extension
behaves: `SKILL.md`, the `mechanics/` hook instructions, and the `references/` description of the
layer.

## Install

```bash
claude plugin marketplace add standards-lab/claude-plugins
claude plugin install marathon-architecture@standards-lab
```

## Enable

Installing makes the extension available. A repository enables it in `.claude/marathon.toml`:

```toml
[project]
extensions = ["marathon-architecture"]   # a top-level architecture/ directory in this repository
```

A workspace coordinator enables it for every member project, which share one layer:

```toml
[workspace]
extensions = ["marathon-architecture"]
```

The coordinator also keeps a second file next to it. The extension owns that file, and marathon
never reads or writes it:

```toml
# .claude/marathon-architecture.toml
repo = "architecture"   # the order key of the member repository that holds the layer
```

A workspace keeps one layer, so the extension is enabled at the coordinator. A member that enables
it on its own is reported as a mismatch and gets no layer of its own.

On a standalone project, the next marathon session creates the `architecture/` directory if it
doesn't exist, and adopts it unchanged if it does. In a workspace, if
`.claude/marathon-architecture.toml` doesn't exist or has no `repo` key, the session agrees the
layer's repository with the architect and writes the file. The architecture repository is a
marathon project of its own, set up with `marathon init` as a `context` project; the hook never
creates it. Requires marathon 0.14 or later.

## What the layer holds

Only knowledge that applies beyond one repository: a principle, a definition, or a convention.
Nothing a reader could learn from a repository's source, so a page that restates an
implementation is a defect. Pages arrive only by promotion: a note in the repository that owns the
knowledge, proven by validated work, becomes a page once the knowledge applies beyond that
repository.
