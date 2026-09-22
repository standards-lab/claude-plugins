# Changelog

All notable changes to the marathon plugin are documented here. Versions follow
[Semantic Versioning](https://semver.org/spec/v2.0.0.html); dates and release links live on the
GitHub releases the tags cut.

## v0.13.0

### Changed

- **Checkpoints replace per-stage review.** A stage commits once its check passes and logs its
  diff stat, check result, and delegation call. The stage list groups stages under checkpoints,
  observable behaviors the architect runs or watches, with Confirmed, Adjust, Re-plan, and
  Interrupt as outcomes. A context project's checkpoint is a scenario walkthrough.
- **`close` opens with a branch review.** The reviewer profile reviews the whole branch against
  the ecosystem's idiom, or a practice the project declares, and writes a gitignored
  `.claude/report.md` (overview, per-change changelog, verification) that closeout deletes. The
  reset's Disposition gains **Validated**, and a handoff records its checkpoint position.
- **Agent profiles replace the agent catalog.** marathon ships `planner`, `executor`, and
  `reviewer` profiles in `agents/`, none pinning a model; the session states the profile, model,
  and reason before each engagement. `[agents]` and `[workspace.agents]` are removed.
- **`context/` is flat.** `design/` and `concepts/` are gone: every note sits under `context/` and
  states its settledness in its opening lines. A note is Integrated once the built work or its
  documentation expresses it, with lasting reasoning moved into the owning repository's docs.
  The ledger is Integrated, Culled, Retained, Add or sharpen, and Cross-repo; Promote leaves core.
- **Experiments are standalone projects.** `experiment` settles the question and where the
  experiment lives (directory and hosting account), runs `init` there with `[experiment] serves`,
  and catalogs it in the served project. The experiment runs its own sessions in parallel; a
  `plan` session takes in its result and archives its remote. Top-level `experiments/` is gone.
- **Waves.** A Next-focus may name steps safe to run concurrently. Each wave session keeps its own
  record at `context/reset/<slug>.md`, and the last to close folds the wave into one commit.
- **The skill is about a third smaller**, with restated rules and cross-file duplication removed.

### Migrating

Move each `context/design/` and `context/concepts/` note to `context/` with a settledness line, or
integrate, cull, or relocate it; move each `experiments/<slug>/` spike to its own repository and
catalog it; drop `[agents]` and `[workspace.agents]` from `marathon.toml`; and add
`.claude/report.md` to `.gitignore`.

## v0.12.0

### Changed

- **The architecture layer is no longer mandatory** — extracted out of core entirely into the new
  `marathon-architecture` extension. `init` no longer scaffolds an `architecture/` directory or
  asks it as a founding decision; `context-engineering.md` no longer describes the layer;
  `close` and `review` lose their architecture-specific clauses; `configuration.md` carries no
  `architecture` key. A project that never generalizes past itself now enables nothing and
  scaffolds nothing, instead of carrying a directory that never gains a second sentence.
- **Extensions never touch core's config schema** — a new rule in `references/extensions.md`: an
  extension's only foothold in `marathon.toml` is the shared `extensions` list; configuration
  beyond enablement lives in a file the extension owns, that only it reads. Named explicitly
  because `marathon-architecture` is the first extension that needed it.
- **Two new context disciplines** in `context-engineering.md`: a pre-write check before anything
  lands in `context/` (does this fact already have exactly one home outside it? cite it, never
  restate it), and a rule against promoting a design note or authoring a skill in the same
  session that designed the shape it documents — promotion waits for a session where something
  real exercised the shape and it held.
- **Context states current truth only** — `design/` and `concepts/` notes, and an architecture
  page, hold what is true now and never a changelog of how it changed; that accounting belongs
  solely to the reset file's Disposition.
- **The stage-loop delegation call is stated, not silently skippable** — every stage now opens by
  stating whether it goes to a declared technical agent or stays with the session, and why,
  mirroring how SETTLE already requires stating an escalation before engaging it.

## v0.11.0

### Added

- **Delegation** — a session may hand one unit of its own work (a stage's implementation, a
  design decision, a call too consequential to settle alone) to another agent and stays its
  owner; marathon states the contract a delegation runs under and names no agent of its own. A
  project catalogs which agents its sessions may reach for under `[agents]` in `marathon.toml`,
  one sub-table per agent with a `delegation` field describing what it's for and why, in the
  project's own words; a session judges fit against that text, never a fixed role lookup. A
  workspace coordinator's `[workspace.agents]` is the baseline, a project's own table overrides
  it key by key. `behavior/delegation.md` holds the contract; `mechanics/configuration.md` holds
  the schema. SETTLE names where a design decision may go to a declared escalation agent, and the
  stage loop of `references/staged-execution.md` names where a stage's implementation may go to
  a declared technical agent; either way the session reports and commits exactly as it would
  have for work it did itself.

- **Sufficiency before building it directly** — before the stage list is settled, a step that
  plans to develop a solution directly asks whether the problem is already resolved by the
  established, idiomatic approach for the language — the standard library, or a dependency the
  ecosystem already treats as standard. If it is, the step needs an adequate reason to build its
  own instead of adopting that approach. Asked at SETTLE, while the answer still changes the
  plan, not discovered at review. A step that proceeds with its own implementation anyway
  carries its reason into the design note it touches, as a rejected alternative; a step with no
  adequate reason adopts the existing approach instead. Stated in `behavior/planning.md`, named
  at the pipeline's SETTLE step. Found by the SQL strategy's own sufficiency rule
  (`standards-lab/context/design/dsl-driven-services.md` §2.4), which names the marathon harness
  as where the question belongs.

- **A stage's check runs the repository's own tooling** — on a code project, a stage that
  produces source now runs whatever conventions or lint tool the repository already wires into
  its own CI or task runner, over the files the stage touched, alongside the language's build,
  vet, and test. The session runs the tool and resolves its findings; it never restates what the
  tool checks. Stated once in `references/staged-execution.md`'s "What a stage is." First
  consumer: sqlate's `sqlint`, which a consuming repository already runs through its own `mise`
  task or CI step.

### Changed

- **The architecture layer, and documentation as the project's own** — every marathon project
  has an architecture layer: the principles, definitions, and conventions that have generalized
  past one repository, the top of the context lifecycle and the target a design note promotes
  to when the built work cannot express it. A standalone project keeps it as a top-level
  `architecture/` directory, scaffolded by `init`; a workspace keeps it as one repository, named
  by the coordinator's `[workspace] architecture` key, whose tree is the architecture with a
  README as every directory's index. It holds nothing a reader could infer from a repository's
  source, and a page that restates a repository is a defect. `docs/` is now only the project's
  own documentation, the guide over its README, API documentation, and source, with the code as
  its source of truth, written by a documentation step of a `start` session like any other
  built work. Under 0.10.0 the docs tier doubled as the workspace's landing zone, named by
  `[workspace] docs`, a member never grew a `docs/` of its own, and the landing zone documented
  each member's implementation, which drifted with every release. `review` gains the
  promotion-candidates check and `close` lands a generalized design note in the architecture
  repository under **Cross-repo**. Found by the `v1.alignment.docs` session, whose module pages
  restated package documentation and were reverted.

### Removed

- **The `docs` command** — project documentation is built work, so a `start` step writes and
  curates it and `review` flags its drift; the architecture repository is a `context` project
  that `start` and `plan` author. What the command's playbook stated about establishing a
  `docs/` directory is one paragraph of `references/context-engineering.md`. The stage rule in
  `references/staged-execution.md` generalizes to carry it: a stage's unit and check follow what
  the stage produces, source or prose, rather than the project kind alone, so a code project's
  `start` runs a documentation step under the prose rule. The Session values of the reset file
  lose `docs`.

## v0.10.0

### Changed

- **The review gate, for every working session** — a stage is reported with the working tree
  uncommitted, so the diff reads cleanly in the architect's tools; it commits only on the
  architect's approval, who then states whether a `reset` follows. Under 0.9.0 the commit
  preceded the report, and the loop applied to a code project's `start` alone. Now every command
  that executes runs it: `start` on either kind, `experiment`, and the context edits of `plan`,
  `review`, and `docs`. A stage's unit follows the project kind: a compilation unit on a code
  project; on a context project, the smallest set of files that must change together, checked by
  the repository's own consistency script and a coherent read. Validation is per kind, and an
  experiment's is its answer to the question it was settled to answer. Stated once in
  `references/staged-execution.md`; the pipeline gains the invariant, and the reset-file schema,
  `start`, `experiment`, `reset`, `close`, `plan`, `review`, `docs`, and the cross-repo stage
  list follow. Found by the `v1.data.sql.prototype` experiment, whose first three stages were
  committed before review.
- **Experiments live at the workspace coordinator** — in a workspace, every experiment sits
  under the coordinator's top-level `experiments/`, on the coordinator's branch, with the reset
  file's Project line naming the coordinator; a standalone project keeps its own `experiments/`.
  A spike depends on member modules as published versions, never through a replace directive,
  and a change it implies for a member's code reaches that repository only by promotion at
  `close`. The reasons are a new section of `references/workspace-coordination.md`;
  `commands/experiment.md`, the tier list of `references/context-engineering.md`, and the
  reset-file schema notes state the rule. Settled by the same experiment, which moved from a
  member repository to the coordinator for exactly this reason.

## v0.9.0

### Changed

- **Staged execution replaces the implementation guide** — on a `code` project the agent
  implements the settled step directly, in stages, and the architect reviews each stage before
  the next begins. A stage is the smallest change set that leaves one compilation unit green —
  its tests and in-source comments included; the check is scoped to that unit, and the module may
  be red between stages because the sequence is in dependency order. The stage list — unit,
  files, one-line why per entry — is the SETTLE artifact, produced in the planning phase and
  approved before any code changes. Each stage lands as its own commit (`on-commit` fires first)
  with a conversational report: `diff --stat`, the check result, prose only on decisions the plan
  didn't spell out. Review outcomes are approve, adjust, and re-plan (the architect enters plan
  mode with findings; invalidated commits revert first; a revised list from stage k is approved
  like the original). Validation — the whole-module build, full test run, and run-and-verify
  check — runs once after the last stage. The new `references/staged-execution.md` holds the
  loop; `start`, `close`, `reset`, the pipeline, and the reset-file schema (stage position in a
  handoff's Next-focus) are rewritten around it.
- **The human is the architect** — vision, direction, and quality control, exercised where the
  processes name them: the SETTLE approval, the per-stage review, the tending confirmation, and
  the pull request. Each playbook states what it expects of the architect at the moment it
  occurs; the standalone role-boundary reference and the git-blame test are retired, and
  "developer" becomes "architect" throughout the skill. `kind` survives on narrower grounds —
  `code` means a build-and-test loop defines stage boundaries and a validation phase; `context`
  has neither.
- **Cross-repo steps replace `coordinate`** — the fan-out model is retired; an orchestrated
  change across workspace projects runs as one working session whose settled step names the
  repos it touches: one spanning stage list in the coordinator's dependency order, a branch per
  touched repo under the step's shared slug, one close that commits and publishes each.
  `references/workspace-coordination.md` is rewritten around what remains: the coordinator and
  its order map, the workspace-holds-no-context rule, continuity at the single reset, downward
  awareness, coordinator conventions.
- **The reset file is purely ephemeral** — its sole purpose is contextual bootstrapping between
  sessions; durable detail belongs in the context layers. Stated outright in
  `mechanics/reset-file.md`, with two fixes: the file is rewritten each session and git is the
  archive (the "accumulates" wording is gone), and an interrupted cross-repo step is expressible —
  the Project line lists the touched repos, the branches share the step's slug, and Next-focus
  carries whatever bootstrap state resuming cold needs. **Cross-repo** joins the declared
  Disposition ledger vocabulary.
- **Close tends the full context scope** — the tending pass establishes the written context the
  step touched before operating on it: standalone, the project's own `context/`; in a workspace,
  each touched repo's `context/`, the coordinator's notes on the changed capability, the docs
  landing zone pages the change moved out from under, and claims about the changed behavior in
  other member repos' context — a stale claim found anywhere in that scope is a defect the pass
  fixes, recorded under **Cross-repo**.
- **Single-source pass** — the decay rule is stated once, in
  `references/context-engineering.md` with its protective qualifier; `close`, `reset`, and
  `review` cite it instead of restating the lossy short form. The `[workspace]` TOML block is
  printed once (`mechanics/configuration.md`), the five-hook list once (`mechanics/hooks.md`),
  and `init`'s hook constraints live only there.

### Added

- **The skill states its version** — a `Version:` line under the SKILL.md title makes the
  extension compatibility check of `mechanics/hooks.md` executable and installed-vs-source skew
  detectable at runtime.
- **Plugin CI** — a push-triggered workflow at the host checks, per plugin, that the manifest
  version, the top CHANGELOG heading, and the SKILL.md version line agree, that
  `marketplace.json` sources resolve, and that every `@`-pointer and `./`-link inside each skill
  resolves to a real file.

### Removed

- `references/implementation-guides.md`, `references/role-boundary.md`, `context/guide.md`,
  `commands/coordinate.md`, and the coordinated fan-out's consolidated guide. The guide-era
  workflow remains available at the `marathon/v0.8.0` tag.

## v0.8.0

### Changed

- **A single workspace reset at the coordinator** — a workspace maintains one reset file, at the
  coordinator; member projects carry none, and a standalone project keeps its own unchanged. LOCATE
  resolves the coordinator's reset from anywhere in the workspace, `reset` and `close` write it
  there (a commit in the coordinator's repository), and the member-reset routing — including the
  resting-point deletion rule — is retired. The schema gains a **Project** line naming the member a
  workspace record concerns. The new `mechanics/reset-file.md` holds the schema, locations, and
  Status semantics; `coordinate` and `references/workspace-coordination.md` re-anchor continuity on
  the coordinator's record plus each project's open branch.
- **Comment-free implementation guides** — stated outright in
  `references/implementation-guides.md`: the guide's code blocks carry no comments of any kind; doc
  comments and API documentation are the agent's closeout work. The "same as the production code"
  clause that pulled sessions toward documented blocks is reworded.
- **Context tending follows validation** — on a code project, the context edits noted at SETTLE
  wait for CONCLUDE, after the developer has applied and validated the guide; EXECUTE produces the
  guide and nothing else. A context project is unchanged. Stated in `mechanics/pipeline.md`
  (SETTLE, CONCLUDE, and a new invariant), `commands/start.md`, and `commands/close.md`.
- **SKILL.md is an index** — the entry file reduces to front matter, the description, and an index
  into the sub-layers: `@` pointers for always-loaded conventions (`behavior/planning.md`,
  `mechanics/pipeline.md`), `./` links for material consulted on demand. The content it restated
  moves to one home each: the reset schema to `mechanics/reset-file.md`, the `marathon.toml`
  canonical layout to the new `mechanics/configuration.md`, the planning philosophy (one step at a
  time, planning is half the work) to the new `behavior/planning.md`, and the hook table lives only
  in `mechanics/hooks.md`, summarized in prose by `references/extensions.md`.
- **Voice is out of scope** — the skill no longer defines or references a writing voice.
  Communication style is identity-level behavior and belongs to user-scoped configuration;
  `behavior/voice.md` is removed and the skill's `behavior/` tier holds workflow conduct only.

### Added

- **Value-of-information test** — `experiment` settles, alongside the question a spike answers, the
  decision the answer changes; a spike that changes no decision isn't run. `plan` weighs
  uncertainty against consequence when choosing the next focus.
- **Assumption annotations** — a `design/` or `concepts/` note names the unverified assumptions it
  rests on, and the reset Disposition records a falsification, so a surprise invalidates identified
  notes instead of forcing a judgment sweep. Stated in `references/context-engineering.md`.
- **Risk-first spiking** — builds proceed in dependency order while `experiment` probes the
  highest-consequence unknown out of band. Stated in `behavior/planning.md`.

## v0.7.0

### Added

- **Mechanics tier** — a new `mechanics/` directory holds the skill's execution specs, written for
  agent action. `mechanics/pipeline.md` defines the five-stage session pipeline every command runs —
  locate, start, settle, execute, conclude — and how each command layers into it; the nine command
  playbooks are slimmed to the content of their stages. `mechanics/hooks.md` is the hook firing spec:
  resolution of the enabled extension set, the firing table, and per-command ordering constraints.
- **Generic extension system** — the platform-tracker hook points (`on-init`, `on-session-start`,
  `on-commit`, `on-closeout`) become five universal hooks bound to the pipeline stages: `on-start`,
  `on-execute`, `on-commit`, `on-reset`, `on-close`, each firing before the moment it names. An
  extension is a separately installed skill; a repository enables it with an `extensions` key in
  `.claude/marathon.toml` — `[project]` for itself, `[workspace]` at a coordinator for every member —
  replacing `init`-time selection. `references/extensions.md` documents the system: installed vs.
  enabled, the extension's SKILL.md declaration, artifact bootstrapping, and the source-of-truth rule
  restated for repo-native extensions (an extension may own its artifact inside the repository;
  anything projected outward is a read-only mirror). `references/extension-hooks.md` is retired.
- **Behavior tier** — a new `behavior/` directory holds always-active conduct, mirroring the
  user-scope `behavior/` convention. The voice standard moves there
  (`references/writing-voice.md` → `behavior/voice.md`) and loads with the skill through an `@`
  reference in SKILL.md's new Behavior section, instead of waiting to be consulted; the standard now
  names its intent — natural language rooted in proper American English grammar, free of the patterns
  identifiable as machine-generated prose.

### Changed

- **SKILL.md reorganized around the sub-structures** — new Behavior and Mechanics sections initialize
  those tiers directly below the direct skill context; the Extensions section replaces Extension
  hooks.
- **`docs/` articulated as the standardized tier of context** — the lifecycle reads
  `concepts/` (volatile) → `design/` (settled) → `docs/` (standardized convention), with the
  established rule unchanged: a `docs/` page itself never decays. Stated in
  `references/context-engineering.md` and `commands/docs.md`.

## v0.6.0

### Added

- **Workspace docs centralization** — a project in a workspace no longer bootstraps its own `docs/`;
  documentation centralizes in one landing-zone project, named by a new optional `[workspace] docs`
  field in the coordinator's `marathon.toml` (an order key, alongside `role` and `order`). `docs` checks
  which case applies — standalone, this project is the landing zone, or another project is — before it
  bootstraps or curates anything. A standalone project is unaffected. `context-engineering.md`'s Decay
  rule gains a second target: a `design/` note also decays once a landing-zone page expresses it, not
  only when the code does. A repository whose own convention narrows or adds to a linked landing-zone
  page records that addition beside the link instead of duplicating the page.

### Fixed

- **Workspace entry routing** — `plan`, `start`, and `experiment` located their project by reading
  `context/reset.md` first, which left a workspace root (no `context/` of its own) an unhandled case
  discovered only by the absence turning up empty. A new `SKILL.md` section, "Finding your project,"
  states the two directory kinds — project and workspace root — as the first check every session-opening
  command makes, and the three commands now point to it ahead of their reset-file logic.

## v0.5.1

### Changed

- **Specificity in the writing voice** — `references/writing-voice.md` adds the discipline of naming
  what is actually happening in a detail rather than reaching for a stock noun or verb, a second
  test (does each term belong to the vocabulary of the discipline it describes?), the
  colon-versus-semicolon rule, and the stock-vocabulary habit in the illustrative aside.
- **Skill prose swept for stock vocabulary** — the commands and references now use the terms their
  subjects already have: the session works on a *step*, the capability map lists *capabilities*,
  the guide preamble states *what the change does*, `marathon.toml` has a canonical *layout*, and
  the role boundary is drawn per kind rather than coming in "shapes".

## v0.5.0

### Added

- **Resting point and workspace-entry routing** — `close` deletes the reset file when the project's
  deliverable is released, it has no next step of its own, and a workspace coordinator carries
  continuity; a session that finds no reset file defers to the coordinator, and a marathon command
  run from the workspace root resolves the coordinator through `[workspace] role` and routes into
  the member repository named by its reset file's Branch line.
- **`[workspace.paths]`** — an optional table in the coordinator's `marathon.toml` that resolves an
  `order` key living outside the workspace root to a directory on this machine.
- **LICENSE** — the repository is licensed Apache-2.0, recorded in the plugin manifest.

### Changed

- **Writing voice restructured** — `references/writing-voice.md` is principle-led: the goal (what a
  capable technical colleague would write), the discipline, and one applicable test, with the former
  habit list demoted to an illustrative aside. Adds the built/planned tense rule and generalizes the
  godoc section to API documentation.
- **Secondary sessions specified** — `plan` and `experiment` open with the same reset-file Status
  check as `start` and resume their own handoffs; `review` and `docs` close through `close`; the
  Session enum grows to `init | plan | start | experiment | review | docs`; branch creation and
  `on-session-start` are sequenced for every session; the session-loop diagram is redrawn so `reset`
  branches off the step.
- **Experiments are tracked** — spikes commit under `experiments/<slug>/` and merge with the branch
  as the durable record of exploratory work; promotion moves proven work out, and stable context
  never cites the directory.
- **Hook contract table** — `references/extension-hooks.md` states each hook's trigger, firing
  commands, and frequency; `init`'s setup commit and `reset`'s WIP commit fire `on-commit`.
- **Canonical configuration** — `SKILL.md` documents the whole `marathon.toml` in one example:
  `[project]`, `[remote]`, and the optional `[workspace]` and `[workspace.paths]`.
- **Coordinator conventions** — an organization-level coordinator's conventions bind member
  repositories through sessions: a member's `review` consults them, a coordinated fan-out applies
  them, and the awareness rule keeps member repos from citing them.
- **Decay rule refined** — a `design/` note decays only when the built work expresses it and the
  note holds no conceptual or pattern detail beyond it; the reset ledger vocabulary (Integrated,
  Retained) is mapped to the operations.
- **Prose normalized** — the skill corpus is rewritten against the restructured voice standard, with
  root-relative cross-references throughout.

## v0.4.0

### Added

- **Writing-voice standard (`references/writing-voice.md`)** — the voice for every piece of prose the
  agent is responsible for: design notes and concepts, reset files, implementation guides, godoc and
  `doc.go`, in-source comments, prose documentation, profiles, and the skill files themselves. Plain
  technical-documentation voice, concrete nouns, objective implementation detail; a list of habits to
  avoid. godoc keeps its idiomatic form. Cited from `SKILL.md`, the role-boundary and
  implementation-guides references, and the commands that author prose.

## v0.3.0

### Added

- **Project kinds (`code` / `context`)** — a project declares its kind at `init` in
  `.claude/marathon.toml`. A `code` project holds production source the developer owns: `start` drafts
  an implementation guide, the developer applies it, and closeout adds tests and documentation. A
  `context` project *is* context — skills, prose, configuration — which the agent authors directly under
  the developer's review, with no guide and no tests. The role boundary now has two shapes, one per kind.
- **`plan` sub-command** — a planning-and-curation session that touches only `context/`: create and
  refine concepts, settle a design, and decide what the next `start` should focus on. Forward-looking,
  where `review` is the backward-looking drift audit. Lands on a branch like `review`.
- **`experiment` sub-command** — a spike in the isolated top-level `experiments/` directory; results are
  concepts, promoted deliberately at closeout, if at all. (Previously `start experiment`.)
- **Workspace coordination (`coordinate`)** — run one change across several marathon projects that live
  as siblings in a workspace. `coordinate` detects the workspace, reads a coordinator project's declared
  dependency `order` (layered, with adjacent peers as sub-arrays), and fans a session out to each project
  in order, honoring its kind. The workspace holds no context; continuity stays per repository.
- **Plugin README** — a quick-reference `README.md` for the plugin directory.

### Changed

- **Sub-command restructure (breaking).** `start <development|context|experiment>` is replaced by
  top-level commands: `plan`, `start`, and `experiment`. `start` no longer takes a type argument — it
  advances the product one step, resolving by project kind. Older `reset.md` files carrying a
  `Session type:` line are still read.
- **Vocabulary** — the notes in `concepts/` are consistently called *concepts*, not *candidate notes*.
- The role-boundary, session-loop, and context-engineering references are generalized so nothing assumes
  a code-only, single-repository model.

## v0.2.0

### Added

- **Human-oriented `docs/` tier** — an optional top-level `docs/` directory for reference documentation
  written for people, a peer to the agent-oriented `context/`. It shares context engineering's
  maintenance discipline (curate it, keep it in sync) but not its lifecycle: a `context/` note decays
  once the code expresses it, while a `docs/` page exists to explain code that already does. A repository
  opts in at any time; most stay `context/`-only.
- **`docs` sub-command** — the deliberate authoring and curation pass for `docs/`. On a repository with
  no `docs/`, the first run bootstraps the tier; later runs extend and restructure it. Plan-mode-driven,
  no code handoff, landed on a branch like `review`.
- **Docs drift in `review`** — `review` now also flags `docs/` pages that have drifted from the code,
  once a `docs/` tier exists. The core build loop is unchanged; documentation is never cram-written at
  closeout and never silently rots.

## v0.1.0

Initial release — the standalone core workflow.

### Added

- **Concept-driven `init`** — evaluate a project-planning concept in plan mode and align with the
  developer, then scaffold a project's top-level `context/` (orientation, `design/`, `concepts/`, and
  single `guide.md`/`reset.md` files) in one pass.
- **Typed sessions** — `start` a `development`, `context`, or `experiment` session. A fresh session
  plans in plan mode and settles scope with the developer before any implementation guide is written;
  a session resuming from a handoff picks up the prior plan in place.
- **Mid-session handoff** — `reset` captures in-flight state and a resume pointer without closing
  the work, so a fresh context window picks up where the last left off.
- **Closeout** — `close` runs the reset transaction (integrate / promote / cull / retain), decays
  design that code now expresses, deletes the spent implementation guide, commits, and publishes the
  branch (pull request, merge request, or the project's equivalent) with its description from the reset
  file.
- **Drift review** — `review` enters plan mode to audit design-vs-code drift and aligns with the
  developer before culling or promoting context.
- **Context engineering** — a volatile-vs-stable context model with deliberate promotion and decay,
  keeping the repository the single source of truth. Any operation that culls or promotes context
  shows its dispositions for developer alignment first.
- **Role boundary** — the developer owns production source; the agent owns tests, documentation,
  and context artifacts.
- **Remote platform** — the remote and its publish command (`gh pr create`, `glab mr create`, or
  another) are declared at init and stored in `.claude/marathon.toml`; the core's git workflow stays
  platform-neutral.
- **Extension hooks** — named, no-op-by-default hook points (`on-init`, `on-session-start`, `on-commit`,
  `on-closeout`) for opt-in platform project-management extensions selected at init.
