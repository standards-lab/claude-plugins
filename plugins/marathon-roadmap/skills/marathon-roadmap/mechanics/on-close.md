# on-close

Fires in `close` only: after the reset file is written, before the closeout commit.

1. Delete the finished task from the manifest. The session record keeps the disposition; the
   roadmap holds only what remains.
2. Check the goals the deletion empties: delete any goal whose criteria now hold, and its
   emptied ancestors in turn.
3. Advance `next`: drop the finished task from its wave, and drop a wave once it is empty. If
   the list is empty, settle what comes next with the architect — `close` agrees the next step in
   any case; this is where it lands.
4. Keep `active` current: an `experiment` session that set up an experiment adds the task it
   serves, and the `plan` session that takes in a closed experiment removes it.
5. Commit the edit where the manifest lives. Same repository: stage it so it rides the closeout
   commit. Coordinator-held, from a member project's session: its own commit at the
   coordinator, recorded under **Cross-repo** in the disposition.
