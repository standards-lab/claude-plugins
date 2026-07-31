# marathon plan

A planning and curation session that works entirely inside `context/`. Use `plan` to create and/or
refine concepts, settle a design, reshape the capability map, or decide what the next `start` should
focus on — the bigger-picture thinking that doesn't fit inside a single build, and getting a concept
ready before you build it.

A `plan` session changes no product: no code, no skill, no prose deliverable — only the agent's written
context. This is what separates it from `start`. It applies to both project kinds: a code project uses
it to think ahead of the build; a context project uses it to work out a design before authoring the
change that expresses it.

## plan is not review

Both `plan` and `review` touch only `context/`, but they look in opposite directions. `review` is a
backward-looking audit: has the context drifted from what the code now expresses? `plan` is
forward-looking: what is the next piece of work, and what shape should it take? Reach for `review` when
the notes feel stale against the code; reach for `plan` when you need to think the next step through.

## 1. Plan in plan mode

Read `context/reset.md` first. If its Status is `handoff` and its Session is `plan`, a previous
planning session stopped mid-work: check out the open branch, read the Next-focus, and pick up from
there — skip the fresh planning below. (A handoff recorded under another Session belongs to that
command.)

Otherwise enter plan mode and settle the thinking with the developer before changing any notes — this
is where the architectural work happens, so give it real depth.

1. Get oriented: the Next-focus in `reset.md`, the capability map in `context/README.md`, and the
   `design/` and `concepts/` notes the topic touches. Load only what the topic needs.
2. Work the topic through with the developer: refine a concept, settle a design decision, or map out
   what the next step should be and how far it should go. Add detail only as far as the upcoming work
   needs it — no planning ahead into steps you haven't earned.
3. The discussion will range wide; capture what belongs to later steps as concepts, and note what the
   next focus looks like.

## 2. Tend the context

With the scope approved, run the `on-session-start` hook and create a branch named after the topic —
the session's changes land on it.

Then bring `context/` in line. Show the developer what you propose to move or delete and get
agreement first.

- **Promote** — move a concept from `concepts/` into `design/` if this session settled it, and say why.
- **Cull** — drop a concept the discussion ruled out.
- **Add or sharpen** — write the new concepts and refine the capability map or the design note the
  upcoming work needs.

Capture only what the upcoming step needs; a note that serves no upcoming work is clutter. Write new and
revised notes per the voice standard in `references/writing-voice.md`.

## 3. Land the changes

A `plan` session's output is context, so land it like any other work — on a branch, published the same
way — so the planning is traceable. Close the session with `close`: it records the dispositions in
`context/reset.md` and sets the Next-focus to the step this planning teed up.
