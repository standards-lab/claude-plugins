# Checkpoint execution and the closing branch review

Captured 2026-09-22 from a workflow-refinement planning session at the coordinator
(`standards-lab`). The workspace has matured past the point where per-stage architect review pays
for itself: the architecture is established, planning already runs deep, and stopping after every
stage to review code still in churn costs more than it catches. The architect wants execution to
run autonomously against behavioral checkpoints instead, with a holistic code-quality pass
deferred to `close`. This concept is settled direction for the `start` session that implements it,
not an open proposal — SETTLE for that session turns the design below into a stage list.

## The gap

`references/staged-execution.md`'s current "Executing a stage" stops every stage uncommitted and
waits for the architect's approval before it commits, on the reasoning that a misstep never
spreads. That granularity is now the exhausting part: the architect would rather validate behavior
— does the result run and do what it should — than referee code that is still moving, and would
rather scrutinize the finished branch once, holistically, than one stage at a time.

## Design

**Stages commit autonomously.** A stage still states its delegation call, executes, and runs its
check, but commits as soon as the check passes — no stop, no wait. The stage report becomes a log
entry (`diff --stat`, check result, delegation call), not a gate.

**Checkpoints are the real gate during execution.** The stage list, settled at SETTLE, groups
stages under checkpoints: an observable behavior the architect can run or watch run, never "the
build passed" or "the tests are green." Every stage list carries at least the final validation as
a checkpoint; an intermediate checkpoint belongs wherever a behavior first becomes observable that
later stages build on. A library-only step with no runnable surface earns its checkpoint through a
consumer exercise written for the purpose — an example program, a conformance run — and the report
shows its output. A stage list that wants today's cadence back can still place a checkpoint after
every stage; the stage list is the dial, so no new configuration key is needed.

**Checkpoint outcomes:**

- **Confirmed** — continue to the next group of stages.
- **Adjust** — the behavior is wrong within the checkpoint's own stages; fix forward in a new
  commit naming the checkpoint it corrects, re-run, re-report.
- **Re-plan** — the finding reaches past the checkpoint's stages; same mechanics as today's
  re-plan (architect enters plan mode, SETTLE re-opens for stages k..N, committed stages revert
  only if a finding invalidates them).
- **Interrupt** — new: the architect breaks in mid-run. The session finishes or abandons the
  in-flight stage, reports, and takes one of the three outcomes above. This is what keeps the
  architect able to stop something that looks wrong without sitting at every stage boundary.

**`close` gains a branch-review step**, before "Tend the context": the deferred code-quality pass
against the whole branch. Who performs it and what it produces is `flat-context.md`'s sibling
concept, `report-pattern.md` — see below.

**Reset file.** A handoff's Next-focus records stage position *and* checkpoint position (for
example: `Stages: 5/9 · checkpoint 1 of 3 confirmed · stage 5 committed`). Disposition gains a
**Validated** entry at closeout naming each checkpoint and its evidence — this is also what would
make a future `marathon-sitrep` cheap to build, without building it now.

## The delegated review and report pattern

The architect has already been hand-instructing this on live sessions: code-producing stages
delegated to an execution agent, no pause at stage boundaries, a holistic review from a review
agent once the last stage lands, then a report for the architect's own final review with `[overview]`, `[changelog]`
(critical files, lowest to highest dependency layer, tests excluded), and `[verification]` (how to
run and interact with the result — omitted when there's no runnable surface). This becomes
`close`'s default rather than something specified per session:

- The stages run through marathon's executor profile and the branch review through its reviewer
  profile, with the session choosing each engagement's model (`agent-profiles.md`). The architect
  can still override either per session.
- **The alignment check defaults to ecosystem idiom, not to an assumed architecture.** Marathon
  stays project-agnostic: core carries no assumption that a project has an `architecture/`
  repository or an Elemental-Architecture-style layer. The review checks the result against
  idiomatic, industry-standard practice for the project's language and ecosystem by default, and
  checks it against an explicit architecture or documented best practices only when the project
  actually declares one — this workspace's own `architecture/` repository, surfaced through the
  optional `marathon-architecture` extension or a plain context citation, never a default core
  supplies. This mirrors `behavior/planning.md`'s existing sufficiency rule (prefer the ecosystem's
  standard approach absent a reason not to), applied at review time instead of SETTLE time.
- The reviewer's report is a scratch markdown file, not committed — it's a rendering for the
  architect's convenience, not project knowledge, so it has no home in `context/`.

## Open questions for the implementing session

- Exact report file location and lifecycle (a scratchpad path, deleted after the architect reads
  it, or kept until the next session starts) — small enough to settle in that session's own
  SETTLE rather than fix here.
- Whether the Interrupt outcome needs anything in `mechanics/hooks.md`, or is purely a
  `staged-execution.md` behavior with no hook implication.

## Files this touches

`references/staged-execution.md`, `mechanics/pipeline.md`, `commands/close.md`,
`mechanics/reset-file.md`, `behavior/delegation.md`, and the restated-gate sentences in
`commands/experiment.md`, `commands/plan.md`, `commands/review.md` (replace with a citation).
