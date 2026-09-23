# marathon plan

A planning session that changes only `context/`. It refines notes, works out a design, reshapes
the capability map, or decides what the next `start` focuses on. It changes no product, which is
the difference from `start`. `review` looks back for drift, and `plan` looks forward.

`plan` runs the session pipeline (`mechanics/pipeline.md`). A handoff under Session `plan`
resumes here.

## Settle

Work through the topic with the architect in real depth, adding detail only as far as the
upcoming work needs. When choosing the next focus, look for the questions whose answers would
change what the project does next, not only the areas with the most unknowns.

Branch slug: the topic.

## Execute

After agreeing the changes with the architect, bring `context/` in line with what the discussion
decided, in stages (`references/staged-execution.md`):

- **Cull**: delete a note the discussion ruled out.
- **Add or sharpen**: write new notes, sharpen a note to what the discussion decided, and refine
  the capability map or the note the upcoming work needs.

Capture only what the upcoming step needs. Taking in a finished experiment also updates its
catalog entry and archives its remote (`commands/experiment.md`).

## Conclude

End with `close`, which records the dispositions and sets Next-focus to the step this planning
prepared.
