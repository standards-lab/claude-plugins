# The session loop

A session is the unit of work in marathon: one focused step, on one branch, from start to a clean
stopping point. This document walks through the loop in full. For the commands themselves, see the
playbooks under `commands/`.

## One session, one step, one branch

Each session does a single concrete thing — the immediate next step, and nothing wider. It happens on
its own branch and ends by publishing that branch to the remote. Keeping sessions this small is what
makes a long project manageable: you're never holding more than one step in your head, and the history
reads as a sequence of clear, finished changes.

Three commands open a working session — `plan`, `start`, and `experiment` — and which one you run says
what the session is for: planning and curating the context, advancing the product, or spiking an idea.
They share the loop below.

## Starting

A working session reads `context/reset.md` and looks at its Status. A missing reset file means the
project is at a resting point: its deliverable is released and continuity sits with the workspace
coordinator (see the Continuity section in `SKILL.md`).

If the Status is `closeout`, the last step is done and this is a new one. This begins with planning,
which matters as much as the build. In plan mode, get
oriented from the Next-focus, the capability map, and the relevant notes, then work the step through
with the developer in real depth: what it involves, how far it should go, and how it fits the wider
design. You want to come out with a clear picture of what you're about to do. The discussion will
naturally range past the single step, so use it to tend the context too — capture later ideas as
concepts in `concepts/`, cull the ones it rules out, and note what the next focus looks like. Once the
scope is agreed, branch and go.

If the Status is `handoff`, the last session stopped mid-work to free up context. The plan already
exists, so skip planning: check out the open branch, read the in-progress state and Next-focus, and pick
up from there.

## Doing the work

What happens next depends on the command and the project's kind.

In a `start` on a **code** project, you write the implementation guide — the concrete changes to make,
in order, with the reasoning where it's needed (see `references/implementation-guides.md`) — and then stop. The
developer applies it. You stay available for fixes but don't get ahead of them. This pause is the role
boundary in action (see `references/role-boundary.md`): the developer owns the implementation, so the agent hands
off the draft and waits.

In a `start` on a **context** project there's no guide and no code to hand off. You author the change
directly — the skills, prose, or configuration the step calls for — because on a context project the
deliverable is the writing itself.

In a `plan` session you edit the notes in `context/` and nothing else — this is where bigger-picture
planning and concept work happen, between builds. In an `experiment` session you work in the isolated
top-level `experiments/` directory and treat whatever you find as a concept, not settled design.

## Handing off mid-session

If the context window fills up before the work is done, don't push through a degraded context. Run
`reset`: tidy the notes you touched, then write `context/reset.md` with `Status: handoff` and a
Next-focus that records exactly where you are and what's next. Leave the branch open. A fresh `start`
resumes it cold from that record. This is what lets a single step span more than one context window
without losing the thread.

## Closing

When the step is finished and validated, `close` ends the session: on a code project's `start` the agent
finishes its part (tests and documentation); you tidy the notes against what now exists; you agree on the
next step with the developer; and the result is published to the remote — a pull request, a merge
request, or the equivalent — with its description from the reset file. The Next-focus you set becomes the
starting point for the next session.

## Why the loop holds together

Every session starts and ends at the same place — `context/reset.md` — so the project always has a
single, current answer to "where are we and what's next," and it lives in the repo rather than in
anyone's memory or a chat log. That's what makes the work resumable by a fresh context, and sustainable
across the length of the project.
