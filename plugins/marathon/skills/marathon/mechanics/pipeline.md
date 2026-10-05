# Session pipeline

A marathon session takes one task of one active goal from plan to merged work. The architect is
involved at two points: approving the task brief, which starts BUILD, and accepting the session
brief, which authorizes publishing and merging. Agents do everything in between.

```
PLAN   architect ⇄ planner, in rounds → task brief + slices     ✔ approve           [touch 1]
BUILD  per slice: implementer → check; then standards-reviewer,
       spec-reviewer (gaps → implementer), editor
BRIEF  session brief                                             ✔ accept / redirect [touch 2]
SHIP   commit the goal record → publish (body = brief) → merge → tag if Release → sync if last
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
   coordinator's. First check the coordinator's lock as step 5 does, then pull its default branch,
   unless the session resumes on its own branch there.
   Resolve every repository name the goal uses through `order`, `[workspace.paths]`, or a spike
   sub-goal's `path`.
3. Find the goal: the one the architect names, which must be in `active` and have a `root`; one
   without a `root` is set up by `plan` first (`commands/plan.md`). With none named, print the
   status digest (`commands/status.md`) and ask. A goal that isn't active is staged by `plan`
   first (`mechanics/goals.md`), except an experiment goal, `experiment.<topic>`, which is never
   staged: only `intake` runs on it (`commands/intake.md`), and `start` refuses and points there.
   A goal no longer listed, because its coordinator sync pull request merged, is found when its
   record in its root is in `handoff` with the next move "merge `sync-<goal>`, then delete the
   record": its root is the one the last version of the manifest that lists the goal names (the
   coordinator's history of `context/roadmap.toml`). START, then RESUME.
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
   default branch with a clean working tree, on the task's branch or its `<slug>-fix` branch
   ("Releasing"), or, resuming a sync, on `sync-<goal>`. A coordinator on a session branch
   ("Branches and pull requests") is held the same way, unless the branch is the resuming
   session's own, the one the table names for this goal or topic: the session resumes on it.
   Anything else means another session holds it: stop and report.

### 2 · START

1. Fire `on-start`.
2. Read the goal record, the goal's and task's manifest entries, the capability map in
   `context/README.md`, and the notes the task cites. Load only what the task needs.

### 3 · PLAN

1. Enter plan mode. The plan file holds the session's planning state: title it with the goal and
   task, and keep each round, its answers, and the drafted brief in it.
2. Run plan rounds with the architect (`behavior/planning.md`, `references/briefs.md`). Print
   each round in full in the reply and end the turn there, never through AskUserQuestion
   (`behavior/planning.md`, "Ask for decisions, never for facts"). The planner profile finds the
   facts and drafts each round, then the task brief and its slices
   (`behavior/delegation.md`). Before round 1, the planner runs `[project] currency` on the
   default branch of each repository the task touches that declares it
   (`mechanics/configuration.md`), and reads the release notes of what trails. Round 1 notes each
   touched repository with no currency command, and reports a command that failed as a fact; a
   failure doesn't block. When the record already holds an approved brief for this task, the
   currency commands still run first: skip the rounds unless something trails that the brief
   doesn't cover, and in that case open one round on it before approval.
3. Present the task brief for approval through ExitPlanMode, with the full brief in the plan
   file. Nothing changes until the architect approves. **[touch 1]**
4. On approval, fire `on-build`.
5. Create the task's branch, named by the task's slug, from the fetched default branch in each
   repository the task touches. The coordinator gets no branch.
6. Log the plan into the goal record, creating the record if it doesn't exist: the brief under
   Task brief, and one line per settled question under Decisions. Set State to `building`, then
   delete the plan file.

### 3R · RESUME

1. Check out the task's branch in each touched repository that hasn't merged it, any
   `<slug>-fix` branch a release left unmerged ("Releasing"), or, resuming a sync, the
   `sync-<goal>` branch in each repository that hasn't merged it.
2. Fire `on-build`, unless the session resumes only a sync or a release. A release fires it
   whenever it opens or resumes a `<slug>-fix` branch ("Releasing").
3. Read the brief, Progress, and Handoff from the goal record. Finish any WIP slice first, then
   continue BUILD from the recorded position. Two next moves resume past BUILD instead:
   - "tag <names>" resumes the release at SHIP ("Releasing"), one repository at a time, from the
     position Handoff records for it: at step 1 for a repository whose task branch hasn't merged;
     at step 2 for one merged with ci not yet green; at its `<slug>-fix` branch's publish, merge,
     or `[remote] ci` step, then the tag step, for one fixing forward; at the fix for one whose
     gate failed; at the failed-release fix for one whose tag was pushed and whose release
     failed; and at the tag step only for one merged and green.
   - "merge `sync-<goal>`, then delete the record" resumes the sync at its remaining steps: merge
     each `sync-<goal>` branch not yet merged, then delete the record and release the lock
     (`mechanics/goals.md`, "Sync", steps 4 to 6).

### 4 · BUILD

Run the build loop in `references/build.md` without stopping. Escalate only for what that file
names. A change the task owes the coordinator goes in the record's pending edits, never into the
coordinator during BUILD. When the context fills, run `reset` on your own (`commands/reset.md`).

### 5 · BRIEF

1. Write the session brief (`references/briefs.md`) to `.claude/briefs/<goal>.md` in the root's
   checkout, and set State to `brief ready`. On the goal's last task, the brief has a Sync
   section. If the root's `.gitignore` doesn't list `.claude/briefs/`, add the line on the task's
   branch first.
2. Show it, and wait for the architect. **[touch 2]**
   - **Accept**: continue with SHIP. On the goal's last task, accepting also authorizes the sync
     and the staging its Sync section shows.
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
   ("merge" or the failing check), commit it on the branch, tell the architect, and stop. When
   the task brief has a Release line, merge and tag one repository at a time instead
   ("Releasing").
6. Switch each repository back to its default branch, pull, and delete the local task branch,
   where Releasing hasn't already. Delete the brief file.
7. When this was the goal's last task, sync (`mechanics/goals.md`, "Sync").

### Releasing

A task brief's Release line lists the tags the task releases (`references/briefs.md`, "Task
brief"), and accepting the session brief authorizes them. SHIP then takes each repository in
turn, lowest layer first, and finishes its release before the next repository's merge:

1. Merge the branch with `[remote] merge`, switch to the default branch, pull, and delete the
   local task branch.
2. When `[remote] ci` is set, run it on the merge commit (`mechanics/configuration.md`).
   Without it, tagging follows the merge directly.
3. Tag each of the repository's Release tags, the base artifact before its sub-modules:
   1. Check that the tag isn't on the remote, unless it is this release's failed tag (below),
      and that its version matches the artifact's
      version as the repository records it, in its manifest or its CHANGELOG's top heading.
   2. When `[project] gate` is set, run it on the merge commit with the tag as its one argument
      (`mechanics/configuration.md`). Exit 0 lets the tag be pushed; nonzero holds it (below).
      A resumed release runs it again before each tag it hasn't pushed.
   3. Create the tag annotated "<artifact> <version>" on the merge commit, and push that tag
      alone.
   4. When the repository has a release workflow, confirm the run the tag started succeeds.

No tag is pushed while the default branch is red or its gate fails. The session fixes forward
until the planned version releases:

- **A red default branch** after the merge, a failed gate, or a version that doesn't match, is
  fixed on a `<slug>-fix` branch from the default branch, published and merged like the task's
  branch, with `[remote] ci` and the gate run again on its merge commit. Then tag. Opening or
  resuming a fix branch fires `on-build` first.
- **A failed release**, a release workflow that fails after the tag is pushed, is fixed the
  same way. Then delete the tag on the remote and locally, and create and push it again at the
  same version on the fix's merge commit.
- **A released tag** is never re-cut. A tag whose release failed is this release's tag when the
  Release line names it and this session or an earlier session of the same release pushed it, as
  Handoff records: it is deleted and pushed again at the same version, as above. Only a tag on the
  remote that no session of this release pushed stops the release with an escalation.

The release always ends at the planned version. It stops for the architect only for a decision
the brief doesn't cover, as an escalation (`references/build.md`). When the context fills, set
State to `handoff` with the next move "tag <names>", naming the tags not yet released and each
Release repository's position: merged with ci not yet green, merged and green, on `<slug>-fix`
at its publish, merge, or ci step, its gate failed, its tag pushed and its release failed, or
released. Commit it where the root stands: on the root's task branch while that branch hasn't
merged, pushed to that branch, whose pull request is already published, so the record lands with
the merge; or on the root's default branch as bookkeeping once it has. Then stop; the next
`start` resumes the release (3R · RESUME).

A resumed release ends at the goal record. Once the last named tag releases, set State to `idle`,
Task and Branch to none, drop Handoff, and check the task if SHIP's record update hadn't landed,
as a bookkeeping commit on the root's default branch. On the goal's last task, sync follows
(`mechanics/goals.md`, "Sync").

## Branches and pull requests

Every session lands its changes through a branch and a pull request: one branch, and one pull
request, in each repository it changes.

| Session | Branch | Pull request body |
|---------|--------|-------------------|
| `start` | the task's slug | the session brief |
| `plan` | `plan-<goal>`, or `plan-<topic>` when it names no goal | the approved round outcome |
| `experiment` | `experiment-<topic>` | the approved round outcome |
| `intake` | `intake-<topic>` | the approved round outcome |
| `start`, fixing a release | `<slug>-fix` | the failure and its fix |
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
(State, Task, Branch) between tasks; State `handoff` with its Handoff section for a release's
"tag <names>" once the root's task branch has merged, and for a stuck sync's "merge
`sync-<goal>`, then delete the record"; the return to `idle` once a resumed release's last tag
releases ("Releasing"); the approved next brief `plan` writes; and deleting the record once its
sync pull request merges (`mechanics/goal-record.md`). Everything else, including the manifest and its arrays, the notes,
the repository catalog, the configuration, and `retro`'s coordinator findings, goes through a
pull request. The one other direct commit is a new repository's first: `init`'s commit, pushed to
create the default branch that later pull requests merge into (`commands/init.md`).

## Committing

Before every commit, any agent confirms the checkout is on the session's branch, or on the root's
default branch for goal-record bookkeeping. A checkout on any other branch means the session stops
and reports.

## How each command uses the pipeline

| Command | Stages it runs |
|---------|----------------|
| `start` | The full pipeline for one task |
| `plan` | LOCATE and START, then plan rounds that end in a `plan-<goal>` pull request in each changed repository (`commands/plan.md`) |
| `experiment` | `plan`'s stages, creating the experiment goal and its planned spikes in an `experiment-<topic>` pull request (`commands/experiment.md`) |
| `intake` | LOCATE and START, then intake rounds that end in an `intake-<topic>` pull request, then archive the spike remotes (`commands/intake.md`) |
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
- No tag is pushed while its repository's default branch is red or its gate fails, and a
  released tag is never re-cut.
- The coordinator changes only through `plan`, `experiment`, `intake`, `retro`, and sync, each
  through its own pull request.
- Every change lands through a session's pull request, except goal-record bookkeeping and a new
  repository's initial commit.
- On a code project, notes state only what validated work proved.
