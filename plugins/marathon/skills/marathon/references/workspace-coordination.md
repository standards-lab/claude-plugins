# Workspaces

marathon manages one repository. But related repositories are often developed together — a set of
capability libraries, a template, and a service that consumes them — and a settled step sometimes
crosses several of them in dependency order. A workspace lets marathon run that step as one
session while keeping each repository a self-contained marathon project. This note is the design.

## The workspace holds no context

A **workspace** is the directory the related projects sit under as siblings. It is not itself a
marathon project: it has no `context/`, and the workspace adds none. That is deliberate. The
knowledge a workspace might seem to want — which repositories exist, how they depend, which one
leads — already has a home in whichever project describes the estate as a whole (typically an
organization-context project). Giving the workspace its own `context/` would be a second
description of the same thing, and the two would drift. So the workspace knowledge splits in two,
and neither part is a workspace file:

- **Mechanism — general, in the skill.** How a session locates the workspace, resolves the
  coordinator's reset, and runs a step that spans member repos. This lives in
  `mechanics/pipeline.md` and the working-session playbooks, and applies to any workspace.
- **Map — specific, in the coordinator project.** Which projects the workspace holds and in what
  dependency order. marathon does not hardcode this; it reads it from the project that declares
  itself coordinator.

## The coordinator and the map

One project in the workspace declares itself the coordinator in its `.claude/marathon.toml`, with
the `[workspace]` block whose canonical layout is `mechanics/configuration.md`. Its `order` is the
map in machine-readable form: a list of **layers**, lowest dependency first. Each entry is either
one project or an array of **adjacent** projects — peers at the same depth with no dependency
between them, so their order within a layer is free. This mirrors a real dependency graph, which
is layered rather than strictly linear.

The coordinator is normally the project that already narrates the estate in prose (its capability
map or references catalog). Declaring `order` there adds no second source of truth: it is the
executable projection of a map the project already keeps. Each key resolves to a checkout — a
sibling directory in the workspace by name, or, for a participant that lives elsewhere on disk,
through the coordinator's optional `[workspace.paths]` table, which maps the key to a directory on
this machine. A key that is neither a sibling nor mapped is resolved by asking the architect. If
no project declares itself coordinator, don't assume an order: enumerate the sibling projects (the
directories with a `context/`) and ask the architect.

## Cross-repo steps

A working session whose settled step spans member repos runs as one session. The step's scope
names the repos it touches, placed in the map's layers; the session works them lowest layer first,
so a higher layer builds against the real change beneath it. The stage list of
`references/staged-execution.md` is one list spanning the repos, its stages grouped by repository
in map order, each stage's unit set by its repository's project kind. The session creates a
branch in each touched repo under the step's shared slug, and its `close` commits and publishes
each repo's branch as that repository's own change proposal.

## Continuity lives at the coordinator

A workspace maintains one reset file, at the coordinator; member projects carry none
(`mechanics/reset-file.md`). One session's story then lives in one place: the record names the
member repos it concerns, and its Next-focus names where the next session continues, so a session
entered anywhere in the workspace routes through the same anchor. An interrupted cross-repo step
resumes from that record plus each touched repo's open branch, and the dependency order on resume
comes from re-reading the coordinator's `order`.

## Experiments live at the coordinator

In a workspace, every experiment lives under the coordinator's top-level `experiments/`, never
inside a member repository. A member repository carries the tooling of a code project: a Go
workspace file, a CI matrix, format sweeps, module lists. A spike placed inside it sits under all
of that and either breaks against it or has to be fenced from each piece in turn. The coordinator
carries no such tooling, so a spike there collides with nothing. The spike depends on the member
modules it needs as published versions, through its own `go.mod` or the equivalent, never
through a replace directive to a sibling checkout. A change the spike implies for a member
repository's code is laid out inside the experiment, and it reaches that repository only by
promotion at `close`, when the architect accepts the result into the workspace's effort. The
session's branch is the coordinator's; if the spike's outcome changes a member repository's
context, that edit is a cross-repo edit on that repository's own branch, recorded under
**Cross-repo** in the reset disposition. A standalone project keeps its own `experiments/`, as
`commands/experiment.md` states.

## Awareness follows the dependency direction

Sessions read the map from the coordinator; they never teach a lower project about the projects
that consume it. A dependency graph's awareness runs downward — a project knows what it builds on,
not what builds on it — and marathon respects that: the estate-wide view lives only in the
coordinator, the one place the whole is legitimately described together. A lower project stays
unaware of its consumers, even inside a cross-repo step.

## Coordinator conventions

An organization-level coordinator often keeps conventions that bind the member repositories —
naming rules, authoring rules, the awareness direction itself. Because awareness runs downward, a
member repository never cites them in its own stable context. The binding runs through sessions
instead: a member's `review` consults the coordinator's conventions as part of its drift check,
and a cross-repo step applies them as it works each repo.
