# goal · harness.optimization

- **State:** building
- **Task:** roadmap-layout
- **Branch:** roadmap-layout

## Tasks

1. [x] version-floor
2. [x] skill-footprint
3. [ ] roadmap-layout
4. [ ] content-budgets
5. [ ] retro-routes
6. [ ] evidence-artifacts
7. [ ] comment-standard
8. [ ] harness-standards
9. [ ] release

## Task brief · roadmap-layout

**Problem**

The coordinator's roadmap is one file: 13,353 words covering 61 goals and 74 tasks. It records no
task order, and every session reads all of it to find a single goal. A goal can take its `root`
and `repos` from an ancestor, so a reader has to walk the tree to learn what a goal locks.
marathon describes this single-file layout in several places: the manifest reference, LOCATE,
plan, experiment, intake, init, sync, status, the README, and the eval scaffold. Each goal should
be a short file of its own that carries its own task order, and an index should hold only the
three arrays.

**Behaviors**

1. The roadmap index keeps its current name. It holds the `active`, `planned`, and `backlog`
   arrays and a format-pointer comment of at most two lines. It holds nothing else: no goal
   tables and no other comments.
2. Every goal has its own TOML file in a roadmap directory beside the index, named by the goal's
   full dotted path. That includes container goals that sit in no array, and the reserved
   `experiment` root. The goal's fields sit at the top level, and each task is a `[tasks.<slug>]`
   table. The file doesn't repeat the goal's name.
3. Every goal file writes `repos`, and `[]` means the goal touches only the coordinator. A file
   writes `root` once the goal has one; a file with no `root` means `plan` sets the goal up first.
   No field takes its value from an ancestor. A spike's file adds `remote` and `path`.
4. A goal file that holds tasks has an `order` array listing every task slug in the file exactly
   once, in run order, and the goal record's task list mirrors it. A task added while the goal is
   active joins `order` when it is written, and joins the record's task list at the goal's next
   `plan` or `start`.
5. LOCATE reads the index, then the named goal's file. For a goal that is no longer listed, it
   finds the root in the history of the goal's deleted file.
6. The roadmap's writers keep the layout whole:
   - staging and pivoting edit only the index;
   - adding a goal creates its file;
   - adding a task adds its table and appends the slug to `order`;
   - `experiment` creates the experiment's file and one file per spike;
   - `intake` deletes the experiment's file, and the file of any ancestor left empty;
   - `init` creates the index, and a file for the first goal once it settles one.
7. Sync deletes the goal's file and the files of its sub-goals. It then deletes the file of each
   ancestor whose criteria now hold, or that is left empty: one with no tasks of its own and no
   file whose name begins with its name.
8. `status <goal>` prints the goal's own tasks in `order`. Each is marked done from the goal
   record, or all are open when the goal has no record. One line follows per sub-goal, with its
   done count. The brief-formats reference shows this table beside the status digest. Bare
   `status` is unchanged and lists active goals only.
9. No mention in either plugin describes the roadmap as a single file holding goal tables. This
   covers:
   - the skill, its references, and SKILL.md;
   - the README;
   - the context-engineering and workspace-coordination references;
   - the Sync example and the pending-edit wording in the brief formats and the goal record.

   Every mention describes the index-plus-goal-files layout. The manifest reference shows its
   example as the index and its goal files.
10. The plan-round eval's scaffold builds an index and a goal file, and the eval passes on it.
11. marathon's CHANGELOG records the layout under Unreleased, with a migration step: split the
    single file into the index and one file per goal, state `root` and `repos` in each file, and
    add `order`. The entry carries no extension-contract marker.
12. Each read-first set this task pushes over its word budget gets a raised budget: its size at
    the end of the task, rounded up to the next 100 words. Every other budget stays as it is, and
    the session brief lists each raise.
13. The repository check passes, and the currency command reports nothing.

**Test seams**

- The repository check
- The plan-round eval, run on the rebuilt scaffold
- The manifest reference's example, parsed with Python's `tomllib`: every file parses, the index
  holds only the arrays, and each `order` matches its file's task tables
- A search across both plugins for the single-file layout, meaning `[goals.` table headers or the
  index described as holding goals or tasks, which finds nothing

**Slices**

1. The manifest reference describes the index and the goal files: their naming, their fields,
   `repos` and `root` with no inheritance, `order`, the comment rule, and an example that parses.
2. LOCATE reads the index and then the goal file, and finds a synced goal's root in its deleted
   file's history. `status <goal>` prints the goal's tasks in order, and the digest is unchanged.
3. The writers: plan, experiment, intake, init, and sync create and delete goal files and keep
   `order`. The goal record mirrors `order`, and the pending edits and the Sync example name goal
   files.
4. The remaining mentions of the single file move to the new layout: SKILL.md, the README, and the
   context-engineering and workspace-coordination references. The eval scaffold builds the new
   layout, and the eval passes.
5. marathon's CHANGELOG gets the migration step under Unreleased, and the budgets are measured a
   final time.

Each slice raises the budgets its own edits push over, so the check passes after every slice.

**Out of scope**

- Splitting the coordinator's roadmap, the one-off script that does it, and moving its dated
  header narrative (task release)
- Layout checks, such as `order` against the task tables and the arrays against the files, which
  belong to `lint-context` (task content-budgets)
- Listing `planned` and `backlog` in bare `status`
- A GitHub Project mirror, and a LOCATE and status CLI (`v1.harness.tooling`)
- Releasing or tagging either plugin (task release)

**Door**

Two-way. Nothing is released. Sessions keep running the installed 0.18.0, which reads the
coordinator's single file until the `release` task.

## Progress

slices 0/5 committed · standards — · spec — · editor —

## Decisions

- Extensions declare a minimum marathon version, not a target; a contract-breaking marathon
  release updates its extensions and bumps their floor in the same goal (plan).
- Nothing releases until the last task, `release`; sessions run the installed 0.18.0 until then
  (plan).
- Content budgets are a `lint-context` check the plugin ships, not prose rules (plan).
- The architecture comment standard becomes a near-ban; `comment-standard` adds architecture to
  the goal's repos through a plan once v1.ai.processing frees it (plan).
- version-floor: applied marathon-architecture 0.3.0 (targets 0.16) under marathon 0.18.0; its
  on-start only locates the layer, and the question itself is the problem this task removes.
- version-floor: bumped CI pins create-gh-release-action v1.12.0 and claude-code 2.1.327 as
  slice 1; neither needed adapting.
- version-floor: a contract change is a CHANGELOG entry starting `**Extension contract:**`;
  marathon's CHANGELOG preamble states the convention and `references/extensions.md`,
  Compatibility, cites it.
- version-floor: unreleased changes sit under `## Unreleased` in each CHANGELOG; `release` renames
  them to marathon v0.19.0 and marathon-architecture's next version and releases both together,
  which satisfies the 0.19.0 floor.
- version-floor: marathon-architecture's CHANGELOG cites marathon's Compatibility section instead
  of restating the rule (standards review).
- roadmap-layout: the roadmap stays in repository files, one `context/roadmap/<goal>.toml` per
  goal, with `context/roadmap.toml` holding only the three arrays; GitHub Issues were rejected for
  losing PR-reviewed planning, offline sessions, and checks (plan).
- roadmap-layout: a goal's task order is an `order` array in its goal file, which the record's
  task list mirrors; every goal file states its own `root` and `repos` (plan).
- roadmap-layout: `status` prints the cross-goal view on demand; a GitHub Project mirror is
  deferred to `v1.harness.tooling` (plan).
- roadmap-layout: the coordinator's roadmap is split during this goal's final sync; the task
  precedes content-budgets so its manifest rules are written once (plan).
- skill-footprint: `@`-includes become links with an explicit read instruction; `@` expands only in
  CLAUDE.md (plan).
- skill-footprint: Each playbook opens with its own read-first list (plan).
- skill-footprint: Word budgets are the task-end size, rounded up to the next 100 (plan).
- skill-footprint: Skill descriptions are capped at 1,024 characters (plan).
- skill-footprint: Releasing and RESUME move to two on-demand files; the "tag <names>" resume moves
  with Releasing (plan).
- skill-footprint: The check resolves backtick file citations, and rejects `@`-include lines in
  every plugin (plan).
- skill-footprint: No Extension contract marker; marathon-architecture keeps its 0.19.0 floor
  (plan).
- skill-footprint: Profiles cite the skill through `${CLAUDE_PLUGIN_ROOT}` paths, not a `skills:`
  preload. They only cite shared rules, with delegation Limits as the home for delegate limits. A
  profile is the home of its own delegate's behavior (plan).
- skill-footprint: Profiles have word budgets (plan).
- skill-footprint: The session quotes relevant standards, the comment rule always included, into the
  slice brief; the implementer never loads `STANDARDS.md` whole. The slice brief gets a written
  format (plan).
- skill-footprint: The standards-reviewer owns doc-comment wording and the sweep for comments that
  excuse; the editor leaves source files (plan).
- skill-footprint: Architect-facing formats are markdown, fenced only around code: rounds as heading
  plus lists; the task brief as bold labels over lists; the session brief as sections; status as a
  table; retro as checkboxes. Goal records store the brief unfenced. All of it is hard-wrapped at
  100 characters (plan).
- skill-footprint: The plan-round eval gains a code-fence grader and a line-width grader (plan).
- skill-footprint: The format work joins skill-footprint; the roadmap redesign is its own task,
  roadmap-layout (plan).
- skill-footprint: The check also verifies `../` links, not only `./` links.
- skill-footprint: The check treats a bare backtick path as a plugin citation only when its first
  segment is a directory in a plugin or skill root. A cross-plugin possessive, as in marathon's
  `references/briefs.md`, resolves against the owning plugin, and released CHANGELOG sections are
  skipped.
- skill-footprint: A playbook's read-first list is a set of links under a "Read first" line.
  experiment's list points at `commands/plan.md`.
- skill-footprint: Releasing's resume detail lives in `mechanics/releasing.md`, "Resuming a
  release", not in `mechanics/resume.md`.
- skill-footprint: The invariant that nothing changes before approval also covers the outcomes of
  plan, experiment, and intake rounds.
- skill-footprint: `behavior/delegation.md`, Limits, holds every delegate limit, including those
  that apply to one profile only.
- skill-footprint: The templates in `references/briefs.md` are blockquotes, and each question's rec
  and changes lines are nested bullets.
- skill-footprint: Word budgets live in `scripts/budgets.txt`, keyed by file path, and a file with
  no budget fails the check.
- skill-footprint: start's read-first set measures 7,898 words against its 7,900 budget, so almost
  any edit to that set needs a budget change.
- skill-footprint: marathon's description shrank from 1,153 to 975 characters to fit the 1,024
  cap.
- harness-standards: harness programming principles live in claude-plugins' STANDARDS.md, not the
  architecture layer, whose pages need more than one repository to express them; the task
  evaluates this goal's practices, those before it, and the Pocock practices 0.16 adopted (plan).
- roadmap-layout: a goal file is named by its full dotted path and sits flat in the roadmap
  directory, as goal records do; nested directories were rejected (plan).
- roadmap-layout: a goal file holds the goal's fields at the top level and its tasks as
  `[tasks.<slug>]`. The full `[goals.<path>]` table was rejected, because the file name is the
  goal's identity (plan).
- roadmap-layout: every goal has a file, container goals and the `experiment` root included, and
  sync deletes an empty ancestor by deleting its file (plan).
- roadmap-layout: `repos` is always written, with `[]` meaning coordinator-only. `root` is written
  once it is known, and its absence means `plan` sets the goal up. Ancestor inheritance is
  dropped, because no goal used it (plan).
- roadmap-layout: `status <goal>` prints the goal's own tasks in `order`, with done marks from the
  record, then one line per sub-goal with its done count (plan).
- roadmap-layout: bare `status` stays the digest of active goals; the index is the cross-goal view
  (plan).
- roadmap-layout: the index's only comment is a format pointer of at most two lines. Narrative
  belongs in notes, and `release`'s sync moves the coordinator's header (plan).
- roadmap-layout: marathon 0.19.0 reads only the per-goal layout. The CHANGELOG has a migration
  step under Unreleased and no extension-contract marker, because no extension reads the roadmap
  (plan).
- roadmap-layout: every mention of the single-file roadmap across the plugins, the eval scaffold
  included, describes the index-plus-goal-files layout. The index keeps its name (plan).
- roadmap-layout: layout checks belong to content-budgets' `lint-context`, which parses with
  Python's standard `tomllib`. This task ships no script (plan).
- roadmap-layout: each budget the task pushes over is raised to its task-end size, rounded up to
  the next 100. Trims wait for content-budgets (plan).
- roadmap-layout: the roadmap stays TOML; Markdown and HTML were rejected because the roadmap is
  data that LOCATE, staging, lint-context, and the planned CLI parse (plan).

## Pending edits

none
