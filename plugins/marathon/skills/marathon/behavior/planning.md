# Planning

How every marathon session decides its scope, at the pipeline's SETTLE stage
(`mechanics/pipeline.md`).

## Plan one step at a time

Plan the next step in detail and nothing beyond it. Start from the lowest-level requirement,
build it, and plan the following step once it's done. Planning further ahead commits the project
to decisions made without evidence, so notes stay brief until their work is close.

Build in dependency order, and investigate risk separately from the builds. When the unknown with
the highest consequence sits above the next step, spike it with `experiment` while the builds
continue from the bottom.

## Check for an existing solution first

Before settling the stage list, ask whether the language's standard library, or a dependency its
ecosystem treats as standard, already solves the problem. If one does, building your own needs a
good reason, recorded as a rejected alternative in the note the step touches. Without a good
reason, adopt the existing solution. Ask this at SETTLE, while the answer can still change the
plan.

## Planning is half the work

In `init`, a new `start`, `plan`, and `review`, planning is where the architectural thinking
happens: what the step involves, how deep it goes, and how it fits the larger design. Plan in
plan mode, and agree the plan with the architect before any consequential action. A `start` that
resumes a handoff uses the plan it inherits.

Planning also maintains the context. Capture ideas for later steps as notes, remove the notes the
discussion rules out, and watch for what the next session should focus on. The pipeline's SETTLE
stage says when those edits are made.
