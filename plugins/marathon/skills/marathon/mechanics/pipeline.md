# Session pipeline

The session mechanics every marathon command runs; each playbook under `commands/` supplies its
stages' content. Hooks fire only at the points named here (`mechanics/hooks.md`).

## Stages

Run the stages in order. `3R · RESUME` replaces `3 · SETTLE` when the reset file records a handoff
for the running command.

### 1 · LOCATE

1. Identify the directory kind. A **standalone project** has its own top-level `context/` and no
   sibling declaring itself coordinator. A
   **workspace** is a directory of projects, one declaring itself coordinator in its
   `.claude/marathon.toml` (`mechanics/configuration.md`); entering at the root or inside a member
   is the same case.
2. Resolve the reset file (`mechanics/reset-file.md`): the project's own `context/reset.md`, or in
   a workspace the coordinator's.
3. Route on its Status:
   - `closeout` → a fresh step, named by Next-focus with, in a workspace, its member project. When
     Next-focus names a wave, first fold it if every lane is finished; otherwise the architect
     names this session's lane, routed on `context/reset/<lane>.md` if one exists. Continue:
     START, SETTLE.
   - `handoff` → resume. Switch to the command the Session line names; the Branch and Project lines
     say where the branch waits. Continue: START, RESUME.
   - Missing → settle a fresh step with the architect.

`init` deviates: its LOCATE checks that neither mark exists.

### 2 · START

1. Fire `on-start`.
2. Orient on the Next-focus, the capability map in `context/README.md`, and the notes the work
   touches. Load only what the session needs.

### 3 · SETTLE

1. Enter plan mode.
2. Work the scope through with the architect to the depth the playbook calls for, weighing it
   against what the project already has (`behavior/planning.md`). A working session expresses it
   as the stage list of `references/staged-execution.md`, which the planner profile may draft
   (`behavior/delegation.md`). Change nothing until the architect approves.
3. Note the context tending the discussion implies: ideas to capture, notes it rules out, the next
   focus taking shape. A **context** project makes these edits in EXECUTE; a **code** project
   waits for CONCLUDE, so the notes record what validation proved. A handoff records them in the
   Disposition only.
4. On approval, fire `on-execute`.
5. Create the branch, named by the playbook's slug rule; a step spanning member repos creates one
   per touched repo under the same slug.

### 3R · RESUME

1. Check out the open branch in each touched repo.
2. Fire `on-execute`.
3. Read the stage list, the stage and checkpoint position, and the next move from Next-focus.
   Continue: EXECUTE.

### 4 · EXECUTE

1. Do the playbook's work as the stage loop of `references/staged-execution.md`: each stage
   commits once its check passes, and the session stops at each checkpoint. A re-plan re-enters
   SETTLE without leaving the branch.
2. Fire `on-commit` before every commit the session makes, here or later.

### 5 · CONCLUDE

- Unfinished, context filling → `reset` (`commands/reset.md`): tidy the touched notes, fire
  `on-reset`, write a handoff record, optionally WIP-commit, and leave the branch open.
- Finished and validated → `close` (`commands/close.md`): review the branch, tidy the notes, agree
  the next step, fire `on-reset`, write the closeout record, fire `on-close`, commit, and publish.

In a workspace the record is the coordinator's, committed in its repository.

## Commands

| Command | Layering |
|---------|----------|
| `plan`, `start`, `experiment` | full pipeline; working sessions |
| `review` | full pipeline; the on-demand pass |
| `init` | LOCATE checks for no marks; SETTLE the founding decisions; EXECUTE the scaffold; CONCLUDE its setup commit |
| `reset`, `close` | CONCLUDE, ending the current session |

## Invariants

- One session, one step, on one branch per touched repo.
- Every session starts and ends at the reset file; the repository carries continuity.
- Nothing changes before the architect approves at SETTLE, except a RESUME of an approved plan.
- A stage commits only once its check passes, and execution stops at every checkpoint.
- On a code project, `context/` asserts only what validated work proved.
