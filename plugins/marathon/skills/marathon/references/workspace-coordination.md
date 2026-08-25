# Workspace coordination

marathon manages one repository. But related repositories are often developed together — a set of
capability libraries, a template, and a service that consumes them — and a single objective sometimes
crosses several of them in dependency order. Workspace coordination lets marathon run that objective as
one orchestrated change while keeping each repository a self-contained marathon project. The `coordinate`
command is the mechanism; this note is the design behind it.

## The workspace holds no context

A **workspace** is the directory the related projects sit under as siblings. It is not itself a marathon
project: it has no `context/`, and coordination adds none. That is deliberate. The knowledge a workspace
might seem to want — which repositories exist, how they depend, which one leads — already has a home in
whichever project describes the estate as a whole (typically an organization-context project). Giving the
workspace its own `context/` would be a second description of the same thing, and the two would drift.
So the coordination knowledge splits in two, and neither part is a workspace file:

- **Mechanism — general, in the skill.** How to detect a workspace, enumerate the participating
  projects, and fan a session out to each. This lives in `commands/coordinate.md` and applies to any
  workspace.
- **Map — specific, in the coordinator project.** Which projects participate and in what order. marathon
  does not hardcode this; it reads it from the project that declares itself coordinator.

## The coordinator and the map

One project in the workspace declares itself the coordinator in its `.claude/marathon.toml`:

```toml
[workspace]
role  = "coordinator"
order = [
  "core-lib",
  ["service-a", "service-b"],
  "gateway",
]
```

`order` is the map in machine-readable form: a list of **layers**, lowest dependency first. Each entry is
either one project or an array of **adjacent** projects — peers at the same depth with no dependency
between them, so their order within the layer is free. This mirrors a real dependency graph, which is
layered rather than strictly linear.

The coordinator is normally the project that already narrates the estate in prose (its capability map or
references catalog). Declaring `order` there adds no second source of truth: it is the executable
projection of a map the project already keeps. Each key resolves to a checkout — a sibling directory in
the workspace by name, or, for a participant that lives elsewhere on disk, through the coordinator's
optional `[workspace.paths]` table, which maps the key to a directory on this machine. A key that is
neither a sibling nor mapped is resolved by asking the developer.

## Detection and degradation

A directory that contains marathon projects but is not one itself (no `context/`) is a workspace.
`coordinate` finds the coordinator by looking for the participating project whose `marathon.toml`
declares `role = "coordinator"`, and reads its `order`.

With no coordinator present, marathon degrades rather than guessing: it enumerates the sibling projects
(the directories that do have a `context/`) and asks the developer which participate and in what order.
Coordination still runs; it just sources the order from the developer for that session instead of from a
declaration.

## Honoring project kind in a fan-out

The participating projects need not be the same kind. A fan-out treats each by its own kind (see
`references/role-boundary.md`): a `code` project gets an implementation-guide slice the developer applies; a
`context` project gets its change authored directly. A single objective — build a capability in a code
library and revise the skill that documents the workflow around it — can therefore span both kinds in one
coordinated run.

## The consolidated guide is ephemeral

For the code-project slices of a coordinated change, the plan is collected into one consolidated guide,
authored at the directory `coordinate` was launched from (typically the workspace root). Because the
workspace versions nothing, this guide is never committed and does not outlive the session — the same
lifetime as a single project's `context/guide.md`. It is a working document for the fan-out, not a record.
The lasting record is what each project commits: its own branch, `reset.md`, and published change.

## Continuity stays per repository

A coordinated change checkpoints in each repository independently. There is no workspace-level state to
resume from, and none is needed: an interrupted fan-out picks up per project, each from its own
`reset.md` Status, and the cross-project order on resume comes from re-reading the coordinator's `order`.
This is why the coordination run itself produces no branch, no reset, and no context — the per-repository
records already carry everything continuity needs.

## Awareness follows the dependency direction

Coordination reads the map from the coordinator; it never teaches a lower project about the projects that
consume it. A dependency graph's awareness runs downward — a project knows what it builds on, not what
builds on it — and coordination respects that: the estate-wide view lives only in the coordinator, the
one place the whole is legitimately described together. A lower project stays unaware of its consumers,
even inside a coordinated change.

## Coordinator conventions

An organization-level coordinator often keeps conventions that bind the member repositories — naming
rules, authoring rules, the awareness direction itself. Because awareness runs downward, a member
repository never cites them in its own stable context. The binding runs through sessions instead: a
member's `review` consults the coordinator's conventions as part of its drift check, and a coordinated
fan-out applies them as it works each project.

## Not yet settled

How a coordinated change should present when several participating projects' extensions each mirror
outward is left open; see `references/extensions.md`. Until it is settled, hooks fire per project and
cross-mirror coordination is deferred.
