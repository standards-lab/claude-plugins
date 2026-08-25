# Session pipeline

Execution spec for every marathon command. The pipeline owns the session mechanics; each command's
playbook under `commands/` supplies the content of the stages — what orientation reads, what settling
weighs, what execution produces. Hook firing is specified in `mechanics/hooks.md` and happens only at
the points named here.

## Stages

Run the stages in order. `3R · RESUME` replaces `3 · SETTLE` when the reset file records a handoff
for the running command.

### 1 · LOCATE

1. Identify the directory kind per `SKILL.md` "Finding your project": a project has its own
   top-level `context/`; a workspace root does not, and one of its member projects declares
   `[workspace] role = "coordinator"` in `.claude/marathon.toml`.
2. At a workspace root: resolve the coordinator, read its `context/reset.md`, and continue in the
   member project it names — Next-focus at `Status: closeout`, the Branch line at `Status: handoff`.
3. Read the project's `context/reset.md` and route on it:
   - File missing → resting point. Continuity is the coordinator's reset file; read it, or settle a
     fresh step with the developer if none names this project.
   - `Status: closeout` → fresh step. Continue: 2 · START, 3 · SETTLE.
   - `Status: handoff` → resume. The Session line names the command that resumes it; if that is not
     the running command, switch to it. Continue: 2 · START, 3R · RESUME.

Two commands deviate. `init` applies only when both checks are empty — no `context/` here, no
sibling declaring a coordinator — and that check is its LOCATE. `coordinate` targets the workspace
itself (`commands/coordinate.md`, step 1).

### 2 · START

1. Fire `on-start`.
2. Orient: the Next-focus in `reset.md`, the capability map in `context/README.md`, and the
   `design/` and `concepts/` notes the work touches. Load only what the session needs.

### 3 · SETTLE

1. Enter plan mode.
2. Work the scope through with the developer, to the depth the command's playbook calls for. Change
   nothing yet.
3. As the discussion ranges wider than the step, note the context tending it implies — ideas to
   capture as concepts, concepts it rules out, the next focus taking shape. These edits are made in
   4 · EXECUTE, after approval.
4. On the developer's approval, fire `on-execute`.
5. Create the branch, named by the slug rule in the command's playbook.

### 3R · RESUME

1. Check out the open branch named in `reset.md`.
2. Fire `on-execute`.
3. Read the in-progress state; the Next-focus is the exact next move. Continue: 4 · EXECUTE.

### 4 · EXECUTE

1. Do the command's work, per its playbook and the project kind.
2. Fire `on-commit` immediately before any commit the session makes, in this stage or a later one.

### 5 · CONCLUDE

The session ends by one of two exits, each with its own playbook:

- Work unfinished and context filling → `reset` (`commands/reset.md`): tidy the touched notes, fire
  `on-reset`, write `context/reset.md` with `Status: handoff`, optionally WIP-commit (`on-commit`),
  leave the branch open.
- Work finished and validated → `close` (`commands/close.md`): finish the agent's part, tidy the
  notes, agree the next step, fire `on-reset`, write `context/reset.md` with `Status: closeout` (or
  delete it at a resting point), fire `on-close`, commit (`on-commit`), publish with the `[remote]`
  publish command.

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
- Every session starts and ends at `context/reset.md`; the repository, not the conversation, carries
  continuity.
- Nothing is created or changed before the developer approves at SETTLE, except by a RESUME picking
  up an approved plan.
