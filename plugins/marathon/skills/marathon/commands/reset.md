# marathon reset

Hand off partway through a step, when the context is filling but the step isn't done, so a new
session can resume the same step on the same branch. Record enough that the new session can
continue without the conversation.

1. **Tidy the notes you touched**, following `references/context-engineering.md`, with a quick
   confirmation from the architect.
2. **Write the record** (`mechanics/reset-file.md`) with `Status: handoff`. In a workspace, write
   the coordinator's record, or the lane's own record during a wave. Next-focus gives the state of
   the work and the exact next move: the file being edited, any pending decision, the approved
   stage list, the stage and checkpoint position, and the revised list after a re-plan.
3. **Keep the work.** Finished stages are already committed. If the context fills partway through
   a stage, make a WIP commit and note that the stage's check hasn't passed. The resuming session
   finishes that stage first. Note a checkpoint that was reported but not confirmed, too; the
   resuming session reports it again first. Leave the branch open and unpublished.
