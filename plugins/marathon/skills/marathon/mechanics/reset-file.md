# The reset file

`context/reset.md` is the session record and the pointer to the next step: written at the end of
one session, read at the start of the next, and rewritten each time; git keeps older versions. It
bootstraps the next session and nothing more; durable detail belongs in the notes.

A standalone project keeps its own. A workspace keeps exactly one, at the coordinator, which every
session in the workspace reads at LOCATE and writes at CONCLUDE.

## Waves

A closeout's Next-focus may name a **wave**: steps the architect judges safe to run at the same
time, with no dependency between them and no repository they share, each in its own session.
While a wave is in flight, each session keeps its own record at `context/reset/<slug>.md`, in the
same schema, and never writes `context/reset.md`. Changes it would make to other shared files
there, such as an extension's artifact, go in its record's Disposition instead of being applied.
The session that closes last, once every other step's closeout record is on the main branch, folds
the wave in one commit: it applies the recorded changes, rewrites `context/reset.md`, and deletes
the wave's records.

## Schema

```markdown
# reset · wire-config-loader

- **Status:** closeout            # handoff | closeout
- **Session:** start              # init | plan | start | experiment | review
- **Project:** core-lib           # workspace only: the member repos the step touched
- **Branch:** wire-config-loader  # one slug, shared across touched repos

## Disposition
- **Integrated:** deleted context/config-loading.md — the loader's package documentation now expresses it.
- **Add or sharpen:** context/config-validation.md now states itself settled.
- **Culled:** deleted context/env-override.md — the loader went another way.
- **Retained:** context/secret-sourcing.md — still unbuilt.
- **Validated:** checkpoint 1, the loader reads a sample file (`go run ./cmd/example`); checkpoint 2, `go test ./...` and the run-and-verify check.

## Next-focus
Add secret sourcing on top of the validated loader.
```

The Disposition uses the ledger of `references/context-engineering.md`, plus **Validated** at a
closeout. In a workspace, Next-focus also names the member project the next session runs in.

## Status

- **closeout** — the session finished and published; the next session starts a fresh step from
  Next-focus on a new branch.
- **handoff** — the session stopped mid-work with its branch open. The Session line names the
  command that resumes it, and Next-focus carries enough to resume cold: the stage list and
  position (`Stages: 5/9 · checkpoint 1 of 3 confirmed · stage 5 committed · list: …`), each
  touched repo's branch state, and the exact next move.
