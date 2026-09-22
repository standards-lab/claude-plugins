# The reset file

`context/reset.md` is the session record and the pointer to the next step: written at the end of
one session, read at the start of the next. It is a purely ephemeral artifact — its sole purpose
is contextual bootstrapping between sessions, and durable detail belongs in the context layers,
never here. It is a single file, rewritten at each `reset` or `close`; git is the archive, and
keeps every older version.

## Where it lives

- A **standalone project** keeps its own `context/reset.md`.
- A **workspace** maintains exactly one reset file, at the coordinator. Member projects carry
  none: a session anywhere in the workspace reads the coordinator's reset at LOCATE and writes it
  at CONCLUDE — a commit in the coordinator's repository. The record names the member projects it
  concerns in its **Project** line. A wave in flight adds per-session records beside it (below).

## Waves

A closeout's Next-focus may name a **wave**: several steps the architect judges safe to run at the
same time, with no dependency between them and no repository they share. The architect runs each
step in its own session. While the wave is in flight, each session keeps its own record,
`context/reset/<slug>.md` beside the reset file, in the same schema, and never writes
`context/reset.md`. Changes it would make to other shared files at the same location, such as an
extension's artifact, go in its record's Disposition instead of being applied.

The session that closes last, once every other step of the wave has a closeout record on the main
branch, folds the wave in one commit: it applies the recorded changes, rewrites `context/reset.md`
with the next Next-focus, and deletes the wave's records.

## Schema

```markdown
# reset · wire-config-loader

- **Status:** closeout            # handoff | closeout
- **Session:** start              # init | plan | start | experiment | review
- **Project:** core-lib           # workspace reset only: the member repo(s) the step touched
- **Branch:** wire-config-loader

## Disposition
- **Integrated:** deleted context/config-loading.md — the loader and its package documentation now express it (config/loader).
- **Add or sharpen:** context/config-validation.md now states itself settled (the validation rules held against the loader).
- **Culled:** deleted context/env-override.md — the loader implementation went another way.
- **Retained:** context/secret-sourcing.md — still unbuilt.
- **Validated:** checkpoint 1, the loader reads a sample file (`go run ./cmd/example`); checkpoint 2, `go test ./...` and the run-and-verify check.

## Next-focus
Add secret sourcing on top of the validated loader. Start here next session.
```

A standalone project's reset omits the Project line. In a workspace reset, Next-focus also names
the member project the next session continues in, so a session entered anywhere in the workspace
finds its way. A step that spanned several member repos lists them all on the Project line; its
branches share the step's slug, so the Branch line stays one value.

## Status semantics

- **closeout** — the recorded session finished and published. The next session starts a fresh
  step from the Next-focus, on a new branch.
- **handoff** — the recorded session stopped mid-work with its branch open. The Session line
  names the command that resumes it; the Branch line (and, in a workspace, the Project line)
  names where; Next-focus records the in-progress state and the exact next move, so a fresh
  context can resume without working it out again. A working session records its stage list and
  position there, and its checkpoint position, for example
  `Stages: 5/9 · checkpoint 1 of 3 confirmed · stage 5 committed · list: …`. A step
  spanning member repos records each touched repo's branch state there too; whatever the shape of
  the interrupted work, Next-focus carries enough bootstrap state to resume it cold.

The Disposition speaks the tending ledger of `references/context-engineering.md`: Integrated,
Culled, Retained, and **Cross-repo**, for an edit the session made in another repository, such as
a member session's edit to an artifact held at the coordinator. **Add or sharpen** records a note
written or changed in place, and an enabled extension may add entries of its own.
A closeout adds **Validated**: each checkpoint the architect confirmed, with its evidence.

Older reset files may carry a `Session type:` line naming `development`/`context`/`experiment`;
read it as the equivalent working session.
