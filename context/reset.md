# reset · marathon-roadmap

- **Status:** closeout
- **Session:** start
- **Branch:** marathon-roadmap

## Disposition

- **Integrated:** `concepts/marathon-roadmap.md` — built as the marathon-roadmap plugin at
  0.1.0 (`plugins/marathon-roadmap/`) and culled per the skill-is-the-source-of-truth rule.
  The skill mirrors marathon's tiers: SKILL.md carries the identity and declaration (artifact
  `context/roadmap.toml`; hooks `on-start`, `on-reset`, `on-close`; targets marathon 0.7);
  `mechanics/pipeline.md` maps marathon's firing points to the per-hook instruction files,
  listing only the hooks the extension connects to; `references/manifest.md` is the full
  manifest format with bootstrap header and worked example. Host surfaces updated:
  `marketplace.json`, the repository README (plugin table, install commands, the
  `marathon.toml` enablement model, structure tree), and the capability map in
  `context/README.md`.
- **Cross-repo:** standards-lab — `[workspace] extensions = ["marathon-roadmap"]` in
  `.claude/marathon.toml`; `context/roadmap.toml`'s convention header slimmed to a pointer at
  the plugin; and the extension's `on-close` applied by hand as its first exercise —
  `backlog.marathon-roadmap` deleted, `next` advanced to `v1.data.reads`.
- **Retained:** `concepts/marathon-functions.md` and `concepts/marathon-optimization.md`,
  untouched by this step. The developer's post-merge steps: tag `marathon-roadmap/v0.1.0` to
  cut the release, and install the plugin so the workspace-enabled extension resolves in
  future sessions.

## Next-focus

`v1.data.reads` — the coordinated reads slice across go-database, go-web-sdk, and
go-web-service, per `go-web-service/context/concepts/data-layer.md`. A `coordinate` session
from the workspace root.
