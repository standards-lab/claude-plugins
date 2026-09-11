# Extension pipeline

Execution spec for marathon-architecture. marathon fires five universal hooks at the points
named in its own `mechanics/hooks.md`; this extension connects to the three below. At each fire,
act from the mechanics file the table names. Nothing here runs outside a marathon session.

| marathon hook | This extension | Act from |
|---------------|----------------|----------|
| `on-start` | layer the architecture conventions into the session; bootstrap or adopt the layer | `mechanics/on-start.md` |
| `on-reset` | land a design note that has generalized, and record the landing | `mechanics/on-reset.md` |
| `on-close` | commit the landing where it landed | `mechanics/on-close.md` |

## Locating the layer

Every hook acts on one layer, resolved once at `on-start` from where the extension is enabled:

- `[project] extensions` on a standalone project → the top-level `architecture/` directory in
  this repository.
- `[workspace] extensions` at the coordinator → the member repository named by `repo` in the
  coordinator's own `.claude/marathon-architecture.toml` — this extension's config, never
  core's — resolved as an order key against the coordinator's `order` and `[workspace.paths]`,
  for every member project's sessions.

A workspace keeps one layer, shared by every member, so two cases resolve to nothing and neither
is guessed:

- The coordinator enables the extension and `.claude/marathon-architecture.toml` does not exist
  yet, or names no `repo`. Bootstrap it per `mechanics/on-start.md`'s step 3: ask the architect
  which member holds the layer, or whether one is still to be created, and continue without the
  layer until it is settled.
- A member project enables the extension under its own `[project] extensions` while the
  coordinator does not. Report the mismatch and ask whether the layer belongs at the
  coordinator. Never create a second layer inside the member.

A session in a member project therefore writes into a repository that is not its own. How that
edit commits is `mechanics/on-close.md`.
