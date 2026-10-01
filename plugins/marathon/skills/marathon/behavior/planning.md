# Planning

These rules govern how every marathon session decides its scope, at the pipeline's PLAN stage
(`mechanics/pipeline.md`).

## Plan one task at a time

Plan the next task in detail and nothing beyond it. Start from the lowest-level requirement,
build it, and plan the following task once it ships. Planning further ahead commits the project
to decisions made without evidence, so notes and roadmap entries stay brief until their work is
close.

Build in dependency order, and investigate risk separately from the builds. When the unknown with
the highest consequence sits above the next task, spike it with `experiment` while the builds
continue from the bottom.

## Check for an existing solution first

Before recommending that a task build something itself, ask whether the language's standard
library, or a dependency its ecosystem treats as standard, already solves the problem. If one
does, building your own needs a good reason, recorded as a rejected alternative in the goal
record. Without a good reason, adopt the existing solution. Ask this in the plan round, while the
answer can still change the brief.

## Ask for decisions, never for facts

PLAN runs in plan mode, in rounds (`references/briefs.md`, "Plan round"). The planner profile
finds the facts and drafts each round (`behavior/delegation.md`); the architect answers every
question in a round at once. The rounds end when no question remains open, and the architect
approves the task brief. Nothing changes before that approval.

## Planning maintains the roadmap

Planning also keeps the path current. Capture ideas for later as roadmap entries or brief notes,
remove what the discussion rules out, and stage, reorder, or drop goals as the architect decides
(`mechanics/goals.md`).
