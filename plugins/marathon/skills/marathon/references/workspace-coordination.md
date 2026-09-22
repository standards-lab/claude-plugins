# Workspaces

Related repositories, such as libraries, a template, and the service that consumes them, are often
developed together, and a step sometimes crosses several of them. A workspace runs that step as
one session while each repository stays a self-contained marathon project.

## The coordinator and the map

A workspace is the directory the projects sit under as siblings. It has no `context/` of its own:
the project that already describes the estate, typically an organization-context project,
declares itself coordinator with a `[workspace]` block (`mechanics/configuration.md`). Its `order`
lists **layers**, lowest dependency first; an array entry is a layer of peers with no dependency
between them. Each key resolves to a sibling directory by name, or through `[workspace.paths]`;
ask the architect about a key that resolves to neither. With no coordinator, enumerate the sibling
projects and ask.

## Cross-repo steps

A step that spans member repos runs as one session, lowest layer first, so a higher layer builds
against the real change beneath it. Its one stage list groups stages by repository in map order,
and it creates a branch per touched repo under a shared slug; `close` publishes each as that
repository's own change proposal. The coordinator's single reset file carries the step's
continuity, and a resumed step re-reads `order`.

## Experiments

An experiment is a standalone project outside the workspace, with its own reset file, so it runs
in parallel with the workspace's sessions. The coordinator catalogs it and takes in its results
(`commands/experiment.md`).

## Awareness follows the dependency direction

A project knows what it builds on, never what builds on it. The estate-wide view lives only in the
coordinator, and a lower project stays unaware of its consumers, even inside a cross-repo step.

Conventions the coordinator keeps for its members, such as naming and authoring rules, bind
through sessions rather than citations: a member's `review` checks against them, and a cross-repo
step applies them.
