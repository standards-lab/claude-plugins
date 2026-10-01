# The goal record

`context/goals/<goal>.md` is a goal's session record, named by the goal's dotted path (for
example, `context/goals/v1.ai.experiment.md`). It lives in the goal's `home` repository, and
every session on the goal reads it at LOCATE and updates it as it works. Each active goal keeps
its own, so goals running at the same time never write the same file. Anything that must outlast
the goal belongs in the notes, the documentation, or the pending coordinator edits that sync
applies.

The goal's first PLAN creates the record, and sync deletes it (`mechanics/goals.md`).

## Schema

```markdown
# goal · v1.ai.experiment

- **State:** building           # idle | planning | brief ready | building | handoff
- **Task:** harness-adapters    # the current task, or none between tasks
- **Branch:** harness-adapters  # the task's branch, one name in every repository it touches

## Tasks

1. [x] harness-driver
2. [ ] harness-adapters
3. [ ] local-subagents
4. [ ] intake

## Task brief · harness-adapters

(the approved brief, in the form of references/briefs.md, "Task brief")

## Progress

slices 2/3 committed · standards — · spec — · editor —

## Handoff

(only while State is handoff: the exact next move, a WIP commit whose check hasn't passed, and an
escalation awaiting an answer)

## Decisions

- harness-adapters: kept stdin JSONL over SSE — simpler, and the spike needs one transport.

## Pending coordinator edits

- roadmap: add backlog goal `harness-resume` (resumable sessions, ruled out of v1).
- references catalog: add the `pi` repository.
```

## Fields

- **State** drives `status` and LOCATE:
  - `idle`: between tasks; the next `start` plans the next unchecked task.
  - `planning`: a plan round is out, waiting on the architect.
  - `brief ready`: BUILD is done and the session brief waits on the architect.
  - `building`: a session is in BUILD.
  - `handoff`: a session stopped partway; the next `start` resumes from Handoff.
- **Tasks** lists the goal's tasks in the order they run, each checked by the SHIP that merges
  it. The manifest keeps a task's table until its goal syncs, so the record is where
  progress shows. A goal with sub-goals lists only its own tasks.
- **Task brief** is the current task's approved brief. The next task's brief replaces it.
- **Progress** records the build loop's position (`references/build.md`).
- **Decisions** records what each task decided without the architect, one line each, so a later
  task doesn't decide it again. It also records rejected alternatives (`behavior/planning.md`).
- **Pending coordinator edits** collects every change the goal owes the coordinator: notes,
  catalog rows, workspace `order`, the manifest, the architecture layer. Sync applies them in one
  commit.

## Where it is committed

- **Home is a member repository**: the record changes on the task's branch, and merges with the
  task's pull request.
- **Home is the coordinator**: the record changes in short, direct commits on the coordinator's
  default branch, like every other coordinator commit (`mechanics/goals.md`).
