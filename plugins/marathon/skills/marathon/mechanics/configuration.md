# Configuration

`.claude/marathon.toml` declares a repository's marathon configuration. `init` writes it once, and
every session reads it.

```toml
[project]
kind = "code"        # production source with a build-and-test loop
# kind = "context"   # the whole repository is context, which the agent writes directly
# Optional: marathon extensions enabled for this project, by skill name.
# extensions = ["<extension>"]

[remote]
platform = "github"  # the platform the project publishes to
publish  = "gh pr create"

# Optional: only a workspace coordinator declares this table.
[workspace]
role  = "coordinator"
order = [
  "core-lib",
  ["service-a", "service-b"],
  "gateway",
]
# Optional: marathon extensions enabled for every project in the workspace.
# extensions = ["<extension>"]

# Optional: the location of an order key that isn't a sibling directory.
[workspace.paths]
core-lib = "~/code/core-lib"

# Optional: only an experiment declares this table.
[experiment]
serves = "org"  # the project, or workspace coordinator, this experiment serves
```

## Keys

- **`[project] kind`**: `code` when the repository holds production source with a build-and-test
  loop, which defines stage boundaries and validation. `context` when the repository is itself
  context, such as prose, configuration, or skills, which the agent writes directly and which has
  no tests. A context project can still version and release what it ships.
- **`[remote]`**: the platform, and the command `close` runs to publish a branch.
- **`[workspace]`**: for the coordinator only. `order` lists layers, lowest first, and an array
  entry is a layer of peers. `[workspace.paths]` gives the location of a key that isn't a sibling
  directory (`references/workspace-coordination.md`).
- **`[experiment]`**: for an experiment only. `serves` names the project, or the workspace's
  coordinator, that reads the experiment's results (`commands/experiment.md`).
- **`extensions`**: the enabled extensions, by skill name. Under `[project]`, they apply to this
  repository. Under `[workspace]`, they apply to every member (`references/extensions.md`).
