# Context engineering

As well as moving the work forward, marathon keeps the project's written context small and
accurate. Notes about finished work, decided questions, or documentation the code has outgrown
take context-window space away from the task, and stale notes start to contradict the code.

Two rules follow:

1. The repository is the only source of truth. An issue tracker and the chat history are not.
2. The built work wins: the code, or on a context project the written deliverable. `context/`
   holds only what the built work and its documentation can't express yet. Once they can, the
   note is deleted.

## What `context/` holds

The directory is flat, with no subdirectories:

- `context/README.md`: stable orientation. A short statement of the vision, and a broad, shallow,
  unordered map of the capabilities the project will need.
- `context/*.md`: notes. A note holds intent the built work doesn't express yet, such as a
  decision with its reasoning and rejected alternatives, or an idea with its open questions.
- `context/reset.md`: the session record and the pointer to the next step
  (`mechanics/reset-file.md`).

The source, the optional `docs/` directory, and any files an enabled extension owns sit outside
`context/` (`references/extensions.md`).

## Writing a note

- **Look for an existing home first.** If the fact already has exactly one home outside
  `context/`, in the built work or its documentation, point to it instead of restating it.
- **State what exists, or what is planned.** Describe what exists in the present tense, and mark
  planned work as planned. A note has no status line, history, dates, or other detail that only
  holds for a while. The roadmap and the reset file track status: what remains, what a session
  decided, and what validation proved. Every document outside that tracking follows this rule,
  including `docs/` and READMEs. CHANGELOGs are exempt.
- **Start short.** A note begins as a sentence or two and gains detail only when its work is
  close.
- **Name the assumptions** the note depends on, one line each ("assumes the loader keeps the
  configuration immutable after start"). When a build proves an assumption false, revisit the
  notes that named it, and record the falsified assumption in the Disposition.
- **Treat a new design as provisional.** A design a session just produced stays provisional until
  a later session's caller, build, or experiment uses it and it holds. The Disposition records
  when that happens. The same applies to writing a skill from the design.

## Project documentation

`docs/` is the project's guide for people who use or contribute to the repository. It is
optional, and it belongs to the built work, not to `context/`. A page that the code has made
wrong is a defect. The change that made it wrong fixes it, or the next documentation step does,
and `review` flags it. A documentation step of a `start` session creates `docs/`. The reasoning
behind a built capability belongs in `docs/`, the package documentation, or the README, written by
the step that builds the capability. Knowledge that applies beyond the repository isn't project
documentation, so it stays a note.

## Tending the notes

The Disposition records each operation on the notes under its name:

- **Integrated**: a note deleted because the built work or its documentation now covers it. A
  note that restates a built interface goes stale, because the interface's documentation is its
  home. When a note's reasoning is still useful after the build, move that reasoning into the
  owning repository's documentation first. Name what now covers the note.
- **Culled**: a note deleted because the work replaced it, abandoned it, or contradicted it.
- **Retained**: a note deliberately kept, with the reason, such as a home that doesn't exist yet.
- **Add or sharpen**: a note written or changed in place.
- **Cross-repo**: an edit the session made in another repository.

An enabled extension may add operations of its own, defined in its skill.

These operations change what the record says about the project, so show the architect what you
propose and get agreement first. Agree them in plan mode for `init`, a new `start`, `plan`, and
`review`, and with a quick confirmation during `reset` and `close`.
