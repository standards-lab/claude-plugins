# The marathon-roadmap plugin

Queued as `backlog.marathon-roadmap` (standards-lab `context/roadmap.toml`). Direction settled in
the workspace roadmap re-plan (2026-08-24); the extension contract it builds on landed with marathon
0.7.0 (`references/extensions.md`, `mechanics/hooks.md`). The session's plan mode settles the final
wording. This note is the trimmed remainder of the `marathon-extensions` concept; the extension
system it also described is built and recorded in the v0.7.0 changelog.

## The plugin

Its own plugin, named `marathon-roadmap` — the name declares the layering in the skill listing,
which is the discovery surface — so anyone can layer a different tracking system on the same hooks
without touching marathon. It codifies the convention the workspace already runs by hand
(standards-lab `context/roadmap.toml` is the lived example):

- **Three primitives**, nested like a filesystem: a goal is a named outcome made of other goals
  and tasks, closing when its criteria hold; a task is one session's requirement — what it is,
  the repositories it touches, its proof; the backlog is the tasks assigned to no goal.
- **One TOML manifest**, `context/roadmap.toml`, at the coordinator in a workspace. Child goals
  nest directly (`[goals.v1.data]`); tasks live under a goal's reserved `tasks` table; backlog
  tasks under the root `backlog` table. Fields: `name`, `summary`, and optional `criteria`,
  `repos`, `proof`, `context` (workspace-relative file links, first segment the repository).
- **Dotted citations.** References cite the dotted slug path, omitting the structural `goals`
  and `tasks` segments: `v1.data.reads`, `backlog.docs-site`. Dots, not slashes, so a task
  reference is never mistaken for a file path; slugs never contain dots.
- **Ephemeral, unordered, proximate.** The manifest captures only what remains — finished
  tasks and closed goals are deleted, the session record keeping the disposition. `next` is the
  only sequence. Resolution matches proximity: the task in front carries detail, everything
  else stays at claim resolution, detail living in linked context files.

## Declaration under the 0.7.0 contract

- Skill-listing description contains "marathon extension"; the SKILL.md declares the artifact
  (`context/roadmap.toml`), the hooks, and the targeted marathon version (0.7).
- Enabled in `marathon.toml`: `[workspace] extensions` at a coordinator, `[project] extensions`
  on a standalone project.
- **Hook behavior**: `on-start` — layer the manifest's conventions in, bootstrapping it if
  enabled and absent; `on-reset` — cite task paths in Next-focus; `on-close` — delete the
  finished task and advance `next`, landing inside the closeout commit; nothing at `on-execute`
  or `on-commit`.

The skill codifies the manifest contract and the hook behavior, and stays non-prescriptive past
them — how a project shapes its goal tree is its own.
