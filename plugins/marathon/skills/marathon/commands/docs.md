# marathon docs

Author and curate the project's human-oriented reference documentation in the top-level `docs/`
directory. `docs/` is optional: a project opts into it the first time you run this command, and most
projects never need it. Use `docs` when the project has reached a point where reference documentation —
the kind a person reads to understand the system — is worth writing and keeping.

Documentation is high-effort work that has to stay coherent as the project evolves, so it gets its own
deliberate pass rather than being folded into the build loop. Run it in plan mode and settle the scope
with the developer before writing.

## docs/ is not context/

`docs/` and `context/` share a posture — keep them curated and in sync, don't let them rot — but they are
different kinds of writing with opposite lifecycles:

- `context/` is agent-oriented working knowledge that decays *toward* the code. A `design/` note is a
  defect once the code expresses it, so `context/` shrinks as the code grows.
- `docs/` is human-oriented reference documentation that *describes* the code. A page is born once the
  code is ready to be explained, and it is durable and accretive.

So the decay rule does not apply to `docs/`. Documenting the code is the point.

## Bootstrapping (first run, no docs/)

If the project has no `docs/` yet, this run establishes it:

1. In plan mode, work out with the developer what the documentation needs to cover and how to structure
   it — the layers or topics a reader moves through, and where to start.
2. Create `docs/` and write the first pages from the code and the `context/` notes as they stand.
3. Keep it shallow where the project is still moving; document what is settled enough to explain.

## Curating (later runs, docs/ exists)

If `docs/` already exists, this run extends and maintains it:

1. In plan mode, decide what to add or rework — a new layer the project grew, a section the code outpaced,
   structure that no longer reads well.
2. Write or revise those pages against the current code. When code and its documentation disagree, the
   code wins; the page is what's wrong.
3. Don't pad it. Add the documentation a reader needs, not coverage for its own sake.

## Land the changes

Documentation is the agent's to write outright (see `references/role-boundary.md`), so there is no code
handoff. Land the changes like any other work — on a branch, published the same way — and record what you
changed in `context/reset.md`, the same as any closeout. The `review` command keeps `docs/` honest between
`docs` passes by flagging pages that have drifted from the code.
