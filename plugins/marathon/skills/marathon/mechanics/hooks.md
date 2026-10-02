# Hooks

This file specifies how marathon finds enabled extensions and fires their hooks
(`references/extensions.md`).

## Finding the enabled extensions

Find the enabled extensions once, at 2 · START, before the first hook fires:

1. Read `[project] extensions` from the project's `.claude/marathon.toml`.
2. In a workspace, also read `[workspace] extensions` from the coordinator's `marathon.toml`.
3. Combine the two lists in order, without duplicates. An extension whose skill is installed is
   active. Tell the architect about any that isn't installed, and continue without it.
4. Read each active extension's SKILL.md declaration. If it targets an incompatible marathon
   version, tell the architect and ask before applying it.

When no extension is enabled, every hook does nothing.

## Firing a hook

To fire a hook, go through the active extensions in the order found above. For each one that
declares the hook, do what its SKILL.md says for it. A hook fires just before the moment it names.

| Hook | When it fires |
|------|---------------|
| `on-start` | At 2 · START, step 1, before the session reads its context |
| `on-build` | At 3 · PLAN, step 4, on approval, or at 3R · RESUME, step 2, after checkout |
| `on-ship` | At 6 · SHIP, step 1, once the architect accepts the session brief |
| `on-record` | Before the goal record is committed, at SHIP and at `reset` |

Order constraints:

- SHIP fires `on-ship`, then `on-record`, then publishes.
- `reset` fires `on-record` and never `on-ship`.
- `plan`, `experiment`, `intake`, and `retro` fire `on-start`. `plan` and `retro` also fire
  `on-record` before any commit that writes a goal record; `experiment` and `intake` write none.
- A sync fires `on-record` before it reads the record's pending edits.
- `init` fires `on-start` only.

An enabled extension whose artifact doesn't exist yet creates it at `on-start`. During a task, a
hook's change to an artifact in another repository goes in the goal record's pending edits, and
sync applies it.
