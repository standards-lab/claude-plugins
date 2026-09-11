# marathon core: three concepts from real usage

Staging note — written from a `personal-agents` session at the architect's request, meant to
move to `~/architecture/claude-plugins/context/concepts/marathon-updates.md` on the
architect's laptop, where the `claude-plugins` project actually lives. Not part of
`personal-agents`; do not commit it there.

## 1. Extract the architecture layer into a `marathon-architecture` extension

### The problem

The architecture layer (`references/context-engineering.md`, "The architecture layer"
section) is currently mandatory, not optional, by marathon core's own words: "It is a
conventional layer of every marathon project, not an option, because it is where a design
note goes when the built work cannot express it." `commands/init.md` scaffolds
`architecture/README.md` unconditionally for a standalone project (or requires a workspace's
`[workspace] architecture` key) as one of its "founding decisions," with no toggle to skip
it.

Not every project needs this layer. A narrow, single-purpose repository — a documentation
set for one piece of infrastructure, a small tool with no principles that will ever
generalize past it — may never produce a design note that outgrows its own `context/design/`
and needs somewhere more general to go. For a project like that, the mandatory layer is
either dead weight (an `architecture/README.md` that never gains a second sentence) or a
temptation to write generalized-sounding content prematurely just to give the scaffold a
reason to exist — the same "committing to a design before you understand it" problem
`context-engineering.md`'s own "Add detail late" section warns against, just at the
project-shape level instead of the note level.

`marathon-roadmap` already demonstrates the right shape for this: a convention a project
opts into by declaring it in `.claude/marathon.toml`'s `extensions` list, inert when not
installed or not enabled, owning its own artifact (`roadmap.toml`) and hook points, per
`references/extensions.md`. The architecture layer should work the same way — present when a
project's shape calls for it, absent otherwise — instead of being baked into `init` and
`context-engineering.md` as a given.

### Proposed shape

Extract everything architecture-layer-specific out of marathon core into a
`marathon-architecture` extension, following the established extension contract:

- **Owns the artifact**: the `architecture/` directory (standalone) or the designated
  workspace architecture repository — bootstrapped at `on-start` per
  `mechanics/hooks.md`'s "Artifact bootstrap" rule, the same way an enabled extension with no
  artifact yet creates one today.
- **Declares its hook points** in its own SKILL.md: most likely `on-reset`/`on-close` (where
  promotion and decay decisions get made and recorded in the reset ledger today) and
  whatever `review`-time hook lets it flag an architecture page that restates a repository —
  currently a bare `review` behavior in `context-engineering.md`'s "Project documentation"
  section, with no hook boundary at all.
- **Leaves core with nothing architecture-specific**: `commands/init.md` drops "Architecture"
  from its founding-decisions checklist and its structure template entirely — a project that
  doesn't enable the extension gets no `architecture/` directory, no README, nothing.
  `context-engineering.md` drops "The architecture layer" section (moves to the extension's
  own reference doc) and softens "Kept outside `context/`: ... `architecture/`" to describe
  it only as something an enabled extension may add.

### What moves, concretely

- `commands/init.md`: remove the "Architecture" founding-decision bullet and the
  `architecture/` line from the structure template; note instead (one line) that an
  extension may add its own scaffolding when enabled, same as `marathon-roadmap` does for
  `roadmap.toml`.
- `references/context-engineering.md`: remove "The architecture layer" section body,
  replacing it with a pointer to the extension's own doc for projects that enable it; keep
  the "Deciding where something goes" decision list's promotion path generic ("promotion
  target defined by an enabled extension, if any" rather than naming `architecture/`
  outright).
- `mechanics/pipeline.md` / `mechanics/hooks.md`: no core change expected — the existing
  five universal hooks plus the `<command>:<moment>` extension point already cover what an
  architecture extension needs; this is the test case for whether that's actually true.
- New `marathon-architecture` skill: SKILL.md declaring the artifact, hook points, and
  targeted marathon version, plus whatever reference doc currently lives in
  `context-engineering.md`'s architecture section, moved and adapted.
- Workspace coordination (`references/workspace-coordination.md`, `mechanics/configuration.md`):
  the `[workspace] architecture` key names which member repository holds the workspace's
  architecture — this only makes sense when the extension is enabled workspace-wide.
  Whether that key moves under an extension-owned config block or stays in core
  `[workspace]` (since a workspace's shape is a core concept even if the layer it points at
  is not) is an open question below, not assumed here.

### Open questions

- Does a workspace *require* the architecture extension enabled at the coordinator, or can
  individual member projects opt in/out independently while the workspace still names one
  architecture repository? The current design treats the architecture repository as a
  workspace-wide given; extracting the layer may mean a workspace can legitimately have no
  architecture repository at all.
- Backward compatibility: does an already-`init`ed project with an `architecture/` directory
  need a migration step (declare the extension retroactively) or does the directory just
  keep existing, unmanaged by core, until someone enables the extension?
- Should `marathon-roadmap`'s own docs be checked for anything that assumes the architecture
  layer exists (a roadmap task promoting into `architecture/`, say) that would break once
  the layer becomes optional?

### Assumptions this rests on

- The five universal hooks (`on-start`, `on-execute`, `on-commit`, `on-reset`, `on-close`)
  plus the `<command>:<moment>` extension mechanism are sufficient to express everything the
  architecture layer currently does inline in core (promotion, decay, the `review`-time
  restatement check) — not verified against `marathon-roadmap`'s actual implementation,
  since this note was written without reading that extension's source.
- `personal-agents` (a documentation-only, single-purpose repository) is a real example of a
  project that doesn't want this layer, not a one-off exception — the motivating case for
  writing this concept down instead of treating it as a one-time workaround.

## 2. Scrutinize context curation harder

### The problem

`references/context-engineering.md` already states the right principle — "the built work
wins," "`context/` is only for things that work can't express yet," "add detail late" — but
gives no concrete, actionable check to run *before* writing a note into `context/`. The
three named operations that keep `context/` accurate (`Promote`, `Decay`, `Cull`) are all
reactive: they clean up duplication and stale claims after the fact, once the built work
catches up or a concept is superseded. Nothing in core currently guards the moment of
*writing* a note in the first place — the point where duplication actually gets created.

This surfaced concretely during `personal-agents`' `init`: a repository with substantial
pre-existing documentation (`setup/`, `models/`, `research/`, `capabilities/`, all written
before marathon was adopted) makes the pull to re-explain that content inside
`context/design/` or a verbose `context/README.md` capability map very easy to fall into,
especially while settling founding decisions and wanting each capability-map entry to feel
complete. The existing rules ("Deciding where something goes") would eventually catch this
at the next `review` — a design note that only restates a repository decays — but that's a
cleanup pass after the duplication already happened and already cost context budget in every
session between `init` and that `review`.

### Proposed addition

Name the discipline explicitly, as a fourth operation alongside `Promote`/`Decay`/`Cull`, or
as an amendment to "Deciding where something goes" — a check a session runs *before* writing
anything into `context/design/`, `context/concepts/`, or the capability map:

> Does this fact already have exactly one home outside `context/`? If yes, write a pointer
> to it (a path, a citation), never a restatement of it.

Concretely, this reads on the same axis `context-engineering.md` already uses (settled vs.
unsettled, expressed vs. not-yet-expressed) but adds a companion axis: *does this already
have a home elsewhere in the repository, settled documentation included, not just the built
work in the strict code sense.* On a `context` project — where "the built work" is prose,
not code — this distinction matters more than it does on a `code` project, since a whole
pre-existing documentation tree can sit right next to `context/` without being "code" in the
sense the current rule seems to assume.

### Worked example from this session

`personal-agents`' `init` practiced this discipline directly, worth citing as the pattern
once this is written up properly:

- `context/roadmap.toml` tasks cite `research/framework-desktop-followups.md` by path
  instead of repeating its reasoning inline.
- The capability map in `context/README.md` reduces each capability to a name, a status
  (done / in progress / not started), and a path — no paragraph-per-capability, since the
  paragraph already exists in `setup/`, `models/`, or `research/`.
- `context/design/` was left empty at `init` rather than seeded with anything, because
  nothing currently lacks a home in the existing doc tree — an empty tier is a legitimate
  outcome of applying the check, not an oversight to fill in later.

### Open questions

- Does this become a named fourth ledger operation (alongside Promote/Decay/Cull in the
  reset file's Disposition vocabulary), or is it purely a pre-write check with no ledger
  entry of its own, since "I chose not to write something" leaves nothing to record?
- Is this specific to `context` projects with pre-existing documentation, or does it apply
  equally to `code` projects (where the equivalent might be: does this fact already live in
  a doc comment, a README, an ADR the code references)?

### Assumptions this rests on

- The gap is real and not just this session being unusually careful — untested against
  other marathon projects with less pre-existing documentation, where the pull toward
  restatement may be weaker to begin with.

## 3. Don't author authoritative context ahead of a proven shape

Surfaced by a separate agent in another marathon session, relayed here rather than
independently derived in `personal-agents`. No worked example from this repository —
record the principle as stated and let it be checked against a real case before it's
written into core.

### The principle, as relayed

Don't author durable, authoritative context — a skill, a settled design note — against
infrastructure that hasn't proven its shape yet. A contract that's just been designed and
not yet built against is still unstable: the first real consumer often forces a change, and
a skill written ahead of that discovery has to be rewritten before it's taught anyone
anything. Authoritative context earns its authority from use. Write it once the shape has
held under a real consumer, not in the same session that designed the shape.

### Where this sits against current core guidance

`references/context-engineering.md`'s `Promote` operation already gates on settledness — "a
decision fixed it, the built work proved it out, or an experiment produced a result" — and
"the built work proved it out" already covers most of this principle for `design/` notes.
What's less covered:

- **Skills specifically.** `context-engineering.md` names skills as "advanced context" but
  never states when authoring one is premature, the way `Promote`'s settledness bar does for
  `design/` notes. A skill is exactly the artifact this principle warns about — it teaches a
  shape to whoever reads it next — and nothing currently guards against writing one in the
  same session that designed the interface it documents.
- **The architecture layer** (or its extracted-extension successor, concept 1 above):
  "generalized past one repository" is a stronger bar than "proven under one real consumer,"
  but the same failure mode applies one level up — promoting to architecture before even a
  single repository's built work has exercised the shape.
- **No explicit session-timing rule.** The closest existing language, "Add detail late," is
  about how much to write, not about timing relative to a consumer — whether to write
  authoritatively at all yet. This principle reads as a session-sequencing rule: design the
  shape in one session (or as a concept), let something real consume it, and only then — in
  a later session — promote the note or author the skill.

### Proposed addition

State it as an explicit rule near `Promote`, distinct from the settledness criteria already
there: a design note or skill documenting an interface, contract, or shape is not promoted
to `design/` (or authored as a skill) in the same session that designed the shape. Promotion
waits for a session where something real — a caller, a build, an experiment — exercised the
shape and it held. Until then, the shape stays in `concepts/`, explicitly provisional, even
if it reads as finished.

### Open questions

- What counts as "a real consumer" precisely — one call site, or the class of situations the
  interface is meant to cover? Too loose a bar reintroduces the problem this principle
  guards against; too strict a bar could stall legitimate work where a shape genuinely has
  one caller by design.
- Does this apply equally to a `code` project's `design/` notes and a `context` project's
  skills, or does the skill case deserve a stricter rule of its own, given how much more a
  skill "teaches" compared to how a `design/` note only informs the next session?
- Is there a concrete pre-write check here, the way concept 2 proposes one ("does this fact
  already have a home elsewhere")? The analogous question would be: "has this shape been
  exercised by anything other than the session that designed it?"

### Assumptions this rests on

- Relayed secondhand from a separate agent's session, not independently verified against a
  concrete failure case in `personal-agents` — worth confirming against a real example
  before it's written into core, not just recorded on report.

## 4. Staged execution's delegation call isn't a stated decision

### The problem

`references/staged-execution.md`'s "Executing a stage" step 1 says a session "may" hand the stage to a
declared agent (`behavior/delegation.md`) — an option the stage loop never forces anyone to actually
weigh. Nothing before Execute asks the session to state its call, so delegation stays silently
skippable: a session can work through every stage of a plan without once considering whether a declared
agent fits better than doing the work directly, and nothing in the loop or the stage report would show
that the question was ever asked.

### Proposed addition

A structural fix: add a checkpoint before Execute in the stage loop requiring the delegation call be
stated as a visible line in the stage report — delegate, and to whom, or not, and why — the same way
SETTLE already requires stating an escalation before engaging it (`mechanics/pipeline.md` 3 · SETTLE
step 2). This turns a silently-skippable option into a decision that has to be made out loud, without
the skill needing to know anything about what agents a given project declares — `marathon.toml` stays
the source of truth for that, same as today.

### Open questions

- Where exactly the checkpoint sits in the "Executing a stage" numbered list — before step 1 (Execute),
  or folded into step 1's own wording.
- Whether the stated line belongs in the stage report only, or also wants a place in the reset file's
  Disposition when a stage hands off mid-flight.

### Assumptions this rests on

- Found live during a real marathon session's staged execution (the auth-strategy step, standards-lab
  workspace), not a hypothetical — the session had a stage a declared agent plausibly fit, and nothing
  in the loop asked the question.
