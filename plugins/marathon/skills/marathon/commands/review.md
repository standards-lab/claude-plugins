# marathon review

Check the health of the `context/` notes when you need to — not every session, but when they feel like
they've drifted from the code, or before starting something big. `review` keeps the repo honest as the
source of truth. It is the backward-looking counterpart to `plan`: `review` asks whether the notes have
fallen behind what the project now expresses, where `plan` works out what comes next.

`review` runs the session pipeline (`mechanics/pipeline.md`), recorded under `Session: review`.
Reviewing means deciding what to delete and what to keep, which matters, so the whole proposal is
settled with the architect before anything changes.

## Settle: look for drift

Read the `context/` notes against the code — or, on a context project, against the deliverable they
document — and look for:

- **Tending candidates** — notes due for the operations of `references/context-engineering.md`,
  under that reference's rules: notes the built work or its documentation now expresses, concepts
  to cull, concepts that have quietly proven out and should state themselves settled, and a
  repository still on the earlier `design/` and `concepts/` layout, to migrate.
- **Orientation** — whether the vision and capability map in `context/README.md` still match where the
  project actually is.
- **Docs drift** — if the project has a `docs/` directory, pages whose code has moved on
  underneath them. Flag them here; the rewriting itself is a documentation step of a `start`
  session, not part of review.
- **Coordinator conventions** — when the project belongs to a workspace, the conventions its
  coordinator keeps for member repositories (naming, authoring, awareness rules), checked against
  this repository's prose and structure. The member repo never cites those conventions; the review
  consults them at session time. See `references/workspace-coordination.md`.

Show the architect what you propose to change and settle it together.

Branch slug: the review's topic.

## Execute: apply what's agreed

Apply only what's agreed: delete the integrated notes (recording that you did), cull the dead
concepts, mark the proven ones settled, and fix the orientation. The edits run in stages under
`references/staged-execution.md`.

End the pass with fewer notes, each one true. Don't invent a home for a note that doesn't have one,
and don't add detail an upcoming step doesn't need.

## Conclude

Close with `close` — it records the dispositions in the reset file under `Session: review` and
publishes the branch.
