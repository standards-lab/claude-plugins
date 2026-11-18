# harness.optimization: marathon checked, not described

The goal is planned work, `harness.optimization`. marathon 0.18.0 carries its rules as prose,
and the content it produces outgrows them.

## Sources

- **Lauren Tan, "All Code Is Slop (Except Mine)"** (AI Engineer Code 2026). Each correction
  becomes a deterministic check rather than another line of rules, because prose rules, docs, and
  comments are where agents write justifications for not doing the right thing, and they go
  stale. Context comes through structure, review reads verification artifacts, a feature map
  gives one vocabulary, and CLIs inside skills replace procedure.
- **Brandon Thomas, "Let's Open Source the World"** (AI Engineer Code 2026). CLI and MCP are
  first-class, and humans review what is hard to reverse. The second point is the backlog task
  `v1.harness.door-scaled-review`.
- **Matt Pocock, "Fixing the PR Bottleneck"** is the baseline marathon 0.16 adopted. Its retro
  routes for tool economy and bloat are not adopted yet.

## What the goal measures against

Read these sizes, in words, to judge the work:

- The skill: 4,289 words in SKILL.md and its three @-includes. The includes appeared as literal
  text when the skill loaded, so they load only when an agent reads them.
- The coordinator's `context/roadmap.toml`: 12,765 words, with dated lines its format forbids.
- A long goal record: docproc's `v1.ai.processing` runs to 14,278 words, 11,283 of them
  Decisions.
- Spike READMEs that became lab notebooks: `spike-html-conventions`, at 24,678 words.
- Coordinator notes that carry history: `ai-hosting.md`, at 8,698 words.

## Settled

- **Extensions declare a floor.** An extension declares the minimum marathon version it
  requires. A release that changes the extension contract updates the affected extensions and
  bumps their floor in the same goal.
- **Budgets are checks.** A `lint-context` script ships in the plugin, and each repository's
  `[project] check` calls it.
- **Comments.**
  - marathon's implementer writes only the comments the standard requires, and never one that
    excuses a workaround.
  - The standards-reviewer sweeps for comments that excuse.
  - The editor stays out of source files.
  - The architecture layer's comment standard becomes a near-ban: package documentation, plus
    the reasons that lint suppressions and pin holds require. That task needs the architecture
    lock.
- **One release at the end.** Nothing releases until the last task, so sessions keep running the
  installed 0.18.0 until then. The sync then cleans the coordinator's roadmap and notes under the
  new check.
- **The LOCATE and status CLI** stays with `v1.harness.tooling` (`tool-based-skills.md`).
