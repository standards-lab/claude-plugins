# Session pipeline

Execution spec for every marathon command. The pipeline owns the session mechanics; each command's
playbook under `commands/` supplies the content of the stages — what orientation reads, what settling
weighs, what execution produces. Hook firing is specified in `mechanics/hooks.md` and happens only at
the points named here.

## Stages

Run the stages in order. `3R · RESUME` replaces `3 · SETTLE` when the reset file records a handoff
for the running command.

### 1 · LOCATE

1. Identify the directory kind. A **standalone project** has its own top-level `context/` and no
   workspace around it. A **workspace** is a directory whose subdirectories are the projects, one
   of which declares `[workspace] role = "coordinator"` in its `.claude/marathon.toml`
   (`mechanics/configuration.md`); entry at the workspace root and entry inside a member project
   are the same case.
2. Resolve the reset file (`mechanics/reset-file.md`): a standalone project's own
   `context/reset.md`, or, anywhere in a workspace, the coordinator's — the workspace's only
   reset; member projects carry none.
3. Read it and route on Status:
   - `Status: closeout` → fresh step. The Next-focus names the step and, in a workspace, the
     member project it runs in. Continue: 2 · START, 3 · SETTLE.
   - `Status: handoff` → resume. The Session line names the command that resumes it; if that is
     not the running command, switch to it. The Branch line — with the Project line in a
     workspace — names where the open branch waits. Continue: 2 · START, 3R · RESUME.
   - File missing → no session has recorded a step here; settle a fresh one with the developer.

Two commands deviate. `init` applies only when both marks are absent — no `context/` here, no
sibling declaring a coordinator — and that check is its LOCATE. `coordinate` targets the workspace
itself (`commands/coordinate.md`, step 1).

### 2 · START

1. Fire `on-start`.
2. Orient: the Next-focus in the reset file, the capability map in the project's
   `context/README.md`, and the `design/` and `concepts/` notes the work touches. Load only what
   the session needs.

### 3 · SETTLE

1. Enter plan mode.
2. Work the scope through with the developer, to the depth the command's playbook calls for. Change
   nothing yet.
3. As the discussion ranges wider than the step, note the context tending it implies — ideas to
   capture as concepts, concepts it rules out, the next focus taking shape. On a **context**
   project these edits are made in 4 · EXECUTE, after approval — there the change is the context.
   On a **code** project they wait for 5 · CONCLUDE: EXECUTE produces only the unvalidated guide,
   and the notes record what the applied code proved, not what a plan intended. A handoff records
   in-flight direction in the reset Disposition without touching the notes.
4. On the developer's approval, fire `on-execute`.
5. Create the branch, named by the slug rule in the command's playbook.

### 3R · RESUME

1. Check out the open branch named in the reset file.
2. Fire `on-execute`.
3. Read the in-progress state; the Next-focus is the exact next move. Continue: 4 · EXECUTE.

### 4 · EXECUTE

1. Do the command's work, per its playbook and the project kind.
2. Fire `on-commit` immediately before any commit the session makes, in this stage or a later one.

### 5 · CONCLUDE

The session ends by one of two exits, each with its own playbook:

- Work unfinished and context filling → `reset` (`commands/reset.md`): tidy the touched notes, fire
  `on-reset`, write the reset file with `Status: handoff`, optionally WIP-commit (`on-commit`),
  leave the branch open.
- Work finished and validated → `close` (`commands/close.md`): finish the agent's part, tidy the
  notes — including the edits noted at SETTLE, now validated — agree the next step, fire
  `on-reset`, write the reset file with `Status: closeout`, fire `on-close`, commit (`on-commit`),
  publish with the `[remote]` publish command.

In a workspace, the reset file both exits write is the coordinator's — a commit in the
coordinator's repository, alongside the session's own.

## Command → pipeline map

| Command | Layering |
|---------|----------|
| `plan`, `start`, `experiment` | full pipeline; working sessions |
| `review`, `docs` | full pipeline; on-demand passes, recorded under their own Session values |
| `init` | LOCATE is the empty checks; SETTLE the founding decisions; EXECUTE the scaffold; CONCLUDE its own setup commit — `on-reset` at the first reset file, `on-commit` at the commit, never `on-close` |
| `reset`, `close` | CONCLUDE invoked directly, ending the current session |
| `coordinate` | the full pipeline per participating project, fanned out in dependency order; the workspace itself runs none of it and keeps no state |

## Invariants

- One session, one step, one branch. A session never widens past the step settled at SETTLE.
- Every session starts and ends at the reset file; the repository, not the conversation, carries
  continuity.
- Nothing is created or changed before the developer approves at SETTLE, except by a RESUME picking
  up an approved plan.
- On a code project, `context/` asserts only what validated work proved; tending follows
  validation.
