# Delegation

A session may hand one unit of its work to another agent and stays its owner: it decides whether
to delegate, reads what comes back firsthand, and reports and commits it as its own.

## The profiles

marathon ships three subagent profiles in the plugin's `agents/` directory:

- **planner** — drafts the stage list at SETTLE from the context the session hands it. It changes
  nothing; the session revises the list with the architect without re-engaging it.
- **executor** — implements one stage against the stage's check.
- **reviewer** — reviews the whole branch once the architect confirms the final checkpoint, and
  returns the architect's report as text for the session to write (`commands/close.md`).

No profile pins a model. Before each engagement, the session states the profile, the model it
chose, and why, so the architect can redirect.

## Limits

- The grain is one unit: one stage list, one stage, one branch review. A delegate that reaches
  past its unit is a finding for a re-plan.
- A delegate never commits, publishes, or writes the reset file, and changes nothing at SETTLE.
  It drafts, and never finishes, prose the repository keeps.
- Delegation is a choice. The session does the work itself when briefing a delegate would cost
  more than the work.
