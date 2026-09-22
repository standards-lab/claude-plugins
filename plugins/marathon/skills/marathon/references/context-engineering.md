# Context engineering

Besides moving the work forward, marathon keeps the project's written context small and accurate.
A context window filled with notes about finished work, settled questions, or documentation the
code has outgrown leaves less room for the task, and stale notes start to contradict the code.

Two rules follow:

1. The repository is the only source of truth — not an issue tracker, not the chat history.
2. The built work wins — the code, or on a context project the authored deliverable. `context/`
   holds only what the built work and its own documentation can't express yet. Once they do, the
   note goes.

## What `context/` holds

The directory is flat, with no subdirectories:

- `context/README.md` — stable orientation: a short vision statement and a broad, shallow,
  unordered map of the capabilities the project will need.
- `context/*.md` — notes: intent the built work doesn't express yet, such as a decision with its
  reasoning and rejected alternatives, or an idea and its open questions.
- `context/reset.md` — the session record and the pointer to the next step
  (`mechanics/reset-file.md`).

Outside `context/` sit the source, the optional `docs/` directory, and whatever an enabled
extension owns (`references/extensions.md`).

## Writing a note

- **Check for a home first.** If the fact already has exactly one home outside `context/`, built
  work or settled documentation, write a pointer to it, never a restatement.
- **State settledness in the opening lines**: settled, and what settled it (a decision, a build,
  an experiment's result), or still open, and what would settle it. When you can't tell, it's
  open. A shape a session just designed stays provisional until a later session's caller,
  build, or experiment exercises it and it holds; the same applies to authoring a skill from it.
- **Start short.** A note begins as a sentence or two and gains detail only when its work is
  close.
- **Name the assumptions** it rests on, a line each ("assumes the loader keeps config immutable
  after start"). When a build falsifies one, the notes that named it are the ones to revisit, and
  the Disposition records the falsification.
- **State current truth, never history.** No amendment notes, dated revision logs, or "replacing
  the earlier approach". What a session did and why belongs in the reset file's Disposition.

## Project documentation

`docs/` is the project's own guide for a person using or contributing to the repository. It is
optional, part of the built work rather than of `context/`, and never decays: a page the code has
moved out from under is a defect, fixed in the change that moved the code or in the next
documentation step, and `review` flags it. A documentation step of a `start` session establishes
it. The reasoning behind a built capability belongs there, or in the package documentation or
README, written by the step that builds the capability. What generalizes past the repository is
not project documentation; it stays a note.

## Tending the notes

The Disposition records each operation under its name:

- **Integrated** — a note deleted because the built work or its documentation now expresses it. A
  note that restates a built interface decays, since the interface's documentation is its home. A
  note whose reasoning outlives the build has that reasoning moved into the owning repository's
  documentation first. Point to what now expresses it.
- **Culled** — a note deleted because it is superseded, abandoned, or contradicted by the work.
- **Retained** — a note deliberately kept, with the reason, such as a home that doesn't exist yet.
- **Add or sharpen** — a note written or changed in place, including one that now states itself
  settled.
- **Cross-repo** — an edit the session made in another repository.

An enabled extension may add entries of its own; its skill defines them.

These operations change what the record says about the project, so show the architect what you
propose and get agreement first: in plan mode for `init`, a fresh `start`, `plan`, and `review`,
and as a quick confirmation during `reset` and `close`.
