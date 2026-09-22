# marathon reset

Hand off mid-session, when the context is filling but the step isn't done, so a fresh context
resumes the same step on the same branch. Capture enough that it can pick up cold.

1. **Tidy the notes you touched** under `references/context-engineering.md`, with a quick
   confirmation from the architect.
2. **Write the record** (`mechanics/reset-file.md`) with `Status: handoff`: the coordinator's in a
   workspace, or the session's own in a wave. Next-focus is the in-progress state and the exact
   next move: the file being edited, the pending decision, the approved stage list, the stage and
   checkpoint position, and after a re-plan the revised list.
3. **Keep the work.** Finished stages are already committed. If the context fills mid-stage, make a
   WIP commit and note that the stage's check hasn't passed; the resuming session finishes it
   first. A checkpoint reported but unconfirmed is noted too, and reported again first. Leave the
   branch open and unpublished.
