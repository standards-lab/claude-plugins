# marathon status

Print the status digest: a table with one row per active goal, with the goals that need the
architect first (`references/briefs/status-digest.md`). Each active spike sub-goal is a row of
its own; an experiment goal is never active, so it has none. `status <goal>` prints that goal's
tasks instead. `status` changes nothing and fires no hook.

Read first, before acting:

- [`../references/manifest.md`](../references/manifest.md)
- [`../mechanics/goal-record.md`](../mechanics/goal-record.md)
- [`../references/briefs.md`](../references/briefs.md)
- [`../references/briefs/status-digest.md`](../references/briefs/status-digest.md)
- [`../references/briefs/goal-status.md`](../references/briefs/goal-status.md)

1. Read `active` from the index (`references/manifest.md`): a standalone project's own, or the
   coordinator's.
2. For each active goal, read its goal record from its root's working tree
   (`mechanics/goal-record.md`). Count the checked tasks against the listed tasks, and take the
   State line. A goal with no record shows `0/n` and the state `no record`, counting its tasks
   in its file's `order`. A plan file under `.claude/plans/` that names the goal shows the State
   `planning`.
3. Mark with `!` each goal that is `planning` or `brief ready`, or whose Handoff holds an
   unanswered escalation or an unmerged pull request, and sort those first.
4. Print the digest, with each goal's locked repositories in its row's last column.

With a goal named (`references/briefs/goal-status.md`):

1. Read the goal's file (`references/manifest.md`) and its goal record, if it has one.
2. List its tasks in `order`, each done when the record checks it, or all open with no record.
3. Add one line per sub-goal, meaning each file whose name extends the goal's by one segment,
   with its done count, counted as the digest's step 2 counts it.
