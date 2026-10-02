# Session pipeline

A marathon session takes one task of one active goal from plan to merged work. The architect is
involved at two points: approving the task brief, which starts BUILD, and accepting the session
brief, which authorizes publishing and merging. Agents do everything in between.

```
PLAN   architect ⇄ planner, in rounds → task brief + slices     ✔ approve           [touch 1]
BUILD  per slice: implementer → check; then standards-reviewer,
       spec-reviewer (gaps → implementer), editor
BRIEF  session brief                                             ✔ accept / redirect [touch 2]
SHIP   commit the goal record → publish (body = brief) → merge → sync if last task
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
   - A **spike** session starts from the coordinator or the workspace, or from the root of the
     standalone project whose roadmap holds the spike's experiment, never from the spike's own
     repository.
2. Read the manifest (`references/manifest.md`): the project's own `context/roadmap.toml`, or the
   coordinator's. First check the coordinator's lock as step 5 does, then pull its default branch.
   Resolve every repository name the
   goal uses through `order`, `[workspace.paths]`, or a spike sub-goal's `path`.
3. Find the goal: the one the architect names, which must be in `active` and have a `root`; one
   without a `root` is set up by `plan` first (`commands/plan.md`). With none named, print
   the status digest (`commands/status.md`) and ask. A goal that isn't active is staged by `plan`
   first (`mechanics/goals.md`).
4. Find where the goal stands, in this order:
   - A plan file under `.claude/plans/` whose title names the goal holds an open plan round:
     continue it. START, then PLAN.
   - Otherwise read the goal record (`mechanics/goal-record.md`) from the root's working tree. If
     the root is on its default branch and a local branch named for the next unchecked task
     exists, read the record from that branch instead. Route on its State:
     - no record, or `idle`: the next unchecked task. START, then PLAN.
     - `building` or `handoff`: resume. START, then RESUME.
     - `brief ready`: START, then BRIEF.
5. Check the lock: every repository the task touches, other than the coordinator, is on its
   default branch with a clean working tree, or on the task's branch. A coordinator on a session
   branch (`plan-`, `experiment-`, `intake-`, `sync-`, or `retro-`) is held the same way. Anything
   else means another session holds it: stop and report.

### 2 · START

1. Fire `on-start`.
2. Read the goal record, the goal's and task's manifest entries, the capability map in
   `context/README.md`, and the notes the task cites. Load only what the task needs.

### 3 · PLAN

1. Enter plan mode. The plan file holds the session's planning state: title it with the goal and
   task, and keep each round, its answers, and the drafted brief in it.
2. Run plan rounds with the architect (`behavior/planning.md`, `references/briefs.md`). The
   planner profile finds the facts and drafts each round, then the task brief and its slices
   (`behavior/delegation.md`). When the record already holds an approved brief for this task,
   skip the rounds.
3. Present the task brief for approval. Nothing changes until the architect approves. **[touch 1]**
4. On approval, fire `on-build`.
5. Create the task's branch, named by the task's slug, from the fetched default branch in each
   repository the task touches. The coordinator gets no branch.
6. Log the plan into the goal record, creating the record if it doesn't exist: the brief under
   Task brief, and one line per settled question under Decisions. Set State to `building`, then
   delete the plan file.

### 3R · RESUME

1. Check out the task's branch in each touched repository.
2. Fire `on-build`.
3. Read the brief, Progress, and Handoff from the goal record. Finish any WIP slice first, then
   continue BUILD from the recorded position.

### 4 · BUILD

Run the build loop in `references/build.md` without stopping. Escalate only for what that file
names. A change the task owes the coordinator goes in the record's pending edits, never into the
coordinator during BUILD. When the context fills, run `reset` on your own (`commands/reset.md`).

### 5 · BRIEF

1. Write the session brief (`references/briefs.md`) to `.claude/briefs/<goal>.md` in the root's
   checkout, and set State to `brief ready`. If the root's `.gitignore` doesn't list
   `.claude/briefs/`, add the line on the task's branch first.
2. Show it, and wait for the architect. **[touch 2]**
   - **Accept**: continue with SHIP.
   - **Redirect**: treat the redirect as gaps, return to BUILD, and brief again. A redirect that
     changes the brief's behaviors returns to PLAN on the same branch.

### 6 · SHIP

1. Fire `on-ship`.
2. Update the goal record: check the task, add its decisions, set State to `idle` and Task to
   none, and clear Progress.
3. Fire `on-record`, then commit the record on the task's branch.
4. Publish each touched repository's branch with its `[remote] publish` command, lowest layer
   first, using the brief file as the body (for GitHub, `--body-file`). A repository with no
   `[remote] publish` merges its branch locally instead ("Branches and pull requests").
5. Merge each one with its `[remote] merge` command, lowest layer first, once its checks pass.
   Without a merge command, or when a check fails, set State to `handoff` with the next move
   ("merge" or the failing check), commit it on the branch, tell the architect, and stop.
6. Switch each repository back to its default branch, pull, and delete the local task branch.
   Delete the brief file.
7. When this was the goal's last task, sync (`mechanics/goals.md`, "Sync").

## Branches and pull requests

Every session lands its changes through a branch and a pull request: one branch, and one pull
request, in each repository it changes.

| Session | Branch | Pull request body |
|---------|--------|-------------------|
| `start` | the task's slug | the session brief |
| `plan` | `plan-<goal>`, or `plan-<topic>` when it names no goal | the approved round outcome |
| `experiment` | `experiment-<topic>` | the approved round outcome |
| `intake` | `intake-<topic>` | the approved round outcome |
| sync | `sync-<goal>` | the Sync section of the last task's session brief |
| `retro` | `retro-<topic>` | the ticked findings |

Each branch starts from the freshly pulled default branch, and each commit stages explicit paths,
never everything. The architect's approval of the outcome, the session brief, or the ticked
findings authorizes publishing and merging. Each repository publishes with its `[remote] publish`
command and merges with its `[remote] merge` command once its checks pass, lowest layer first,
then switches back to its default branch, pulls, and deletes the local branch. A merge that can't
happen leaves the branch checked out, which holds the repository until it merges: tell the
architect and stop.

A project with no `[remote] publish` branches the same way and, once the architect approves,
merges the branch locally into its default branch. A standalone project's own manifest, notes,
and configuration follow the same rule as a coordinator's.

Only goal-record bookkeeping commits straight to the root's default branch: the header lines
(State, Task, Branch) between tasks, the approved next brief `plan` writes, and deleting the record
once its sync pull request merges (`mechanics/goal-record.md`). Everything else, including the
manifest and its arrays, the notes, the repository catalog, the configuration, and `retro`'s
coordinator findings, goes through a pull request.

## Committing

Before every commit, any agent confirms the checkout is on the session's branch, or on the root's
default branch for goal-record bookkeeping. A checkout on any other branch means the session stops
and reports.

## How each command uses the pipeline

| Command | Stages it runs |
|---------|----------------|
| `start` | The full pipeline for one task |
| `plan` | LOCATE and START, then plan rounds that end in a `plan-<goal>` pull request in each changed repository (`commands/plan.md`) |
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
- The coordinator changes only through `plan` and sync, each through its own pull request.
- Every change lands through a session's pull request, except goal-record bookkeeping.
- On a code project, notes state only what validated work proved.
