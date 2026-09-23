# marathon-roadmap

A [marathon](../marathon/) extension that keeps a roadmap manifest, `context/roadmap.toml`, as the
one record of what remains on a project's or workspace's path to its target end state. Goals, tasks,
and a backlog nest in one TOML file. marathon's sessions keep it current through the extension
hooks: the session loads the manifest's conventions at start, cites tasks by dotted path in the
session record, and at closeout deletes the finished task and advances `next`.

This README is a quick reference. The skill's files under
[`skills/marathon-roadmap/`](./skills/marathon-roadmap/) define how the extension behaves:
`SKILL.md`, the `mechanics/` hook instructions, and the `references/` manifest format.

## Install

```bash
claude plugin marketplace add standards-lab/claude-plugins
claude plugin install marathon-roadmap@standards-lab
```

## Enable

Installing makes the extension available. A repository enables it in `.claude/marathon.toml`:

```toml
[project]
extensions = ["marathon-roadmap"]     # this project keeps its own context/roadmap.toml
```

A workspace coordinator enables it for every member project, and the members share one manifest:

```toml
[workspace]
extensions = ["marathon-roadmap"]     # one context/roadmap.toml at the coordinator
```

The next marathon session creates the manifest if it doesn't exist. Requires marathon 0.14 or
later.

## The manifest

The manifest has three kinds of entry, nested like directories:

- A **goal** is a named outcome made of other goals and tasks. It closes when its criteria hold.
- A **task** is one session's work.
- The **backlog** holds the tasks that belong to no goal.

The manifest holds only what remains: finished work is deleted. `next` is the only entry that
holds a sequence, including waves of lanes that run at the same time. Everything else cites an
entry by its dotted path, such as `v1.data.reads` or `backlog.docs-site`.
