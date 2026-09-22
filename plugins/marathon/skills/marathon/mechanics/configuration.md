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

# Optional: only an isolated experiment that serves a workspace declares this block.
[experiment]
workspace = "org"  # the name of the coordinator repository this experiment serves
```

## Project kind

`[project] kind` declares which of the two kinds the repository is. It decides how `start` and
`close` behave.

- **code** — the repository contains production source code: the implementation logic that makes
  a program behave. The agent implements each step in stages, each committed once its check
  passes, with the architect confirming behavior at checkpoints; closeout follows validation.
- **context** — the repository *is* context: prose, configuration, and skills (which are advanced
  context, not source). The agent authors the whole repository directly — no tests. The architect
  sets direction, confirms each checkpoint, and approves the pull request, which is where
  ownership of the change passes to the project. A context project can still version and
  release what it ships (a plugin, a document set); it just has no code layer.

When in doubt, ask whether there is a build-and-test loop that defines stage boundaries and a
validation phase; if there is, it's `code`.

## Remote

`[remote]` names the platform the project publishes to and the command that proposes a change
there — `gh pr create` on GitHub, `glab mr create` on GitLab, the equivalent elsewhere. Closeout
runs it to publish the finished branch.

## Workspace

Only a coordinator declares `[workspace]`. `order` is the dependency map a cross-repo step
flows through — a list of layers, lowest first; an array entry is a layer of adjacent peers.
`[workspace.paths]` resolves an order key that is not a sibling directory. The design is
`references/workspace-coordination.md`.

## Experiment

Only an isolated experiment declares `[experiment]` (`commands/experiment.md`). `workspace` names
the coordinator repository the experiment serves. The experiment is not a workspace member: it
keeps its own reset file, and the key records only which workspace reads its results. The goal
it serves is cited in its `context/README.md`.

## Extensions

`extensions` lists enabled extensions by skill name: under `[project]` for this repository, under
`[workspace]` at a coordinator for every member project. Enabling, resolution, and the hook
points are `references/extensions.md` and `mechanics/hooks.md`.
