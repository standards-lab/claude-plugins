# Tool-based skills: the tooling layer beneath the plugins

The tooling layer is planned work, tracked as `v1.harness.tooling`. It starts from marathon's four
Go tools, in the [`marathon/` module](../marathon/), which repositories pin through mise. The rest
of marathon's procedure, and its extensions', is prose, and no schema or `spec/` directory exists.

## The idea

A skill encodes judgment: when to act, why, and what a result means. A tool encodes procedure:
the schema a document must satisfy, the check a change must pass, the transformation a step
performs. Procedure written as prose in a skill costs context on every load and drifts from what
runs; a tool is tested, versioned, and the same on every call. So the plan is:

- A tool carries the checks, as `lint`, `schema`, or `validate` commands. A tool that reconciles
  a document against an external system offers `plan` to show the difference and `apply` to
  perform it.
- The skill says "run the tool and resolve its findings" and never restates what the tool checks.
  A skill that grows a procedural section has found a candidate tool.
- A domain is built in three layers, each depending only on the one below: the specification (the
  document formats, as language-neutral schemas in a versioned `spec/` directory), the tooling
  that makes it operable, and the skill on top. Code validates against a schema and never defines
  it.
- An extension invokes the tooling's commands and never re-encodes them. Behavior the tooling
  lacks becomes a new command.
- Custom code stays limited to scaffold, lint, schema, plan, apply, and export. Anything else is
  configuration, an adapter, or a document.

The first consumers are marathon's roadmap manifest and goal records, whose formats the skill
documents in prose and `lint-context` checks only for shape and budgets, and
marathon-extraction's deterministic checks (`marathon-extraction.md`, "Tooling").

## Deferred guards

`v1.harness.tooling` takes up two guards that `harness.optimization` deferred:

- A hook that refuses a chained `gh pr merge` or `git push`.
- A required `check` status on claude-plugins' main, once goal-record bookkeeping lands by pull
  request. With the admin bypass it may not stop an admin's merge; without one it rejects
  marathon's bookkeeping pushes to main.

## Assumptions

- The manifest and goal-record formats stay stable enough to fix as schemas.

## Open questions

- Whether the schemas live in this repository or in a specification repository of their own,
  with instances pinned to a version.
- Whether structured, repetitive generation, such as a rollup to a fixed template, moves to a
  local model behind a schema-validated tool, and what would justify it.
- How the first slice is cut: one document type and its full command loop is the working
  assumption.
