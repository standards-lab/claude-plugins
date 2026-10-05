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

## Start from current dependencies

Before round 1, the planner runs each touched repository's currency command
(`mechanics/pipeline.md`, 3 · PLAN). What trails is planned with the task, as the brief's first
slices (`references/briefs.md`, "Task brief"), so the task builds on current dependencies.

## Ask for decisions, never for facts

PLAN runs in plan mode, in rounds (`references/briefs.md`, "Plan round"). The planner profile
finds the facts and drafts each round (`behavior/delegation.md`); the architect answers every
question in a round at once. The rounds end when no question remains open, and the architect
approves the task brief. Nothing changes before that approval.

A round is printed in full in the reply, in the plan round's format, and the plan file keeps a
copy. The architect answers it inline, in their next message. A round is never presented or
collected through AskUserQuestion: its headlines drop each question's options, `rec:`, and
`changes:`, and the architect can't decide from them. Plan mode tells the session to end each
turn with AskUserQuestion or ExitPlanMode. While a round is open, the turn ends with the round
printed instead. Only the task brief's approval, once the rounds end, goes through ExitPlanMode,
with the full brief in the plan file. Intake rounds and BUILD escalations follow the same rule.

## Planning maintains the roadmap

Planning also keeps the path current. Capture ideas for later as roadmap entries or brief notes,
remove what the discussion rules out, and stage, reorder, or drop goals as the architect decides
(`mechanics/goals.md`).
