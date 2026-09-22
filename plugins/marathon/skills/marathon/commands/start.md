# marathon start

Advance the product one concrete step, or resume one in progress. How the step executes follows
the project kind (`mechanics/configuration.md`). For planning that touches only `context/`, use
`plan`; for a spike, use `experiment`.

`start` runs the session pipeline (`mechanics/pipeline.md`). A handoff under Session `start`
resumes the stage loop from its recorded position.

## Settle

Settle the single concrete step and how far it goes, in real depth (`behavior/planning.md`). Add
detail to the note it touches only as far as the step needs. Express the scope as the stage list
with its checkpoints (`references/staged-execution.md`), approved before the branch is created.

Branch slug: the step.

## Execute

Run the stage loop, then validation. A **code** project's stages are compilation units brought to
green with their tests, and context tending waits for `close`. A **context** project's stages are
the deliverable itself plus any `context/` the change settles.

## Conclude

`reset` when the context fills first; `close` when the step is finished and validated.
