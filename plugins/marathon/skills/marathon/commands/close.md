# marathon close

Close out a session whose work is finished and validated. It's the same idea as `reset`, but for
completed work: review the branch, tidy the notes, record what happened, and publish the work.
`close` is the pipeline's CONCLUDE stage for a finished step (`mechanics/pipeline.md`).

`close` assumes the work is validated per `references/staged-execution.md`: every stage
committed, every checkpoint confirmed, and the validation for the project kind confirmed last. A
`plan` session has nothing to build. For an `experiment`, validation is the spike's answer to its
question; decide with the architect what, if anything, the spike earned. A `review` session's
cleanup is itself the deliverable. A documentation step validates its pages against the current
code, and where they disagree, the page is what's wrong. Don't close on a failure; fix it first.

## 1. Review the branch

The branch review is the holistic code-quality pass over the whole branch, made once the last
stage has landed rather than stage by stage. It runs on a `start` or an `experiment` whose branch
changed code or a deliverable; a `plan` or `review` session, whose branch changed only
`context/`, skips it.

1. Prepare the report's location. The report is `.claude/report.md`: the project's own when it
   stands alone, the coordinator's in a workspace. It is never committed, so before anything
   writes it, confirm that repository's `.gitignore` lists `.claude/report.md`. If it doesn't,
   add the line now; the line rides the closeout commit.
2. Engage the reviewer profile (`behavior/delegation.md`), stating the model and why, or do the
   review directly when briefing the reviewer would cost more. The brief carries the approved
   stage list, the branch in each touched repository, the confirmed checkpoints, the report's
   path, and any architecture or documented practice the project declares, cited by path.
   Without a declared one, the review checks the result against the ecosystem's idiomatic
   practice. The report's format is in the reviewer profile, the plugin's `agents/reviewer.md`.
3. Read the report firsthand, correct anything it gets wrong, and tell the architect it's ready
   to read.
4. Settle each finding with the architect: fix it in a new commit on the branch and re-run the
   affected checks, or record it for a later step. A finding that reaches past the step is a
   re-plan (`references/staged-execution.md`).

## 2. Tend the context

First establish the scope of written context the step touched, then bring it in line with what
now exists, with the tending operations of `references/context-engineering.md` — integrate, cull,
retain, under that reference's rules. This pass absorbs the edits noted at SETTLE: on a code
project they waited here for the validated stages, so the notes record what the work proved, not
what the plan intended.

The scope follows from where the project sits:

- **Standalone** — the project's own `context/`.
- **In a workspace** — each touched repo's `context/`; the coordinator's notes that describe the
  changed capability; the repository's own `docs/` pages the change moved out from under, fixed
  in this change or flagged for a documentation step; and claims about the changed behavior in
  other member repos' context. A stale claim found in any of them is a defect this pass fixes,
  recorded under **Cross-repo** in the Disposition.

These changes matter, so show the architect and confirm before applying them.

## 3. Agree on the next step

Talk through what comes next — don't set the direction on your own. What the session turned up
often changes the priority, and the conversation is usually where the next step comes from. Settle
together on the single concrete next step. That becomes the Next-focus in step 4.

## 4. Record what happened

Rewrite the reset file with `Status: closeout`, filling every field of the schema in
`mechanics/reset-file.md`. In a workspace that is the coordinator's `context/reset.md` — the
workspace's only reset, committed in the coordinator's repository; a standalone project rewrites
its own. The judgment fields are the Disposition — what you integrated, culled, and retained,
and under **Validated**, each checkpoint with its evidence — and the Next-focus, set
to the step you agreed on in step 3: the handoff the next session reads.

## 5. Commit and publish

Delete `.claude/report.md`; the architect has read it, and it holds nothing the repository
keeps. Then stage everything and commit — a step that spanned member repos commits in each
touched repo, and in a workspace the coordinator's reset rewrite is its own commit in the
coordinator's repository.
Then publish each branch with the remote platform's command from that repo's
`.claude/marathon.toml` — `gh pr create`, `glab mr create`, or the project's equivalent — using
the change description from the reset file. If a project declared no remote, stop after its
commit.
