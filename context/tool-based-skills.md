# Tool-based skills: the tooling layer beneath the plugins

The tooling layer is planned work, tracked as `v1.harness.tooling`. No code expresses it yet:
marathon and its extensions carry their procedure in prose, and no plugin calls a tool of its
own.

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
documents only in prose today, and marathon-extraction's deterministic checks
(`marathon-extraction.md`, "Tooling").

## Assumptions

- The manifest and goal-record formats stay stable enough to fix as schemas.
- A tool can ship inside a plugin and run wherever the plugin is installed.

## Open questions

- The language and packaging of the tools, and how a plugin declares and pins them.
- Whether the schemas live in this repository or in a specification repository of their own,
  with instances pinned to a version.
- Whether structured, repetitive generation, such as a rollup to a fixed template, moves to a
  local model behind a schema-validated tool, and what would justify it.
- How the first slice is cut: one document type and its full command loop is the working
  assumption.
