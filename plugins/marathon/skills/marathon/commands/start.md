# marathon start

Run one task of an active goal: plan it if it has no approved brief, build it, brief the
architect, and ship it. When the task is the goal's last, the same session syncs the goal.
`start` runs the full pipeline (`mechanics/pipeline.md`); the argument after `start` names the
goal, as in `start v1.data`.

On a spike, `start experiment.<topic>.<spike>` runs the spike's tasks, its path steps, one per
session. An experiment goal has only its intake, which `marathon intake` runs
(`commands/intake.md`): `start` on `experiment.<topic>` refuses and points there.

## Plan

LOCATE routes an `idle` goal to PLAN. Without an approved brief for the next task, plan the
task in rounds to an approved brief and slice list (`behavior/planning.md`,
`references/briefs.md`). With one, PLAN runs the currency commands first and then presents the
brief, opening a round only when something trails that the brief doesn't cover
(`mechanics/pipeline.md`, 3 · PLAN). Add detail to the notes the task touches only as far as the
brief needs.

Branch slug: the task.

## Build

Run the build loop (`references/build.md`) without stopping:

- On a **code** project, each slice adds behavior through the test seam, with the check passing.
- On a **context** project, each slice is the deliverable itself, with the consistency script
  passing and the prose read for coherence.

The editor keeps the goal record, the notes, and the documentation current as part of the loop, so
the brief describes a finished state.

## Brief

Write and show the session brief. Accepting it authorizes publishing and merging; a redirect
returns to BUILD, or to PLAN when it changes the brief's behaviors.

On the goal's last task, the brief has a Sync section, and accepting it also authorizes the sync
(`references/briefs.md`, "Session brief"). A spike's last brief gives its answer
(`references/briefs.md`, "Answer section").

## Ship

Update and commit the goal record, then publish with the brief as the pull request's body and
merge (`mechanics/pipeline.md`, 6 · SHIP). When the task brief has a Release line, SHIP also tags
the release, fixing forward until the planned version releases (`mechanics/pipeline.md`,
"Releasing"). On the goal's last task, sync (`mechanics/goals.md`).

If the context fills before BRIEF, `reset` runs on its own and the next `start` on the goal
resumes.
