# Context engineering

Besides moving code forward, marathon's main job is keeping the project's written context small and
accurate. This document explains how the `context/` directory is organized, how to decide where
something goes, and when to move or delete it.

## Why bother

A model's context window is limited. When it fills with notes about finished work, settled questions,
or documentation the code has outgrown, there is less room for the task at hand, and the stale notes
start to contradict the code. So marathon maintains its written context instead of just accumulating
it.

Two rules follow:

1. The repository is the only source of truth — not an issue tracker, not the chat history. Project
   knowledge lives in the repo, either in `context/` or in the code.
2. Within the repo, the code wins. `context/` is only for things the code can't express yet. Once the
   code expresses something, the note describing it is redundant and should go.

## How `context/` is organized

The directory holds a few kinds of files, grouped by how often they change.

Stable (changes rarely, and on purpose):

- `context/README.md` — the project's orientation: a short vision statement and a map of the
  capabilities the project will need. Keep the map broad and shallow; it is not an ordered plan.
- `context/design/` — design notes: validated intent the code can't express yet. The reasoning behind a
  decision, constraints, alternatives that were rejected, intent for parts not yet built.

Volatile (changes often, and gets cleaned up):

- `context/concepts/` — candidate notes: ideas and open questions that aren't settled. Most of the
  churn lives here.
- `context/guide.md` — the current session's implementation guide. One file, deleted when the session
  closes.
- `context/reset.md` — the latest session's record and the pointer to the next step. One file; git
  keeps the older versions.

Kept outside `context/`:

- `experiments/` — a top-level directory for spikes, isolated so throwaway work doesn't mix into the
  real tree. Created when an experiment session needs it.
- the source code — the implementation, and the final word on what the project does.

## Deciding where something goes

- Orientation (vision or the capability map)? → `context/README.md`
- Settled intent the code doesn't express yet? → `context/design/`
- Still a candidate or an open question? → `context/concepts/`
- Already expressed by the code? → it doesn't belong in `context/`; delete it.

When you can't tell whether something is settled, treat it as a candidate. Promoting it later is cheap;
walking back a design note you committed to too early is not.

## Add detail late

Notes start short. A capability in the map, or a topic in `design/`, begins as a sentence or two and
gets filled in only when you are about to work on it. Writing detailed design for distant work has the
same problem as writing the code too early: you commit to a shape before you understand it, then have to
maintain or delete the note when reality turns out different. Add detail at the point you need it.

## Moving and removing notes

Three operations keep `context/` accurate:

- **Promote** — move a candidate from `concepts/` to `design/` once it is settled: a decision fixed it,
  code proved it out, or an experiment produced a result. Don't do it silently — move the file and note
  why in `context/reset.md`. Organizing `concepts/` and `design/` the same way makes it obvious where a
  note should land.
- **Decay** — delete a `design/` note once the code fully expresses what it described. At that point the
  note is a weaker second copy of the code, and the two will drift apart. Record the removal (and point
  to the code) in the reset file.
- **Cull** — delete a candidate in `concepts/` when it is no longer viable: superseded, abandoned, or
  contradicted by the way the implementation actually went.

The point of all three is to keep `context/` short and true, rather than letting it grow into a parallel
description of the project that slowly disagrees with the code.

## Check before you cut

Promoting, decaying, and culling change what the documentation says about the project, so don't do them
on your own. Show the developer what you propose to move or delete and get agreement first — in plan
mode for `init`, a fresh `start`, and `review`, and as a quick confirmation during `reset` and `close`.
