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

- The skill: SKILL.md links each file instead of `@`-including it, and each command's playbook
  opens with a "Read first" list of the files it needs. `start`'s read-first set, SKILL.md and the
  playbook included, is 7,981 words against a budget of 8,000. `mechanics/pipeline.md` is about
  1,970 words, down from 2,715, since Releasing and RESUME moved to their own files.
- The planner profile: 273 words, down from 554, since profiles cite the skill's rules instead of
  restating them.
- The coordinator's `context/roadmap.toml`: one file of 13,353 words over 61 goals and 74 tasks,
  with no task order and dated lines its format forbids. It stays one file until `release`.
- A long goal record: docproc's `v1.ai.processing` runs to 14,278 words, 11,283 of them
  Decisions.
- Spike READMEs that became lab notebooks: `spike-html-conventions`, at 24,678 words.
- Coordinator notes that carry history: `ai-hosting.md`, at 8,698 words.

## Settled

- **Budgets are checks.** A `lint-context` script ships in the plugin, and each repository's
  `[project] check` calls it.
- **Comments** (task `comment-standard`). The architecture layer's comment standard becomes a
  near-ban: package documentation, plus the reasons that lint suppressions and pin holds require.
  That task needs the architecture lock. marathon's implementer and standards-reviewer profiles
  hold marathon's own comment rules, and marathon's `behavior/delegation.md` keeps the editor out
  of source files.
- **One release at the end.** Nothing releases until the last task, so sessions keep running the
  installed 0.18.0 until then. The sync then cleans the coordinator's roadmap and notes under the
  new check.
- **Roadmap layout** (task `roadmap-layout`). Unreleased marathon reads an index of the three
  arrays plus one `context/roadmap/<goal>.toml` per goal, holding its own `root` and `repos`, an
  `order` array, and its tasks (marathon's `references/manifest.md`). The installed 0.18.0 still
  reads the coordinator's single file, so `release`'s sync splits it and moves its dated header
  into notes.
- **Harness standards** (task `harness-standards`). claude-plugins gains a `STANDARDS.md` for the
  judgement calls of harness programming, which the standards-reviewer applies. The architecture
  layer's `harness/` pages left in its architecture-diet because one repository can't promote a
  page; their history (`b927dd8^:harness/`) and the Pocock practices of 0.16 are inputs.
- **The LOCATE and status CLI** stays with `v1.harness.tooling` (`tool-based-skills.md`).

## Findings for later

A retro or `retro-routes` can take these up.

- claude-plugins has no `STANDARDS.md`, so nothing points to the architecture pages the standards
  review applies: `principles/rolling-currency.md`, `principles/independent-releases.md`, and
  `principles/context-architecture.md` in the architecture repository.
- The check finds a command's read-first set by the "Read first" line. A playbook whose line is
  renamed silently yields a smaller set that still passes its budget. The check could require one
  "Read first" line per playbook.
- The check skips a bare citation whose first directory exists in no plugin or skill root, so a
  typo in that directory goes uncaught.
- marathon's manifest reference says `repos = []` means "only `plan` works on it", but `intake`
  runs on experiment goals, which write `repos = []`. This predates `roadmap-layout`.
- The plan-round eval fails the release gate's 1.0 threshold on `main`, scoring 0.83–0.89
  because the model's reply lines run past 100 characters. `release` needs it fixed, or the
  threshold settled, before tagging.
