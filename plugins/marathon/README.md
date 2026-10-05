# marathon

A workflow for long-running development, run as a software factory. marathon treats the
repository as the only source of truth: a top-level `context/` directory of notes, a roadmap
manifest of goals and tasks, and one record per active goal. Each session takes one task from an
approved plan to merged work. The architect approves the task brief and accepts the session brief;
subagents implement, review, and keep the context current in between. The pipeline is inspired by
Matt Pocock's talk "Fixing the PR Bottleneck" and his skills repository
(`github.com/mattpocock/skills`).

This README is a quick reference. The skill's files under [`skills/marathon/`](./skills/marathon/)
define how marathon behaves: `SKILL.md`, the `commands/` playbooks, the `mechanics/` specifications,
the `behavior/` rules that apply in every session, and the `references/` details.

## Install

```bash
claude plugin marketplace add standards-lab/claude-plugins
claude plugin install marathon@standards-lab
```

Then run `marathon init` in a repository to set it up, or `marathon status` in a repository that
already has a top-level `context/` directory.

## Commands

Run a command as `marathon <command>` or `/marathon:marathon <command>`.

| Command | What it does |
|---------|--------------|
| `init` | Sets up marathon on a repository, once, from a planning concept. |
| `plan` | Defines goals and makes them ready: creates goals and tasks, sets up repositories and records, stages and orders them. |
| `start <goal>` | Runs the goal's next task: plan, build, brief, ship, and sync on the last task. |
| `experiment` | Creates an `experiment.<topic>` goal, its `intake` task, and its planned spikes, once per topic. |
| `intake <experiment>` | Takes in a finished experiment: decides what the served goals build from the spikes' answers, closes it, and archives the spikes. |
| `status` | Prints one line per active goal, with what needs the architect first. |
| `retro` | Turns comments on past work into checks, standards, skill changes, or notes. |
| `reset` | Hands off a task partway; runs on its own when the context fills. |

## The pipeline

```
PLAN   plan rounds → task brief + slices          ✔ approve           [touch 1]
BUILD  implementer per slice → check; standards-reviewer; spec-reviewer; editor
BRIEF  session brief                              ✔ accept / redirect [touch 2]
SHIP   goal record → pull request (body = brief) → merge → tag if Release → sync if last
```

BUILD stops for the architect only for a one-way door, a decision the brief doesn't cover, or
scope beyond the task. When the task brief has a Release line, SHIP tags each release once main
is green and the repository's optional release gate passes, fixing forward until the planned
version releases. Every session, planning and sync
included, lands through a branch and a pull request in each repository it changes. See
[`mechanics/pipeline.md`](./skills/marathon/mechanics/pipeline.md) and
[`references/build.md`](./skills/marathon/references/build.md).

## Goals

The roadmap manifest, `context/roadmap.toml`, holds goals (outcomes) and their tasks (one
session's work each). Every goal is `active`, `planned`, or `backlog`, except an experiment goal,
which is never listed. Active goals run side by side, and each locks the repositories it touches,
so no two share one and no worktrees are needed.
A goal's record, `context/goals/<goal>.md` in its root repository, holds its tasks, the current
brief, progress, decisions, and the edits it owes other repositories, which sync applies when its
last task ships. The coordinator changes only through `plan`, `experiment`, `intake`, `retro`, and
sync. See [`mechanics/goals.md`](./skills/marathon/mechanics/goals.md).

## Experiments

An experiment, `experiment.<topic>`, spikes an idea before a goal commits to it. Each spike is a
sub-goal, `experiment.<topic>.<spike>`, that `plan` sets up in a new repository of its own; spikes
of one experiment can run at once. A spike's sync lands its answer in the note the experiment cites
and proposes the next planned spike to stage. Once every spike has synced, `intake` decides what
the served goals build. See [`commands/experiment.md`](./skills/marathon/commands/experiment.md)
and [`commands/intake.md`](./skills/marathon/commands/intake.md).

## Project kinds

`init` declares a project `code` or `context` in `.claude/marathon.toml`, along with its one
deterministic `check` command. Two commands are optional: a read-only `currency` command, which
reports what trails its latest release, and a release `gate` command, such as a plugin's eval
suite, which SHIP runs before each tag. On a code project a slice adds behavior with its tests; on a
context project a slice is the deliverable prose, checked by the consistency script and a read.

## Currency

PLAN runs each touched repository's `currency` command before round 1, so a task begins current.
Round 1 lists what trails and asks only about what needs a decision: a breaking change to adapt
to, a feature to adopt, or a new major to upgrade, hold, or give its own task. The task's slices
then begin with one upgrade slice per trailing repository, and a held item is recorded as a
decision. CI never runs the command. See
[`mechanics/configuration.md`](./skills/marathon/mechanics/configuration.md) and
[`references/briefs.md`](./skills/marathon/references/briefs.md).

## Workspaces

A workspace is a directory of marathon projects that sit side by side. One of them, the
coordinator, holds the roadmap and the dependency order. A task may change several members; it
runs lowest layer first, each on a branch with the same name, and ships one pull request per
repository. The coordinator changes only through the sessions listed under Goals, each through its
own pull request. See
[`references/workspace-coordination.md`](./skills/marathon/references/workspace-coordination.md).

## Extensions

A separately installed skill can extend marathon by acting at the hooks every session fires. A
repository enables an extension in `.claude/marathon.toml`, for itself under `[project]`, or for
a whole workspace under the coordinator's `[workspace]`. See
[`references/extensions.md`](./skills/marathon/references/extensions.md) and
[`mechanics/hooks.md`](./skills/marathon/mechanics/hooks.md).

## Releases

Releases are cut from [`CHANGELOG.md`](./CHANGELOG.md) by tag. Before SHIP pushes a
`marathon/v<version>` tag, the repository's release gate, `scripts/gate.sh`, runs the consistency
check and marathon's eval suite. Pushing the tag starts the host's release workflow.
