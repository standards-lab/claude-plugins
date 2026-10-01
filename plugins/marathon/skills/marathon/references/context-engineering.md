# Context engineering

marathon keeps the project's written context small and accurate. Notes about finished work,
decided questions, or documentation the code has outgrown take context-window space away from the
task, and stale notes start to contradict the code.

Two rules follow:

1. The repository is the only source of truth. An issue tracker and the chat history are not.
2. The built work wins: the code, or on a context project the written deliverable. `context/`
   holds only what the built work and its documentation can't express yet. Once they can, the
   note is deleted.

## What `context/` holds

The directory is flat, with no subdirectories, except `context/goals/`, which holds the records
of the goals homed in this repository (`mechanics/goal-record.md`):

- `context/README.md`: stable orientation. A short statement of the vision, and a broad, shallow,
  unordered map of the capabilities the project will need.
- `context/*.md`: notes. A note holds intent the built work doesn't express yet, such as a
  decision with its reasoning and rejected alternatives, or an idea with its open questions.
- `context/roadmap.toml`: the roadmap manifest, at a standalone project or the coordinator
  (`references/manifest.md`).
- `context/goals/<goal>.md`: a goal record, the session state of one active goal.

The source, the optional `docs/` directory, and any files an enabled extension owns sit outside
`context/` (`references/extensions.md`).

## Writing a note

- **Look for an existing home first.** If the fact already has exactly one home outside
  `context/`, in the built work or its documentation, point to it instead of restating it.
- **State what exists, or what is planned.** Describe what exists in the present tense, and mark
  planned work as planned. A note has no status line, history, dates, or other detail that only
  holds for a while. The roadmap and the goal records track status: what remains, what a task
  decided, and what validation proved. Every document outside that tracking follows this rule,
  including `docs/` and READMEs. CHANGELOGs are exempt.
- **Start short.** A note begins as a sentence or two and gains detail only when its work is
  close.
- **Name the assumptions** the note depends on, one line each ("assumes the loader keeps the
  configuration immutable after start"). When a build proves an assumption false, revisit the
  notes that named it, and record the falsified assumption in the goal record's Decisions.
- **Treat a new design as provisional.** A design a session just produced stays provisional until
  a later session's caller, build, or experiment uses it and it holds. The goal record's Decisions
  records when that happens. The same applies to writing a skill from the design.

## Project documentation

`docs/` is the project's guide for people who use or contribute to the repository. It is
optional, and it belongs to the built work, not to `context/`. A page that the code has made
wrong is a defect. The change that made it wrong fixes it, or a later task does, and `retro`
flags it. The first task that documents the repository creates `docs/`. The reasoning
behind a built capability belongs in `docs/`, the package documentation, or the README, written by
the task that builds the capability. Knowledge that applies beyond the repository isn't project
documentation, so it stays a note.

## Tending the notes

The editor profile tends the notes during BUILD (`behavior/delegation.md`), and names each
operation it performs:

- **Integrated**: a note deleted because the built work or its documentation now covers it. A
  note that restates a built interface goes stale, because the interface's documentation is its
  home. When a note's reasoning is still useful after the build, move that reasoning into the
  owning repository's documentation first. Name what now covers the note.
- **Culled**: a note deleted because the work replaced it, abandoned it, or contradicted it.
- **Retained**: a note deliberately kept, with the reason, such as a home that doesn't exist yet.
- **Add or sharpen**: a note written or changed in place.
- **Cross-repo**: an edit the session made in another repository.

An enabled extension may add operations of its own, defined in its skill.

The session lists these operations in the session brief's Decided without you, so the architect
sees them at the brief before anything merges. `retro` proposes them for the architect to tick.
