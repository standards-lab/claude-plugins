# marathon reset

Hand off partway through a task, so the next `start` on the goal resumes it without the
conversation or the architect. The orchestrating session runs `reset` on its own when its context
fills before BRIEF; the architect can also ask for it.

1. **Keep the work.** Finished slices and review fixes are already committed. If a slice is partway
   done, make a WIP commit on the task's branch and note that its check hasn't passed.
2. **Write the handoff** into the goal record (`mechanics/goal-record.md`): set State to
   `handoff`, bring Progress up to date, and write Handoff with the exact next move, the WIP
   slice, and any escalation still waiting on the architect.
3. **Commit the record.** Fire `on-record`, then commit it on the task's branch in the root.
4. **Leave the branches open** and unpublished, and stop. The next `start` on the goal resumes at
   3R · RESUME (`mechanics/pipeline.md`).

A session that stops during PLAN needs no reset: the plan file keeps the open round. One whose
context fills while SHIP releases hands off the same way, with the next move "tag <names>"
naming the tags not yet released, committed where `mechanics/pipeline.md` ("Releasing") says.
