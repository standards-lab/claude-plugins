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

# Optional: the agents this project's sessions may delegate a unit of work to
# (behavior/delegation.md). One sub-table per agent, keyed by its name; delegation states what
# it's for, when it's warranted, and what it isn't for, in the project's own words — enough
# for a session to judge fit without reading anything else.
[agents.fable]
delegation = '''
Technical delegate for one stage's implementation or one design decision: a unit of work
technically complex enough that working it out directly would cost more than briefing this
agent to do it, handed over with full context. Not for documentation, commit messages, or
context notes — the session drafts and finishes those itself.
'''

[agents.opus]
delegation = '''
Escalation delegate for a high-stakes, hard-to-reverse decision: wide blast radius across the
project or workspace, or more context than the step itself needs held at once to reason
through. Not for implementation or routine technical work — fable is for that. Reach for this
only when getting the call right matters more than speed or cost.
'''

# Optional: only a workspace coordinator declares this block.
[workspace]
role  = "coordinator"
order = [
  "core-lib",
  ["service-a", "service-b"],
  "gateway",
]
architecture = "architecture"   # the order key of the workspace's architecture repository
# Optional: marathon extensions enabled for every project in the workspace.
# extensions = ["<extension>"]

# Optional: resolves an order key that is not a sibling directory.
[workspace.paths]
core-lib = "~/code/core-lib"

# Optional: the baseline agent declarations for every project in the workspace, resolved key by
# key against each project's own [agents] table — the project's value wins on a shared key.
[workspace.agents.fable]
delegation = '''
Technical delegate for one stage's implementation or one design decision: a unit of work
technically complex enough that working it out directly would cost more than briefing this
agent to do it, handed over with full context. Not for documentation, commit messages, or
context notes — the session drafts and finishes those itself.
'''
```

## Project kind

`[project] kind` declares which of the two kinds the repository is. It decides how `start` and
`close` behave.

- **code** — the repository contains production source code: the implementation logic that makes
  a program behave. The agent implements each step directly in stages, each reviewed before the
  next; closeout follows validation.
- **context** — the repository *is* context: prose, configuration, and skills (which are advanced
  context, not source). The agent authors the whole repository directly — no tests. The architect
  sets direction and reviews and approves each change; the pull request is the ownership seam. A
  context project can still version and release what it ships (a plugin, a document set); it just
  has no code layer.

When in doubt, ask whether there is a build-and-test loop that defines stage boundaries and a
validation phase; if there is, it's `code`.

## Remote

`[remote]` names the platform the project publishes to and the command that proposes a change
there — `gh pr create` on GitHub, `glab mr create` on GitLab, the equivalent elsewhere. Closeout
runs it to publish the finished branch.

## Agents

`[agents]` declares the agents this project's sessions may delegate to
(`behavior/delegation.md`): one sub-table per agent, named by its key, with a `delegation` field
describing what it's for. A workspace coordinator's `[workspace.agents]` is the baseline for
every member project; a project's own `[agents]` table adds an entry or overrides an existing
key's `delegation` text, resolved key by key. A project or workspace that declares none, or none
whose `delegation` fits the work at hand, changes nothing — the session does the work itself.

## Workspace

Only a coordinator declares `[workspace]`. `order` is the dependency map a cross-repo step
flows through — a list of layers, lowest first; an array entry is a layer of adjacent peers.
`architecture` names, by order key, the workspace's architecture repository: the project whose
whole tree is the architecture layer (`references/context-engineering.md`). Every workspace
declares one. `[workspace.paths]` resolves an
order key that is not a sibling directory. The design is
`references/workspace-coordination.md`.

## Extensions

`extensions` lists enabled extensions by skill name: under `[project]` for this repository, under
`[workspace]` at a coordinator for every member project. Enabling, resolution, and the hook
points are `references/extensions.md` and `mechanics/hooks.md`.
