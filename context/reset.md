# reset · user-invocable

- **Status:** closeout
- **Session:** start
- **Branch:** user-invocable

## Disposition

- **Integrated:** marathon-roadmap 0.1.1 — `user-invocable: false` in the SKILL.md frontmatter
  makes the extension skill passive: out of the slash-command menu, still resolvable by
  marathon's hook firing and loadable by the model. Changelog entry added; no behavior change
  past the menu.
- **Cross-repo:** docs — `context/concepts/harness-engineering.md` gains the passive-skill
  convention in its candidate list.
- **Retained:** `concepts/marathon-functions.md` and `concepts/marathon-optimization.md`,
  untouched. The developer's post-merge step: tag `marathon-roadmap/v0.1.1` and update the
  installed plugin.

## Next-focus

Unchanged from the last closeout: `v1.data.reads` — the coordinated reads slice across
go-database, go-web-sdk, and go-web-service, per
`go-web-service/context/concepts/data-layer.md`. A `coordinate` session from the workspace
root.
