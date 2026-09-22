# Extension pipeline

| marathon hook | This extension | Act from |
|---------------|----------------|----------|
| `on-start` | layer the conventions in; bootstrap or adopt the layer | `mechanics/on-start.md` |
| `on-reset` | land a generalized note and record it | `mechanics/on-reset.md` |
| `on-close` | commit the landing | `mechanics/on-close.md` |

## Locating the layer

Resolved once at `on-start`:

- Enabled under `[project]` → the top-level `architecture/` directory of this repository.
- Enabled under `[workspace]` at the coordinator → the member repository named by `repo` in the
  coordinator's `.claude/marathon-architecture.toml`, resolved against `order` and
  `[workspace.paths]`, shared by every member's sessions.

Never guess a missing layer. With no config file or no `repo`, bootstrap per `on-start.md`. A
member enabling the extension under `[project]` while its coordinator doesn't is a mismatch to
report; never create a second layer inside a member.
