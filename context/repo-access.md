# Repository access: which repositories a goal holds

Repository access is planned work, tracked as `marathon.repo-access`. It changes which
repositories a goal holds.

## Task `shared-repos`

A goal owns its root repository, locked for the goal's life. Its other repositories are not locked
by the goal: each is held only while it is off its default branch, with a task's branch checked
out. Several goals can then change the same non-root repository, queued by whether it is on its
default branch. Today marathon's `mechanics/goals.md`, "The repository lock", locks every
repository in a goal's `repos`.

The task also carries the goal's single release, `marathon/v0.23.0` and
`marathon-architecture/v0.4.1`. With the release, it applies the CHANGELOG's "Migrating from
0.22.0" step to claude-plugins' own `.claude/marathon.toml`, moving the CI wait into
`[remote] checks`; the installed 0.22.0 copy and binary don't know that key.

## Assumptions

- A repository's default branch is the only branch a goal finishes on, so "on its default branch"
  means no goal is mid-task there.

## Open questions

- What a queued goal does when its non-root repository is held.
- How the coordinator's `[workspace] exclusive` groups and a goal's `order` interact with the new
  lock.
- How LOCATE and `status` report a held non-root repository.
- Whether a held non-root repository needs the holding goal named on disk.
