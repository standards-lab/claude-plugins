# marathon review

Check the health of the `context/` notes when you need to — not every session, but when they feel like
they've drifted from the code, or before starting something big. `review` keeps the repo honest as the
source of truth. It is the backward-looking counterpart to `plan`: `review` asks whether the notes have
fallen behind what the project now expresses, where `plan` works out what comes next.

Reviewing means deciding what to delete and what to promote, which matters, so run it in plan mode and
agree with the developer before changing anything.

## 1. Look for drift

Read the `context/` notes against the code — or, on a context project, against the deliverable they
document — and look for:

- **Decay** — `design/` notes the project now fully expresses. These should go; the design is describing
  the present instead of the unbuilt.
- **Cull** — concepts the work invalidated or replaced, or that have just gone stale.
- **Promote** — concepts that have quietly proven out and belong in `design/`.
- **Orientation** — whether the vision and capability map in `context/README.md` still match where the
  project actually is.
- **Docs drift** — if the project has a `docs/` tier, pages whose code has moved on underneath them.
  Flag them here; the rewriting itself is deliberate work for a `docs` pass, not part of review.

## 2. Agree, then apply

Show the developer what you propose to change and settle it together. Apply only what's agreed: delete
the decayed notes (recording that you did), cull the dead concepts, promote the proven ones, and fix
the orientation.

Keep the bar high and the footprint small. Don't invent a home for a note that doesn't have one, and
don't add detail a coming step doesn't need. The goal is fewer stale notes, not more structure.

## 3. Land the changes

Because `review` edits the agent's context notes, land them like any other work — on a branch, published
the same way — so the cleanup is traceable. Record what you changed in `context/reset.md`, the same as
any closeout.
