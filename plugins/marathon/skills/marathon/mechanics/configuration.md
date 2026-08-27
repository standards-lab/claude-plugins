# Configuration

The canonical layout of `.claude/marathon.toml`, the file that declares a repository's marathon
configuration. `init` writes it once; every session afterward reads it.

```toml
[project]
kind = "code"        # production source the developer authors and answers for
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
docs  = "core-lib"   # optional: the order key that is this workspace's docs landing zone
# Optional: marathon extensions enabled for every project in the workspace.
# extensions = ["<extension>"]

# Optional: resolves an order key that is not a sibling directory.
[workspace.paths]
core-lib = "~/code/core-lib"
```

## Project kind

`[project] kind` declares which of the two kinds the repository is. It decides how `start` and
`close` behave, and how `coordinate` treats the project in a fan-out; the role boundary each kind
draws is `references/role-boundary.md`.

- **code** — the repository contains production source code: the implementation logic that makes
  a program behave. The agent drafts each change as an implementation guide, the developer
  applies it, and closeout adds tests and documentation.
- **context** — the repository *is* context: prose, configuration, and skills (which are advanced
  context, not source). The agent authors the whole repository directly — no implementation
  guide, no tests. The developer sets direction and reviews and approves each change; the pull
  request is the ownership seam. A context project can still version and release what it ships (a
  plugin, a document set); it just has no code layer.

When in doubt, apply the git-blame test in `references/role-boundary.md`: whom should the
deliverable's history point at?

## Remote

`[remote]` names the platform the project publishes to and the command that proposes a change
there — `gh pr create` on GitHub, `glab mr create` on GitLab, the equivalent elsewhere. Closeout
runs it to publish the finished branch.

## Workspace

Only a coordinator declares `[workspace]`. `order` is the dependency map a coordinated change
flows through — a list of layers, lowest first; an array entry is a layer of adjacent peers.
`docs` names the workspace's docs landing zone by order key. `[workspace.paths]` resolves an
order key that is not a sibling directory. The design is
`references/workspace-coordination.md`.

## Extensions

`extensions` lists enabled extensions by skill name: under `[project]` for this repository, under
`[workspace]` at a coordinator for every member project. Enabling, resolution, and the hook
points are `references/extensions.md` and `mechanics/hooks.md`.
