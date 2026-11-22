# claude-plugins

claude-plugins is the Standards Lab organization's Claude Code plugin marketplace. Its plugins
codify the organization's development processes, so every project applies them the same way. The
repository is developed with marathon, so the workflow runs against its own source.

## Capability map

- **marathon** (built): the long-haul development workflow, run as a software factory. The
  architect is involved twice per task, at the task brief and at the session brief, and subagents
  implement, review against standards and the spec, and keep the context current in between.
  Goals and tasks in the roadmap manifest are the units of work. PLAN runs each repository's
  optional `currency` command first, so a task begins by upgrading what trails, and SHIP waits on
  a pull request's checks only where `[remote] checks` declares CI, and runs an optional release
  gate before each tag. Each step that runs a configured command calls the
  plugin's `marathon` command, which runs it exactly as declared. Its design lives in
  [plugins/marathon](../plugins/marathon/), and its tools are Go commands in the
  [`marathon/` module](../marathon/).
- **marathon-architecture** (built): the extension that finds the architecture layer each
  repository's `STANDARDS.md` points to, so the standards-reviewer can apply it. Like every
  extension, it declares the minimum marathon version it requires (marathon's
  `references/extensions.md`). Its design lives in
  [plugins/marathon-architecture](../plugins/marathon-architecture/).
- **harness.optimization** (built): marathon checks its rules instead of describing them. Its
  `lint-context` holds each repository's `context/` and `roadmap/` to their file kinds, word
  budgets, and shapes, and its `lint-references` resolves the markdown references in comments.
  See [harness-optimization.md](harness-optimization.md).
- **marathon-extraction** (planned): carries patterns a consumer proves back to the blueprint.
  See [marathon-extraction.md](marathon-extraction.md).
- **marathon-references** (planned): a proposed extension that holds the references catalog.
  See [marathon-references-extension.md](marathon-references-extension.md).
- **Repository access** (planned): a goal locks only its root, holding its other repositories
  while they are off their default branch, and each command reads only the brief formats it
  writes ([repo-access.md](repo-access.md)).
- **Tooling layer** (planned): deterministic tools the skills call in place of procedure written
  as prose. See [tool-based-skills.md](tool-based-skills.md).
- **Harness testing** (built): `scripts/check.sh` runs on every push. It checks the plugins'
  versions, CHANGELOG headings, links, citations, word budgets, the brief formats each Read first
  list must name, and manifests, fails on any tracked Python file, and runs the Go check over the
  `marathon/` module, the `lint-context` and `lint-references` it builds, and the tests in
  `scripts/tests/`. Each failure message names its fix. Before each tag, the release gate, `scripts/gate.sh`, runs the check, cross-builds
  marathon's release binaries for a marathon tag, and runs the plugin's `claude plugin eval`
  suite. marathon's suite has one case, drawn from an observed failure. See
  [harness-testing.md](harness-testing.md).
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
