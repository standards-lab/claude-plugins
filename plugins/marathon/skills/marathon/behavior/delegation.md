# Delegation

A session may hand one unit of its own work to another agent instead of doing it directly: the
initial stage list, a stage's implementation, or the review of the finished branch. The session
stays the delegation's owner. It decides whether to delegate, reviews what comes back, and
reports and commits exactly as it would have for work it did itself.

## The profiles

marathon ships three subagent profiles, one per purpose, in the plugin's `agents/` directory.
Each profile's own file holds its instructions.

- **planner** — designs the session's initial stage list at SETTLE, from the full context the
  session hands it. It changes nothing and returns a recommendation. The session revises the
  list with the architect afterward; a revision during SETTLE or a re-plan needs no new planner
  engagement.
- **executor** — implements one technical stage of the approved stage list against the stage's
  own check.
- **reviewer** — reviews the whole branch once the final technical stage lands and writes the
  architect's report, the branch review of `commands/close.md`.

The profiles are fixed by purpose, so a project declares nothing to use them.

## Choosing the model

No profile pins a model. The session chooses the model for each engagement, by what the work in
front of it needs, and passes it when it engages the profile. Before engaging a profile, the
session states which profile, which model, and why, so the architect can redirect.

## The grain

One coherent unit of work: one stage list, one stage of the approved list, or one branch review.
Never a whole session, and never a sub-step inside a unit. A delegate that reaches further than
the unit it was given is a finding for a re-plan, not a wider delegation.

## What a delegate never does

A delegate never commits, never publishes, and never writes the reset file. Engaged at SETTLE,
it changes nothing and returns a recommendation, the same as any other planning-phase work; plan
mode's rule holds through the delegation. It drafts, and never finishes, prose the repository
keeps: commit messages, context notes, documentation.

## What the session owes

The session reads what the delegate produced firsthand, meaning the diff, the check output, and
the reasoning, before it reports or writes anything from it. A stage's log entry and the report
the architect reads are the session's own responsibility.

## Doing the work directly

Delegation is a choice, never a requirement. A session does the work itself when briefing a
delegate would cost more than the work, and nothing about the pipeline changes either way.
