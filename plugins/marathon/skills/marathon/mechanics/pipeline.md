# Session pipeline

A marathon session takes one task of one active goal from plan to merged work. The architect is
involved at two points: approving the task brief, which starts BUILD, and accepting the session
brief, which authorizes publishing and merging. Agents do everything in between.

```
PLAN   architect ⇄ planner, in rounds → task brief + slices     ✔ approve           [touch 1]
BUILD  per slice: implementer → check; then standards-reviewer,
       spec-reviewer (gaps → implementer), editor
BRIEF  session brief                                             ✔ accept / redirect [touch 2]
SHIP   update the goal record (sync if last task) → publish (body = brief) → merge
```

Each playbook under `commands/` says which stages it runs. Extension hooks fire only at the points
this file names (`mechanics/hooks.md`).

## Stages

### 1 · LOCATE

1. Identify the directory:
   - A **standalone project** has its own top-level `context/`, and no sibling project declares
     itself coordinator.
   - A **workspace** is a directory of projects, one of which declares itself coordinator in its
     `.claude/marathon.toml` (`mechanics/configuration.md`). Starting at the workspace root and
     starting inside a member are the same case.
2. Read the manifest (`references/manifest.md`): the project's own `context/roadmap.toml`, the
   coordinator's, or for an experiment, that of the project its `[experiment] serves` names. First
   pull the default branch of the repository that holds the manifest.
3. Find the goal: the one the architect names, which must be in `active`. With none named, print
   the status digest (`commands/status.md`) and ask. A goal that isn't active is staged by `plan`
   first (`mechanics/goals.md`).
4. Read the goal record (`mechanics/goal-record.md`) from its home's working tree, and route on
   its State:
   - no record, or `idle`: the next unchecked task. Continue with START, then PLAN.
   - `planning`: continue the open plan round. START, then PLAN.
   - `building` or `handoff`: resume. START, then RESUME.
   - `brief ready`: START, then BRIEF.
5. Check the lock: every repository the task touches is on its default branch with a clean
   working tree, apart from this goal's own uncommitted record, or on the task's branch. Anything else means another session holds it: stop and
   report.

### 2 · START

1. Fire `on-start`.
2. Read the goal record, the goal's and task's manifest entries, the capability map in
   `context/README.md`, and the notes the task cites. Load only what the task needs.

### 3 · PLAN

1. Enter plan mode and set State to `planning` in the goal record's working tree. A record in a
   member home stays uncommitted until step 5's branch carries it.
2. Run plan rounds with the architect (`behavior/planning.md`, `references/briefs.md`). The
   planner profile finds the facts and drafts each round, then the task brief and its slices
   (`behavior/delegation.md`).
3. Present the task brief for approval. Nothing changes until the architect approves. **[touch 1]**
4. On approval, fire `on-build`.
5. Create the task's branch, named by the task's slug, from the fetched default branch in each
   repository the task touches, except the coordinator, which takes direct commits.
6. Write the brief into the goal record, creating the record if it doesn't exist, and set State to
   `building`.

### 3R · RESUME

1. Check out the task's branch in each touched repository.
2. Fire `on-build`.
3. Read the brief, Progress, and Handoff from the goal record. Finish any WIP slice first, then
   continue BUILD from the recorded position.

### 4 · BUILD

Run the build loop in `references/build.md` without stopping. Escalate only for what that file
names. When the context fills, run `reset` on your own (`commands/reset.md`).

### 5 · BRIEF

1. Write the session brief (`references/briefs.md`) to `.claude/brief.md` in the home
   repository's checkout, and set State to `brief ready`. If that repository's `.gitignore`
   doesn't list `.claude/brief.md`, add the line on the task's branch first.
2. Show it, and wait for the architect. **[touch 2]**
   - **Accept**: continue with SHIP.
   - **Redirect**: treat the redirect as gaps, return to BUILD, and brief again. A redirect that
     changes the brief's behaviors returns to PLAN on the same branch.

### 6 · SHIP

1. Fire `on-ship`.
2. Update the goal record: check the task, add its decisions, set State to `idle` and Task to
   none, and clear Progress. When this is the goal's last task, sync instead
   (`mechanics/goals.md`, "Sync").
3. Fire `on-record`, then commit the record where it lives (`mechanics/goal-record.md`).
4. Publish each touched repository's branch with its `[remote] publish` command, lowest layer
   first, using `.claude/brief.md` as the body (for GitHub, `--body-file`).
5. Merge each one with its `[remote] merge` command, lowest layer first, once its checks pass.
   Without a merge command, stop here and tell the architect the pull requests are open.
6. Switch each repository back to its default branch, pull, and delete the local task branch.
   Delete `.claude/brief.md`. A last task finishes its sync at the coordinator now.

## Committing

Before every commit, any agent confirms the checkout is on the task's branch, on the
coordinator's default branch for a coordinator commit, or on the `retro-<topic>` branch of a
`retro`. A checkout on any other branch means the
session stops and reports.

## How each command uses the pipeline

| Command | Stages it runs |
|---------|----------------|
| `start` | The full pipeline for one task |
| `plan` | LOCATE and START, then plan rounds that end in direct coordinator commits (`commands/plan.md`) |
| `experiment` | `plan`'s stages, plus setting up the spike repository (`commands/experiment.md`) |
| `retro` | LOCATE and START, then the retro (`commands/retro.md`) |
| `status` | Reads only; no stage changes anything (`commands/status.md`) |
| `reset` | Writes the handoff during BUILD (`commands/reset.md`) |
| `init` | Its own setup (`commands/init.md`) |

## Invariants

- A session advances one task of one goal, on one branch name in each touched repository.
- Every session starts and ends at the goal record, so the next session learns where to begin
  from the repository, not the conversation.
- Nothing changes before the architect approves the task brief, except a RESUME of an approved
  one.
- Nothing is published before the architect accepts the session brief.
- On a code project, notes state only what validated work proved.
