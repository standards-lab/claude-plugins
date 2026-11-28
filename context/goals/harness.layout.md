# goal · harness.layout

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] lint-references
2. [x] standards
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

- plan: rejected `_agent/` for the agent area: Go and the web tools treat both alike, and only the
  dot form is hidden, as `.claude/`, `.state/` and `.build/` are.
- plan: an existing `.state/` becomes the ignored `.agent/state/`.
- plan: the session brief's Verify keeps the check's one-line result; its steps never cite agent
  harnesses or internal test commands, with a good and bad example.
- plan: `archive <repo> <category/path>` is a Go command in marathon's module and works in any
  directory: it records a clone in its parent's `index.md` under a hierarchical path (one heading
  per segment, at most four, matched case-insensitively, created sorted), and for a catalogued
  repository also sets `archived = true` and drops its local path. Rejected the workspace's
  `references.md` as the index: it covers only workspace repositories.
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

## Pending edits

- coordinator · roadmap/harness.layout.toml: drop "Releases marathon." from the
  lint-references and marathon-layout summaries; `archive` keeps the goal's one marathon release.
- coordinator · roadmap/harness.layout.toml: end the goal's summary with "The architecture's
  agent-separation principle and the Go and Web Elemental standards state the layout." in place
  of "The goal record's Decisions hold the layout."
- coordinator · roadmap/harness.layout.toml: task `web`'s summary adds that a package's `demo/`
  becomes `examples/` (web-components and web-foundation have a `demo/`).
