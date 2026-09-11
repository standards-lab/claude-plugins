# on-reset

Fires before the reset file is written — a standalone project's own `context/reset.md`, or in a
workspace the coordinator's — in `reset` and in `close`.

1. On a closeout (`Status: closeout`), take from the session's tending pass any design note the
   work showed to have generalized past this repository, and land it by the shape the layer takes:
   - **Standalone** — the knowledge is written as a page under `architecture/` in this
     repository, in this session. The design note then decays under marathon's rule, its removal
     pointing at the page.
   - **In a workspace** — the note lands as a concept in the architecture repository's
     `context/concepts/`, and that repository authors the page in a session of its own. The
     member's design note does not decay yet: it decays when the page exists, not when the
     concept lands.

2. Record the landing in the Disposition, in marathon's ledger vocabulary: **Promoted** for the
   note that moved, **Integrated** for a design note the new page let go, and **Cross-repo** for
   anything a member project's session wrote in the architecture repository.

3. On a handoff (`Status: handoff`), nothing is promoted. A candidate the session turned up is
   in-flight direction: it goes in the Disposition as the session's own note, and the promotion
   waits for the closeout that finishes the step.
