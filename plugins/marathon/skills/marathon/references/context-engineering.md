# Context engineering

Besides moving code forward, marathon's main job is keeping the project's written context small and
accurate. This document explains how the `context/` directory is organized, how to decide where
something goes, and when to delete it.

## Why bother

A model's context window is limited. When it fills with notes about finished work, settled questions,
or documentation the code has outgrown, there is less room for the task at hand, and the stale notes
start to contradict the code. So marathon maintains its written context instead of just accumulating
it.

Two rules follow:

1. The repository is the only source of truth — not an issue tracker, not the chat history. Project
   knowledge lives in the repo, either in `context/` or in the built work and its documentation.
2. Within the repo, the built work wins — the code, or on a context project the authored deliverable
   (its skills, prose, configuration). `context/` is only for things that work and its own
   documentation can't express yet. Once they do, the note describing it is redundant and should go.

## How `context/` is organized

The directory is flat: every file sits directly under `context/`, with no subdirectories.

Stable (changes rarely, and on purpose):

- `context/README.md` — the project's orientation: a short vision statement and a map of the
  capabilities the project will need. Keep the map broad, shallow, and unordered.

Volatile (changes often, and gets cleaned up):

- `context/*.md` — notes: intent the built work doesn't express yet. A note may hold a settled
  decision with its reasoning, constraints, and rejected alternatives, or an idea and its open
  questions. Most of the churn lives here.
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

### A note states its own settledness

Each note states in its opening lines whether it is settled: settled direction, and what settled
it (a session's decision, a build, an experiment's result), or a
concept still open, and what would settle it. A note that describes an interface, a contract, or
a shape a session just designed states itself provisional. It states itself settled only in a
later session, once something real, such as a caller, a build, or an experiment, exercised the
shape and it held. The same rule applies to authoring a skill from a note.

## Project documentation

`docs/` is the project's own documentation: the guide a person reads to use or contribute to the
repository, indexed by the README in reading order, shipped with the code where the ecosystem
packages a repository's tree. It is optional. Many projects never need it, because the README,
the API documentation, and the source are the documentation until a reader needs more, and a
project establishes it when a documentation step of a `start` session settles what the guide
covers and how a reader moves through it. It is not a tier of the context lifecycle. It sits
beside the code as part of the built work, the code is its source of truth, and a page the code
has moved out from under is a defect fixed in the change that moved the code or in the next
documentation step.

The reasoning behind a built capability belongs in the owning repository's documentation, such as
its package documentation, its README, or `docs/`, and the step that builds the capability writes
it there. What generalizes past the repository, a principle, a convention, a definition, is not
project documentation: it stays a note, and no documentation step writes it.

The decay rule below applies to `context/` and never to `docs/`. `context/` is agent-oriented
and decays toward the built work; project documentation describes the code and is durable and
accretive. `review` flags a `docs/` page the code has moved out from under.

## Deciding where something goes

- Orientation (vision or the capability map)? → `context/README.md`
- Anything the built work and its own documentation don't express yet? → a note in `context/`,
  with its settledness stated in its opening lines.
- Already expressed by the built work or its documentation? → it doesn't belong in `context/`;
  delete it.

When you can't tell whether something is settled, write it as a concept. Stating it settled later
is cheap; walking back a note you committed to too early is not.

Before writing any of it, check whether the fact already has exactly one home outside
`context/` — settled documentation included, not only the built work in the strict code sense.
If it does, write a pointer to it (a path, a citation), never a restatement. This is what keeps
`context/` from becoming a second, drifting description of material a repository's own docs
already carry.

## Add detail late

Notes start short. A capability in the map, or a new note, begins as a sentence or two and gets
filled in only when you are about to work on it. Writing detailed design for distant work has the
same problem as writing the code too early: you commit to a design before you understand it, then have
to maintain or delete the note when reality turns out different. Add detail at the point you need it.

## Name the assumptions

A note names the unverified assumptions it rests on — the claims that, if a build falsifies them,
invalidate the note. A line or two is enough ("assumes the loader keeps config immutable after
start"). The payoff comes at the surprise: when work falsifies an assumption, the notes that named
it are the ones to revisit, and the reset Disposition records the falsification — instead of a
judgment sweep over the whole tree at the next `review`. A note that rests only on what is
already built needs no annotation.

## Tending the notes

Four operations keep `context/` accurate, and the reset file's Disposition records each under its
own name:

- **Integrated** — a note deleted because the built work or its own documentation now expresses
  what it described. A note that restates a built interface decays even when it reads as design,
  since the interface's own documentation is the one home for it. A note whose reasoning outlives
  the build moves that reasoning into the owning repository's documentation, written by the step
  that builds the capability, and the note is deleted in the same close. Record the removal and
  point to what now expresses it.
- **Culled** — a note deleted because it is no longer viable: superseded, abandoned, or
  contradicted by the way the work actually went.
- **Retained** — a note deliberately left in place, with the reason stated. A legitimate reason is
  that its home doesn't exist yet, such as an owning repository whose `docs/` a later step
  establishes.
- **Cross-repo** — an edit the session made in another repository's notes or documentation.

A note that settles in place is not moved anywhere: its settledness line changes, and the
Disposition records the note under **Add or sharpen**. An enabled extension may add ledger
entries of its own for what it moves out of `context/`; its own skill defines them.

The point of all four is to keep `context/` short and true, rather than letting it grow into a
parallel description of the project that slowly disagrees with the built work.

## State current truth, not history

A note states what is true now — never how it changed, when, or why. A changelog entry embedded
in a note's own prose ("amended: X caused Y", a dated revision log, "replacing the earlier Z
approach") is exactly the kind of parallel description `context/` exists to avoid: it makes the
note larger without making it truer, and the next reader has to separate the current claim from
its history to use it. Provenance — what a session did, and why — belongs in the reset file's
Disposition, the session's own ledger; the note it concerns keeps only the fact the session
settled. A settledness line names what settled the note, never the sequence of states before it.

## Check before you cut

Integrating, culling, and moving reasoning into a repository's documentation change what the
written record says about the project, so don't do them on your own. Show the architect what you
propose to change or delete and get agreement first — in plan mode for `init`, a fresh `start`,
`plan`, and `review`, and as a quick confirmation during `reset` and `close`.
