# Workspaces

Related repositories, such as libraries, a template, and the service that uses them, are often
developed together, and a step sometimes changes several of them. A workspace runs that step as
one session, while each repository stays a complete marathon project of its own.

## The coordinator and the dependency order

A workspace is the directory that holds the projects as siblings. It has no `context/` of its
own. Instead, the project that already describes the whole set of repositories, usually an
organization-context project, declares itself coordinator with a `[workspace]` table
(`mechanics/configuration.md`).

The coordinator's `order` lists **layers**, lowest dependency first. An array entry is a layer of
peers with no dependencies between them. Each key is the name of a sibling directory, or a
location given in `[workspace.paths]`. Ask the architect about a key that matches neither. When
no project declares itself coordinator, list the sibling projects and ask.

## Steps that span repositories

A step that changes several member repositories runs as one session, lowest layer first, so each
higher layer builds against the real change below it. The step has one stage list, grouped by
repository in `order`. It creates a branch in each touched repository, all with the same name,
and `close` publishes each branch as that repository's own change proposal. The coordinator's
reset file records the step's state between sessions, and a resumed step reads `order` again.

## Experiments

An experiment is a standalone project outside the workspace, with its own reset file, so it runs
in parallel with the workspace's sessions. The coordinator lists it in a catalog and takes in its
results (`commands/experiment.md`).

## Projects know only what they depend on

A project knows what it builds on, and never what builds on it. Only the coordinator sees the
whole set of repositories. A lower project stays unaware of its consumers, even during a step that
spans repositories.

The conventions the coordinator keeps for its members, such as naming and writing rules, apply
through sessions, not citations: a member's `review` checks the member against them, and a step
that spans repositories applies them.
