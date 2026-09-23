# marathon review

An on-demand check of the `context/` notes against the code, or against the deliverable on a
context project. Run it when the notes seem stale or before a large change. `review` runs the
session pipeline (`mechanics/pipeline.md`) and is recorded under `Session: review`.

## Settle: look for drift

- **Notes to tend** (`references/context-engineering.md`): notes the built work or its
  documentation now covers, notes to cull, and notes that describe history or status instead of
  what exists or what is planned.
- **Orientation**: whether `context/README.md` still describes the project as it is.
- **Documentation drift**: `docs/` pages the code has made wrong, flagged for a documentation
  step.
- **Coordinator conventions**: in a workspace, whether this repository follows the conventions
  the coordinator keeps for its members (`references/workspace-coordination.md`).

Agree the whole proposal with the architect before changing anything.

Branch slug: the review's topic.

## Execute

Apply only what was agreed, in stages (`references/staged-execution.md`). The session ends with
fewer notes, each of them true.

## Conclude

End with `close`.
