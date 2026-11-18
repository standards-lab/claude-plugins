# goal · harness.optimization

- **State:** building
- **Task:** content-budgets
- **Branch:** content-budgets

## Tasks

1. [x] version-floor
2. [x] skill-footprint
3. [x] roadmap-layout
4. [ ] content-budgets
5. [ ] retro-routes
6. [ ] evidence-artifacts
7. [ ] comment-standard
8. [ ] harness-standards
9. [ ] release

## Task brief · content-budgets

**Problem**

marathon states the size and shape of what it writes as prose, nothing checks it, and the content
has outgrown it. docproc's `v1.ai.processing` record holds 12,334 words of Decisions. A spike
README runs to 24,678 words, and a note to 24,934. Decisions only accumulate. Spikes keep their
answer in either README, in both, or in neither, and the capability map is unstructured prose.
`roadmap-layout` also left the roadmap's layout rules unchecked. Each rule should be a
deterministic check that every repository's check runs, and the skill should point to that check
instead of restating it.

**Behaviors**

1. marathon ships `lint-context`, one executable on the plugin's PATH. It needs only Python 3.11
   or later and its standard library. It takes no arguments and checks the repository at the
   current directory. It prints one line per failure, naming the file, the rule, and the measured
   size against the limit where there is one. It exits nonzero when anything fails, and exits 0
   silently otherwise. A word is a whitespace-separated token.
2. A file in the context directory is the context README, a note directly in that directory, a
   goal record in its goals directory, or the roadmap index or a goal file in its roadmap
   directory. Any other file there fails.
3. Word budgets are fixed, with no configuration:
   - a goal record's Task brief 1,200, Decisions 2,000, and Pending edits 1,000, with its other
     sections unbudgeted;
   - each note 1,500;
   - the context README 1,000;
   - the root README 1,500.

   Documentation pages are unbudgeted. A spike repository gets the same rules.
4. In the roadmap:
   - each goal's and task's summary is at most 100 words, each proof 40, and each criteria
     statement 60;
   - the index holds only the three arrays and a comment of at most two lines;
   - every goal file writes `repos`;
   - a goal file with tasks has an `order` that lists each task table exactly once;
   - a goal record's task list matches its goal file's `order` when that file is in the same
     repository.
5. The context README has a `## Capability map` heading. Each entry under it is a list item that
   opens with a bold feature name, then `(built)` or `(planned)`. A built entry carries at least
   one relative link, and every link it carries resolves. A planned entry may carry none.
6. Every relative link outside code in the root README, the context directory, and the
   documentation pages resolves.
7. claude-plugins' check runs `lint-context` from the working tree and runs its tests, and both
   pass on this repository: its map has the new shape, and this goal's record fits its budgets.
8. The skill says each repository's check calls `lint-context`. A repository whose CI lacks the
   plugin fetches the script from the pinned `marathon/vX.Y.Z` tag by its tag URL. `init` writes
   a check that calls it, and describes the map as built and planned entries.
9. Decisions keep only what binds later tasks: settlements they rely on, rejected alternatives,
   holds, and falsified assumptions. At BUILD the editor drops the Decisions that no longer bind,
   on the task's branch, and the session brief lists each drop.
10. A spike keeps its question, evidence, and answer in its context README. Each survey or result
    gets its own documentation page, linked from its evidence item. The answer section links to
    the context README's answer, and spike setup, `experiment`, and the answer section all say so.
11. The skill's prose for each rule `lint-context` checks is cut to a pointer to it. Each
    read-first set a trim shrinks gets a lowered budget, and each set pushed over gets a raised
    one. Either way the new budget is the set's size at task end, rounded up to the next 100
    words. The session brief lists each change.
12. marathon's CHANGELOG records the change under Unreleased, with no extension-contract marker.
    Its migration step says to call `lint-context` from the check, reshape the map, move a
    spike's answer into its context README, and trim whatever fails.
13. The goal record's pending edits include setting standards-lab's check to `lint-context` at
    `release`'s sync.
14. The repository check passes, and the currency command reports nothing.

**Test seams**

- `lint-context` run on fixture trees by stdlib `unittest` tests, which assert its exit status and
  its failure lines
- The repository check

**Slices**

1. `lint-context` classifies the files in the context directory and holds the READMEs, the notes,
   and the record sections to their budgets. claude-plugins' check runs it from the working tree,
   along with the tests. This goal's Decisions are pruned to what binds later tasks, so the
   record passes with this task's approved Decisions in it.
2. The roadmap checks: entry budgets, the index's contents, `repos`, `order`, and the record's
   task list.
3. The capability map's shape and the relative links. This repository's map takes the new shape.
4. The skill: the check calls `lint-context`, CI fetches it by tag, and `init` writes both. The
   prose of each checked rule becomes a pointer, and the budgets are measured.
5. The Decisions rule and the spike layout in the skill. The editor's pass applies the rule to
   this task's own Decisions.
6. The CHANGELOG entry, the standards-lab pending edit, and a final budget measure.

Each slice adjusts the budgets its own edits move, so the check passes after every slice.

**Out of scope**

- Running `lint-context` on other repositories or trimming their content. Each repository does
  that when it adopts the script, and `release`'s sync does it for the coordinator.
- Checks across repositories, such as a member's record against the coordinator's goal file
- Budget options or per-repository overrides, and budgets on documentation pages or other record
  sections
- CLAUDE.md, and STANDARDS.md (task `harness-standards`)
- Releasing or tagging either plugin (task `release`)

**Door**

Two-way. Nothing is released. Sessions keep running the installed 0.18.0 until `release`, and
consumers adopt `lint-context` only after it.

## Progress

slices 2/6 committed · standards — · spec — · editor —

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
- roadmap-layout: a task's `repos` stays optional, and a task without one uses its goal's `repos`.
  That default works inside one file, so it is not ancestor inheritance.
- roadmap-layout: goal files keep the `name` title field. "Doesn't repeat the goal's name" means
  the file doesn't repeat its dotted path.
- roadmap-layout: container goals write `repos = []`, and the index comment is one line.
- roadmap-layout: prose always tells the two arrays apart as "the goal's `order`" and "the
  workspace `order`".
- roadmap-layout: a sub-goal's file name "extends" its parent's by one segment, not "begins with"
  it, so `v10` is not a child of `v1`.
- roadmap-layout: `experiment` also creates the file for the reserved root goal `experiment` when
  none exists, because intake's empty-ancestor deletion can remove it.
- roadmap-layout: in `status <goal>`, a done task is marked `x`, and a sub-goal with no record
  counts 0/n.
- roadmap-layout: the CHANGELOG puts the migration steps inside the Changed bullet, not in a
  separate "Migrating from" section.
- roadmap-layout: budgets raised: experiment 9300→9700, init 3500→3800, intake 9500→9900, plan
  8900→9300, retro 6300→6400, start 7900→8000, status 5100→5500.
- roadmap-layout: the plan-round eval scores 0.83–0.89 against the gate's 1.0 threshold on both
  `main` and this branch, because the model's reply lines run past 100 characters; that predates
  this task. The scaffold builds the new layout, and in every run the model found the goal
  through it.
- content-budgets: `lint-context` ships as one stdlib-only Python executable in marathon's `bin/`,
  called by each repository's `[project] check` (plan).
- content-budgets: CI without the plugin fetches `lint-context` from the pinned `marathon/vX.Y.Z`
  tag; skipping when absent was rejected (plan).
- content-budgets: Decisions keep only what binds later tasks (relied-on settlements, rejected
  alternatives, holds, falsified assumptions); the editor drops the rest on the task branch at BUILD
  and the session brief lists each drop (plan).
- content-budgets: a spike's surveys and results live in `docs/`, one page per evidence item linked
  from it; its question, evidence, and answer live in `context/README.md` (plan).
- content-budgets: the capability map is named features marked `(built)` or `(planned)`, under a
  required `## Capability map` heading; built entries need a resolving relative link, planned ones
  may have none (plan).
- content-budgets: budgets are fixed in `lint-context` with no configuration or options;
  `lint-context` checks only its own repository and takes no arguments (plan).
- content-budgets: budgets are Task brief 1,200, Decisions 2,000, Pending edits 1,000; notes 1,500;
  context README 1,000; root README 1,500; manifest summary 100, proof 40, criteria statement 60;
  `docs/` pages unbudgeted (plan).
- content-budgets: spikes get the same rules with no marker; file kinds are fixed paths, and any
  other file under `context/` fails (plan).
- content-budgets: relative links outside code resolve in the root README, `context/`, and `docs/`;
  the roadmap index holds only the arrays and a comment of at most two lines, and every goal file
  writes `repos` (plan).
- content-budgets: skill prose for each rule `lint-context` checks is cut to a pointer, and
  read-first budgets move to their task-end size (plan).
- content-budgets: standards-lab's check becomes `lint-context` at `release`'s sync, after the
  roadmap split (plan).
- content-budgets: budget numbers live only in `lint-context`, whose failure lines name each limit;
  it checks the current directory, not the git top level.
- content-budgets: a map entry reads `- **Name** (built|planned): text`; a link resolves when its
  file or directory exists, and anchors, absolute URLs, and code are skipped.
- content-budgets: tests are stdlib `unittest` on fixture trees built in temporary directories, kept
  outside the plugin; CI uses the runner's `python3`.
- content-budgets: no extension-contract marker, because no extension reads the files `lint-context`
  checks.

## Pending edits

none
