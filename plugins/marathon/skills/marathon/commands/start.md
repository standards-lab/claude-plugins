# marathon start

Advance the product one concrete step — or resume a step already in progress. `start` takes no
argument; how it hands off the step depends on the project's kind (see `references/role-boundary.md`),
not on anything you pass it. For planning that touches only `context/`, use `plan`; for a spike, use
`experiment`.

`start` runs the session pipeline (`mechanics/pipeline.md`). A handoff recorded under Session `start`
resumes here; on a **code** project, `context/guide.md` is then still the live implementation guide —
follow it from where the Next-focus leaves off.

## Settle

The scope to settle is the single concrete step to take now, and how far it should go. Work it
through with the developer in enough depth to come out with a clear picture of what to build —
planning matters as much as the build; this is where the architectural thinking happens, so give it
real depth and don't rush to the code. Add detail to the relevant `design/` note only as far as this
step needs — no further.

Branch slug: the step.

## Execute

Do the step, according to the project kind:

- **code project** — write `context/guide.md`, the implementation guide: the full code for each change,
  in the order to apply it, with prose only where the reasoning needs it, and a short run-and-verify at
  the end. It does not include the tests or documentation the agent adds at closeout. See
  `references/implementation-guides.md` for how to write it well. Then stop and let the developer apply
  it, staying available for fixes; don't run ahead.
- **context project** — there is no guide and no code handoff. Author the change directly: the skills,
  prose, or configuration the step calls for, plus any `context/` the change settles. You are producing
  the deliverable itself, under the developer's review. Stay within the one step.

## Conclude

`reset` when the context fills before the step is done; `close` when it's finished and validated.
