# Extension pipeline

| marathon hook | What this extension does | Instructions |
|---------------|--------------------------|--------------|
| `on-start` | Finds the layer, creates or adopts it, and gives the session its location | `mechanics/on-start.md` |

## Finding the layer

Find the layer once, at `on-start`:

- Enabled under `[project]`: the layer is this repository's top-level `architecture/` directory.
- Enabled under `[workspace]` at the coordinator: the layer is the member repository named by
  `repo` in the coordinator's `.claude/marathon-architecture.toml`, located through `order` and
  `[workspace.paths]`. Every member's sessions share it.

Never guess where a missing layer is. When the configuration file or its `repo` key is missing,
follow `on-start.md`. When a member enables the extension under `[project]` and its coordinator
doesn't, report the mismatch. Never create a second layer inside a member.
