# marathon start

Advance the product one concrete step — or resume a step already in progress. `start` takes no
argument; how it executes the step depends on the project's kind (see
`mechanics/configuration.md`), not on anything you pass it. For planning that touches only
`context/`, use `plan`; for a spike, use `experiment`.

`start` runs the session pipeline (`mechanics/pipeline.md`). A handoff recorded under Session
`start` resumes here; on a **code** project, the reset file's Next-focus carries the approved
stage list and the stage position — pick up the loop from there.

## Settle

The scope to settle is the single concrete step to take now, and how far it should go. Work it
through with the architect in enough depth to come out with a clear picture of what to build —
planning matters as much as the build; this is where the architectural thinking happens, so give
it real depth and don't rush to the code. Add detail to the relevant `design/` note only as far as
this step needs — no further.

On a **code** project, the settled scope is expressed as the stage list
(`references/staged-execution.md`), approved by the architect before the branch is created. In a
workspace, a step may span member repos: the stage list then groups by repository, ordered by the
coordinator's `order` map, and the session creates a branch in each touched repo under the step's
shared slug.

Branch slug: the step.

## Execute

Do the step, according to the project kind:

- **code project** — run the stage loop of `references/staged-execution.md`: implement a stage,
  bring its compilation unit to green, commit, report, and stop for the architect's review;
  repeat through the list, then run validation. Context tending — including the edits noted at
  SETTLE — is closeout work, done after validation.
- **context project** — author the change directly: the skills, prose, or configuration the step
  calls for, plus any `context/` the change settles. You are producing the deliverable itself,
  under the architect's review. Stay within the one step.

## Conclude

`reset` when the context fills before the step is done; `close` when it's finished and validated.
