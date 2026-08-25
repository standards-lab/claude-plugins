# Changelog

All notable changes to the marathon plugin are documented here. Versions follow
[Semantic Versioning](https://semver.org/spec/v2.0.0.html); dates and release links live on the
GitHub releases the tags cut.

## [v0.7.0]

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

## [v0.5.1]

### Changed

- **Specificity in the writing voice** — `references/writing-voice.md` adds the discipline of naming
  what is actually happening in a detail rather than reaching for a stock noun or verb, a second
  test (does each term belong to the vocabulary of the discipline it describes?), the
  colon-versus-semicolon rule, and the stock-vocabulary habit in the illustrative aside.
- **Skill prose swept for stock vocabulary** — the commands and references now use the terms their
  subjects already have: the session works on a *step*, the capability map lists *capabilities*,
  the guide preamble states *what the change does*, `marathon.toml` has a canonical *layout*, and
  the role boundary is drawn per kind rather than coming in "shapes".

## [v0.5.0]

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

## [v0.4.0]

### Added

- **Writing-voice standard (`references/writing-voice.md`)** — the voice for every piece of prose the
  agent is responsible for: design notes and concepts, reset files, implementation guides, godoc and
  `doc.go`, in-source comments, prose documentation, profiles, and the skill files themselves. Plain
  technical-documentation voice, concrete nouns, objective implementation detail; a list of habits to
  avoid. godoc keeps its idiomatic form. Cited from `SKILL.md`, the role-boundary and
  implementation-guides references, and the commands that author prose.

## [v0.3.0]

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

## [v0.2.0]

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

## [v0.1.0]

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
