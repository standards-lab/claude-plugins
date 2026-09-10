# Delegation

A session may hand one unit of its own work to another agent instead of doing it directly — a
stage's implementation, a design decision, a call the session judges too consequential to settle
alone — and stays the delegation's owner: it decides whether to delegate, reviews what comes
back, and reports and commits exactly as it would have for work it did itself.

## Cataloging agents

A project declares the agents its sessions may delegate to in `.claude/marathon.toml`, under
`[agents]`: one sub-table per agent, each with a `delegation` field describing, in the project's
own words, what the agent is for and why. A session reads that catalog and judges which agent,
if any, fits the work in front of it — never a fixed role lookup, since what an agent is for is
entirely the project's declaration. A `delegation` field earns its keep only if it carries
enough for that judgment on its own: what the agent is for, what makes the work warrant it, and
what it isn't for — a field that states only a unit of grain gives the session nothing to weigh
delegating against doing the work directly. `mechanics/configuration.md` holds the schema,
including how a workspace coordinator's declarations layer under a project's own.

## The grain

One coherent unit of work: one stage of the approved list, one settled design question, one bug.
Never a whole session, and never a sub-step inside a unit — a delegate that reaches further than
the unit it was given is a finding for a re-plan, not a wider delegation.

## What a delegate never does

A delegate never commits, never publishes, and never writes the reset file. Engaged at SETTLE,
it changes nothing and returns a recommendation, the same as any other planning-phase work; plan
mode's rule holds through the delegation. It drafts, and never finishes, prose the repository
keeps — commit messages, context notes, documentation.

## What the session owes

The session reads what the delegate produced firsthand — the diff, the check output, the
reasoning — before it reports or writes anything from it. The stage report is the session's own,
and the architect reviews the working tree, not a summary of it. Before engaging a delegate for a
consequential call, the session states which one and why, so the architect can redirect.

## Where none fits

A project that declares no agent, or none whose `delegation` fits the work at hand, changes
nothing about the pipeline: the session does the work itself.
