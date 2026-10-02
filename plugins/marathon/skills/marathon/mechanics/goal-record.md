# The goal record

`context/goals/<goal>.md` is a goal's session record, named by the goal's dotted path (for
example, `context/goals/experiment.ai.spike-local-subagents.md`). It lives in the goal's `root`
repository, and every session on the goal reads it at LOCATE and updates it as it works. Each
active goal keeps its own, so goals running at the same time never write the same file. An
experiment goal keeps none; each of its spike sub-goals keeps its own in the spike's repository.
Anything that must outlast the goal belongs in the notes, the documentation, or the pending edits
that sync applies.

The goal's first approved brief creates the record, and sync deletes it last, once its sync pull
request merges (`mechanics/goals.md`).

## Schema

```markdown
# goal · experiment.ai.spike-local-subagents

- **State:** building    # idle | building | brief ready | handoff
- **Task:** transport    # the current task, or none between tasks
- **Branch:** transport  # the task's branch, one name in every repository it touches

## Tasks

1. [x] harness
2. [ ] transport
3. [ ] fan-out

## Task brief · transport

(the approved brief, in the form of references/briefs.md, "Task brief")

## Progress

slices 2/3 committed · standards — · spec — · editor —

## Handoff

(the exact next move, a WIP commit whose check hasn't passed, and an escalation awaiting an answer)

## Decisions

- transport: kept stdin JSONL over SSE — simpler, and the spike needs one transport.

## Pending edits

- coordinator · roadmap: add backlog goal `harness-resume` (resumable sessions, ruled out of v1).
- coordinator · references catalog: add the `pi` repository.
- architecture · a page on model tiers, from `context/model-tiers.md` in this repository.
```

The header lines, Tasks, Decisions, and Pending edits are always present. Task brief and Progress
appear while a task is planned or in progress, and Handoff only while State is `handoff`.

## Fields

- **State** drives `status` and LOCATE:
  - `idle`: between tasks. The next `start` plans the next unchecked task, or, when Task brief
    already holds that task's brief, presents it for approval.
  - `building`: a session is in BUILD.
  - `brief ready`: BUILD is done and the session brief waits on the architect.
  - `handoff`: a session stopped partway; the next `start` resumes from Handoff.

  An open plan round lives in the session's plan file, not here (`mechanics/pipeline.md`, 3 ·
  PLAN).
- **Tasks** lists the goal's tasks in the order they run, each checked by the SHIP that merges
  it. The manifest keeps a task's table until its goal syncs, so the record is where progress
  shows. A goal with sub-goals lists only its own tasks.
- **Task brief** is the current task's approved brief. The next task's brief replaces it.
- **Progress** records the build loop's position (`references/build.md`).
- **Handoff** gives the exact next move while State is `handoff`: the BUILD position `reset`
  wrote (`commands/reset.md`); "merge", or the failing check, when SHIP can't merge; "tag
  <names>" when the context fills during a release, naming the tags not yet released
  (`mechanics/pipeline.md`, "Releasing"); "merge, then sync" or "tag <names>, then sync" on the
  goal's last task; or "merge `sync-<goal>`, then delete the record" when a sync pull request
  can't merge (`mechanics/goals.md`, "Sync").
- **Decisions** logs what each plan round settled and what each task decided without the
  architect, one line each, so a later task doesn't decide it again. It also records rejected
  alternatives (`behavior/planning.md`).
- **Pending edits** collects every change the goal owes a repository outside its own work, each
  prefixed with that repository: the coordinator's notes, catalog, workspace `order`, and
  manifest, or a page in the architecture repository. Sync applies them.

## Where it is committed

The root is a repository the goal locks, so the record changes on the task's branch and merges
with the task's pull request. Only bookkeeping commits it straight to the root's default branch:
the header lines between tasks, the next brief `plan` writes, and deleting the record once its
sync pull request merges (`mechanics/pipeline.md`, "Branches and pull requests").
