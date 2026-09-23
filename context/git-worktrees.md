# Git worktrees for concurrent sessions

This note covers a planned rule for marathon: sessions that run at the same time each work in a
git worktree of their own, instead of sharing one checkout. A short maintenance session lands
it in the marathon plugin as a patch release.

## The failure it prevents

A wave's lanes share no member repository, but every lane commits its record,
`context/reset/<lane>.md`, to the coordinator. The coordinator is the one repository every lane
touches, and all its lanes share one working tree.

On 2026-09-23, two lane sessions ran against standards-lab at once:

1. The messaging-experiment session committed to its branch, `spike-messaging`.
2. The ai-experiment session created `spike-harness-driver` from that checkout, so the new branch
   started from `spike-messaging`, not `main`. It then checked the new branch out.
3. The messaging session's next commit landed on `spike-harness-driver` and was pushed to the
   ai lane's PR (standards-lab/org #47).

The recovery was a cherry-pick onto the right branch from a separate worktree, then a rebase of
the ai lane's branch after the messaging PR merged. Nothing was lost, but each lane's PR carried
the other lane's commits until the cleanup.

## The rule

- **A lane session works in its own worktree** of each repository it touches, including the
  coordinator. It creates its branch there at SETTLE, and RESUME checks the branch out there.
- **The main checkout stays on the default branch** while a wave runs. No lane checks a branch
  out in it.
- **`close` removes the worktree** after publishing. **`reset` leaves it in place** for the
  resuming session, and Next-focus records its path beside the Branch line.
- **A session checks the branch before every commit.** If the checkout is on a branch the
  session didn't create, it stops and reports rather than committing. This backstop also catches
  a session that runs without a worktree.
- **The harness's worktree support comes first.** Claude Code can put a session in a worktree
  itself, and where it can, the playbooks defer to it instead of prescribing
  `git worktree add` by hand.

## Why not every session

A session working alone gains nothing from a worktree and pays its costs. A fresh worktree holds
only tracked files, and the workspace depends on gitignored ones: go-web-service's
`secrets*.json` and `mise.local.toml`, standards-lab's `references.local.toml`, and the `go.work`
files some modules ignore. Docker Compose also names its project after the directory, so a
worktree's stack gets its own containers and volumes and can collide on ports with the main
checkout's stack. The rule therefore applies to concurrent sessions, which today means wave
lanes. A lane that needs gitignored files or a compose stack has to supply them in its worktree,
and the playbook should say so.

## Where it lands

In `plugins/marathon/skills/marathon/`:

- `mechanics/reset-file.md`, "Waves": state that lanes share the coordinator's repository, and
  require a worktree per lane.
- `mechanics/pipeline.md`: branch creation at 3 · SETTLE step 5, and checkout at 3R · RESUME
  step 1.
- `commands/close.md` and `commands/reset.md`: removing the worktree, or keeping it and recording
  its path.
- `references/workspace-coordination.md`: the coordinator's working tree during a wave.

Release it as marathon 0.14.1, with the version agreeing across the plugin and the marketplace
manifest, and with `scripts/check.sh` passing.

## Open questions

- Where lane worktrees live: beside each repository (`<repo>.worktrees/<branch>`), or under one
  workspace directory. The location has to be predictable, because a handoff records it.
- Whether the branch check before each commit belongs in `on-commit` guidance or in the stage
  loop in `references/staged-execution.md`.

## Assumptions

- Claude Code's worktree support can place a session in an existing worktree, not only a new
  one. A resumed lane needs that.
- Wave lanes are the only concurrent sessions for now. Experiments run in their own repositories
  and need no worktree.
