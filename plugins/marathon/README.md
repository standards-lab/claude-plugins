# marathon

A workflow for long-running development built on context engineering. marathon treats the
repository as the only source of truth. It keeps a top-level `context/` directory of notes,
deletes each note once the code expresses it, and runs the work as small sessions, each of which
finishes one step on its own branch.

This README is a quick reference. The skill's files under [`skills/marathon/`](./skills/marathon/)
define how marathon behaves: `SKILL.md`, the `commands/` playbooks, the `mechanics/` specifications,
the `behavior/` rules that apply in every session, and the `references/` details.

## Install

```bash
claude plugin marketplace add standards-lab/claude-plugins
claude plugin install marathon@standards-lab
```

Then run `marathon init` in a repository to set it up, or `marathon start` in a repository that
already has a top-level `context/` directory.

## Commands

Run a command as `marathon <command>` or `/marathon:marathon <command>`.

| Command | What it does |
|---------|--------------|
| `init` | Sets up marathon on a repository, once, from a planning concept. |
| `plan` | Runs a planning session that changes only `context/`: refines notes and decides the next step. |
| `start` | Advances the product one concrete step. |
| `experiment` | Sets up a spike as a standalone project. |
| `reset` | Hands off partway through a step, so a new session can resume the same branch. |
| `close` | Finishes and publishes a completed session. |
| `review` | Checks the notes for drift from the code and cleans them up. |

## Sessions

A session does one step, on one branch. A working session (`plan`, `start`, or `experiment`)
begins by reading the reset file: a standalone project's `context/reset.md`, or in a workspace
the one reset file at the coordinator. A `closeout` status means the session plans a new step in
plan mode. A `handoff` status means it resumes the open branch. The session ends with `close`,
which publishes finished work, or `reset`, which hands off unfinished work. Both rewrite the reset
file, and its Next-focus tells the next session where to begin.

## Project kinds

`init` declares a project `code` or `context` in `.claude/marathon.toml`.

Both kinds work in stages. Each stage commits once its check passes, and the architect confirms
the results at the checkpoints the stage list defines. When the branch changed prose, an editor
pass runs before the final checkpoint. `close` reviews the whole branch before publishing.
[`references/staged-execution.md`](skills/marathon/references/staged-execution.md) defines a
stage for each kind and how a step is validated.

## Workspaces

A workspace is a directory of marathon projects that sit side by side. A step may change several
of them. One session works through the touched repositories in the coordinator's dependency
order, each on its own branch with the same name. The workspace has no context of its own. The
reset file at the coordinator records its state between sessions. An experiment is a standalone
project outside the workspace, so it runs in parallel with the workspace's sessions. See
[`references/workspace-coordination.md`](./skills/marathon/references/workspace-coordination.md).

## Waves

A wave runs several lanes at the same time, in a standalone project or a workspace. Each lane owns
one thing: the standalone project, a workspace member, the coordinator, or an experiment. Every
lane keeps its record in the repository that holds the reset file. A lane that doesn't own that
repository changes it only in a git worktree of its own, which holds nothing but the lane's
`context/` changes. See [`mechanics/waves.md`](./skills/marathon/mechanics/waves.md).

## Extensions

A separately installed skill can extend marathon by acting at the five hooks every session
fires. A repository enables an extension in `.claude/marathon.toml`, for itself under `[project]`,
or for a whole workspace under the coordinator's `[workspace]`. The repository stays the source
of truth, and anything an extension copies elsewhere is a read-only mirror. See
[`references/extensions.md`](./skills/marathon/references/extensions.md) for the system and
[`mechanics/hooks.md`](./skills/marathon/mechanics/hooks.md) for when each hook fires.

## Releases

Releases are cut from [`CHANGELOG.md`](./CHANGELOG.md) by tag. Pushing a tag
`marathon/v<version>` starts the host's release workflow.
