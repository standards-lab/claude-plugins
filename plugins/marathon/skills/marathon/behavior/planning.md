# Planning conduct

Always-active discipline for how marathon sessions plan. It governs every command that settles
scope; the pipeline (`mechanics/pipeline.md`) is where it acts.

## One step at a time

marathon plans the immediate next step in detail and nothing past it, and a session's focus never
spreads beyond that one area. Start from the lowest-level requirement, build it, and let the next
step come into view once it's done. Over many sessions, these small finished steps stack up into
the complete solution.

This is why the notes in `context/` stay shallow until the work on them is close. Planning far
ahead commits you to decisions you haven't earned yet, and it's usually the planning-ahead, not
the building, that turns out wrong. Keep the focus narrow and let the design grow out of working
code.

Builds proceed in dependency order; risk is probed out of band. When the highest-consequence
unknown sits above the step in front, `experiment` is the risk-first instrument: spike it cheaply
while the builds continue from the bottom.

## Planning is half the work

Planning matters as much as building. In `init`, in a fresh `start`, in a `plan` session, and in
`review`, the planning phase is where the real architectural thinking happens — you work out what
the step involves, how deep it needs to go, and how it fits the larger design. The quality of the
implementation is largely set here, so plan with enough depth and clarity that you come out with a
clear picture of what you're about to build. Rushing planning to get to the code is how a session
builds the wrong thing, or the right thing at the wrong depth.

Because the discussion ranges wider than the single step, planning is also where you tend the
context: capture ideas that belong to later steps as concepts in `concepts/`, cull the ones the
discussion has ruled out, and start to spot what the next session's focus should be. When those
edits land depends on the project kind; the pipeline's SETTLE stage holds the rule.

marathon plans in plan mode and settles things with the architect before any consequential
action — `init` before scaffolding, a fresh `start` before implementing the first stage or
authoring the change, `plan` and `review` before changing notes. (A `start` that resumes a handoff skips this
and picks up the existing plan.)
