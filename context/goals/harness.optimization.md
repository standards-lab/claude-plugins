# goal · harness.optimization

- **State:** building
- **Task:** retro-routes
- **Branch:** retro-routes

## Tasks

1. [x] version-floor
2. [x] skill-footprint
3. [x] roadmap-layout
4. [x] content-budgets
5. [ ] retro-routes
6. [ ] evidence-artifacts
7. [ ] comment-standard
8. [ ] harness-standards
9. [ ] release

## Task brief · retro-routes

**Problem**

`retro` gathers the architect's redirects and escalations from goal-record Decisions, which the
editor now prunes. Redirects never reach a lasting record at all: a redirected brief is
overwritten, and only the final one becomes the pull request body. retro's routes start at a
check, so a mistake that a design change could make impossible gets a check instead. It has no
route for costly tool calls or for instructions that change nothing, both of which Pocock's retro
covers. A check message that doesn't name its fix leaves the agent to guess: 7 of lint-context's
messages don't, and about 14 of the repository check's don't. The repository check also has two
silent gaps. A renamed "Read first" line shrinks a command's budgeted set without failing, and a
typo in a bare citation's first directory goes uncaught.

**Behaviors**

1. The session brief has a "Redirects and escalations" section after Decided without you. It
   lists each BUILD escalation and each BRIEF redirect in the task, one line each with the
   architect's answer, or reads "none". The session records each one in the goal record's
   Progress as it happens, so a handoff loses none. A brief written after a redirect lists the
   earlier ones.
2. `retro` gathers from four sources:
   - the merged pull request bodies in the range, including their Redirects and escalations and
     their Decided without you
   - the standards-reviewer's commits
   - the touched notes, for drift
   - the touched docs, for drift

   It no longer reads goal-record Decisions or pull request comments.
3. `retro` also reads the local session transcripts of the range's sessions when they are
   present. The retro's header names each repository whose transcripts were missing, and the
   retro continues without them.
4. `retro` routes in this order: design, check/eval, standard, skill/profile, context.
   - A design finding changes the work so the mistake can't be made.
   - A ticked design finding becomes a task, written on the coordinator's branch, in the goal
     that owns the work. When no active goal owns the repository, the finding line proposes a
     goal, and the tick confirms it.
   - `retro` doesn't make the change itself.
5. `retro` looks for two more kinds of finding and routes each to the existing layers:
   - tool economy: a costly tool call that a cheaper call or a tool could replace
   - instruction trim: an instruction that changes no behavior
6. The skill says a check's failure message names its fix. Each check that `retro` proposes
   carries its message.
7. Every failure message of `lint-context` and of the repository check names what to change for
   it to pass. The tests assert the new lines.
8. The repository check fails a command playbook that has no "Read first" line, or more than
   one. The failure names the playbook and the fix.
9. The repository check fails a bare citation whose first directory is one edit away from a
   directory in its plugin's root or a skill's root. A citation inside a skill also resolves
   against its plugin's root, so citations to `agents/*.md` are checked.
10. Each read-first set a trim shrinks gets a lowered budget, and each set pushed over gets a
    raised one. Either way, the new budget is the set's size at task end, rounded up to the next
    100 words. The session brief lists each change.
11. marathon's CHANGELOG records the change under Unreleased. It carries no extension-contract
    marker, because no extension reads the session brief or `retro`'s sources and routes. Its
    migration step says to expect the new session-brief section.
12. The goal record's Decision that left `retro`'s source to this task is retired. The goal
    note's Findings for later loses that item and the two check gaps.
13. The repository check passes, and the currency command reports nothing.

**Test seams**

- `lint-context`: stdlib `unittest` tests run it on fixture trees and assert its exit status and
  its failure lines.
- The repository check: stdlib `unittest` tests run it on seeded copies of the repository and
  assert its failure lines. It also runs on the repository itself.
- A dry retro over this goal's merged pull requests, shown in the session brief's Evidence.

**Slices**

1. The repository check's two gaps. Each playbook has exactly one "Read first" line. A citation
   directory one edit away from a real one fails. Citations inside a skill also resolve against
   the plugin's root.
2. Messages that name their fix: lint-context's seven and the repository check's own, with their
   tests.
3. `retro`'s routes:
   - design first, and a ticked design finding becomes a task, with a proposed goal when no goal
     owns the repository
   - the tool economy and instruction trim kinds
   - proposed checks carry their messages, and the skill states the message rule
   - the retro format shows a design finding
4. `retro`'s sources: the session brief's Redirects and escalations section, recorded in
   Progress as each one happens. The gather reads pull request bodies, and transcripts when
   present, instead of Decisions and comments.
5. The CHANGELOG entry, the retired Decision, the goal note's findings, and a final budget
   measure.

Each slice adjusts the budgets its own edits move, so the check passes after every slice.

**Out of scope**

- Applying any finding the dry retro turns up. A later `retro` proposes them.
- A `retro` eval case
- A tool to search transcripts. `retro` reads them as they are stored.
- The manifest's `repos = []` wording, and spike `docs/` links to a goal record that sync
  deletes. Both stay in Findings for later.
- STANDARDS.md (task `harness-standards`), and releasing or tagging either plugin (task
  `release`)

**Door**

Two-way. Nothing is released. Sessions keep running the installed 0.18.0 until `release`, so the
new section and routes reach no session before then.

## Progress

slices 5/5 committed · standards ✓ · spec — · editor —

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
- retro-routes: retro reads redirects and escalations from a session-brief section kept in the
  pull request body; a goal-record section (deleted at sync) and transcripts alone were rejected,
  and pull request comments are dropped as empty in practice (plan).
- retro-routes: transcripts are optional evidence; tool economy and instruction trim are finding
  kinds, not route layers; a ticked design finding becomes a task, never a retro-branch change,
  and proposes a goal when none owns the repository (plan).
- retro-routes: a citation's first directory fails only one edit from a real one; an allowlist and
  owner-marking were rejected. No retro eval case (plan).

## Pending edits

- standards-lab · `.claude/marathon.toml`: set `[project] check` to `lint-context`, applied at
  `release`'s sync after the roadmap split (content-budgets).
