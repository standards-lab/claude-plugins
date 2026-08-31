# marathon reset

Hand off in the middle of a session. Use this when the context window is filling up but the work
isn't finished, and you want a fresh context to keep going on the same step and the same branch.
`reset` is the pipeline's CONCLUDE stage taken mid-work (`mechanics/pipeline.md`).

The written record is what the next session resumes from, so capture enough that a new context can
pick up cold.

## 1. Tidy the notes you touched

Bring the `context/` notes you changed this session in line with where things actually stand, so
the handoff doesn't carry stale notes forward — the tending operations of
`references/context-engineering.md` (promote, decay, cull), applied under that reference's rules
to the notes this session touched.

These changes matter, so show the architect what you plan to change and get a quick confirmation
before you do it.

## 2. Write the reset file

Rewrite the reset file with `Status: handoff`, filling every field of the schema in
`mechanics/reset-file.md` — the coordinator's `context/reset.md` in a workspace, the project's own
when standalone. The Disposition records what step 1 integrated, promoted, culled, or retained.
The judgment field is **Next-focus** — the in-progress state and the exact next move. Be specific:
name the file you're editing, the decision that's pending, the next thing to do. On a code
project's `start`, Next-focus carries the approved stage list, which stage the pointer is on, and
whether that stage is implemented-unreviewed or reviewed; after a re-plan it carries the revised
list. This is what a cold context resumes from.

## 3. Keep the work

On a code project's `start`, the finished stages already sit on the branch as their own commits; a
WIP commit is only for an in-progress stage. Otherwise, optionally make a WIP commit so
in-progress work isn't lost. Leave the branch open and don't publish it — the work isn't finished.
A later session will see `Status: handoff`, check out the branch, and keep going.
