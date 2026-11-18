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

slices 6/6 committed · standards done · spec done · editor done

## Decisions

- Nothing releases until the last task, `release`; sessions run the installed 0.18.0 until then
  (plan).
- The architecture comment standard becomes a near-ban; `comment-standard` adds architecture to
  the goal's repos through a plan once v1.ai.processing frees it (plan).
- harness-standards: harness programming principles live in claude-plugins' STANDARDS.md, not the
  architecture layer, whose pages need more than one repository to express them; the task
  evaluates this goal's practices, those before it, and the Pocock practices 0.16 adopted (plan).
- version-floor: unreleased changes sit under `## Unreleased` in each CHANGELOG; `release` renames
  them to marathon v0.19.0 and marathon-architecture's next version and releases both together,
  which satisfies the 0.19.0 floor.
- skill-footprint: No Extension contract marker; marathon-architecture keeps its 0.19.0 floor
  (plan).
- roadmap-layout: GitHub Issues were rejected for the roadmap, for losing PR-reviewed planning,
  offline sessions, and checks; Markdown and HTML were rejected because the roadmap is data that
  LOCATE, staging, lint-context, and the planned CLI parse (plan).
- roadmap-layout: nested goal directories, a full `[goals.<path>]` table in a goal file, and
  ancestor inheritance of `root` and `repos` were rejected (plan).
- roadmap-layout: a GitHub Project mirror of the cross-goal view is deferred to
  `v1.harness.tooling` (plan).
- roadmap-layout: the coordinator's roadmap is split during this goal's final sync, which also
  moves the index's narrative header into notes; marathon 0.19.0 reads only the split layout (plan).
- roadmap-layout: the plan-round eval scores 0.83–0.89 against the gate's 1.0 threshold on both
  `main` and this branch, because the model's reply lines run past 100 characters; that predates
  this task.
- content-budgets: CI without the plugin fetches `lint-context` from the pinned `marathon/vX.Y.Z`
  tag; skipping when absent was rejected. The fetch can't be verified until `release` tags (plan).
- content-budgets: budget options and per-repository overrides were rejected (plan).
- content-budgets: standards-lab's check becomes `lint-context` at `release`'s sync, after the
  roadmap split (plan).
- content-budgets: no extension-contract marker, because no extension reads the files `lint-context`
  checks.
- content-budgets: `lint-context` skips dotfiles and dot directories under `context/`, looser than
  the brief's "any other file fails".
- content-budgets: a capability-map entry requires `: text` after its status, and the map section
  allows no lead-in prose.
- content-budgets: the editor's drops are listed under the session brief's "Decided without you".
- content-budgets: `commands/retro.md` still reads redirects and escalations from Decisions, which
  pruning removes; `retro-routes` decides where retro gets them.

## Pending edits

- standards-lab · `.claude/marathon.toml`: set `[project] check` to `lint-context`, applied at
  `release`'s sync after the roadmap split (content-budgets).
