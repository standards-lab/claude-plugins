# Session pipeline

Every marathon command runs this pipeline. Each playbook under `commands/` supplies the content of
its stages. Extension hooks fire only at the points this file names (`mechanics/hooks.md`).

## Stages

Run the stages in order. When the reset file records a handoff for the running command,
`3R · RESUME` replaces `3 · SETTLE`.

### 1 · LOCATE

1. Identify what kind of directory the session is in:
   - A **standalone project** has its own top-level `context/`, and no sibling project declares
     itself coordinator.
   - A **workspace** is a directory of projects, one of which declares itself coordinator in its
     `.claude/marathon.toml` (`mechanics/configuration.md`). Starting at the workspace root and
     starting inside a member are the same case.
2. Find the reset file (`mechanics/reset-file.md`): the project's own `context/reset.md`, or the
   coordinator's in a workspace.
3. Route on the reset file's Status:
   - `closeout`: start a new step, the one Next-focus names, together with its member project in
     a workspace. When Next-focus names a wave, fold the wave first if every lane is finished. If
     a lane is still open, the architect names this session's lane, and the session routes on
     `context/reset/<lane>.md` if that file exists. Continue with START, then SETTLE.
   - `handoff`: resume. Switch to the command the Session line names. The Branch and Project lines
     say where the open branch is. Continue with START, then RESUME.
   - No reset file: settle a new step with the architect.

`init` is the exception: its LOCATE checks that marathon isn't set up, with no `context/` here and
no sibling declaring itself coordinator.

### 2 · START

1. Fire `on-start`.
2. Read Next-focus, the capability map in `context/README.md`, and the notes the work touches.
   Load only what the session needs.

### 3 · SETTLE

1. Enter plan mode.
2. Work through the scope with the architect to the depth the playbook asks for, and weigh it
   against what the project already has (`behavior/planning.md`). A working session expresses
   the scope as a stage list (`references/staged-execution.md`), which the planner profile may
   draft (`behavior/delegation.md`). Change nothing until the architect approves.
3. Note the context edits the discussion implies: ideas to capture, notes it rules out, and the
   next focus as it takes shape. A **context** project makes these edits during EXECUTE. A
   **code** project waits until CONCLUDE, so the notes record only what validation proved. A
   handoff records them in its Disposition and makes no edits.
4. When the architect approves, fire `on-execute`.
5. Create the branch, named by the playbook's slug rule. A step that spans member repositories
   creates one branch in each touched repository, all with the same name. A wave's lane creates
   each branch in a worktree of its own, finalizes the worktree, and works there
   (`references/workspace-coordination.md`).

### 3R · RESUME

1. Check out the open branch in each touched repository. A wave's lane enters the worktree that
   Next-focus records instead.
2. Fire `on-execute`.
3. Read the stage list, the stage and checkpoint position, and the next move from Next-focus.
   Continue with EXECUTE.

### 4 · EXECUTE

1. Do the playbook's work as the stage loop in `references/staged-execution.md`. Each stage
   commits once its check passes, and the session stops at each checkpoint. A re-plan returns to
   SETTLE on the same branch.
2. Before every commit the session makes, in this stage or later, confirm that the checkout is on
   the session's branch, then fire `on-commit`. If the checkout is on a branch the session didn't
   create or resume, stop and report to the architect rather than commit.

### 5 · CONCLUDE

- When the step is unfinished and the context is filling, run `reset` (`commands/reset.md`): tidy
  the touched notes, fire `on-reset`, write a handoff record, make a WIP commit if needed, and
  leave the branch open.
- When the step is finished and validated, run `close` (`commands/close.md`): review the branch,
  tidy the notes, agree on the next step, fire `on-reset`, write the closeout record, fire
  `on-close`, commit, and publish.

In a workspace, the record is the coordinator's and is committed in the coordinator's repository.

## How each command uses the pipeline

| Command | Stages it runs |
|---------|----------------|
| `plan`, `start`, `experiment` | The full pipeline, as working sessions |
| `review` | The full pipeline, as an on-demand check |
| `init` | LOCATE checks that marathon isn't set up; SETTLE agrees the founding decisions; EXECUTE creates the structure; CONCLUDE makes the setup commit |
| `reset`, `close` | CONCLUDE, which ends the current session |

## Invariants

- A session does one step, on one branch in each touched repository.
- Every session starts and ends at the reset file, so the next session learns where to begin
  from the repository, not the conversation.
- Nothing changes before the architect approves at SETTLE, except when RESUME continues an
  approved plan.
- A stage commits only once its check passes, and execution stops at every checkpoint.
- On a code project, `context/` states only what validated work proved.
