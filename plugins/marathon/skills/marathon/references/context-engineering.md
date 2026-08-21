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
2. Within the repo, the built work wins — the code, or on a context project the authored deliverable
   (its skills, prose, configuration). `context/` is only for things that work can't express yet. Once
   it does, the note describing it is redundant and should go.

## How `context/` is organized

The directory holds a few kinds of files, grouped by how often they change.

Stable (changes rarely, and on purpose):

- `context/README.md` — the project's orientation: a short vision statement and a map of the
  capabilities the project will need. Keep the map broad, shallow, and unordered.
- `context/design/` — design notes: validated intent the code can't express yet. The reasoning behind a
  decision, constraints, alternatives that were rejected, intent for parts not yet built.

Volatile (changes often, and gets cleaned up):

- `context/concepts/` — concepts: ideas and open questions that aren't settled. Most of the churn lives
  here.
- `context/guide.md` — a code project's implementation guide for the current `start`. Created during the
  session and deleted when it closes; a context project never has one.
- `context/reset.md` — the latest session's record and the pointer to the next step. One file; git
  keeps the older versions.

Kept outside `context/`:

- `experiments/` — a top-level directory for spikes, isolated so throwaway work doesn't mix into the
  real tree. Created when an experiment session needs it.
- `docs/` — an optional top-level directory for human-oriented reference documentation. See below.
- the source code — the implementation, and the final word on what the project does.

## The docs/ tier

`docs/` is reference documentation written for people: the explanation a reader works through to
understand the system. It is optional — most projects stay `context/`-only — and a project opts in at
any time by running the `docs` command, which bootstraps the tier on its first run and curates it after
that.

Where `docs/` lives depends on whether the project stands alone or belongs to a workspace. A standalone
project keeps its own `docs/`, built and curated the way described above. A project in a workspace never
grows a `docs/` of its own; documentation centralizes in one landing-zone project instead, named by the
coordinator's `[workspace] docs` field. A member repository links to the landing zone's pages rather
than restating them, and, where its own convention adds to or narrows what a linked page states, records
that addition beside the link — in its README or its `context/` — rather than duplicating the page.
`commands/docs.md` checks which case applies before it bootstraps or curates anything.

`docs/` borrows the posture of context engineering — keep it curated, keep it in sync, don't let it rot —
but not the lifecycle, because it is the opposite kind of writing:

- `context/` is agent-oriented and decays *toward* the code. A `design/` note is a defect once the code
  expresses it, so `context/` shrinks as the code grows.
- `docs/` is human-oriented and *describes* the code. A page is born once the code is ready to be
  explained, and it is durable and accretive.

So the decay rule below applies to `context/`, never to `docs/`: a `docs/` page restating the code is
doing its job, not duplicating it. The `docs` command does the deliberate authoring; `review` flags
`docs/` pages the code has moved out from under, and a later `docs` pass rewrites them.

## Deciding where something goes

- Orientation (vision or the capability map)? → `context/README.md`
- Settled intent the code doesn't express yet? → `context/design/`
- Still a concept or an open question? → `context/concepts/`
- Already expressed by the built work? → it doesn't belong in `context/`; delete it.

When you can't tell whether something is settled, treat it as a concept. Promoting it later is cheap;
walking back a design note you committed to too early is not.

## Add detail late

Notes start short. A capability in the map, or a topic in `design/`, begins as a sentence or two and
gets filled in only when you are about to work on it. Writing detailed design for distant work has the
same problem as writing the code too early: you commit to a design before you understand it, then have
to maintain or delete the note when reality turns out different. Add detail at the point you need it.

## Moving and removing notes

Three operations keep `context/` accurate:

- **Promote** — move a concept from `concepts/` to `design/` once it is settled: a decision fixed it,
  the built work proved it out, or an experiment produced a result. Don't do it silently — move the file
  and note why in `context/reset.md`. Organizing `concepts/` and `design/` the same way makes it obvious
  where a note should land.
- **Decay** — delete a `design/` note once the built work, or (in a workspace) the docs landing zone,
  fully expresses what it described, and the note holds no conceptual or pattern detail beyond it. A
  note that still explains a pattern, a boundary, or a style neither can state on its own is doing
  design work and stays; duplication in API documentation or a landing-zone page alone is not decay. A
  note that does decay is a weaker second copy the built work will drift from — record the removal (and
  point to the code, deliverable, or landing-zone page) in the reset file.
- **Cull** — delete a concept in `concepts/` when it is no longer viable: superseded, abandoned, or
  contradicted by the way the work actually went.

The reset file records these operations in its ledger vocabulary: a decayed note is recorded as
**Integrated**, a note deliberately left in place as **Retained**; Promoted and Culled record
themselves.

The point of all three is to keep `context/` short and true, rather than letting it grow into a parallel
description of the project that slowly disagrees with the code.

## Check before you cut

Promoting, decaying, and culling change what the documentation says about the project, so don't do them
on your own. Show the developer what you propose to move or delete and get agreement first — in plan
mode for `init`, a fresh `start`, `plan`, and `review`, and as a quick confirmation during `reset` and
`close`.
