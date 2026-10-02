# Configuration

`.claude/marathon.toml` declares a repository's marathon configuration. `init` writes it once, and
every session reads it.

```toml
[project]
kind  = "code"         # production source with a build-and-test loop
# kind = "context"     # the whole repository is context, which the agent writes directly
check = "mise run check"
# Optional: marathon extensions enabled for this project, by skill name.
# extensions = ["<extension>"]

[remote]
platform = "github"    # the platform the project publishes to
publish  = "gh pr create"
# Optional: without it, SHIP stops at the open pull request.
merge    = "gh pr checks --watch && gh pr merge --merge --delete-branch"
# Optional: waits for the default branch's CI on the merge commit, and fails if it fails.
ci       = "gh run watch --exit-status $(gh run list --branch main --commit \"$(git rev-parse HEAD)\" --limit 1 --json databaseId --jq '.[0].databaseId')"

# Optional: only a workspace coordinator declares this table.
[workspace]
role  = "coordinator"
order = [
  "core-lib",
  ["service-a", "service-b"],
  "gateway",
]
# Optional: groups of repositories at most one active goal may touch at a time.
# exclusive = [["core-lib", "service-a", "service-b", "gateway"]]
# Optional: marathon extensions enabled for every project in the workspace.
# extensions = ["<extension>"]

# Optional: the location of an order key that isn't a sibling directory.
[workspace.paths]
core-lib = "~/code/core-lib"

# Optional: the hosting convention `experiment` and `plan` propose for a new spike.
[workspace.experiments]
path   = "~/experiments/spike-<slug>"
remote = "https://github.com/<owner>/spike-<slug>.git"
```

## Keys

- **`[project] kind`**: `code` when the repository holds production source with a build-and-test
  loop. `context` when the repository is itself context, such as prose, configuration, or skills,
  which the agent writes directly and which has no tests. A context project can still version and
  release what it ships.
- **`[project] check`**: the repository's one deterministic check, which every slice and review
  runs (`references/build.md`). A context project names its consistency script, if it has one.
- **`[remote]`**: the platform, the command SHIP runs to publish a branch, and optionally the
  command it runs to merge the published branch once its checks pass.
- **`[remote] ci`**: optional. A command that waits for the default branch's CI run on the merge
  commit, with that commit checked out, and exits nonzero if the run fails. CI often runs more
  after a merge than on the pull request, such as integration tests only on the default branch,
  so the merge command's wait on the pull request's checks isn't enough. SHIP runs it before
  tagging a release; without it, tagging follows the merge directly (`mechanics/pipeline.md`,
  6 · SHIP).
- **`[workspace]`**: for the coordinator only. `order` lists layers, lowest first, and an array
  entry is a layer of peers. `exclusive` lists groups of repositories that at most one active goal
  may touch at a time, narrowing the repository lock (`mechanics/goals.md`). `[workspace.paths]`
  gives the location of a key that isn't a sibling directory
  (`references/workspace-coordination.md`). `[workspace.experiments]` gives the local path and
  remote a new spike takes, with `<slug>` replaced (`commands/experiment.md`).
- **`extensions`**: the enabled extensions, by skill name. Under `[project]`, they apply to this
  repository. Under `[workspace]`, they apply to every member (`references/extensions.md`).
