# Extension pipeline

Execution spec for marathon-roadmap. marathon fires five universal hooks at the points named in
its own `mechanics/hooks.md`; this extension connects to the three below. At each fire, act from
the mechanics file the table names. Nothing here runs outside a marathon session.

| marathon hook | This extension | Act from |
|---------------|----------------|----------|
| `on-start` | layer the roadmap conventions into the session; bootstrap a missing manifest | `mechanics/on-start.md` |
| `on-reset` | cite roadmap tasks in the session record | `mechanics/on-reset.md` |
| `on-close` | delete the finished task and advance `next` | `mechanics/on-close.md` |

## Locating the manifest

Every hook acts on one manifest, resolved once at `on-start` from where the extension is
enabled:

- `[workspace] extensions` at the coordinator → the coordinator's `context/roadmap.toml`, for
  every member project's sessions.
- `[project] extensions` → the project's own `context/roadmap.toml`.

A session in a member project therefore edits a manifest in another repository — the
coordinator's. How those edits commit is each hook's own concern; `on-close.md` states the rule
where it matters.
