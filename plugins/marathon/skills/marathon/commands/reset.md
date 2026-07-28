# marathon reset

Hand off in the middle of a session. Use this when the context window is filling up but the work isn't
finished, and you want a fresh context to keep going on the same piece of work and the same branch.

The idea is that the written record, not the conversation, carries the work forward. Capture enough that
a new context can pick up cold.

## 1. Tidy the notes you touched

Bring the `context/` notes you changed this session in line with where things actually stand, so the
handoff doesn't carry stale notes forward:

- **Decay** — if the code (or the authored deliverable) now expresses a `design/` note, delete that note.
- **Cull** — drop a concept this session's work has invalidated or replaced.
- **Promote** — move a concept from `concepts/` into `design/` if it proved out, and say why.

These changes matter, so show the developer what you plan to change and get a quick confirmation before
you do it.

## 2. Write context/reset.md

Rewrite the reset file with `Status: handoff`. Record:

- the branch and which session this is (`plan`/`start`/`experiment`),
- what you integrated, promoted, culled, or retained in step 1, and
- **Next-focus** — the in-progress state and the exact next move. Be specific: name the file you're
  editing, the decision that's pending, the next thing to do. This is what a cold context resumes from.

Write it per the voice standard in `references/writing-voice.md`.

## 3. Keep the work

Optionally make a WIP commit on the branch so in-progress work isn't lost. Leave the branch open and
don't publish it — the work isn't finished. A later `start` will see `Status: handoff`, check out the
branch, and keep going.
