# The skill is the source of truth

Each plugin's design is expressed in its own files under `plugins/<name>/` — the skill's `SKILL.md`, its
`commands/`, `mechanics/`, `behavior/`, and `references/` tiers. Those files are the authority on how the
plugin behaves.

`context/` does not restate that design. It contains what the skill files cannot express: the
repository's vision and capability map (`README.md`) and candidate ideas not yet built
(`concepts/`). The session record is the workspace's reset file at the coordinator. A design note
here records intent the skill files do not yet capture; once the skill expresses it, the note is
removed, by the decay rule marathon applies everywhere.

This keeps the repository free of the duplication a plugin host invites: a plugin documented in two places
drifts in one of them. The skill documents the plugin; the context documents the project.
