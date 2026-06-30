# marathon close

Close out a session whose work is finished and validated. It's the same idea as `reset`, but for
completed work: tidy the notes, record what happened, and publish the work.

## 1. Finish the agent's part

In a development session, the developer has applied the production code. Now add the parts the role
boundary gives the agent:

- tests for the new behavior, and
- in-source comments and API docs, plus any prose documentation.

Then validate — run the tests and the build. Don't close on a failure; fix it (within the agent's part)
or hand it back to the developer.

A context session has no code to finish; for an experiment, decide what, if anything, the spike earned.
Either way, go on to step 2.

## 2. Tidy the notes

Bring `context/` in line with what now exists:

- **Decay** — delete `design/` notes the code now expresses, and record that you did.
- **Cull** — drop `concepts/` candidates the implementation invalidated or replaced.
- **Promote** — move proven candidates from `concepts/` into `design/`, and say why.

These changes matter, so show the developer and confirm before applying them.

## 3. Agree on the next step

Talk through what comes next — don't set the direction on your own. What the session turned up often
changes the priority, and the conversation is usually where the next step comes from. Settle together on
the single concrete next piece of work. That becomes the Next-focus in the next step.

## 4. Record what happened

Rewrite `context/reset.md` with `Status: closeout`. Record what you integrated, promoted, culled, and
retained, and set Next-focus to the next piece of work you agreed on in step 3 — the handoff the next
`start` reads.

## 5. Drop the guide

Delete `context/guide.md`. It's done its job; what remains is the code and the reset file. A finished
guide left lying around is exactly the kind of already-built documentation the workflow is trying to
avoid.

## 6. Commit and publish

Stage everything and commit; this fires the `on-commit` hook. Then publish the branch with the remote
platform's command from `.claude/marathon.toml` — `gh pr create`, `glab mr create`, or the project's
equivalent — using the change description from `context/reset.md`. If the project declared no remote,
stop after the commit.

Run the `on-closeout` hook so any project-management extension can mirror the result outward. With no
extension, it does nothing.
