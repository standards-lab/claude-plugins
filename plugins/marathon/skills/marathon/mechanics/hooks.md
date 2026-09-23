# Hooks

How marathon finds enabled extensions and fires their hooks (`references/extensions.md`).

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
| `on-execute` | At 3 · SETTLE step 4, on approval, or at 3R · RESUME step 2, after checkout |
| `on-commit` | Immediately before each commit the session makes |
| `on-reset` | Before the reset file (`mechanics/reset-file.md`) is written, in both `reset` and `close` |
| `on-close` | In `close` only, after the reset file is written and before the closeout commit |

Order constraints:

- `close` fires `on-reset`, then `on-close`, then `on-commit`, and then publishes.
- `reset` writes a handoff and never fires `on-close`.
- `init` fires `on-start`, `on-execute`, `on-reset`, and `on-commit`, and never `on-close`.

An enabled extension whose artifact doesn't exist yet creates it at `on-start`.
