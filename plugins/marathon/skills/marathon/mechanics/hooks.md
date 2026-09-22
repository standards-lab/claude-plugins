# Hook execution

Firing spec for marathon's extension hooks (`references/extensions.md`).

## Resolution

Resolve the enabled extension set once, at 2 · START, before the first fire:

1. Read `[project] extensions` from the project's `.claude/marathon.toml`.
2. In a workspace, also read `[workspace] extensions` from the coordinator's `marathon.toml`.
3. Union the two lists, preserving order. An installed name is active; report a missing one to the
   architect and continue without it.
4. Read each active extension's SKILL.md declaration, and report an incompatible version before
   applying it.

An empty set makes every fire a no-op.

## Firing

To fire a hook: for each active extension declaring that point, in resolution order, do what its
SKILL.md specifies there. A hook fires before the moment it names.

| Hook | Fires |
|------|-------|
| `on-start` | pipeline 2 · START, step 1, before orientation |
| `on-execute` | 3 · SETTLE step 4, on approval; or 3R · RESUME step 2, after checkout |
| `on-commit` | immediately before each commit the session makes |
| `on-reset` | before the reset file (`mechanics/reset-file.md`) is written — in `reset` and in `close` |
| `on-close` | in `close` only: after the reset file is written, before the closeout commit |

Ordering constraints:

- `close`: `on-reset` → `on-close` → `on-commit` → publish.
- `reset` (handoff): never fires `on-close`.
- `init`: fires `on-start`, `on-execute`, `on-reset`, `on-commit`; never `on-close`.

An enabled extension whose artifact doesn't exist yet creates it at `on-start`.
