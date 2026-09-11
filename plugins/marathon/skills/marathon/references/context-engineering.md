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
- `context/reset.md` — the latest session's record and the pointer to the next step, kept by a
  standalone project and by a workspace coordinator; a workspace member carries none. One file;
  git keeps the older versions (`mechanics/reset-file.md`).

Kept outside `context/`:

- `experiments/` — a top-level directory for spikes, isolated so throwaway work doesn't mix into the
  real tree. Created when an experiment session needs it. A standalone project keeps its own; in
  a workspace, the coordinator keeps the only one (`references/workspace-coordination.md`).
- `docs/` — an optional top-level directory for the project's own documentation, the guide a
  person reads to use or contribute to the repository. See below.
- whatever an enabled extension owns — a directory or file outside `context/` that the extension
  maintains and bootstraps in the session after it is enabled (`references/extensions.md`). The
  core adds none.
- the source code — the implementation, and the final word on what the project does.

## Project documentation

`docs/` is the project's own documentation: the guide a person reads to use or contribute to the
repository, indexed by the README in reading order, shipped with the code where the ecosystem
packages a repository's tree. It is optional. Many projects never need it, because the README,
the API documentation, and the source are the documentation until a reader needs more, and a
project establishes it when a documentation step of a `start` session settles what the guide
covers and how a reader moves through it. It is not a tier of the context lifecycle. It sits
beside the code as part of the built work, the code is its source of truth, and a page the code
has moved out from under is a defect fixed in the change that moved the code or in the next
documentation step. What generalizes past the repository, a principle, a convention, a
definition, is not project documentation: it settles in `design/` like any other intent, and no
documentation step writes it.

The decay rule below applies to `context/` and never to `docs/`. `context/` is agent-oriented
and decays toward the code; project documentation describes the code and is durable and
accretive. `review` flags a `docs/` page the code has moved out from under.

## Deciding where something goes

- Orientation (vision or the capability map)? → `context/README.md`
- Settled intent the code doesn't express yet? → `context/design/`
- Still a concept or an open question? → `context/concepts/`
- Already expressed by the built work? → it doesn't belong in `context/`; delete it.

When you can't tell whether something is settled, treat it as a concept. Promoting it later is cheap;
walking back a design note you committed to too early is not.

Before writing any of it, check whether the fact already has exactly one home outside
`context/` — settled documentation included, not only the built work in the strict code sense.
If it does, write a pointer to it (a path, a citation), never a restatement. This is what keeps
`context/` from becoming a second, drifting description of material a repository's own docs
already carry.

## Add detail late

Notes start short. A capability in the map, or a topic in `design/`, begins as a sentence or two and
gets filled in only when you are about to work on it. Writing detailed design for distant work has the
same problem as writing the code too early: you commit to a design before you understand it, then have
to maintain or delete the note when reality turns out different. Add detail at the point you need it.

## Name the assumptions

A note in `design/` or `concepts/` names the unverified assumptions it rests on — the claims that,
if a build falsifies them, invalidate the note. A line or two is enough ("assumes the loader keeps
config immutable after start"). The payoff comes at the surprise: when work falsifies an
assumption, the notes that named it are the ones to revisit, and the reset Disposition records the
falsification — instead of a judgment sweep over the whole tree at the next `review`. A note that
rests only on what is already built needs no annotation.

## Moving and removing notes

Three operations keep `context/` accurate:

- **Promote** — move a concept from `concepts/` to `design/` once it is settled: a decision fixed it,
  the built work proved it out, or an experiment produced a result. Don't do it silently — move the file
  and note why in the reset file. Organizing `concepts/` and `design/` the same way makes it obvious
  where a note should land. A note documenting an interface, a contract, or a shape a session just
  designed is not promoted in that same session — the same rule applies to authoring a skill.
  Promotion waits for a session where something real, a caller, a build, an experiment, exercised the
  shape and it held; until then it stays in `concepts/`, explicitly provisional, even if it reads as
  finished.
- **Decay** — delete a `design/` note once the built work fully expresses what it described, and
  the note holds no conceptual or pattern detail beyond it. A note that still explains a pattern,
  a boundary, or a style the built work cannot state on its own is doing design work and stays;
  duplication in API documentation alone is not decay. A note that does decay is a weaker second
  copy the built work will drift from — record the removal (and point to the code or deliverable
  that expresses it) in the reset file.
- **Cull** — delete a concept in `concepts/` when it is no longer viable: superseded, abandoned, or
  contradicted by the way the work actually went.

The reset file records these operations in its ledger vocabulary: a decayed note is recorded as
**Integrated**, a note deliberately left in place as **Retained**; Promoted and Culled record
themselves.

The point of all three is to keep `context/` short and true, rather than letting it grow into a parallel
description of the project that slowly disagrees with the code.

## State current truth, not history

A note in `design/` or `concepts/` states what is true now — never how it changed, when, or why.
A changelog entry embedded in a note's own prose ("amended: X caused Y", a dated revision log,
"replacing the earlier Z approach") is exactly the kind of parallel description `context/` exists
to avoid: it makes the note larger without making it truer, and the next reader has to separate
the current claim from its history to use it. Provenance — what a session did, and why — belongs
in the reset file's Disposition, the session's own ledger; the note it concerns keeps only the
fact the session settled.

## Check before you cut

Promoting, decaying, and culling change what the documentation says about the project, so don't do them
on your own. Show the architect what you propose to move or delete and get agreement first — in plan
mode for `init`, a fresh `start`, `plan`, and `review`, and as a quick confirmation during `reset` and
`close`.
