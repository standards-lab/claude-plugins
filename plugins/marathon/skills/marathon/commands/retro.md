# marathon retro

Turn what the architect had to say about past work into changes to the system, so no comment is
written twice. `retro` reviews a task, a goal, or a date range, and includes the drift pass: notes
the code now expresses or contradicts. It runs LOCATE and START (`mechanics/pipeline.md`).

## Gather

For the range, read:

- the merged pull requests and their bodies, which are the session briefs, and the architect's
  comments on them
- the redirects and escalations the goal records' Decisions recorded
- the commits the standards-reviewer made, which show what the implementer kept missing
- the notes, `context/README.md`, and `docs/` pages the range touched, for drift

## Route

Route each finding to the cheapest layer that would have caught it, in this order:

1. **check/eval**: a rule a check can enforce, or for a plugin, an eval case
2. **standard**: a judgement call for the repository's `STANDARDS.md`
3. **skill/profile**: a change to a skill or a subagent profile
4. **context**: a note to cull, sharpen, or write, or a pointer to add

Drift findings route the same way: a note the built work expresses is culled, a note it
contradicts is fixed, and a `docs/` page it made wrong becomes a task.

Show the findings in the retro format (`references/briefs.md`, "Retro"), each one line with its
proposed change, and let the architect tick the ones to apply.

## Apply

- A ticked finding in the coordinator's `context/` is a direct commit there, as in `plan`.
- A ticked finding in a repository no active goal locks is applied on a branch named
  `retro-<topic>`, then published and merged as SHIP does, with the ticked findings as its body.
- A ticked finding in a repository another active goal locks goes into that goal's record as a
  decision for its next task, or into the roadmap as a task.
