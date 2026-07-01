# reset · adopt-marathon

- **Status:** closeout
- **Session type:** development
- **Branch:** adopt-marathon

## Disposition

- **Adopted marathon** on `claude-plugins`, the harness level of the reference architecture — the plugin
  marketplace is now managed the same way as the other levels, and marathon is battle-tested against its own
  source. Scaffolded the `context/` tree (`README.md`, `design/`, `concepts/`, `reset.md`) and the Claude
  configuration (`CLAUDE.md`, `.claude/settings.json`, `.claude/marathon.toml`; github / `gh pr create`,
  standalone).
- **Settled** (`design/`): `skill-is-the-source-of-truth` — plugin design lives in the skill files;
  `context/` does not restate it.
- **Candidate** (`concepts/`): `workspace-coordination` — running marathon from a contextless workspace that
  coordinates the nested marathon projects.

## Next-focus

Design and build **workspace coordination** in marathon: a workspace layer that holds no context and
coordinates the nested marathon projects in dependency order. Start from
`concepts/workspace-coordination.md` — settle the mechanism-versus-map split and how the coordinator
declares itself, then work out the first concrete step (likely workspace detection and project
enumeration). Start here next session with `marathon start`.
