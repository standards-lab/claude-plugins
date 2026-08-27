# marathon plan

A planning and curation session that works entirely inside `context/`. Use `plan` to create and/or
refine concepts, settle a design, reshape the capability map, or decide what the next `start` should
focus on — the bigger-picture thinking that doesn't fit inside a single build, and getting a concept
ready before you build it.

A `plan` session changes no product: no code, no skill, no prose deliverable — only the agent's written
context. This is what separates it from `start`. It applies to both project kinds: a code project uses
it to think ahead of the build; a context project uses it to work out a design before authoring the
change that expresses it.

`plan` runs the session pipeline (`mechanics/pipeline.md`). A handoff recorded under Session `plan`
resumes here.

## plan is not review

Both `plan` and `review` touch only `context/`, but they look in opposite directions. `review` is a
backward-looking audit: has the context drifted from what the code now expresses? `plan` is
forward-looking: what is the next step, and what scope should it be constrained against? Reach for
`review` when the notes feel stale against the code; reach for `plan` when you need to think the next
step through.

## Settle

The scope to settle is the topic: work it through with the developer — refine a concept, settle a
design decision, or map out what the next step should be and how far it should go. Add detail only as
far as the upcoming work needs it — no planning ahead into steps you haven't earned. This is where the
architectural work happens, so give it real depth.

When the topic is choosing the next focus, weigh uncertainty against consequence together: probe
where an answer would change what the project does next, not merely where the most is unknown. An
unknown that changes no decision can stay unknown.

Branch slug: the topic.

## Execute: tend the context

Bring `context/` in line with what the discussion settled. Show the developer what you propose to move
or delete and get agreement first.

- **Promote** — move a concept from `concepts/` into `design/` if this session settled it, and say why.
- **Cull** — drop a concept the discussion ruled out.
- **Add or sharpen** — write the new concepts and refine the capability map or the design note the
  upcoming work needs.

Capture only what the upcoming step needs; a note that serves no upcoming work is clutter.

## Conclude

A `plan` session's output is context, so land it like any other work — on a branch, published the same
way — so the planning is traceable. Close with `close`: it records the dispositions in the reset
file and sets the Next-focus to the step this planning teed up.
