# marathon coordinate

Run one change across several marathon projects that live together in a workspace. Use `coordinate`
when a single objective genuinely spans more than one project in dependency order — build a capability
in a lower project and take it up into the one that consumes it, in the same effort. Each project is
handled as a normal marathon session with its own branch, `reset.md`, and published change;
`coordinate` is the orchestration wrapper around them, and it holds no state of its own.

For the design behind this — why the workspace holds no context, and how the mechanism splits from the
map — see `references/workspace-coordination.md`.

## What a workspace is

A **workspace** is a directory that contains marathon projects as siblings but is not itself one: it
has no `context/` of its own. It is an ordinary, often unversioned, container. The projects under it
are independent marathon repositories, each with its own `context/` and `.claude/marathon.toml`.

One project may declare itself the **coordinator** in its `.claude/marathon.toml`:

```toml
[workspace]
role  = "coordinator"
# Each entry is a dependency layer, lowest first. An entry is either one
# project, or an array of projects that share a layer with no dependency
# between them.
order = [
  "core-lib",
  ["service-a", "service-b"],
  "gateway",
]
```

`order` is the sequence a change flows through, expressed as **layers**. A layer is one entry in the
list; the layers are strictly ordered, lowest first. An entry that is a single key is a layer of one
project. An entry that is an array is a layer of **adjacent** projects — peers that sit at the same
depth with no dependency between them, so their order within the layer is free and they can even
proceed in parallel. In the example, `service-a` and `service-b` both sit above `core-lib` and below
`gateway`, but neither depends on the other.

The coordinator is normally the workspace's context project (the one that already describes the estate
as a whole), so declaring the order there adds no second source of truth; it is the machine-readable
form of the map that project already keeps in prose.

## 1. Detect the workspace and enumerate

Run from the workspace directory (or name it). Then:

1. Confirm it's a workspace: the directory has no `context/`, but its subdirectories include marathon
   projects (each has a `context/`).
2. Find the coordinator: the participating project whose `.claude/marathon.toml` has
   `[workspace] role = "coordinator"`. Read its `order` and flatten it into layers.
3. Resolve each key to a checkout — a sibling directory in the workspace by name, or, when a
   participant lives outside the workspace root, through the optional `[workspace.paths]` table in
   the coordinator's `marathon.toml`, which maps the key to a directory on this machine. Ask the
   developer about a key that is neither a sibling nor mapped.
4. **Degrade gracefully.** If no project declares itself coordinator, don't assume an order: enumerate
   the sibling projects (the directories with a `context/`) and ask the developer which participate and
   in what order.

## 2. Take the objective

`coordinate` needs to know which projects the change touches and what each one does. It accepts the
objective in any of three forms:

- **a set of reset files** — the participating projects are already checkpointed; their `reset.md`
  Next-focus lines are the per-project instructions.
- **a single reset** the orchestration points to — one project's Next-focus that names the cross-repo
  work; expand it into the participants.
- **direct instructions** — described in the session, in place of pre-planned per-project detail.

From whichever form, settle the set of involved projects and what each contributes, placed in the
coordinator's layers. Projects not touched by the objective are skipped, and a layer with no involved
project collapses.

## 3. Plan the change across the projects

Plan in plan mode, spanning the projects together, before touching any of them. Work out how the change
decomposes across the layers — what lands in each project, and what a higher layer depends on from a
lower one. Where a project already has a `reset.md` Next-focus for its part, resume that plan; where the
objective is fresh, plan that project's step as a normal fresh `start` would.

On a **code** project the plan produces an implementation-guide slice; on a **context** project it is
direct authoring. Collect the guide slices into one **consolidated guide** authored at the directory
`coordinate` was launched from — for example `guide.md` at the workspace root. This guide is ephemeral:
the workspace holds no context and versions nothing, so the file is never committed and does not survive
the session, exactly like a single project's `context/guide.md`.

## 4. Fan out, layer by layer

Work the layers in order, lowest first. Finish a layer before starting the next, so a higher layer
builds against the real change beneath it and not a promise. Within a layer, the peer projects have no
dependency between them: do them in any order. For each project, run its part as the marathon session it
would be on its own, honoring the project's kind. Each project branches at the start of its own part,
under the coordinated change's shared slug; the plan approved collectively in step 3 substitutes for
that project's own plan-mode gate; and `on-session-start` fires per project as its part begins (see
`references/extension-hooks.md`).

- **code project** — hand the developer that project's slice of the consolidated guide and let them
  apply it; stay available for fixes. Then, at that project's closeout, the agent adds its tests and
  documentation and validates.
- **context project** — author the change directly in that project; there is no guide slice and no
  handoff.

## 5. Close each project

Context management, validation, and closeout happen **per project**, not across the workspace. In each
project, run `close` as usual: tidy that project's `context/`, write its `reset.md`, commit, and publish
its own branch as its own change proposal. Each project ends with its own `reset.md`, branch, and PR.

The coordination run itself produces nothing to commit: no workspace branch, no workspace `reset.md`, no
workspace context. When every participating project has closed, the objective is done and the ephemeral
guide is discarded.

## Continuity and interruption

Because each project checkpoints independently, an interrupted fan-out resumes per project: a project
mid-work has left its `reset.md` at `Status: handoff` and is picked up by a normal `start` on its
branch. The cross-project order on resume comes from the coordinator's `order`, not from any workspace
record — there is none. Re-running `coordinate` re-reads the layers and continues from wherever each
project's `reset.md` says it is.

## Not yet handled

When a coordinated change touches several projects that each mirror to a project-management tracker, how
the extension hooks fire across them is not yet defined. For now the hooks fire per project as they
normally do; cross-tracker coordination is left to a later step. See `references/extension-hooks.md`.
