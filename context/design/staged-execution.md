# marathon v0.9.0 — staged execution replaces the implementation guide

Settled at the 2026-08-31 workspace retrospective; `backlog.marathon-staged-execution` in the
coordinator's roadmap cites this note. Per `skill-is-the-source-of-truth.md` this note records
intent the skill files do not yet capture and is removed once the v0.9.0 session expresses it.
Everything below was settled in planning; §"Changes by file" is the change list in the order to
apply, and §"Folded retrospective items" carries the additional scope the retrospective folded
into the same session.

## Decision record

- **The implementation guide is retired.** On a `code` project the agent implements directly,
  in stages, and the developer reviews each stage. The developer owns direction and quality;
  the pull request and the per-stage review are the ownership seams, as they already are on a
  `context` project.
- **The git-blame test is retired** with the guide. Its rationale (fluency by applying every
  change) is replaced by a different mechanism: a review gate per stage, directed at decisions
  rather than transcription. `role-boundary.md` says so explicitly rather than leaving the old
  argument standing.
- **`kind` survives on narrower grounds.** A `code` project has a build-and-test loop that
  defines stage boundaries and a validation phase; a `context` project has neither. The role
  boundary is now the same on both kinds.
- **Stage** = the smallest change set that leaves one compilation unit green — the package, in
  Go — with the stage's tests and in-source comments included. The check is scoped to that unit
  (`go build`, `go vet`, `go test` on the package path), not the module. The module is allowed
  to be red between stages; the stage sequence is in dependency order so every broken caller is
  a later stage.
- **The stage list is the SETTLE artifact.** It replaces the guide as the thing approved before
  any code changes. Each entry: the stage's compilation unit, the files it touches, one line on
  why.
- **One commit per stage**, firing `on-commit`. The commit message carries the stage's decision
  lines. A finding at stage k that invalidates earlier stages is handled by reverting those
  commits.
- **Stage report is conversational**, not a file: `diff --stat`, the build/test result, then
  prose only on decisions the plan didn't spell out and anything the agent is least sure of.
- **Review outcomes**: approve and adjust happen in conversation. **Re-plan** is the
  convention: developer switches to plan mode and sends findings; the agent re-enters SETTLE
  for stages k..N, reverting invalidated commits first, and produces a revised stage list
  approved the same way.
- **Validation** = the whole-module build and full test run (`./...`), plus the run-and-verify
  behavior check the guide used to end with. It runs once, after the last stage is approved.
- **Closeout** loses "finish the agent's part" and "drop the guide." Context tending still
  follows validation.
- **No config switch** to keep the guide mode. Anyone who wants it has the `marathon/v0.8.0`
  tag.

## Changes by file

All paths are under `plugins/marathon/` unless noted.

### `skills/marathon/references/implementation-guides.md` → delete

Replace with a new `skills/marathon/references/staged-execution.md` (below). Nothing from the
old file survives verbatim; two of its rules move: "code carries the what, prose carries the
why" becomes the shape of the stage report, and "end with how to run and verify" becomes the
validation phase.

### `skills/marathon/references/staged-execution.md` → new

Sections, in this order:

1. **What a stage is.** The definition from the decision record. State it language-neutrally:
   "the smallest unit the language compiles on its own — a package in Go, a crate in Rust, a
   module in Python" — and that a stage includes its tests and in-source comments, so nothing
   is left invalidated for later. The green check is scoped to that unit only; the module may
   be red between stages.
2. **The stage list.** Written at SETTLE, approved before execution. Dependency order, lowest
   first, so a stage that changes an exported interface is followed by the stages that consume
   it. Each entry: unit, files, one-line why. The list is the plan's artifact and is not
   committed anywhere — it lives in the conversation and, on handoff, in the reset file's
   Next-focus.
3. **Executing a stage.** Implement, run the unit-scoped build/lint/test, fix until green,
   commit with the stage's decision lines in the message (`on-commit` fires first), then stop
   and report. Never run ahead into the next stage.
4. **The stage report.** Conversational. `diff --stat`; the check result; then prose only on
   the decisions the plan didn't spell out and the parts the agent is least confident of. Code
   carries the what; the report carries the why. No restating what the diff shows.
5. **Review outcomes.** Approve → next stage. Adjust → edit in place, re-run green, amend or
   add a commit, re-report. Re-plan → the convention: developer enters plan mode and sends
   findings; the agent revisits stages k..N; landed stages stay landed unless a finding
   invalidates them, in which case reverting those commits is the first act of the re-plan; a
   revised stage list from k onward is approved like the original, then execution resumes at k.
6. **Validation.** After the last stage is approved: whole-module build and full test run, then
   the run-and-verify behavior check — the concrete commands and what to look for. Don't close
   on a failure.
7. **Stay within the step.** Carried over: no opportunistic refactors, no unrelated cleanups;
   anything else is a concept or a future step.

### `skills/marathon/references/role-boundary.md` → rewrite

- Retitle the code-project section. Replace "the developer owns the production code" and the
  git-blame paragraph with: on a `code` project the agent implements the whole change —
  production code, tests, comments, docs, context — in stages, and the developer owns the work
  by directing it and reviewing every stage before the next begins. Ownership is in the review
  gate and the pull request.
- Remove the "everything around the code" list; the split it described no longer exists.
- Context-project section: unchanged in substance; drop the phrase "there is nothing to hand
  off through a guide."
- **The test** section: reduce to the one question that still differs by kind — is there a
  build-and-test loop that defines stages and a validation phase (`code`) or not (`context`)?
  Who authors is no longer kind-dependent.
- **Why split it this way**: rewrite the code-project rationale. Say plainly that the git-blame
  test is retired and why: applying drafted code spent developer attention on transcription;
  the stage review spends it on judgment, and staging keeps any one review small enough to
  catch a detail before it derails the rest. Keep the sentence about tests and docs being the
  first things dropped under pressure — that's still why the agent writes them in the same
  stage as the code.
- Keep "Keep it language-neutral."

### `skills/marathon/commands/start.md` → rewrite the Execute section and the resume note

- Line 9 area: a `start` resumed on a code project picks up the stage list recorded in the
  reset Next-focus; there is no `context/guide.md`.
- **Settle**: add that on a code project the settled scope is expressed as the stage list
  (`references/staged-execution.md` §2), approved before the branch is created.
- **Execute**, code project: replace the guide bullet with the stage loop — implement stage,
  unit-scoped green, commit, report, stop for review; repeat; then validation. Point at
  `staged-execution.md` §3–6. Keep: context tending waits for closeout.
- **Execute**, context project: drop "there is no guide and no code handoff"; the rest stands.

### `skills/marathon/commands/close.md` → cut steps 1 and 5, renumber

- Delete **1. Finish the agent's part**. Replace with a one-paragraph precondition: on a code
  project, `close` assumes validation (`staged-execution.md` §6) has passed; on the other
  kinds, the existing sentences about what "validated" means for context, plan, experiment,
  review, docs sessions stay.
- Step 2 (tidy the notes): change "they waited here for the applied code to validate them" to
  "they waited here for the validated stages."
- Delete **5. Drop the guide** entirely.
- Renumber: precondition, tidy, agree next step, record, commit and publish.

### `skills/marathon/commands/reset.md` → extend step 2

In **Write the reset file**, add: on a code project's `start`, Next-focus carries the stage
list, which stage the pointer is on, and whether that stage is implemented-unreviewed or
reviewed; after a re-plan it carries the revised list. Step 3 (WIP commit) becomes: stage
commits already exist on the branch; a WIP commit is only for an in-progress stage.

### `skills/marathon/mechanics/reset-file.md` → stage position, plus the simplification below

Under **Status semantics**, `handoff`: note that a code-project `start` handoff records stage
position in Next-focus, e.g. `Stages: 3/7 · stage 3 implemented, unreviewed · list: …`. Kept
prose, not a new field. (This file also takes the ephemeral-reset simplification in
§"Folded retrospective items".)

### `skills/marathon/mechanics/pipeline.md` → two edits

- SETTLE step 3: replace "EXECUTE produces only the unvalidated guide, and the notes record
  what the applied code proved" with "EXECUTE produces the staged commits, and the notes
  record what validation proved."
- Add to **4 · EXECUTE**: on a code project's `start`, EXECUTE is the stage loop of
  `references/staged-execution.md`; a re-plan re-enters 3 · SETTLE for the remaining stages
  without leaving the branch.
- Invariant "On a code project, `context/` asserts only what validated work proved" —
  unchanged.

### `skills/marathon/mechanics/configuration.md` → project kind bullets

- `code`: "The agent implements each step directly in stages, each reviewed before the next;
  closeout follows validation."
- `context`: drop "no implementation guide"; keep "no tests."
- Replace the git-blame sentence ("When in doubt, apply the git-blame test…") with: when in
  doubt, ask whether there is a build-and-test loop; if there is, it's `code`.

### `skills/marathon/commands/init.md` → two passages

- Project-kind bullet (~line 24): remove "whether sessions hand off an implementation guide"
  and the "whose hands should the deliverable's history show" sentence; replace with the
  build-and-test-loop question.
- Scaffold section: `CLAUDE.md` role-boundary line for a `code` project becomes "the agent
  implements in reviewed stages; the developer directs and reviews." Delete the "Don't create
  `context/guide.md`" bullet.

### `skills/marathon/commands/coordinate.md` → drop the consolidated guide

- Step 3: replace "produces an implementation-guide slice … consolidated guide … `guide.md` at
  the workspace root" with: on a code project the plan produces that project's stage list; the
  collective plan is the set of stage lists across projects, held in the conversation, never
  written to the workspace.
- Step 4, code project bullet: run that project's stage loop and validation; the developer
  reviews each stage as in a standalone `start`.
- Step 5: remove "the ephemeral guide is discarded."

### `skills/marathon/references/workspace-coordination.md`

- **Honoring project kind in a fan-out**: a `code` project runs its stage loop; a `context`
  project is authored directly.
- Delete **The consolidated guide is ephemeral**. If a replacement sentence is wanted: the
  collective plan is stage lists in conversation; the lasting record is each project's stage
  commits and the coordinator's reset.

### `skills/marathon/references/context-engineering.md`

Remove the `context/guide.md` bullet from the volatile tier.

### `skills/marathon/mechanics/hooks.md`

The example moment `start:guide-written` no longer names a real moment. Replace with
`start:stage-committed` (still undefined — the sentence says none are defined; only the
example changes).

### `skills/marathon/behavior/planning.md`

Line 37: "a fresh `start` before writing the guide or authoring the change" → "a fresh `start`
before implementing the first stage or authoring the change."

### `skills/marathon/SKILL.md` → References index

- `role-boundary.md` entry: drop "the git-blame test"; describe as "how the work divides
  between developer and agent, and what still differs by project kind."
- Replace the `implementation-guides.md` entry with `staged-execution.md`: "how a code
  project's `start` executes: stages, the stage list, the report, review outcomes, re-plan,
  validation."

### `README.md` (plugin)

Project kinds section: `code` → "the agent implements in reviewed stages, each a green
compilation unit with its tests; the developer directs and reviews every stage." `context` →
drop "there is no implementation guide."

### `.claude-plugin/plugin.json`

`version` → `0.9.0`. Description: replace nothing structural; optionally add "staged execution
with per-stage review" after "plan/start/experiment sessions."

### `CHANGELOG.md` (plugin)

New top entry covering: staged execution replacing the implementation guide (the decision
record above, condensed); the role boundary uniform across kinds with the git-blame test
retired and `kind` deciding only the build-and-test loop; the removals
(`references/implementation-guides.md`, `context/guide.md`, the coordinated fan-out's
consolidated guide) with the previous workflow available at `marathon/v0.8.0`; plus the folded
items below.

### Host repo: `CLAUDE.md` (repo root)

Role boundary section, line 19: drop "There is no implementation guide and" → "There are no
tests." The rest already describes the model code projects now share.

### Host repo: `context/`

Normal `close` tending: any `design/` or `concepts/` note about the guide, the role boundary,
or the git-blame test decays or is culled — this note included; record it in the reset
Disposition.

### Member repos (beyond the original change list)

Five sibling repos restate the guide-era role boundary verbatim and take the same rewrite:
`go-core/CLAUDE.md`, `go-database/CLAUDE.md`, `go-web-sdk/CLAUDE.md`,
`go-web-sdk-template/CLAUDE.md`, `go-web-service/CLAUDE.md` — each carries "The developer owns
the production Go source — they apply it and answer for it. The agent writes everything else:
… the implementation guide, and the reset file." Their `.claude/marathon.toml` kind comments
repeat it ("The agent drafts changes as an implementation guide"). Per the
context-architecture principle (standards-lab `design/context-architecture.md`), the rewrite
replaces the restatement with a link to the skill's single role-boundary statement rather than
a new restatement.

## Folded retrospective items

Additional v0.9.0 scope settled at the retrospective, in the same session:

1. **The reset file is simplified, not extended.** It is a purely ephemeral artifact: one file
   — a standalone project's `context/reset.md`, or in a workspace the coordinator's — whose
   sole purpose is contextual bootstrapping between sessions. Durable detail belongs in the
   context layers, never the reset. Within that frame, fix the two defects found:
   `commands/coordinate.md` claims the session record "accumulates" at the coordinator while
   `mechanics/reset-file.md` says the file is rewritten each close — git is the archive; say
   so. And an interrupted fan-out is unrepresentable (single Project/Branch lines) — the
   handoff of a coordinated change must be expressible as bootstrap state.
2. **The single-source pass over the skill** (context-architecture principle). The decay rule
   is stated once, in `references/context-engineering.md`, with its protective qualifier ("a
   note that still explains a pattern, a boundary, or a style neither the code nor a page can
   state on its own is doing design work and stays"); `commands/close.md`, `commands/reset.md`,
   and `commands/review.md` cite it instead of restating the lossy short form. The
   `[workspace] role/order` TOML block is printed once (it appears in
   `mechanics/configuration.md`, `commands/coordinate.md`, and
   `references/workspace-coordination.md`); the five-hook list likewise (it appears in
   `mechanics/pipeline.md`, `mechanics/hooks.md`, and `references/extensions.md`). The
   extension's `Cross-repo` disposition entry, which the live reset already uses, joins the
   declared ledger vocabulary in `mechanics/reset-file.md`.
3. **The coordinate × roadmap on-close rule.** A coordinated change runs one task across
   several project closes; today marathon-roadmap's `on-close` ("delete the finished task;
   advance `next`") would fire at every one. The rule: the coordinated change's final close
   deletes the task and advances `next`; earlier closes cite the task. Lands in
   `commands/coordinate.md` and marathon-roadmap's `mechanics/on-close.md` (a small
   marathon-roadmap version bump alongside).
4. **The skill states its version.** A `version:` line in `SKILL.md` makes
   `mechanics/hooks.md`'s extension compatibility check executable (marathon-roadmap declares
   "Targets: marathon 0.8" against nothing today) and makes the installed-vs-source skew
   acknowledged in `CLAUDE.md` detectable at runtime.
5. **Plugin CI and first evals.** A push-triggered workflow checking: `plugin.json` version ==
   top CHANGELOG heading (== tag at release), `marketplace.json` sources resolve, and every
   `@`-pointer and `./`-link inside each skill resolves to a real file (a dead link currently
   fails silently at model-read time). Then one eval fixture per command as a regression
   floor — `skill-creator` is available in the harness.

## Verification before close

- `grep -rn -i "guide" plugins/marathon CLAUDE.md README.md` returns only CHANGELOG history
  entries.
- `grep -rn -i "blame" plugins/marathon` returns only the CHANGELOG.
- Every `./` link in `SKILL.md` resolves to a file (now also enforced by the new CI).
- Read `start.md` → `staged-execution.md` → `close.md` end to end as one path and confirm no
  step assumes a file the workflow no longer creates.
- The member-repo `CLAUDE.md` sweep leaves no restated role boundary — links only.

## Closeout

Commit, PR, tag `marathon/v0.9.0` after merge (marathon-roadmap's bump tags with it). Docs
landing zone (`standards-lab/docs`): update whatever page describes the marathon code-project
loop to match the README's project-kinds paragraph.
