# Planning conduct

How every marathon session settles scope, at the pipeline's SETTLE (`mechanics/pipeline.md`).

## One step at a time

Plan the immediate next step in detail and nothing past it. Start from the lowest-level
requirement, build it, and let the next step come into view once it's done. Planning far ahead
commits you to decisions you haven't earned, so notes stay shallow until their work is close.

Builds proceed in dependency order; risk is probed out of band. When the highest-consequence
unknown sits above the step in front, spike it with `experiment` while the builds continue from
the bottom.

## Sufficiency first

Before settling the stage list, ask whether the problem is already solved by the language's
idiomatic approach: the standard library, or a dependency the ecosystem treats as standard. If it
is, building your own needs an adequate reason, recorded in the note the step touches as a
rejected alternative. Without one, adopt the existing approach. Ask at SETTLE, while the answer
still changes the plan.

## Planning is half the work

In `init`, a fresh `start`, `plan`, and `review`, planning is where the architectural thinking
happens: what the step involves, how deep it goes, and how it fits the larger design. Plan in plan
mode, and settle with the architect before any consequential action. A `start` that resumes a
handoff picks up the existing plan instead.

Planning also tends the context: capture ideas for later steps as open notes, cull the ones the
discussion rules out, and watch for the next session's focus. The pipeline's SETTLE says when
those edits land.
