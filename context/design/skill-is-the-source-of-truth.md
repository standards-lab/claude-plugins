# The skill is the source of truth

Each plugin's design is expressed in its own files under `plugins/<name>/` — the skill's `SKILL.md`, its
`commands/`, and its `references/`. Those files are the authority on how the plugin behaves.

`context/` does not restate that design. It holds what the code cannot yet express: the repository's vision
and capability map (`README.md`), candidate ideas not yet built (`concepts/`), and the session record
(`reset.md`) while one is open. A design note here records intent the skill files do not yet capture; once the skill
expresses it, the note is removed — the same decay rule marathon applies everywhere.

This keeps the repository free of the duplication a plugin host invites: a plugin documented in two places
drifts in one of them. The skill documents the plugin; the context documents the project.
