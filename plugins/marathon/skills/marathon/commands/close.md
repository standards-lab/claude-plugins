# marathon close

Close out a session whose work is finished and validated. It's the same idea as `reset`, but for
completed work: tidy the notes, record what happened, and publish the work.

## 1. Finish the agent's part

On a **code** project's `start`, the developer has applied the production code for this step. Now add
the parts the role boundary gives the agent:

- tests for the new behavior, and
- in-source comments and API docs, plus any prose documentation — written per the voice standard in
  `references/writing-voice.md`.

Then validate — run the tests and the build. Don't close on a failure; fix it (within the agent's part)
or hand it back to the developer.

On a **context** project there is no code to finish — the change was authored directly, so validation is
that it reads coherently and stays consistent. A `plan` session has nothing to build either. For an
`experiment`, decide what, if anything, the spike earned. A `review` session closes the same way as a
context change: the cleanup itself is the deliverable. A `docs` session validates its pages against the
current code — where they disagree, the page is what's wrong. Either way, go on to step 2.

## 2. Tidy the notes

Bring `context/` in line with what now exists:

- **Decay** — delete `design/` notes the code (or the authored deliverable) now expresses, and record
  that you did.
- **Cull** — drop the concepts the work invalidated or replaced.
- **Promote** — move proven concepts from `concepts/` into `design/`, and say why.

These changes matter, so show the developer and confirm before applying them.

## 3. Agree on the next step

Talk through what comes next — don't set the direction on your own. What the session turned up often
changes the priority, and the conversation is usually where the next step comes from. Settle together on
the single concrete next step. That becomes the Next-focus in step 4.

## 4. Record what happened

Rewrite `context/reset.md` with `Status: closeout`, filling every field of the schema in `SKILL.md`.
The judgment fields are the Disposition — what you integrated, promoted, culled, and retained — and
the Next-focus, set to the step you agreed on in step 3: the handoff the next session reads.

When the project has instead reached a resting point — its deliverable is released, it has no next
step of its own, and a workspace coordinator carries continuity — delete `context/reset.md` rather
than rewriting it. The coordinator's reset file becomes the pointer to what happens next.

## 5. Drop the guide

If this session created `context/guide.md` — a `code` project's `start` — delete it. It's done its job;
what remains is the code and the reset file. A finished guide left lying around is exactly the kind of
already-built documentation the workflow is trying to avoid. A `context` project and a `plan` session
have no guide to drop.

## 6. Commit and publish

Stage everything and commit; this fires the `on-commit` hook. Then publish the branch with the remote
platform's command from `.claude/marathon.toml` — `gh pr create`, `glab mr create`, or the project's
equivalent — using the change description from `context/reset.md`. If the project declared no remote,
stop after the commit.

Run the `on-closeout` hook so any project-management extension can mirror the result outward. With no
extension, it does nothing.
