# on-close

Fires after the closeout record is written, before the commit. In a wave, a session that isn't the
last to close records steps 1–3 in its own Disposition instead; the session that folds the wave
applies every record's edits.

1. Delete the finished task.
2. Delete any goal the deletion leaves with its criteria holding, and its emptied ancestors.
3. Drop the task from its `next` entry, and the entry once it is empty. If `next` is empty, settle
   what comes next with the architect.
4. Commit where the manifest lives: in the closeout commit, or as its own coordinator commit from
   a member's session, recorded under **Cross-repo**.
