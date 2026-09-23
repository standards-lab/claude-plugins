---
name: marathon-roadmap
user-invocable: false
description: >
  A marathon extension that keeps a roadmap manifest, context/roadmap.toml, as the one record of
  what remains on a project's or workspace's path to its target end state. The manifest holds goals,
  tasks, and a backlog in one TOML file, and everything else cites its entries by dotted path.
  marathon sessions keep it current through the extension hooks: the session loads the manifest's
  conventions at start, cites tasks in the reset file, and at closeout deletes the finished task and
  advances `next`, including its waves. Load this skill when a marathon session finds it enabled, or
  when the architect asks about the roadmap, roadmap.toml, goals and tasks, the backlog, or what
  comes next.
---

# marathon-roadmap

Version: 0.2.1

marathon plans one step at a time and keeps notes brief beyond it, so the path itself goes
unwritten: the outcomes the project is working toward and the tasks that remain. This extension
keeps that path in one manifest, `context/roadmap.toml`, which the sessions keep current as they
work. The manifest holds goals (named outcomes), tasks (one session's requirement each), and a
backlog of tasks that belong to no goal.

## Declaration

- **Artifact:** `context/roadmap.toml`. Enabled under `[workspace]` at a coordinator, one manifest
  at the coordinator serves every member project. Enabled under `[project]`, the project keeps its
  own.
- **Hooks:** `on-start`, `on-reset`, `on-close`.
- **Targets:** marathon 0.14.

## Mechanics

The hook map, and where the manifest is found, load with this skill:

@mechanics/pipeline.md

Each hook's instructions:

- [`mechanics/on-start.md`](./mechanics/on-start.md): load the conventions, and create a missing
  manifest.
- [`mechanics/on-reset.md`](./mechanics/on-reset.md): cite tasks in the session record.
- [`mechanics/on-close.md`](./mechanics/on-close.md): delete the finished task and advance `next`.

## References

- [`references/manifest.md`](./references/manifest.md): the `roadmap.toml` format. Read it before
  editing or creating the manifest.

Each project decides how to shape its goal tree.
