# marathon start

Advance the product one concrete step — or resume a step already in progress. `start` takes no
argument; how it hands off the step depends on the project's kind (see `references/role-boundary.md`),
not on anything you pass it. For planning that touches only `context/`, use `plan`; for a spike, use
`experiment`.

The first thing `start` does is read `context/reset.md` and check its Status. If the file doesn't
exist, the project is at a resting point: its deliverable was released and continuity sits with the
workspace coordinator. Read the coordinator's reset file for the next step, or settle a fresh step
with the developer if none names this project.

## Resuming a handoff (Status: handoff)

A previous session stopped mid-work. The plan already exists, so don't re-plan from scratch.

First check the Session line. If it records `plan` or `experiment`, the handed-off work belongs to
that command — resume it with `plan` or `experiment` instead, which pick up their own kind of session
the same way. If it records `start`:

1. Check out the open branch named in `reset.md`.
2. Read the in-progress state and the Next-focus; the Next-focus is your exact next move.
3. Run the `on-session-start` hook (does nothing if no extension).
4. Pick up the work. On a **code** project, `context/guide.md` is still the live implementation guide —
   follow it. On a **context** project, resume authoring the change directly from where the Next-focus
   leaves off.

## Starting fresh (Status: closeout)

The previous session finished, so this is a new step. Plan it in plan mode and agree on the
scope with the developer before writing anything. Planning matters as much as the build — this is where
the architectural thinking happens, and it sets how deep and how wide the step should go. Give it real
depth; don't rush to the code.

1. Enter plan mode. Get oriented: read the Next-focus in `reset.md`, the capability map in
   `context/README.md`, and the `design/` and `concepts/` notes the upcoming step touches. Load only what
   the step needs.
2. With the developer, work through the step in enough depth to come out with a clear picture of what to
   build: the single concrete step to take now, and how far it should go. Add detail to the relevant
   `design/` note only as far as this step needs — no further.
3. The discussion will range wider than the step. Tend the context as it does: capture ideas for later
   steps as concepts in `concepts/`, cull the ones it rules out, and note anything that looks like the
   next session's focus.
4. Once it's approved, run the `on-session-start` hook, then create a new branch named after the step.

Then do the step, according to the project kind:

- **code project** — write `context/guide.md`, the implementation guide: the full code for each change,
  in the order to apply it, with prose only where the reasoning needs it, and a short run-and-verify at
  the end. It does not include the tests or documentation the agent adds at closeout. See
  `references/implementation-guides.md` for how to write it well. Then stop and let the developer apply
  it, staying available for fixes; don't run ahead.
- **context project** — there is no guide and no code handoff. Author the change directly: the skills,
  prose, or configuration the step calls for, plus any `context/` the change settles. You are producing
  the deliverable itself, under the developer's review; write its prose per the voice standard in
  `references/writing-voice.md`. Stay within the one step.

## Ending the session

When the work pauses with the context window filling up but the piece isn't done, run `reset` to hand
off. When the piece is finished and validated, run `close`.
