# The reset file

`context/reset.md` records the last session and names the next step. A session writes it at its
end, and the next session reads it at its start. Each session rewrites the whole file, and git
keeps the earlier versions. The file only gets the next session started. Anything that needs to
last belongs in the notes.

A standalone project keeps its own reset file. A workspace keeps exactly one, at the coordinator.
Every session in the workspace reads it at LOCATE and writes it at CONCLUDE.

## Schema

```markdown
# reset · wire-config-loader

- **Status:** closeout            # handoff | closeout
- **Session:** start              # init | plan | start | experiment | review
- **Project:** core-lib           # workspace only: the member repos the step touched
- **Branch:** wire-config-loader  # one name, shared by every touched repo

## Disposition
- **Integrated:** deleted context/config-loading.md; the loader's package documentation now covers it.
- **Add or sharpen:** context/config-validation.md now names the rules the loader enforces.
- **Culled:** deleted context/env-override.md; the loader took a different approach.
- **Retained:** context/secret-sourcing.md; not built yet.
- **Validated:** checkpoint 1, the loader reads a sample file (`go run ./cmd/example`); checkpoint 2, `go test ./...` and the run-and-verify check.

## Next-focus
Add secret sourcing on top of the validated loader.
```

The Disposition uses the entries defined in `references/context-engineering.md`. A closeout adds
**Validated**. In a workspace, Next-focus also names the member project the next session runs in.

## Status

- **closeout**: the session finished and published. The next session starts the step Next-focus
  names, on a new branch.
- **handoff**: the session stopped partway through, with its branch open. The Session line names
  the command that resumes it. Next-focus holds everything needed to resume without the
  conversation:
  - the stage list and the position in it, for example
    `Stages: 5/9 · checkpoint 1 of 3 confirmed · stage 5 committed · list: …`
  - the state of the branch in each touched repository
  - the exact next move

## Waves

In a workspace, a closeout's Next-focus may name a **wave**: a set of **lanes** the architect
judges safe to run at the same time, each in sessions of its own. A lane is a single step, or a
sequence of steps run in order. Lanes share no member repository, and no file at the coordinator
other than their own records.

While a wave is running, each lane keeps its own record at `context/reset/<lane>.md`, in the same
schema, named after the slug of the lane's first step. The record's Next-focus names the lane's
next step, and after the lane's last step it reads `Lane finished.` No lane writes
`context/reset.md`. A change a lane would make to any other shared file at the coordinator, such
as an extension's artifact, goes in the lane's Disposition instead of being applied.

A wave is folded in one commit once every lane's record reads `Lane finished.` on the main branch.
If the other lanes are already merged, the session that finishes the last lane folds the wave.
Otherwise, the next session whose LOCATE finds every lane finished folds it first. Folding applies
the changes the lanes recorded, rewrites `context/reset.md`, and deletes the lane records.
