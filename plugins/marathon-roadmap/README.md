# marathon-roadmap

A [marathon](../marathon/) extension that keeps a roadmap manifest — `context/roadmap.toml` — as
the single source of what remains on a project's or workspace's path to its target end state.
Goals, tasks, and a backlog nest in one TOML file; marathon's sessions keep it current through
the extension hooks: conventions layered in at session start, dotted task citations in the
session record, the finished task deleted and `next` advanced at closeout.

This README is a quick reference. The skill itself is the source of truth for how the extension
behaves; its files under [`skills/marathon-roadmap/`](./skills/marathon-roadmap/) — `SKILL.md`,
the `mechanics/` execution specs, and the `references/` manifest format — describe the full
detail.

## Install

```bash
claude plugin marketplace add standards-lab/claude-plugins
claude plugin install marathon-roadmap@standards-lab
```

## Enable

Installation makes the extension available; a repository enables it in `.claude/marathon.toml`:

```toml
[project]
extensions = ["marathon-roadmap"]     # this project keeps its own context/roadmap.toml
```

or, at a workspace coordinator, for every member project against one shared manifest:

```toml
[workspace]
extensions = ["marathon-roadmap"]     # one context/roadmap.toml at the coordinator
```

The next marathon session bootstraps the manifest if it doesn't exist yet. Requires marathon
0.7 or later.

## The manifest

Three primitives, nested like a filesystem: a **goal** is a named outcome made of other goals
and tasks, closing when its criteria hold; a **task** is one session's requirement; the
**backlog** holds the tasks assigned to no goal. The manifest is ephemeral — finished work is
deleted, only what remains stays — and `next` is the only sequence it asserts. References
elsewhere cite dotted slug paths: `v1.data.reads`, `backlog.docs-site`.
