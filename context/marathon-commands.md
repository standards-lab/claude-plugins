# marathon commands after the factory redesign

This note covers how marathon's commands change when the pipeline in `marathon-factory.md` and
the goal model in `marathon-goals.md` land. `factory.pipeline` changes the playbooks, and
`factory.goals` changes everything that reads the roadmap.

## The command set

| Command | Role |
|---|---|
| `init` | Set up a repository or workspace |
| `plan` | Run plan rounds; stage, reorder or drop goals; take in a spike; write task briefs into goal records |
| `start` | Run one task: build, brief, ship, and sync when it is the goal's last task |
| `experiment` | Set up a spike as a goal homed in its own repository |
| `status` | Print the status digest (new) |
| `retro` | Review a session, a goal or a date range, including the context-drift pass (new) |
| `reset` | Hand off by hand; it happens automatically when the context fills |

Retired:

- `close`: its review, tending and publishing become `start`'s BUILD and SHIP.
- `review`: its drift check becomes part of `retro`.

## What changes in each playbook

- **init**
  - Writes `roadmap.toml` with `active`, `planned` and `backlog`.
  - Writes a `STANDARDS.md` stub and a pointers-only `CLAUDE.md`.
  - Records the workspace's `check` command.
- **plan**
  - SETTLE becomes plan rounds (`marathon-briefs.md`).
  - Its output is a task brief and slices in the goal record, or edits to the roadmap's three
    states.
  - It commits directly to the coordinator's default branch, and stays short.
- **start**
  - Runs PLAN only when the goal record has no approved brief.
  - Runs BUILD without stopping.
  - Shows the session brief, then publishes and merges once the architect accepts.
  - Syncs the goal when this task is its last.
- **experiment**
  - Sets up the spike repository and stages a spike goal homed there.
  - The single question, the evidence list and the narrated CLI stay as they are; they work.
- **status (new)**
  - Reads `active` and each goal record and prints the digest.
  - Changes nothing.
- **retro (new)**
  - Gathers sessions, briefs and merged PRs over the range.
  - Proposes changes routed by layer, and applies the ones the architect ticks.
- **reset**
  - Writes handoff state into the goal record and leaves the task's branch open.
  - The orchestrating session calls it on its own when its context fills.

## Retired mechanics

- `mechanics/waves.md`: lanes, lane records, worktrees and folding.
- `mechanics/reset-file.md` and `context/reset.md`: the goal record holds the session state.
- Checkpoints and their outcomes in `references/staged-execution.md`; slices replace stages.
- `.claude/report.md`: the session brief replaces the branch report.
- Delegation statements: profiles inherit the session's model, and the session announces no
  model.
- The `executor` profile: renamed `implementer`. The `reviewer` profile: split into
  `standards-reviewer` and `spec-reviewer`.
- marathon-roadmap as a separate extension: folded into core (`marathon-goals.md`).
- The architecture extension's promotion hooks: the extension shrinks to the pointers each
  `STANDARDS.md` follows. Promotion happens only from validated code.
