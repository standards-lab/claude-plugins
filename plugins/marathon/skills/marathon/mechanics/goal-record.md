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

(the approved brief's fields, as markdown in the shape of references/briefs.md, "Task brief")

## Progress

slices 2/3 committed · standards — · spec — · editor —
- escalation, slice 2: raise pi's floor to 0.9 for RPC mode? yes

## Handoff

(the exact next move, a WIP commit whose check hasn't passed, and an escalation awaiting an answer)

## Decisions

- transport: kept stdin JSONL over SSE — simpler, and the spike needs one transport.
- transport: held pgx at v5.6.0: v6.0.0 changes Query; its own task, pgx-v6.

## Pending edits

- coordinator · roadmap/harness-resume.toml: create the goal (resumable sessions, ruled out of
  v1); roadmap.toml: add it to `backlog`.
- coordinator · references catalog: add the `pi` repository.
- architecture · a page on model tiers, from `context/model-tiers.md` in this repository.
```

The header lines, Tasks, Decisions, and Pending edits are always present. Task brief and Progress
appear while a task is planned or in progress, and Handoff only while State is `handoff`.
`lint-context` holds Tasks to the goal file's `order`, and Task brief, Decisions, and Pending edits
to their budgets.

## Fields

- **State** drives `status` and LOCATE:
  - `idle`: between tasks. The next `start` plans the next unchecked task, or, when Task brief
    already holds that task's brief, takes it to approval (`mechanics/pipeline.md`, 3 · PLAN).
  - `building`: a session is in BUILD.
  - `brief ready`: BUILD is done and the session brief waits on the architect.
  - `handoff`: a session stopped partway; the next `start` resumes from Handoff.

  An open plan round lives in the session's plan file, not here (`mechanics/pipeline.md`, 3 ·
  PLAN).
- **Tasks** mirrors the `order` of the goal's file (`references/manifest.md`), each task checked
  by the SHIP that merges it. The goal's file keeps a task's table until the goal syncs, so the
  record is where progress shows. A goal with sub-goals lists only its own tasks. A task that
  joins `order` while the goal is active, such as one `intake` writes, joins Tasks at the goal's
  next `plan` or `start`.
- **Task brief** is the current task's approved brief, stored as markdown in the shape of
  `references/briefs.md`, "Task brief", unfenced. The next task's brief replaces it. Briefs
  already stored fenced stay as they are until the record is next rewritten.
- **Progress** records the build loop's position (`references/build.md`), and each escalation
  and redirect as it happens, one line each with the architect's answer, so a handoff loses none
  and the session brief lists them all.
- **Handoff** gives the exact next move while State is `handoff`: the BUILD position `reset`
  wrote (`commands/reset.md`); "merge", or the failing check, when SHIP can't merge; "tag
  <names>" when the context fills during a release (`mechanics/releasing.md`); "merge, then
  sync" or "tag <names>, then sync" on the goal's last task; or "merge `sync-<goal>`, then
  delete the record" when a sync pull request can't merge (`mechanics/goals.md`, "Sync").
- **Decisions** keeps only what binds later tasks, one line each, so a later task doesn't decide
  it again: a settled question a later task relies on, a rejected alternative
  (`behavior/planning.md`), a hold, and a falsified assumption
  (`references/context-engineering.md`). A hold is a decision of the form "<task>: held <item>
  at <pin>: <reason>", whose reason names the latest version it declined; PLAN lists it as a fact
  and asks again only once a later version is out. A decision the merged code or its
  documentation now expresses no longer binds, and git history keeps it: at BUILD the editor
  drops it on the task's branch (`references/build.md`, "The loop"), and the session brief lists
  the drop (`references/briefs.md`, "Session brief").
- **Pending edits** collects every change the goal owes a repository outside its own work, each
  prefixed with that repository: the coordinator's notes, catalog, workspace `order`, roadmap
  index, and goal files, or a page in the architecture repository. Sync applies them.

## Where it is committed

The root is a repository the goal locks, so the record changes on the task's branch and merges
with the task's pull request. Only bookkeeping commits it straight to the root's default branch
(`mechanics/pipeline.md`, "Branches and pull requests").
