# Isolated experiments: parallel work outside the workspace tree

Captured 2026-09-22 from a workflow-refinement planning session at the coordinator
(`standards-lab`). The architect wants to run genuinely parallel, non-interfering work — for
example `v1.messaging` and `v1.ai`, two roadmap goals with no dependency between them — instead of
sitting idle between review checkpoints on one goal while another waits its turn. This concept is
settled direction for the `start` session that implements it.

## The gap

A workspace has exactly one `reset.md`, at the coordinator (`mechanics/reset-file.md`), and every
session — including today's in-tree experiments — routes through it (LOCATE resolves the single
reset file before anything else). That means only one workspace session is ever in flight at a
time. `references/workspace-coordination.md` already keeps experiments in-tree at the coordinator
deliberately, so spike code doesn't collide with member-repository tooling — `blobfs`'s 150 KB Go
module sat directly under the public `standards-lab/experiments/blobfs/`, alongside a `context`-kind
repository's own prose and configuration. The isolation was right; putting the spike inside the
org-context repository was the compromise, and nothing requires it once the spike can be its own
repository.

## Design

**An isolated experiment is a standalone marathon project**, outside the workspace tree entirely —
not a workspace member, not mapped through `[workspace.paths]`. That alone gives it independent
session state: its own `context/reset.md`, its own branches, its own `marathon.toml`, none of
which the coordinator's single reset file ever sees.

- **Location.** `~/experiments/<slug>` is the architect's own machine convention; marathon does
  not fix the path.
- **Declaration.** `init` gains one founding decision on a code project: does this repository
  serve a workspace goal? If yes, `.claude/marathon.toml` records a new core table:

  ```toml
  [experiment]
  workspace = "org"   # the references key of the coordinator this experiment serves
  ```

  The goal it serves is cited by dotted roadmap path in `context/README.md` and the reset file,
  under `marathon-roadmap`'s existing citation convention — core never learns dotted paths itself.
- **The connection back to the workspace** reuses the references convention already in production
  use at `standards-lab` (`context/repo-references.md`, once `flat-context.md` lands): a committed
  `references.toml` naming the workspace repositories the experiment draws on, and a gitignored
  `references.local.toml` pointing at local checkouts. The experiment's sessions read those
  checkouts — architecture principles, sibling library concepts, `go-web-service`'s reference
  composition — and never write them. Code dependencies on member modules stay published versions
  through `go.mod`, never a `replace` directive. This is `marathon-references-extension.md`'s
  proposal applied before the extension exists; the convention works today without a hook, and
  making the extension a prerequisite would stall this on that concept's own open questions.
- **Entry point.** `commands/experiment.md`'s SETTLE gains a grain decision: in-tree (today's
  model — one session, evidence for the step directly in front) or isolated (its own project,
  several sessions, a candidate for a new repository, or work meant to run in parallel with the
  workspace). Choosing isolated runs `init` at the new location, seeds the references files and
  `context/README.md`, writes the experiment's first reset, adds its key to the coordinator's
  references catalog, and marks the roadmap task `active` (`roadmap-waves.md`). The coordinator
  session then closes; the architect opens a session in the new repository and works it with
  `start`.
- **Graduation is a read, not a write.** The experiment's last session is an ordinary `close`; its
  Disposition makes the record self-contained, the way `blobfs`'s did. A coordinator `plan`
  session then reads the closed experiment and does the intake — captures findings as notes, adds
  the tasks the goal now needs. For a goal that earns a new repository, the tasks follow the shape
  `goals.blobfs.tasks.build` already established: a `build` task whose session runs `init` in the
  new sibling directory, re-derives the library from the experiment rather than copying it, adds
  the key to the coordinator's `order` map at the layer SETTLE decides, and adds it to the
  references catalog. The experiment never edits the workspace; the workspace reads the
  experiment.
- **The existing in-tree `experiments/` stays**, narrowed to genuinely one-session spikes.
  `standards-lab/experiments/` itself (`blobfs`, `sql-dsl`) is frozen as historical record — no new
  entries — since anything the workspace needs a code spike for is now either large enough to
  isolate or small enough not to need a repository at all.

## Files this touches

`commands/experiment.md` (grain decision, isolated Execute and Conclude), `commands/init.md`
(`[experiment]` founding decision), `mechanics/configuration.md` (`[experiment]` table),
`references/workspace-coordination.md` (experiments section notes the isolated form),
`references/context-engineering.md`'s "Kept outside `context/`" list.

At `standards-lab`: `context/workspace-structure.md` notes `experiments/` is frozen; the
references catalog gains an "Experiments" grouping.

## Open questions for the implementing session

- Whether a per-session reset file at the coordinator (`context/reset/<slug>.md`) is ever needed
  as a fallback — if isolated experiments turn out to need mid-flight shims into member libraries
  often enough that promotion backlogs grow long (as `blobfs` did into `sqlate`), the honest fix
  may be per-session resets inside the workspace instead of further isolation.
- Whether `[experiment]` needs anything beyond `workspace`, or whether the served-goal citation in
  `context/README.md` carries enough.
