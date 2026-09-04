---
name: marathon-roadmap
user-invocable: false
description: >
  A marathon extension that keeps a roadmap manifest — context/roadmap.toml — as the single
  source of what remains on a project's or workspace's path to its target end state: goals,
  tasks, and a backlog in one TOML file, cited by dotted slug paths. marathon sessions keep it
  current through the extension hooks — conventions layered in at session start, task citations
  in the reset file, the finished task deleted and `next` advanced at closeout. Load this skill
  when a marathon session resolves it as an enabled extension, or when the architect asks about
  the roadmap, roadmap.toml, goals and tasks, the backlog, or what comes next.
---

# marathon-roadmap

Version: 0.1.4

marathon plans one step at a time and deliberately keeps `context/` shallow past the step in
front. What that leaves unwritten is the path: the outcomes the project is working toward and the
tasks that remain on the way. This extension carries that path in one manifest,
`context/roadmap.toml`, and binds its upkeep to marathon's session hooks, so the manifest stays
current as a side effect of the sessions themselves.

The manifest holds three primitives, nested like a filesystem:

- **goal** — a named outcome made of other goals and tasks; it closes when its criteria hold.
- **task** — one session's requirement: what it is, the repositories it touches, its proof.
- **backlog** — the tasks assigned to no goal.

The roadmap is ephemeral, unordered, and proximate. It captures only what remains — a finished
task is deleted, the session record keeping the disposition. The tree asserts no order; the root
`next` list is the only sequence. Resolution matches proximity: the task in front carries detail,
everything else stays at claim resolution, and depth lives in linked context files rather than
here.

## Declaration

- **Artifact:** `context/roadmap.toml`. Enabled under `[workspace]` at a coordinator, one
  manifest at the coordinator serves every member project; enabled under `[project]`, the
  project keeps its own.
- **Hooks:** `on-start`, `on-reset`, `on-close`.
- **Targets:** marathon 0.10.

## Mechanics

The execution layer, loaded with this skill: the map from marathon's firing points to this
extension's hooks, and where the manifest resolves.

@mechanics/pipeline.md

The hook instructions the pipeline acts from:

- [`mechanics/on-start.md`](./mechanics/on-start.md) — layer the conventions in; bootstrap a
  missing manifest.
- [`mechanics/on-reset.md`](./mechanics/on-reset.md) — cite roadmap tasks in the session record.
- [`mechanics/on-close.md`](./mechanics/on-close.md) — delete the finished task; advance `next`.

## References

- [`references/manifest.md`](./references/manifest.md) — the `roadmap.toml` format: structure,
  fields, dotted citations, the bootstrap header, and a worked example. Load it before editing or
  creating the manifest.

## Scope

The extension codifies the manifest contract and the hook behavior, and stays non-prescriptive
past them. How a project shapes its goal tree — what its root goals are, how deep they nest, what
sits in the backlog — is its own.
