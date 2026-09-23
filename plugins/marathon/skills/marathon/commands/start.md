# marathon start

Advance the product one concrete step, or resume a step in progress. How the step runs depends
on the project kind (`mechanics/configuration.md`). For planning that changes only `context/`,
use `plan`. For a spike, use `experiment`.

`start` runs the session pipeline (`mechanics/pipeline.md`). A handoff under Session `start`
resumes the stage loop at its recorded position.

## Settle

Settle the one concrete step and how far it goes, in real depth (`behavior/planning.md`). Add
detail to the note the step touches only as far as the step needs. Express the scope as a stage
list with its checkpoints (`references/staged-execution.md`). The architect approves the list
before the branch is created.

Branch slug: the step.

## Execute

Run the stage loop, then validation:

- On a **code** project, each stage is a compilation unit brought to passing with its tests.
  Context edits wait for `close`.
- On a **context** project, the stages are the deliverable itself, plus any `context/` edits the
  change settles.

## Conclude

Run `reset` if the context fills first. Run `close` when the step is finished and validated.
