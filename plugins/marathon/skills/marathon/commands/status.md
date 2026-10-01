# marathon status

Print the status digest: one line per active goal, with the goals that need the architect first
(`references/briefs.md`, "Status digest"). `status` changes nothing and fires no hook.

1. Read `active` from the manifest (`references/manifest.md`): a standalone project's own, or the
   coordinator's.
2. For each active goal, read its goal record from its home's working tree
   (`mechanics/goal-record.md`). Count the checked tasks against the listed tasks, and take the
   State line. A goal with no record shows `0/n  no record`, counting its tasks in the manifest.
3. Mark with `!` each goal whose State is `planning` or `brief ready`, or whose Handoff holds an
   unanswered escalation, and sort those first.
4. Print the digest, with each goal's locked repositories at the end of its line.
