---
name: marathon-roadmap
user-invocable: false
description: >
  A marathon extension that keeps a roadmap manifest — context/roadmap.toml — as the single
  source of what remains on a project's or workspace's path to its target end state: goals,
  tasks, and a backlog in one TOML file, cited by dotted slug paths. marathon sessions keep it
  current through the extension hooks — conventions layered in at session start, task citations
  in the reset file, the finished task deleted and `next`'s waves advanced at closeout. Load this skill
  when a marathon session resolves it as an enabled extension, or when the architect asks about
  the roadmap, roadmap.toml, goals and tasks, the backlog, or what comes next.
---

# marathon-roadmap

Version: 0.1.6

marathon plans one step at a time and keeps notes shallow past it, which leaves the path unwritten:
the outcomes the project is working toward and the tasks that remain. This extension keeps that
path in one manifest, `context/roadmap.toml`, current as a side effect of the sessions: goals
(named outcomes), tasks (one session's requirement each), and a backlog of tasks under no goal.

## Declaration

- **Artifact:** `context/roadmap.toml`. Enabled under `[workspace]` at a coordinator, one
  manifest at the coordinator serves every member project; enabled under `[project]`, the
  project keeps its own.
- **Hooks:** `on-start`, `on-reset`, `on-close`.
- **Targets:** marathon 0.13.

## Mechanics

Loaded with this skill: the hook map and where the manifest resolves.

@mechanics/pipeline.md

- [`mechanics/on-start.md`](./mechanics/on-start.md) — the conventions; bootstrap.
- [`mechanics/on-reset.md`](./mechanics/on-reset.md) — cite tasks in the session record.
- [`mechanics/on-close.md`](./mechanics/on-close.md) — delete the finished task; advance `next`.

## References

- [`references/manifest.md`](./references/manifest.md) — the `roadmap.toml` format. Load it
  before editing or creating the manifest.

How a project shapes its goal tree is its own.
