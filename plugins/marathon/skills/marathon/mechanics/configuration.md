# Configuration

The canonical layout of `.claude/marathon.toml`, the file that declares a repository's marathon
configuration. `init` writes it once; every session afterward reads it.

```toml
[project]
kind = "code"        # production source with a build-and-test loop
# kind = "context"   # the whole repository is context — the agent authors it directly
# Optional: marathon extensions enabled for this project, by skill name.
# extensions = ["<extension>"]

[remote]
platform = "github"  # the platform the project publishes to
publish  = "gh pr create"

# Optional: only a workspace coordinator declares this block.
[workspace]
role  = "coordinator"
order = [
  "core-lib",
  ["service-a", "service-b"],
  "gateway",
]
# Optional: marathon extensions enabled for every project in the workspace.
# extensions = ["<extension>"]

# Optional: resolves an order key that is not a sibling directory.
[workspace.paths]
core-lib = "~/code/core-lib"

# Optional: only an experiment declares this block.
[experiment]
serves = "org"  # the project, or workspace coordinator, this experiment serves
```

## Keys

- **`[project] kind`** — `code` when the repository holds production source with a build-and-test
  loop that defines stage boundaries and validation; `context` when the repository is itself
  context (prose, configuration, skills) that the agent authors directly, with no tests. A context
  project can still version and release what it ships.
- **`[remote]`** — the platform and the command `close` runs to publish a branch.
- **`[workspace]`** — coordinator only: `order` lists layers lowest first, an array entry a layer
  of peers, and `[workspace.paths]` maps a key that isn't a sibling directory
  (`references/workspace-coordination.md`).
- **`[experiment]`** — experiment only: `serves` names the project, or the workspace's
  coordinator, that reads its results (`commands/experiment.md`).
- **`extensions`** — enabled extensions by skill name, under `[project]` for this repository or
  `[workspace]` for every member (`references/extensions.md`).
