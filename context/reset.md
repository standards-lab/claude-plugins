# reset · marathon-extensions

- **Status:** closeout
- **Session:** start
- **Branch:** marathon-extensions

## Disposition

- **Integrated:** the extension-system and voice-wiring halves of `concepts/marathon-extensions.md`
  — built into marathon 0.7.0. The session's plan-mode work also widened the settled design:
  extensions enable through a `marathon.toml` `extensions` key (`[project]`, or `[workspace]` at a
  coordinator) rather than artifact presence; no separate hook-contract version — an extension
  targets a marathon version; every hook fires before the moment it names, with `on-close` ahead of
  the closeout commit so extension finalization lands inside it. Two new skill tiers carry the
  result: `mechanics/` (`pipeline.md`, the five-stage session pipeline all nine commands now layer
  into; `hooks.md`, the firing spec) and `behavior/` (`voice.md`, always active via SKILL.md's
  Behavior section, naming the framing intent). `references/extensions.md` documents the system;
  `extension-hooks.md` and `session-loop.md` retired into the mechanics tier. `docs/` articulated as
  the standardized tier of the context lifecycle (`references/context-engineering.md`). The
  user-scope voice companion landed in ~/claude-settings (`69a7fcf`).
- **Retained:** the concept's remaining half, renamed `concepts/marathon-roadmap.md` and updated to
  the 0.7.0 contract — the next session's input.
- **Added:** `concepts/marathon-optimization.md` — an unscheduled pass over data flow and context
  budget once the restructure has seen real use.
- **Cross-repo:** standards-lab `context/roadmap.toml` — `backlog.marathon-extensions` deleted,
  `next` advanced to `backlog.marathon-roadmap`, its context link updated to the renamed concept.
  docs gained `context/concepts/harness-engineering.md`, capturing a landing-zone layer for harness
  engineering conventions.

## Next-focus

`backlog.marathon-roadmap`: build the marathon-roadmap extension plugin on the 0.7.0 contract, per
`concepts/marathon-roadmap.md` — the `context/roadmap.toml` manifest convention as a marathon
extension declaring `on-start`, `on-reset`, and `on-close`. After it, `v1.data.reads` (go-core).
