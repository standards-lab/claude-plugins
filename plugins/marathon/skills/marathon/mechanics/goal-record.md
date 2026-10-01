# The goal record

`context/goals/<goal>.md` is a goal's session record, named by the goal's dotted path (for
example, `context/goals/experiment.ai.md`). It lives in the goal's `root` repository, and every
session on the goal reads it at LOCATE and updates it as it works. Each active goal keeps its own,
so goals running at the same time never write the same file. Anything that must outlast the goal
belongs in the notes, the documentation, or the pending edits that sync applies.

The goal's first approved brief creates the record, and sync deletes it last
(`mechanics/goals.md`).

## Schema

```markdown
# goal · experiment.ai

- **State:** building                # idle | building | brief ready | handoff
- **Task:** spike-local-subagents    # the current task, or none between tasks
- **Branch:** spike-local-subagents  # the task's branch, one name in every repository it touches

## Tasks

1. [x] spike-harness-driver
2. [ ] spike-local-subagents
3. [ ] personal-agents
4. [ ] intake

## Task brief · spike-local-subagents

(the approved brief, in the form of references/briefs.md, "Task brief")

## Progress

slices 2/3 committed · standards — · spec — · editor —

## Handoff

(the exact next move, a WIP commit whose check hasn't passed, and an escalation awaiting an answer)

## Decisions

- spike-local-subagents: kept stdin JSONL over SSE — simpler, and the spike needs one transport.

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
- **Decisions** logs what each plan round settled and what each task decided without the
  architect, one line each, so a later task doesn't decide it again. It also records rejected
  alternatives (`behavior/planning.md`).
- **Pending edits** collects every change the goal owes a repository outside its own work, each
  prefixed with that repository: the coordinator's notes, catalog, workspace `order`, and
  manifest, or a page in the architecture repository. Sync applies them.

## Where it is committed

The root is a repository the goal locks, so the record changes on the task's branch and merges
with the task's pull request. Between tasks, `plan` may write the next brief into it as a direct
commit on the root's default branch, and sync deletes it the same way.
