# Roadmap waves: parallel-safe sequencing in `next`

Captured 2026-09-22 from a workflow-refinement planning session at the coordinator
(`standards-lab`). Isolated experiments (`isolated-experiments.md`) let independent roadmap goals
run in true parallel; the roadmap manifest needs a way to express which goals may run together and
which must wait, and a way to track more than one effort in flight at once. Scoped to the
`marathon-roadmap` extension, not core marathon. This concept is settled direction for the
extension's own implementing session.

## The gap

`context/roadmap.toml`'s `next` is a single flat sequence today, read and advanced by
`marathon-roadmap`'s hooks as one item at a time. That models exactly one effort in flight, which
matched a workspace with exactly one session possible at a time — no longer true once an isolated
experiment can proceed independently of the coordinator's own session.

## Design

`next` becomes a list of waves, in the same shape `[workspace] order` already uses in
`standards-lab/.claude/marathon.toml` — a string, or an array of concurrent peers. No new idiom,
and a flat list still reads correctly as a sequence of one-entry waves, so nothing already in
`next` needs to change shape the day this lands:

```toml
next = [
  "blobfs.sources",
  ["blobfs.build", "v1.messaging.experiment", "v1.ai.experiment"],
  ["blobfs.admin", "v1.storage.service"],
  "v1.storage.suite",
  # ...
]

active = ["v1.messaging.experiment"]  # tasks in flight outside the coordinator's own session
```

The rule for sharing a wave — no dependency between the tasks, no repository the tasks share — is
the architect's judgment, narrated in the header comment the way sequencing decisions already are.
It is not a computed or validated field: the workspace already resequences by hand, and a
scheduler would fight that. A deterministic check that two tasks in one wave share no repository
is a candidate for `v1.harness.tasks.tooling`, later, not a hook here.

Hook changes:

- `on-start` reads the current wave as the first non-empty `next` entry, and reports which of its
  members are already `active`.
- `on-reset` is unchanged — the coordinator's own reset keeps a single Next-focus, since it's
  still the pointer for the coordinator's one session; the multi-effort view lives in `active` and
  the current wave, not in the reset schema.
- `on-close` deletes the finished task from its wave and from `active`, drops a wave once it's
  empty, and adds a task to `active` when an isolated experiment's grain starts one
  (`isolated-experiments.md`).

## Files this touches

`marathon-roadmap`'s `references/manifest.md`, its hook files, `SKILL.md`. Version bump to 0.2.0.

## Open questions for the implementing session

- Whether `active` needs a status beyond presence (e.g., which repository/branch it's on) or
  whether that already lives in that effort's own `context/reset.md`, one lookup away via the
  references catalog.
- Whether a wave that goes stale (one member finishes, the others haven't, and nothing is actually
  blocked) needs any manifest signal, or whether that's simply visible in `active` shrinking.
