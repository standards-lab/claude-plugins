# Configuration

`.claude/marathon.toml` declares a repository's marathon configuration. `init` writes it once, and
every session reads it.

```toml
[project]
kind     = "code"      # production source with a build-and-test loop
# kind = "context"     # the whole repository is context, which the agent writes directly
check    = "mise run check"
# Optional: reports what trails its latest release; PLAN runs it, CI never does.
currency = "mise run currency"
# Optional: decides whether a Release tag may be pushed; SHIP runs it, CI never does.
gate     = "mise run gate"
# Optional: marathon extensions enabled for this project, by skill name.
# extensions = ["<extension>"]

[remote]
platform = "github"    # the platform the project publishes to
publish  = "gh pr create"
# Optional: without it, SHIP stops at the open pull request.
merge    = "gh pr checks --watch && gh pr merge --merge --delete-branch"
# Optional: waits for the default branch's CI on the merge commit, and fails if it fails.
# The runs can take a few seconds to appear after the merge, so it polls for them first.
ci       = '''
sha=$(git rev-parse HEAD)
for _ in $(seq 30); do
  runs=$(gh run list --commit "$sha" --event push --json databaseId --jq '.[].databaseId')
  [ -n "$runs" ] && break || sleep 10
done
[ -n "$runs" ] || exit 1
for run in $runs; do gh run watch --exit-status "$run" || exit 1; done
'''

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
- **`[project] currency`**: optional. The repository's read-only command that reports what trails
  its latest release: the direct dependencies, the toolchain, and the pins, such as CI actions and
  images. Indirect dependencies are left to the ecosystem's resolver. PLAN runs it at development
  time; CI never does, since a release upstream would fail a build nothing in the repository
  changed. It exits 0 with empty stdout when everything is current. It exits nonzero when something
  trails, with one stdout line per trailing item, suggested as `<where>: <item> <pin> -> <latest>`,
  where the latest is what the ecosystem's tooling reports. Nonzero with no stdout lines means the
  command itself failed, which PLAN reports as a fact and doesn't block on. Diagnostics go to
  stderr. marathon prescribes no tool.
- **`[project] gate`**: optional. The repository's release gate, the release-time part of its
  automated checks, such as a plugin's eval suite (`references/build.md`). SHIP runs it at
  development time on the merge commit, before each Release tag, with that tag as its one
  argument, as in `mise run gate v1.4.0` (`mechanics/pipeline.md`, "Releasing"). It exits 0
  when the tag may be pushed, and nonzero to hold it. CI never runs it. marathon fixes only this
  contract and prescribes no tool.
- **`[remote]`**: the platform, the command SHIP runs to publish a branch, and optionally the
  command it runs to merge the published branch once its checks pass.
- **`[remote] ci`**: optional. A command that waits for the default branch's CI runs on the merge
  commit, with that commit checked out, and exits nonzero if one fails or none starts. CI often
  runs more after a merge than on the pull request, such as integration tests only on the default
  branch, so the merge command's wait on the pull request's checks isn't enough. SHIP runs it
  before tagging a release; without it, tagging follows the merge directly
  (`mechanics/pipeline.md`, "Releasing").
- **`[workspace]`**: for the coordinator only. `order` lists layers, lowest first, and an array
  entry is a layer of peers. `exclusive` lists groups of repositories that at most one active goal
  may touch at a time, narrowing the repository lock (`mechanics/goals.md`). `[workspace.paths]`
  gives the location of a key that isn't a sibling directory
  (`references/workspace-coordination.md`). `[workspace.experiments]` gives the local path and
  remote a new spike takes, with `<slug>` replaced (`commands/experiment.md`).
- **`extensions`**: the enabled extensions, by skill name. Under `[project]`, they apply to this
  repository. Under `[workspace]`, they apply to every member (`references/extensions.md`).
