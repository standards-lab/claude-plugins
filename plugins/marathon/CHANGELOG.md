# Changelog

All notable changes to the marathon plugin are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
