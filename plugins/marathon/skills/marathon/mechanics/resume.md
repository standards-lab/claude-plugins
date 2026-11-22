# 3R · RESUME

LOCATE routes a goal whose record is in `building` or `handoff` here, after START
(`mechanics/pipeline.md`, 1 · LOCATE). RESUME takes the task up where the goal record left it,
then returns to the pipeline at BUILD, or past it for a release or a sync.

1. Check out the task's branch in each touched repository that hasn't merged it, any
   `<slug>-fix` branch a release left unmerged (`mechanics/releasing.md`), or, resuming a sync,
   the `sync-<goal>` branch in each repository that hasn't merged it.
2. Fire `on-build`, unless the session resumes only a sync or a release. A release fires it
   whenever it opens or resumes a `<slug>-fix` branch (`mechanics/releasing.md`).
3. Read the brief, Progress, and Handoff from the goal record. Finish any WIP slice first, then
   continue BUILD from the recorded position. Three next moves resume past BUILD instead:
   - "tag <names>" resumes the release at SHIP, from each repository's recorded position
     (`mechanics/releasing.md`, "Resuming a release").
   - "merge `sync-<goal>`, then delete the record" resumes the sync at its remaining steps: merge
     each `sync-<goal>` branch not yet merged, then delete the record, which releases the goal's
     root (`mechanics/goals.md`, "Sync", steps 4 to 6).
   - "sync <repo>, then delete the record" resumes the sync at step 3 for <repo> alone: apply
     the waiting edits on a `sync-<goal>` branch there, merge it, then delete the record
     (`mechanics/goals.md`, "Sync", steps 3 to 6). While <repo> is still held, name its holder as
     LOCATE step 5 does, and stop.
