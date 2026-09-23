# on-close

This hook fires after the closeout record is written and before the commit. During a wave, a
session that doesn't fold the wave records steps 1 to 3 in its lane's record instead of applying
them, and the session that folds the wave applies every lane's recorded edits.

1. Delete the finished task.
2. Delete any goal whose criteria now hold because of the deletion, and any ancestor goal left
   empty.
3. Remove the task from `next`. Then remove its lane if the lane is now empty, and its wave if the
   wave is now empty. If `next` is empty, agree with the architect what comes next.
4. Commit the manifest where it lives: in the closeout commit, or, from a member's session, as a
   separate commit at the coordinator, recorded under **Cross-repo**.
