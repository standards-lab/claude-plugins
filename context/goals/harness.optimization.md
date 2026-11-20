# goal · harness.optimization

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] version-floor
2. [x] skill-footprint
3. [x] roadmap-layout
4. [x] content-budgets
5. [x] retro-routes
6. [x] evidence-artifacts
7. [x] comment-standard
8. [x] harness-standards
9. [x] release

## Task brief · harness.optimization · release

**Problem**

marathon 0.19.0's changes, and marathon-architecture's floor that depends on them, reach no
session until they are tagged. The coordinator still holds the single-file roadmap 0.19.0 can't
read, its check doesn't run lint-context, and the plan-round eval fails the release gate.

**Behaviors**

1. marathon's eval suite scores 1.0 over three runs at the gate's settings, with no threshold or
   grader relaxed.
2. marathon's Unreleased section releases as `## v0.19.0`, with a `### Migrating from 0.18.0`
   section of ordered steps; marathon-architecture's releases as `## v0.4.0`. Each manifest's
   version matches its heading, and the check passes.
3. The goal note states what the goal built; its findings for later stay, and every pointer to it
   resolves. The tooling task's context note records the two deferred guards.
4. `marathon/v0.19.0`, then `marathon-architecture/v0.4.0`, are pushed after the gate passes on
   the merge commit. Each release run succeeds, and the pinned-tag fetch of lint-context passes
   on claude-plugins.
5. The coordinator's roadmap is an index of the three arrays plus a format pointer, and one goal
   file per goal with explicit `repos` and `root`, an `order` for each goal with tasks, and no
   comments. harness.optimization is gone from it and from `active`.
6. The header's plan survives as an undated coordinator note, and the dividers are gone. Each
   goal-attached comment block sits in the note it cites, or in its goal's coordinator note.
7. The coordinator's `[project] check` runs lint-context. Its only failures are the over-budget
   notes and entries and the missing capability map, which the Sync section lists for
   harness.adoption.
8. The Sync section proposes harness.adoption for a plan to stage. The session merges the
   coordinator's `sync-harness.optimization` pull request itself, then deletes the goal record;
   reinstalling 0.19.0 and stopping and resuming other sessions is the architect's.

**Test seams**

- claude-plugins' check, and `scripts/gate.sh <tag>` on the merge commit.
- lint-context on the coordinator's sync branch, its output matching Behavior 7 exactly.
- A TOML parse of the old roadmap compared with the index and goal files: the same goals, tasks,
  fields, and arrays.

**Slices**

claude-plugins, then the coordinator, with no upgrade slice.

1. The plan-round width fix. Demo: the eval suite at the gate's settings scores 1.0.
2. The CHANGELOG releases, manifest versions, goal note, and tooling note. Demo: the check
   passes and both versions match their headings.
3. Ship and release. Demo: both tags pushed in order, both runs green, the fetched lint-context
   passes.
4. The coordinator sync branch: split, header, dividers, comments, check. Demo: the parse
   comparison is identical, and lint-context prints only the expected failures.
5. Merge the coordinator's sync pull request and delete the goal record. Demo: standards-lab's
   main holds the split roadmap, and the record is gone.

**Out of scope**

- Trimming the coordinator's notes and entries, and its capability map: harness.adoption.
- Committing the one-off split script.
- Staging harness.adoption.
- The goal-record order mismatches in v1.ai.processing and v1.client.stack.
- Any change to architecture.

**Door**

One-way for the two tags. The coordinator split is a revertible merge, but 0.18.0 sessions can't
read the roadmap once it lands; the architect coordinates the reinstall. The rest is revertible
commits.

**Release**

`marathon/v0.19.0`, then `marathon-architecture/v0.4.0`.

## Decisions

- Nothing releases until the last task, `release`; sessions run the installed 0.18.0 until then
  (plan).
- harness-standards: harness programming principles live in claude-plugins' STANDARDS.md, not the
  architecture layer, whose pages need more than one repository to express them; the task
  evaluates this goal's practices, those before it, and the Pocock practices 0.16 adopted (plan).
- version-floor: unreleased changes sit under `## Unreleased` in each CHANGELOG; `release` renames
  them to marathon v0.19.0 and marathon-architecture's next version and releases both together,
  which satisfies the 0.19.0 floor.
- roadmap-layout: GitHub Issues were rejected for the roadmap, for losing PR-reviewed planning,
  offline sessions, and checks; Markdown and HTML were rejected because the roadmap is data that
  LOCATE, staging, lint-context, and the planned CLI parse (plan).
- roadmap-layout: nested goal directories, a full `[goals.<path>]` table in a goal file, and
  ancestor inheritance of `root` and `repos` were rejected (plan).
- roadmap-layout: a GitHub Project mirror of the cross-goal view is deferred to
  `v1.harness.tooling` (plan).
- roadmap-layout: the coordinator's roadmap is split during this goal's final sync, which also
  moves the index's narrative header into notes; marathon 0.19.0 reads only the split layout (plan).
- content-budgets: CI without the plugin fetches `lint-context` from the pinned `marathon/vX.Y.Z`
  tag; skipping when absent was rejected. The fetch can't be verified until `release` tags (plan).
- content-budgets: budget options and per-repository overrides were rejected (plan).
- retro-routes: retro reads redirects and escalations from the session brief in the pull request
  body; a goal-record section, which sync deletes, and transcripts alone were rejected (plan).
- retro-routes: a citation's first directory fails only when one edit from a real one; an
  allowlist and owner-marking were rejected. No retro eval case (plan).
- evidence-artifacts: an artifact is text in the session brief; stored files were rejected
  because the brief is gitignored and deleted at SHIP (plan).
- comment-standard: moving provenance lines into the context README was rejected (plan).
- comment-standard: a hold list read by the currency command, docproc's form, was rejected for
  pin-hold reasons (plan).
- comment-standard: a portable comment checker beside lint-context and a go/ast analyzer for
  free comments were rejected; stock lint rules and standards review enforce comments (plan).
- comment-standard: making the near-ban marathon's default in every repository was rejected; the
  standards-reviewer applies the repository's declared rule (plan).
- comment-standard: claude-plugins' scripts adopt the near-ban in harness-standards (plan).
- harness-standards: claude-plugins only; STANDARDS.md points to architecture's pages without
  editing them (plan).
- harness-standards: STANDARDS.md is one line per judgement call or pointer; rationale sections
  were rejected as the prose that goes stale (plan).
- harness-standards: staged execution and `@`-loading were dropped from the old harness pages as
  superseded (plan).
- harness-standards: the comment near-ban covers TOML and YAML configuration, which drifts like
  source. The convention lives in the STANDARDS pointer; amending architecture's comments page
  was deferred to harness.adoption (plan).
- harness-standards: `[remote] merge` keeps `gh pr checks --watch &&`; `--auto` was rejected
  because auto-merge is off and it returns before the merge (plan).
- harness-standards: a required `check` status on main was dropped: with no bypass it rejects the
  goal-record bookkeeping commits marathon pushes to main, and with the admin bypass it may not
  stop an admin's merge; `marathon merge` ends the retyping that merged #49 red (escalation).
- harness-standards: configured commands run through `marathon <key>`, built here rather than
  deferred to v1.harness.tooling. `marathon-run`, `marathon-exec`, `marathon-command`, and
  per-table tools were rejected; the root command takes the key as its subcommand (plan).
- harness-standards: the command runs all six keys through bash with appended arguments; refusing
  multi-line values and arguments was rejected (plan).
- harness-standards: a hook refusing chained merges was deferred; a marathon hook fires in every
  session (plan).
- harness-standards: the one-command-per-call rule stays out of marathon as user-scope
  behavior (plan).
- harness-standards: the comment-standard Decision stands; lint-context's goal-file rule shapes
  the manifest and is not a comment checker for source (plan).
- harness-standards: release's sync moves goal-attached roadmap comments into the note each cites;
  a block-by-block judgement pass was rejected (plan).
- harness-standards: `marathon` reads `.claude/marathon.toml` at the git top-level and runs the
  command from there, so a step can call it from any subdirectory (build).
- harness-standards: `marathon` exits 2 on a usage or configuration failure; otherwise it exits
  with the command's status, or 128 plus the signal number when a signal ends the command (build).
- harness-standards: a multi-line value gets the arguments as positional parameters, not appended.
  Any value that holds a newline is multi-line, including one written with a TOML `\n` escape
  (build).
- harness-standards: an eval case's description says what each grader matches, because a
  grader's body is its pattern or criteria and its front matter carries no comments (build).

- release: the coordinator's over-budget notes and entries and its missing capability map go to
  harness.adoption; its check runs lint-context with those failures showing (plan).
- release: the release session merges the coordinator's sync pull request itself; a handoff for
  the architect to merge was rejected. The architect reinstalls 0.19.0 and stops and resumes the
  other sessions (plan).
- release: the roadmap header survives only as its undated plan in a coordinator note (plan).
- release: the two deferred guards are recorded in `tool-based-skills.md`, not the roadmap (plan).
- release: a `rec:` gives its reason and a question restates the fact it leans on, because the
  self-contained-questions grader failed runs that lacked them (build).
- release: the configuration example's lint-context fetch pins `marathon/v0.19.0`, the first tag
  that ships lint-context (build).
- release: the release-heading test inserts its Unreleased heading above an older fixed release,
  so it doesn't depend on the release under way (build).
- release: the gate failed twice on a stochastic eval; retries and a lower threshold were rejected
  for a root fix (escalation).
- release: a marathon Stop hook wraps the formats deterministically. This reverses
  harness-standards' objection that a marathon hook fires in every session, for reply formatting
  only, since the hook exits at once on a reply without a format heading (escalation).
- release: the judge's criteria mirror briefs.md's plan-round rules, and the gate passes
  `--judge-model sonnet`, because eval's default judge, haiku, flipped verdicts on identical text
  (escalation).
- release: the judge is named by the `sonnet` alias, not a pinned ID, following rolling currency;
  the stored rounds recalibrate it if a release shifts verdicts (escalation).
- release: eval's own sonnet judge failed a round that the same model passes outside eval, and
  eval neither stores reasoning nor re-grades stored output; the gate rests on deterministic
  graders, with a back-reference regex, and the judge moves to `evals-judged/`, run by hand,
  because eval can't mark a grader non-scoring; an in-eval calibration case was rejected
  (escalation).
- release: a postscript after a round's **reply:** line stays allowed; the judge reads only the
  round (escalation).
- release: a format heading over 100 characters is wrapped like any other line (build).
- release: the wrapper never starts a continuation line on a Markdown block start, which had
  turned `step 1. Then` into a list item and failed numbered-questions; the gate keeps every
  run's trace and final message in its results directory, so a failed run can be diagnosed
  (build).

## Pending edits

- standards-lab · `.claude/marathon.toml`: set `[project] check` to `lint-context`, applied at
  `release`'s sync after the roadmap split (content-budgets).
- standards-lab · `context/roadmap.toml`: at release's sync, move the header narrative into
  notes, drop the section dividers, and move each goal-attached comment into the note it cites,
  or its goal's note when it cites none (harness-standards).
