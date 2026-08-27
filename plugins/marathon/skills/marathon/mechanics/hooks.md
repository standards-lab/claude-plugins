# Hook execution

Firing spec for marathon's extension hooks. What an extension is, what it declares, and how a
repository enables one: `references/extensions.md`. Hooks fire only at the pipeline points named in
`mechanics/pipeline.md`.

## Resolution

Resolve the enabled extension set once, at 2 · START, before the first fire:

1. Read `[project] extensions` from the project's `.claude/marathon.toml`.
2. In a workspace, also read `[workspace] extensions` from the coordinator's `marathon.toml`.
3. Union the two lists, preserving order. For each name:
   - skill installed (its skill-listing description contains "marathon extension") → active;
   - not installed → report the missing plugin to the developer and continue without it.
4. For each active extension, read its SKILL.md declaration: the artifact it owns, the hook points
   it acts at, the marathon version it targets. On an incompatible version, report the mismatch and
   ask before applying.

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
- `coordinate`: fires nothing itself; each fanned-out project session fires its own.

## Artifact bootstrap

An enabled extension whose artifact does not exist yet: create the artifact at `on-start`, as the
extension's SKILL.md directs.

## Session-specific moments

Moments beyond the five universal hooks are named `<command>:<moment>` — for example
`start:guide-written`, or `close:published` for the point after publish when the change proposal's
URL is known. None are defined; one is added to this spec when an extension earns it.
