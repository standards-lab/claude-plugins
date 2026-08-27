# on-reset

Fires before the reset file is written — a standalone project's own `context/reset.md`, or in a
workspace the coordinator's — in `reset` and in `close`.

- Next-focus names the roadmap task the next session advances by its dotted path, alongside
  whatever else the record says about picking it up.
- The Disposition carries any manifest edits the session made — a task added, a stale claim
  corrected, detail added to the task in front — under its usual entries; edits to a
  coordinator-held manifest from a member project's session go under **Cross-repo**.
- On a handoff (`Status: handoff`), the manifest does not advance: the task in flight stays
  until a closeout finishes it.
