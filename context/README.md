# claude-plugins

claude-plugins is the Standards Lab organization's Claude Code plugin marketplace. Its plugins
codify the organization's development processes, so every project applies them the same way. The
repository is developed with marathon, so the workflow runs against its own source.

## Capability map

- **marathon** (built): the long-haul development workflow, run as a software factory. The
  architect is involved twice per task, at the task brief and at the session brief, and subagents
  implement, review against standards and the spec, and keep the context current in between.
  Goals and tasks in the roadmap manifest are the units of work. PLAN runs each repository's
  optional `currency` command first, so a task begins by upgrading what trails, and SHIP runs an
  optional release gate before each tag. Its design lives in
  [plugins/marathon](../plugins/marathon/).
- **marathon-architecture** (built): the extension that finds the architecture layer each
  repository's `STANDARDS.md` points to, so the standards-reviewer can apply it. Like every
  extension, it declares the minimum marathon version it requires (marathon's
  `references/extensions.md`). Its design lives in
  [plugins/marathon-architecture](../plugins/marathon-architecture/).
- **harness.optimization** (planned): the active goal that holds marathon and the content it
  produces to budgets by checks. See [harness-optimization.md](harness-optimization.md).
- **marathon-extraction** (planned): carries patterns a consumer proves back to the blueprint.
  See [marathon-extraction.md](marathon-extraction.md).
- **marathon-references** (planned): a proposed extension that holds the references catalog.
  See [marathon-references-extension.md](marathon-references-extension.md).
- **Tooling layer** (planned): deterministic tools the skills call in place of procedure written
  as prose. See [tool-based-skills.md](tool-based-skills.md).
- **Harness testing** (built): `scripts/check.sh` runs on every push. It checks consistency,
  rejects `@`-include lines, resolves file links and citations, and fails a citation directory one
  edit from a real one. It requires one "Read first" line per command playbook, holds each
  command's read-first set and each agent profile to a word budget, caps skill descriptions at
  1,024 characters, and runs `claude plugin validate`. It also runs marathon's `lint-context` from
  the working tree, with its tests. Each failure message names its fix. Before each tag, the
  release gate, `scripts/gate.sh`, runs the check and the plugin's `claude plugin eval` suite.
  marathon's suite has one case, drawn from an observed failure. See [harness-testing.md](harness-testing.md).
- **Marketplace host** (built): the marketplace manifest,
  [.claude-plugin/marketplace.json](../.claude-plugin/marketplace.json), with each plugin
  versioned and released on its own.
- **Further plugins** (planned): the set grows as processes prove worth codifying, such as
  diagram generation, dev-blog infrastructure, and skills for the organization's architecture.

## How this repository works

- Each plugin's design lives in its own files under `plugins/<name>/`, not in `context/`.
  `scripts/check.sh` verifies that every pointer inside them resolves.
- Notes for planned work sit directly under `context/`. A goal rooted here keeps its record in
  `context/goals/<goal>.md` while it is active; the workspace roadmap is at the coordinator
  (standards-lab).
