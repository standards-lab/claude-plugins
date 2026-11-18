# Releasing

SHIP releases when the task brief has a Release line (`mechanics/pipeline.md`, 6 · SHIP), and
RESUME comes here when the next move is "tag <names>" (`mechanics/resume.md`). The line lists the
tags the task releases (`references/briefs.md`, "Task brief"). Accepting the session brief
authorizes them, so SHIP tags them without escalating again. SHIP takes each repository in turn,
lowest layer first, and finishes its release before the next repository's merge:

1. Merge the branch with `[remote] merge`, switch to the default branch, pull, and delete the
   local task branch.
2. When `[remote] ci` is set, run it on the merge commit (`mechanics/configuration.md`).
   Without it, tagging follows the merge directly.
3. Tag each of the repository's Release tags, the base artifact before its sub-modules:
   1. Check that the tag isn't on the remote, unless it is this release's failed tag (below),
      and that its version matches the artifact's version as the repository records it, in its
      manifest or its CHANGELOG's top heading.
   2. When `[project] gate` is set, run it on the merge commit with the tag as its one argument
      (`mechanics/configuration.md`). Exit 0 lets the tag be pushed; nonzero holds it (below).
      A resumed release runs it again before each tag it hasn't pushed.
   3. Create the tag annotated "<artifact> <version>" on the merge commit, and push that tag
      alone.
   4. When the repository has a release workflow, confirm the run the tag started succeeds.

No tag is pushed while the default branch is red or its gate fails. The session fixes forward
until the planned version releases:

- **A red default branch** after the merge, a failed gate, or a version that doesn't match, is
  fixed on a `<slug>-fix` branch from the default branch, published and merged like the task's
  branch, with `[remote] ci` and the gate run again on its merge commit. Then tag. Opening or
  resuming a fix branch fires `on-build` first.
- **A failed release**, a release workflow that fails after the tag is pushed, is fixed the
  same way. Then delete the tag on the remote and locally, and create and push it again at the
  same version on the fix's merge commit.
- **A released tag** is never re-cut. A tag whose release failed is this release's tag when the
  Release line names it and this session or an earlier session of the same release pushed it, as
  Handoff records: it is deleted and pushed again at the same version, as above. Only a tag on the
  remote that no session of this release pushed stops the release with an escalation.

The release always ends at the planned version. It stops for the architect only for a decision
the brief doesn't cover, as an escalation (`references/build.md`). When the context fills, set
State to `handoff` with the next move "tag <names>", naming the tags not yet released and each
Release repository's position: merged with ci not yet green, merged and green, on `<slug>-fix`
at its publish, merge, or ci step, its gate failed, its tag pushed and its release failed, or
released. Commit it where the root stands: on the root's task branch while that branch hasn't
merged, pushed to that branch, whose pull request is already published, so the record lands with
the merge; or on the root's default branch as bookkeeping once it has. Then stop; the next
`start` resumes the release (`mechanics/resume.md`).

## Resuming a release

The next move "tag <names>" resumes the release at SHIP, one repository at a time, from the
position Handoff records for it: at step 1 for a repository whose task branch hasn't merged; at
step 2 for one merged with ci not yet green; at its `<slug>-fix` branch's publish, merge, or
`[remote] ci` step, then the tag step, for one fixing forward; at the fix for one whose gate
failed; at the failed-release fix for one whose tag was pushed and whose release failed; and at
the tag step only for one merged and green.

A resumed release ends at the goal record. Once the last named tag releases, set State to `idle`,
Task and Branch to none, drop Handoff, and check the task if SHIP's record update hadn't landed,
as a bookkeeping commit on the root's default branch. On the goal's last task, sync follows
(`mechanics/goals.md`, "Sync").
