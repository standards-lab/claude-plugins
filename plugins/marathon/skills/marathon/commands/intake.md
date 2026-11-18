# marathon intake

Take in a finished experiment: consolidate its spikes' answers, decide with the architect what each
goal it serves builds from them, and close the experiment. `intake` runs the experiment's `intake`
task, its only one; the argument names the experiment, as in `intake experiment.ai`. A spike that
works is evidence, and the intake is where it becomes a decision.

Read first, before acting:

- [`../mechanics/pipeline.md`](../mechanics/pipeline.md)
- [`../behavior/planning.md`](../behavior/planning.md)
- [`../behavior/delegation.md`](../behavior/delegation.md)
- [`../references/briefs.md`](../references/briefs.md)
- [`../mechanics/goals.md`](../mechanics/goals.md)
- [`../references/manifest.md`](../references/manifest.md)
- [`../references/context-engineering.md`](../references/context-engineering.md)

## When it runs

`intake` runs only when no spike sub-goal of the experiment has a file left: each spike's sync
has deleted it (`mechanics/goals.md`, "Sync"). When any remains, stop and report each one
with its array and, when it has a record, its State.

`intake` runs LOCATE and START (`mechanics/pipeline.md`), which fires `on-start`. LOCATE finds
the experiment's file, not an entry in `active`: an experiment goal is never listed. The
experiment keeps no goal record, and the intake writes none (`mechanics/hooks.md`).

## State

The plan file holds the intake's whole state: titled with the experiment, it keeps each round,
the approved outcome, and how far the edits and the archiving have gone. A broken intake resumes
from it. LOCATE finds a plan file under `.claude/plans/` whose title names the experiment, and
the intake continues from the first step not yet done: the open round, the `intake-<topic>`
branch, its merge, or the remotes still to archive. A coordinator left on `intake-<topic>` is the
intake's own branch, and it resumes there (`mechanics/pipeline.md`, 1 · LOCATE).

## Rounds

Enter plan mode and run plan rounds with the architect, led by the spikes' answers
(`references/briefs.md`, "Intake round"), each presented as `behavior/planning.md`, "Ask for
decisions, never for facts", says. The planner finds the facts.

- **The first round** names the goals the experiment serves, which no manifest field names: the
  experiment's summary and its intake task's entry inform the proposal. It consolidates every
  spike's answer section from the note the experiment cites in its `context`
  (`references/briefs.md`, "Answer section").
- **Later rounds** decide, for each served goal, what it builds from the answers: the tasks to
  write into it, and the note each answer section folds into.
- **The rounds also name** the spike remotes to archive and what the repository catalog records
  of each spike. A spike the round keeps open, such as a repository that outlives the
  experiment, stays open.

## Edits

Once the architect approves the outcome (`mechanics/pipeline.md`, "Invariants"), at the
coordinator, in one commit on an `intake-<topic>` branch:

- Write the decided tasks into each served goal's file (`references/manifest.md`). The intake
  edits only goal files' tasks and never another goal's record: an active served goal takes them
  into its record at its next `plan` or `start` (`mechanics/goal-record.md`, "Fields").
- Fold each answer section into the served goals' notes at the coordinator, creating a note a
  served goal then cites in its `context` when it has none there, and remove the answer section
  from the experiment's note. Sharpen or delete that note as the round decided
  (`references/context-engineering.md`).
- Delete the experiment's file, and the file of any ancestor left empty (`mechanics/goals.md`,
  "Sync").
- Update the repository catalog as the round decided, recording each spike it archives as
  archived.

Publish and merge the branch with the approved round outcome as its body (`mechanics/pipeline.md`,
"Branches and pull requests").

## Archive

After the pull request merges, archive each spike remote the round named. Check each one first
and skip any already archived: on GitHub, `gh repo view <remote> --json isArchived`, then
`gh repo archive <remote> --yes`; on another platform, its equivalent. Record each archived
remote in the plan file as it goes. When one can't be archived, tell the architect and stop; the
next `intake` resumes there.

Delete the plan file once every named remote is archived.
