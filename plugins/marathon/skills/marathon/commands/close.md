# marathon close

Close out a session whose work is finished and validated. It's the same idea as `reset`, but for
completed work: tidy the notes, record what happened, and publish the work. `close` is the
pipeline's CONCLUDE stage for a finished step (`mechanics/pipeline.md`).

`close` assumes the work is validated. On a **code** project's `start`, that means validation
(`references/staged-execution.md`) has passed: the whole-module build, the full test run, and the
run-and-verify check. On a **context** project the change was authored directly, so validation is
that it reads coherently and stays consistent. A `plan` session has nothing to build. For an
`experiment`, decide what, if anything, the spike earned. A `review` session's cleanup is itself
the deliverable. A `docs` session validates its pages against the current code — where they
disagree, the page is what's wrong. Don't close on a failure; fix it first.

## 1. Tend the context

First establish the scope of written context the step touched, then bring it in line with what
now exists, with the tending operations of `references/context-engineering.md` — promote, decay,
cull, under that reference's rules. This pass absorbs the edits noted at SETTLE: on a code
project they waited here for the validated stages, so the notes record what the work proved, not
what the plan intended.

The scope follows from where the project sits:

- **Standalone** — the project's own `context/`.
- **In a workspace** — each touched repo's `context/`; the coordinator's notes that describe the
  changed capability; the docs landing zone pages the change moved out from under (flag them —
  the rewriting itself is a `docs` pass); and claims about the changed behavior in other member
  repos' context. A stale claim found in any of them is a defect this pass fixes, recorded under
  **Cross-repo** in the Disposition.

These changes matter, so show the architect and confirm before applying them.

## 2. Agree on the next step

Talk through what comes next — don't set the direction on your own. What the session turned up
often changes the priority, and the conversation is usually where the next step comes from. Settle
together on the single concrete next step. That becomes the Next-focus in step 3.

## 3. Record what happened

Rewrite the reset file with `Status: closeout`, filling every field of the schema in
`mechanics/reset-file.md`. In a workspace that is the coordinator's `context/reset.md` — the
workspace's only reset, committed in the coordinator's repository; a standalone project rewrites
its own. The judgment fields are the Disposition — what you integrated, promoted, culled, and
retained — and the Next-focus, set to the step you agreed on in step 2: the handoff the next
session reads.

## 4. Commit and publish

Stage everything and commit — a step that spanned member repos commits in each touched repo, and
in a workspace the coordinator's reset rewrite is its own commit in the coordinator's repository.
Then publish each branch with the remote platform's command from that repo's
`.claude/marathon.toml` — `gh pr create`, `glab mr create`, or the project's equivalent — using
the change description from the reset file. If a project declared no remote, stop after its
commit.
