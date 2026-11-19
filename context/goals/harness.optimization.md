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
7. [ ] comment-standard
8. [ ] harness-standards
9. [ ] release

## Task brief · evidence-artifacts

**Problem**

The architect reviews the session brief for outcomes: what changed, and whether each new
behavior works. Today Evidence is a loose list of commands with no tie to the task brief's
behaviors. It mostly repeats the check and the tests, which leaves the architect to work out how
to see each behavior. The spec-reviewer already exercises each behavior but returns only "the
commands you ran". The brief also carries detail the architect doesn't review by: what each
reviewer changed, and a code excerpt for every Core change. The read-first sets that hold the
brief's format have 13 to 59 words of room, so this change has to replace text, not add it.

**Behaviors**

1. The session brief opens with what changed: Summary, then Core changes. A Core change carries a
   code excerpt only where the code is the change, and every entry keeps its one line on why.
2. A Verify section replaces Evidence. It opens with the check's one-line result. Then comes one
   item per task-brief behavior, numbered as the brief numbers them. Each item gives the steps the
   architect can follow, outside the check and the tests, and what they should see. The check and
   the tests aren't repeated per behavior.
3. An item reads "tests only", naming the test, only when no step outside the tests shows the
   behavior. It is the exception, not the default.
4. Decided without you lists the decisions the brief didn't cover and each Decision the editor
   dropped. It no longer lists what the reviewers changed.
5. The spec-reviewer runs the check, then writes and follows the steps for each numbered
   behavior. It returns the steps and what it saw in Verify's form, citing the session-brief
   format rather than restating it. Its gap report is unchanged.
6. Validate runs the check and follows each behavior's steps that the spec-reviewer returned. A
   failing step blocks BRIEF the same way a failing check does. The results are Verify.
7. On a spike's last task, Verify carries the spike's numbered evidence.
8. No read-first budget or profile budget changes. The session brief lists each set's size at
   task end.
9. marathon's CHANGELOG records the change under Unreleased. It carries no extension-contract
   marker, because no extension reads the session brief. Its migration step says to expect Verify
   in place of Evidence, and reviewers' changes no longer in Decided without you.
10. The goal note records verification artifacts as settled, citing the Lauren Tan source.
11. The repository check passes, and the currency command reports nothing.

**Test seams**

- The repository check: citations resolve and every budget holds.
- The branch's spec-reviewer, engaged through `claude -p --plugin-dir` on this task's branch. It
  returns one Verify item per behavior of this brief. This is needed because the session itself
  runs the installed 0.18.0, whose profile doesn't know the Verify form.
- This task's own session brief, written in the new form.

**Slices**

1. The session-brief format:
   - the opening line
   - Core changes' excerpt rule
   - Verify, with its "tests only" exception
   - the trimmed Decided without you and its examples
   - the spike line

   Every mention of Evidence as the brief's section becomes Verify.
2. The spec-reviewer writes, follows, and returns each behavior's steps. Validate follows them and
   blocks BRIEF on a failing step.
3. The CHANGELOG entry, the goal note's settled line, and a final budget measure.

Each slice keeps every budget and passes the check.

**Out of scope**

- A check or eval that enforces one item per behavior (round 1, question 3). A retro proposes one
  once a brief misses an item.
- Stored artifact files such as screenshots and logs (round 1, question 1).
- Budget ceilings, and reviewers reading `git diff --stat` first. Both stay in Findings for later.
- Releasing or tagging, which is task `release`.

**Door**

Two-way. Nothing is released. Sessions keep running the installed 0.18.0 until `release`, so
Verify reaches no session before then.

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
- evidence-artifacts: one item per behavior is a skill rule; a lint-context check and an eval case
  wait for a retro finding (plan).

## Pending edits

- standards-lab · `.claude/marathon.toml`: set `[project] check` to `lint-context`, applied at
  `release`'s sync after the roadmap split (content-budgets).
