# marathon docs

Author and curate the project's human-oriented reference documentation in the top-level `docs/`
directory. `docs/` is optional: a project opts into it the first time you run this command, and most
projects never need it. Use `docs` when the project has reached a point where reference documentation —
the kind a person reads to understand the system — is worth writing and keeping.

Documentation is high-effort work that has to stay coherent as the project evolves, so it gets its own
deliberate pass rather than being folded into the build loop. `docs` runs the session pipeline
(`mechanics/pipeline.md`), recorded under `Session: docs`. Branch slug: the pass.

## The standardized tier of context

`docs/` is the top layer of the context lifecycle: standardized context, above the volatile
`concepts/` and the settled `design/`. A design note graduates by being expressed — by the built work,
or by a `docs/` page that standardizes it — while a `docs/` page itself never decays: it describes
what exists, and it is durable and accretive. The tiers and their rules are laid out in
`references/context-engineering.md`.

## Find the landing zone

Before bootstrapping or curating anything, work out which project `docs/` belongs to. Check whether
this project is a standalone project or sits in a workspace (the LOCATE stage of
`mechanics/pipeline.md` makes the same check).

- **Standalone.** `docs/` belongs here; continue below as usual.
- **In a workspace, and the coordinator's `[workspace] docs` names this project.** This project is the
  landing zone; continue below as usual.
- **In a workspace, and `[workspace] docs` names a different project.** Don't bootstrap or curate a
  `docs/` here. Tell the developer the workspace centralizes documentation in that project, and point
  them there.
- **In a workspace, and the coordinator declares no `[workspace] docs`.** No landing zone is settled.
  Ask the developer whether this project should be the one — and add the field to the coordinator's
  `marathon.toml` if so — or whether another project holds it.

## Bootstrapping (first run, no docs/)

If the project has no `docs/` yet, this run establishes it:

1. Settle with the developer what the documentation needs to cover and how to structure it — the
   layers or topics a reader moves through, and where to start.
2. Create `docs/` and write the first pages from the code and the `context/` notes as they stand.
3. Keep it shallow where the project is still moving; document what is settled enough to explain.

## Curating (later runs, docs/ exists)

If `docs/` already exists, this run extends and maintains it:

1. Settle what to add or rework — a new layer the project grew, a section the code outpaced,
   structure that no longer reads well.
2. Write or revise those pages against the current code. When code and its documentation disagree, the
   code wins; the page is what's wrong.
3. Don't pad it. Add the documentation a reader needs, not coverage for its own sake.

## Conclude

Close with `close` — it records the pass in the reset file under `Session: docs` and publishes the
branch. The `review` command keeps `docs/` honest between `docs` passes by flagging pages that have
drifted from the code.
