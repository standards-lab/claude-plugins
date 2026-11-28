# goal · harness.layout

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] lint-references
2. [x] standards
3. [x] marathon-layout
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

- plan: rejected `_agent/` for the agent area: Go and the web tools treat both alike, and only the
  dot form is hidden, as `.claude/`, `.state/` and `.build/` are.
- plan: an existing `.state/` becomes the ignored `.agent/state/`.
- marathon-layout: `archive` is reopened: one central archive reference, in source control in
  one place, sharded so no file outgrows one context, and no per-parent `index.md` scattered across
  the filesystem. `archive`'s PLAN settles its location, sharding, and command shape.
- plan: building the Go `examples/` modules and the web examples is the `examples` backlog goal,
  not this goal's.
- plan: each migration takes its repositories' currency; go-web-service adapts go-core v0.13's
  `EventKind.Raise` at its 12 calls. Only the templates and claude-plugins release; library and
  web moves ride their next feature release.
- plan: `capstones` runs while the architect pauses the capstone goals.
- lint-references: marathon releases once, at `archive`, after the goal's plugin tasks;
  lint-references and marathon-layout leave Unreleased entries.
- standards: this goal's architecture pages run ahead of the code: its criteria commit every
  repository to the layout, and each migration's standards review checks against the pages.
  Rejected writing them after `tools`: every migration would be reviewed against pages that
  forbid it.
- marathon-layout: no layout lint during the migrations; review holds the layout.
- marathon-layout: held claude-plugins' `.claude/briefs/` ignore line until an installed
  marathon writes `.agent/state/briefs/`: 0.24.1 still writes `.claude/briefs/`. The `archive`
  release lifts it.

## Pending edits

- coordinator · roadmap/harness.layout.toml: drop "Releases marathon." from the
  lint-references and marathon-layout summaries; `archive` keeps the goal's one marathon release.
- coordinator · roadmap/harness.layout.toml: end the goal's summary with "The architecture's
  agent-separation principle and the Go and Web Elemental standards state the layout." in place
  of "The goal record's Decisions hold the layout."
- coordinator · roadmap/harness.layout.toml: task `web`'s summary adds that a package's `demo/`
  becomes `examples/` (web-components and web-foundation have a `demo/`).
- coordinator · roadmap/harness.layout.toml: add task `architecture` (the architecture repository
  takes the layout: ignore lines, `mise.agent.toml`, settings `env`; repos = ["architecture"])
  after `archive` in `order`; the record's Tasks gains it at the same position.
- coordinator · roadmap/harness.layout.toml: rewrite `archive`'s summary, and the goal summary's
  "`archive` into an `index.md`", to the central, sharded archive reference.
- coordinator · roadmap: a backlog goal for a layout lint, after `tools`: the ignore lines,
  `mise.agent.toml`, the settings `env`, and Go test placement.
