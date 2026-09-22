# Extension pipeline

| marathon hook | This extension | Act from |
|---------------|----------------|----------|
| `on-start` | layer the conventions in; bootstrap a missing manifest | `mechanics/on-start.md` |
| `on-reset` | cite roadmap tasks in the session record | `mechanics/on-reset.md` |
| `on-close` | delete the finished task and advance `next` | `mechanics/on-close.md` |

The manifest is the coordinator's `context/roadmap.toml` when enabled under `[workspace]`, so a
member's session edits it in another repository, or the project's own when enabled under
`[project]`.
