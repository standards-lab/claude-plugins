# The docs command as an extension

Promote marathon's `docs` command from a core optional feature to an extension — the 2026-08-31
retrospective's recommendation, kept out of v0.9.0 as its own backlog task
(`backlog.marathon-docs-extension` in the coordinator's roadmap).

## The case for promotion

- `docs` is the only core feature bound to an org-specific artifact — the landing-zone binding
  (`docs = "docs"` in the coordinator's `marathon.toml`). That is exactly the shape
  marathon-roadmap externalized: core marathon stays generic, the extension carries the org's
  convention.
- The always-loaded skill index lightens: the `docs` command and its landing-zone rules currently
  load for every session in the workspace.
- It gives the extension taxonomy its second built member alongside the sitrep concept's proposed
  enhancement facet, with marathon-roadmap's layout as the de facto extension template — nowhere
  yet named as one; worth stating if this lands.

## Evaluate rather than presume

- The docs tier participates in `close`/`review` context tending, so the split must keep that
  behavior hook-reachable (`on-close` firing, the tending ledger).
- The `docs` command is user-invoked (an on-demand pass), while the existing extension spec is
  hook-fired — this is the enhancement-facet question `concepts/marathon-sitrep.md` raises;
  docs-as-extension may need that facet defined first or alongside.
- References that name docs in core (the SKILL.md command table, the pipeline's command map, the
  workspace configuration's `docs =` key, the docs-tier section of
  `references/context-engineering.md`) need the same single-source treatment as the rest of
  v0.9.0.

## Assumptions

- Assumes an extension can carry a user-invoked command (the enhancement facet) without changes
  to the harness's skill routing.
