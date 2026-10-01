# marathon as a software factory

This note covers how a marathon session turns an approved task into merged work. The
architect reviews twice, once at the plan and once at the result, and agents do everything
in between. It replaces the stage-and-checkpoint loop of marathon 0.15. It is tracked as
`factory.pipeline`. The goal model it runs on is `marathon-goals.md`, the formats the architect
reads are `marathon-briefs.md`, and the command set is `marathon-commands.md`.

Inspired by Matt Pocock's talk "Fixing the PR Bottleneck" (AI Engineer Paris 2026) and his
skills repository, `github.com/mattpocock/skills`.

## The problem it solves

Under 0.15, one `start` plus its `close` stops for the architect about seven times:

- SETTLE approval
- each checkpoint
- the branch report
- each review finding
- the context confirmation
- the next-step decision
- the PR

The session also announces a model before every stage. An architect watching more than two
sessions ends up micromanaging all of them, and no session gets the chance to bring its work to
a finished state before asking for attention.

## The pipeline

```
PLAN   architect ⇄ planner
       round 1..n: numbered questions + recommendations (facts found by agents)
       → task brief + slice list                   ✔ approve          [touch 1]
BUILD  autonomous; the session orchestrates, holding pointers only
       for each slice: implementer (fresh context) → check ✓
       standards-reviewer → commits fixes → check ✓
       spec-reviewer → gaps? → implementer → re-review
       editor → goal record, notes, docs, prose
BRIEF  session brief                               ✔ accept / redirect [touch 2]
SHIP   PR (body = brief) → CI green → merge → update goal record; sync if last task
```

- **The architect is involved at two points.** Approving the task brief starts BUILD. Accepting the
  session brief authorizes publishing and merging. Merging follows the brief mechanically, so it
  is not a separate review.
- **Everything happens in the terminal.** The PR body is the session brief, so the record on the
  remote reads the same as what the architect accepted.
- **The session escalates during BUILD only for:**
  - a one-way door: an irreversible migration, data loss, or an action with outside effects
  - a decision the task brief doesn't cover
  - scope beyond the task

  An escalation uses the plan-round format and resumes BUILD once it is answered.
- **Handoff is automatic.** The orchestrating session holds only pointers to the brief, the goal
  record and the commits. When its context fills anyway, it writes the handoff state into the
  goal record and stops. The next session resumes from there without asking the architect.
- **Removed:** checkpoints, stage approvals, per-stage delegation statements, and the
  post-validation findings, tending and next-step stops.

## Subagent profiles

| Profile | Tools | Loads | Does |
|---|---|---|---|
| planner | read-only | context notes, roadmap, code | Drafts round questions, the brief and slices; checks for an existing standard solution first |
| implementer | all | the slice brief and the check command only | Makes it work with the check passing |
| standards-reviewer | edit, commit | the repo's `STANDARDS.md`, the architecture pages it points to, the diff | Applies the standards and commits fixes; sweeps for tests that lie |
| spec-reviewer | read-only | the task brief, the diff, the running behavior | Finds missing, wrong or out-of-scope behavior and returns gaps to the implementer |
| editor | edit | goal record, notes, README and docs | Keeps the context current: updates the goal record, culls notes the code now expresses, drafts pending coordinator edits, edits prose |

- **The implementer never loads standards or architecture.** Implementing already fills a
  context window: exploring, editing and debugging. Review is a lighter load, so the conventions
  go there. The implementer makes it work, and the standards-reviewer makes it good.
- **The reviewers commit fixes.** The architect then reads a finished artifact rather than a list of
  comments. Only what a reviewer can't resolve reaches the session brief, under Needs you.
- **The two review axes stay apart.** Standards and spec run in separate contexts, so code that
  follows every convention but builds the wrong thing can't pass on its conventions.
- **No profile pins a model.** Every profile inherits the session's model. New models don't
  reach every tier at the same time, and the architect picks the current generation as each one
  lands, which a pinned profile would hold back.

## Checks come first

Quality comes in three layers, cheapest first: automated checks, automated review, then human
review. Each layer exists to make the next one lighter, and the human layer concentrates on
one-way doors.

- **Code repositories.** Each one exposes a single deterministic `check`, such as
  `mise run check`, that runs build, vet, tests, lint, and the repository's own tools (sqlint,
  `split-check`). A mechanical rule becomes a check rather than prose. Examples are import
  boundaries (depguard, or the spikes' `split-check`) and version currency (a script).
- **Tests can lie.** A green check proves nothing when the tests are:
  - **tautological:** they restate the implementation, such as asserting a constant's value
  - **structure-sensitive:** they read source layout or internals rather than behavior
  - **unable to fail:** mocks remove the failure modes that matter

  The standards-reviewer sweeps for all three. The design cure is deep modules: small interfaces
  over large implementations, tested only at their seams.
- **Standards.** Each repository keeps a short `STANDARDS.md` of judgement calls only. It points
  to the few architecture pages that apply, and only the standards-reviewer reads it. `CLAUDE.md`
  holds navigation pointers only.

### Plugin evals

For claude-plugins, the automated checks are `claude plugin eval` suites. `harness-testing.md`
covers the details:

- Each plugin keeps an `evals/` directory.
- Cases come from observed failures, which the retro turns into eval cases.
- Graders are deterministic where possible: `tool_used`, `tool_order`, `regex` and
  `file_exists`. An `llm` grader judges only short outputs.
- `scripts/check.sh` stays the per-push CI check. `claude plugin eval --threshold 1.0` is a
  required step of the release script.
- The seed cases cover the new contract:
  - the implementer never reads `STANDARDS.md`
  - the standards-reviewer commits its fixes
  - the session brief has its required sections
  - plan-round questions are numbered, each with a recommendation

## The retro

The retro (`marathon-commands.md`) turns the architect's comments into changes to the system,
so no comment is written twice. It routes each finding to the cheapest layer that would have
caught it, in this order:

1. a check, or for a plugin an eval case
2. a judgement call in `STANDARDS.md`
3. a skill or profile change
4. a context pointer

## Walkthrough: one task

1. **Plan round.** `plan` on `v1.ai.experiment` puts two numbered questions to the architect, each
   with a recommendation. The architect answers "1 ok, 2 yes but drop resume". The planner
   writes the task brief and three slices into the goal record, and the architect approves.
2. **Build.**
   - **Implementer:** each slice gets a fresh implementer that sees the slice brief and
     `mise run check`, nothing more.
   - **Standards-reviewer:** it reads `STANDARDS.md`, commits four fixes, and deletes two
     tautological tests.
   - **Spec-reviewer:** it finds one missing behavior. The implementer adds it, and the
     spec-reviewer passes the re-review.
   - **Editor:** it updates the goal record to task 3/5 and culls a note the code now expresses.
3. **Session brief.** One screen covering Summary, Core changes, Evidence, Decided without you,
   and Merge danger (two-way door, spike repository only). Needs you reads "(none)". The
   architect accepts.
4. **Ship.** The PR opens with the brief as its body, CI passes, and it merges.
5. **Sync.** This task isn't the goal's last, so the goal record moves on to task 4. When the last
   task ships, its session syncs the goal into the coordinator (`marathon-goals.md`).
