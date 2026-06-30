# marathon start

Begin — or resume — a session. Usage: `start [development|context|experiment]`. The type defaults to
`development` and decides what closeout does, not how the session opens.

The first thing `start` does is read `context/reset.md` and check its Status.

## Resuming a handoff (Status: handoff)

A previous session stopped mid-work. The plan already exists, so don't re-plan from scratch.

1. Check out the open branch named in `reset.md`.
2. Read the in-progress state and the Next-focus; the Next-focus is your exact next move.
3. Run the `on-session-start` hook (does nothing if no extension).
4. Pick up the work. In a development session, `context/guide.md` is still the live implementation
   guide — follow it.

## Starting fresh (Status: closeout)

The previous session finished, so this is a new piece of work. Plan it in plan mode and agree on the
scope with the developer before writing anything. Planning matters as much as the build — this is where
the architectural thinking happens, and it sets how deep and how wide the implementation should go. Give
it real depth; don't rush to the code.

1. Enter plan mode. Get oriented: read the Next-focus in `reset.md`, the capability map in
   `context/README.md`, and the `design/` and `concepts/` notes the coming step touches. Load only what
   the step needs.
2. With the developer, work through the step in enough depth to come out with a clear picture of what to
   build: the single concrete step to take now, and how far it should go. Add detail to the relevant
   `design/` note only as far as this step needs — no further.
3. The discussion will range wider than the step. Tend the context as it does: capture ideas for later
   steps as candidate notes in `concepts/`, cull candidates it rules out, and note anything that looks
   like the next session's focus.
4. Once it's approved, run the `on-session-start` hook, then create a new branch named after the step.

Then continue by type:

- **development** — write `context/guide.md`, the implementation guide: the full code for each change,
  in the order to apply it, with prose only where the reasoning needs it, and a short run-and-verify at
  the end. It does not include the tests or documentation the agent adds at closeout. See
  `references/implementation-guides.md` for how to write it well. Then stop and let the developer apply
  it, staying available for fixes; don't run ahead.
- **context** — no code handoff. Edit the files in `context/` directly to do the planning or writing the
  session is for.
- **experiment** — make the top-level `experiments/<slug>/` directory and spike there. Treat results as
  candidates; nothing moves into `design/` without a deliberate promotion at closeout.

## Ending the session

When the work pauses with the context window filling up but the piece isn't done, run `reset` to hand
off. When the piece is finished and validated, run `close`.
