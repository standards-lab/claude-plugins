# Changelog

All notable changes to the marathon plugin are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
