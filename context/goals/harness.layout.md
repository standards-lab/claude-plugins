# goal · harness.layout

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [ ] lint-references
2. [ ] standards
3. [ ] marathon-layout
4. [ ] archive
5. [ ] templates
6. [ ] capstones
7. [ ] core
8. [ ] docproc
9. [ ] libraries
10. [ ] go-ai
11. [ ] messaging
12. [ ] docinfer
13. [ ] web
14. [ ] webappgen
15. [ ] service
16. [ ] tools

## Decisions

- plan: `.agent/` is the project's agent area in every repository; `.claude/` stays Claude
  Code's. Goal records stay tracked in `context/goals/`. Rejected `_agent/`: Go and the web tools
  treat both alike, and only the dot form is hidden, as `.claude/`, `.state/` and `.build/` are.
- plan: only new agent work goes in `.agent/` (scratch tests, probes, ad-hoc harnesses,
  measurement scripts, generated artifacts); `.state/` becomes the ignored `.agent/state/`.
- plan: Go unit tests move into their package's `test/` (`<pkg>/test/*_test.go`, package
  `<pkg>_test`), fixtures into `<pkg>/test/testdata/`. Rejected a root test tree: it can't import
  `cmd/<x>/internal/` packages or reach nested modules. `export_test.go` hooks and the tests using
  them stay beside the source, the one exception; `package main` tests become subprocess tests in
  the root `integration` tier, which stays.
- plan: web tests stay in `test/` and web builds in `dist/`; a web `tsconfig.test.json` includes
  `.agent/`, and ESLint and Prettier keep checking it.
- plan: Go commands build into git-ignored `bin/` (`go build -o bin/ ./...` where a module has
  mains); libraries keep `go build ./...`.
- plan: a multi-line mise task is a file in visible `mise-tasks/<group>/<member>`; one-liners
  stay inline. Agent tasks live in `.agent/mise-tasks/`, loaded only by `mise.agent.toml`
  (`[task_config] includes = ["mise-tasks", ".agent/mise-tasks"]`), with `MISE_ENV=agent` in each
  repository's `.claude/settings.json` `env`. The check's sub-steps are agent tasks wrapping step
  scripts; a human `check` runs the scripts directly. Rejected `hide = true` for agent tasks.
  Templates ship no profile beyond `mise.agent.toml`, and `.gitignore` lists `mise.local.toml`.
- plan: the session brief's Verify keeps the check's one-line result; its steps never cite agent
  harnesses or internal test commands, with a good and bad example.
- plan: lint-references finds the coordinator through a tracked `coordinator` key in a member's
  `.claude/marathon.toml`, resolving a first segment in the repository, beside it, then through
  `[workspace.paths]`; CI keeps skipping. Rejected an environment variable and a check flag.
- plan: `archive <repo> <category/path>` is a Go command in marathon's module and works in any
  directory: it records a clone in its parent's `index.md` under a hierarchical path (one heading
  per segment, at most four, matched case-insensitively, created sorted), and for a catalogued
  repository also sets `archived = true` and drops its local path. Rejected the workspace's
  `references.md` as the index: it covers only workspace repositories.
- plan: Go infrastructure's scenario examples are a nested `examples/` module; web examples get a
  web-elemental page of their own. Building them is the `examples` backlog goal.
- plan: each migration takes its repositories' currency; go-web-service adapts go-core v0.13's
  `EventKind.Raise` at its 12 calls. Only the templates and claude-plugins release; library and
  web moves ride their next feature release.
- plan: `capstones` runs while the architect pauses the capstone goals.

## Pending edits
