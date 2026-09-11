# on-close

Fires in `close` only: after the reset file is written, before the closeout commit.

1. **Same repository** — a page landed under a standalone project's `architecture/`: stage it so
   it rides the closeout commit, like any other context edit the session made.
2. **Another repository** — a concept landed in the workspace's architecture repository: commit
   it there, on a branch under the step's slug, creating the branch if the session had not
   already touched that repository. It publishes with the session's other branches, through that
   repository's own `[remote]` publish command, and is recorded under **Cross-repo**.
3. A session that landed nothing does nothing here.
