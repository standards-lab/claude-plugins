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
- `architecture/` — the architecture layer: the principles, definitions, and conventions that
  have generalized past one repository. A conventional layer, present in every marathon
  project: a standalone project keeps it as a top-level directory, and a workspace keeps it as
  one repository named by the coordinator. See below.
- `docs/` — an optional top-level directory for the project's own documentation, the guide a
  person reads to use or contribute to the repository. See below.
- the source code — the implementation, and the final word on what the project does.

## The architecture layer

The architecture layer is the top of the context lifecycle: the principles, the definitions, and
the conventions that have generalized past one repository, written for a general reader and
relied on from outside the project. It is a conventional layer of every marathon project, not
an option, because it is where a design note goes when the built work cannot express it.

- A **standalone project** keeps its architecture in a top-level `architecture/` directory,
  scaffolded by `init` with a README as its index.
- A **workspace** keeps its architecture in one repository, named by the coordinator's
  `[workspace] architecture` field (`mechanics/configuration.md`). That repository is a
  `context` project whose whole tree is the architecture, with a README at its root and in
  every directory as the index GitHub renders, and it runs sessions like any other project.

The architecture holds only what generalizes. Nothing a reader could infer from a repository's
source belongs in it, so it stays stable while the repositories change beneath it, and a page
that restates a repository's implementation is a defect, reduced to the principle it states or
removed. A repository links the architecture's principles from its README and states beside the
link any convention of its own that narrows a principle; it never restates the page. The
architecture may catalog the repositories that implement it, with a description and a link for
each, and goes no deeper.

Knowledge reaches the architecture through the promotion sequence, and only by it. A concept
proves out and is promoted into `design/` in the repository that owns it. A design note is
expressed either by the built work, at which point it decays under the rule below, or by an
architecture page, once the knowledge has generalized past that one repository. In a workspace
the second case is a cross-repository step: the member's `close` or `review` finds that a design
note has generalized and lands it as a concept in the architecture repository, recorded under
**Cross-repo**, and the architecture repository authors the page in a session of its own. A
design note that describes one repository's implementation is expressed by the code and the
repository's own documentation, never by the architecture.

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
definition, is not project documentation: it belongs to the architecture layer and reaches it by
promotion from `design/`, never by a documentation step.

The decay rule below applies to `context/` and never to the architecture or to `docs/`.
`context/` is agent-oriented and decays toward the code; project documentation describes the
code and is durable and accretive; the architecture states what generalizes and changes only
when a principle does. `review` flags a `docs/` page the code has moved out from under, and, in
an architecture repository, a page that restates a repository.

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
  where a note should land.
- **Decay** — delete a `design/` note once the built work, or an architecture page, fully expresses
  what it described, and the note holds no conceptual or pattern detail beyond it. A note that still
  explains a pattern, a boundary, or a style neither can state on its own is doing design work and
  stays; duplication in API documentation or an architecture page alone is not decay. A note that
  does decay is a weaker second copy the built work will drift from — record the removal (and point
  to the code, deliverable, or architecture page) in the reset file.
- **Cull** — delete a concept in `concepts/` when it is no longer viable: superseded, abandoned, or
  contradicted by the way the work actually went.

The reset file records these operations in its ledger vocabulary: a decayed note is recorded as
**Integrated**, a note deliberately left in place as **Retained**; Promoted and Culled record
themselves.

The point of all three is to keep `context/` short and true, rather than letting it grow into a parallel
description of the project that slowly disagrees with the code.

## Check before you cut

Promoting, decaying, and culling change what the documentation says about the project, so don't do them
on your own. Show the architect what you propose to move or delete and get agreement first — in plan
mode for `init`, a fresh `start`, `plan`, and `review`, and as a quick confirmation during `reset` and
`close`.
