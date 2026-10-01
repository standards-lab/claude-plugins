# Workspaces

Related repositories, such as libraries, a template, and the service that uses them, are often
developed together, and a task sometimes changes several of them. A workspace runs such a task as
one session, while each repository stays a complete marathon project of its own.

## The coordinator and the dependency order

A workspace is the directory that holds the projects as siblings. It has no `context/` of its
own. Instead, the project that already describes the whole set of repositories, usually an
organization-context project, declares itself coordinator with a `[workspace]` table
(`mechanics/configuration.md`). The coordinator holds the workspace's roadmap manifest.

The coordinator's `order` lists **layers**, lowest dependency first. An array entry is a layer of
peers with no dependencies between them. Each key is the name of a sibling directory, or a
location given in `[workspace.paths]`. Every repository name a goal uses, in `root` and `repos`,
resolves the same way, or through a spike task's `path`. Ask the architect about a name that
resolves through none of them. When no project declares itself coordinator, list the sibling
projects and ask.

## Goals running side by side

Each active goal locks the repositories in its `repos`, so goals that run at the same time never
share a member repository, and each works on the main checkout of its own
(`mechanics/goals.md`). Each goal's record lives in its root, one of the repositories it locks.
The coordinator is the one shared repository, and it changes only through `plan`'s planning and
administrative edits and through syncs, as short, direct commits on its default branch.

## Tasks that span repositories

A task that changes several member repositories runs as one session, lowest layer first, so each
higher layer builds against the real change below it. Its slices are ordered by `order`. It
creates a branch with the same name in each touched repository, and SHIP publishes and merges each
one as that repository's own pull request, lowest layer first.

## Experiments

An experiment goal's spikes are repositories outside the workspace, each listed in the roadmap with
its remote and local path, so the experiment runs beside the workspace's goals. Its intake task
decides what the served goal takes from the result, and its sync moves the spikes into the
workspace's repository catalog as archived (`commands/experiment.md`).

## Projects know only what they depend on

A project knows what it builds on, and never what builds on it. Only the coordinator sees the
whole set of repositories. A lower project stays unaware of its consumers, even during a task that
spans repositories.

The conventions the coordinator keeps for its members, such as naming and writing rules, reach the
members through each repository's `STANDARDS.md`, which the standards-reviewer applies, and through
`retro`, which routes a convention that keeps being missed into a check.
