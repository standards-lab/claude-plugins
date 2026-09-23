# Extension pipeline

| marathon hook | What this extension does | Instructions |
|---------------|--------------------------|--------------|
| `on-start` | Loads the conventions, and creates a missing manifest | `mechanics/on-start.md` |
| `on-reset` | Cites roadmap tasks in the session record | `mechanics/on-reset.md` |
| `on-close` | Deletes the finished task and advances `next` | `mechanics/on-close.md` |

Enabled under `[workspace]`, the manifest is the coordinator's `context/roadmap.toml`, so a
member's session edits it in another repository. Enabled under `[project]`, the manifest is the
project's own.
