# marathon start

Run one task of an active goal: plan it if it has no approved brief, build it, brief the
architect, and ship it. When the task is the goal's last, the same session syncs the goal.
`start` runs the full pipeline (`mechanics/pipeline.md`); the argument after `start` names the
goal, as in `start v1.data`.

On a spike, `start experiment.<topic>.<spike>` runs the spike's tasks, its path steps, one per
session. An experiment goal has only its intake, which `marathon intake` runs
(`commands/intake.md`): `start` on `experiment.<topic>` refuses and points there.

Read first, before acting:

- [`../mechanics/pipeline.md`](../mechanics/pipeline.md)
- [`../behavior/planning.md`](../behavior/planning.md)
- [`../behavior/delegation.md`](../behavior/delegation.md)
- [`../mechanics/goal-record.md`](../mechanics/goal-record.md)
- [`../references/briefs.md`](../references/briefs.md)
- [`../references/briefs/plan-round.md`](../references/briefs/plan-round.md)
- [`../references/briefs/task-brief.md`](../references/briefs/task-brief.md)
- [`../references/briefs/slice-brief.md`](../references/briefs/slice-brief.md)
- [`../references/briefs/session-brief.md`](../references/briefs/session-brief.md)
- [`../references/briefs/answer-section.md`](../references/briefs/answer-section.md)
- [`../references/build.md`](../references/build.md)

## Plan

LOCATE routes an `idle` goal to PLAN. Without an approved brief for the next task, plan the task in
rounds to an approved brief and slice list (`behavior/planning.md`,
`references/briefs/plan-round.md`, `references/briefs/task-brief.md`). With one, PLAN takes it to
approval (`mechanics/pipeline.md`, 3 · PLAN). Add detail to the notes the task touches only as far
as the brief needs.

Branch slug: the task.

## Build

Run the build loop (`references/build.md`) without stopping, briefing each slice
(`references/briefs/slice-brief.md`):

- On a **code** project, each slice adds behavior through the test seam, with the check passing.
- On a **context** project, each slice is the deliverable itself, with the consistency script
  passing and the prose read for coherence.

The editor keeps the goal record, the notes, and the documentation current as part of the loop, so
the brief describes a finished state.

## Brief

Write and show the session brief, and act on the architect's accept or redirect
(`mechanics/pipeline.md`, 5 · BRIEF). On the goal's last task, the brief has a Sync section
(`references/briefs/session-brief.md`). A spike's last brief gives its answer
(`references/briefs/answer-section.md`).

## Ship

Update and commit the goal record, then publish with the brief as the pull request's body and
merge (`mechanics/pipeline.md`, 6 · SHIP). When the task brief has a Release line, SHIP also tags
the release, fixing forward until the planned version releases (`mechanics/releasing.md`). On
the goal's last task, sync (`mechanics/goals.md`).

Once the task merges, end the last message with the prompt that starts the next session, fenced as
code: `/marathon:marathon start <goal>` for the goal's next task or the goal the sync staged;
`/marathon:marathon plan <goal>` for a goal the sync only proposed; or
`/marathon:marathon intake experiment.<topic>` when a spike's sync leaves no planned spike. A
session that stops in `handoff` ends without one, and the session brief never carries it.

If the context fills before BRIEF, the session hands off (`references/build.md`, "Handoff").
