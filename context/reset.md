# reset · project-kinds-and-coordination

- **Status:** closeout
- **Session:** start
- **Branch:** project-kinds-and-coordination

## Disposition

- **Integrated workspace coordination into the skill.** Culled `context/concepts/workspace-coordination.md`
  — the marathon skill now expresses it, per `design/skill-is-the-source-of-truth.md`, via the new
  `commands/coordinate.md` playbook and `references/workspace-coordination.md` deep-dive (workspace holds
  no context; coordinator declares a layered `[workspace] order`; per-project fan-out honoring kind;
  ephemeral consolidated guide; continuity per repository).
- **Shipped project kinds and restructured the session API** in marathon 0.3.0. Added the `code`/`context`
  project kind (declared in `.claude/marathon.toml`), and replaced `start <development|context|experiment>`
  with top-level `plan`, `start`, and `experiment` commands. Harmonized the whole skill around both —
  `SKILL.md`, every `commands/*.md`, and `role-boundary`/`session-loop`/`context-engineering`/
  `implementation-guides`/`extension-hooks` — and standardized the notes-in-`concepts/` vocabulary on
  "concepts". Bumped `plugin.json` to 0.3.0 with a `CHANGELOG.md` entry, and added a quick-reference
  `plugins/marathon/README.md`.
- **Dogfood.** Declared `claude-plugins` itself `kind = "context"` and rewrote its `CLAUDE.md` role
  boundary to match — the agent authors the skill/plugin files directly under developer review.
- **Retained:** `design/skill-is-the-source-of-truth.md` — the governing principle, unchanged.

## Next-focus

Exercise workspace coordination for real. In `standards-lab` (the coordinator), add
`[workspace] role = "coordinator"` with a layered `order`, and declare `[project] kind` in the workspace
projects (`standards-lab` context, `go-libraries` code). Then run `marathon coordinate` on a genuine
cross-repo objective — e.g. a `go-libraries` capability taken up into `go-service` for validation — to
prove the mechanism end-to-end and surface what the design missed. Note: the active marathon skill is the
installed marketplace copy, so reinstall `marathon@standards-lab` once 0.3.0 releases before coordinating.
Start here next session.
