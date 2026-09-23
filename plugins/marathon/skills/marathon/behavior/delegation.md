# Delegation

A session may hand one unit of its work to another agent. The session still owns the work: it
decides whether to delegate, reads the result firsthand, and reports and commits it as its own.

## The profiles

marathon ships four subagent profiles in the plugin's `agents/` directory:

- **planner**: drafts the stage list at SETTLE from the context the session provides. It changes
  nothing, and the session revises the list with the architect without engaging it again.
- **executor**: implements one stage and brings the stage's check to passing.
- **editor**: edits the prose a branch changed, once, after the last stage commits and before the
  final checkpoint. It edits files in place and never changes their meaning.
- **reviewer**: reviews the whole branch after the architect confirms the final checkpoint, and
  returns the architect's report as text for the session to write (`commands/close.md`).

No profile sets a model. Before each engagement, the session states the profile, the model it
chose, and why, so the architect can redirect it.

## Limits

- A delegate takes one unit: one stage list, one stage, one branch edit, or one branch review. A
  delegate that goes beyond its unit has found something that needs a re-plan.
- A delegate never commits, publishes, or writes the reset file, and it changes nothing at
  SETTLE. The session reads any prose a delegate writes before the repository keeps it.
- Delegation is optional. The session does the work itself when briefing a delegate would cost
  more than the work.
