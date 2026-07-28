# Role boundary

marathon divides the work between the developer and the agent. Where the line falls depends on the
project's kind (declared at `init`; see the project-kind section of `SKILL.md`). There are two shapes,
one per kind.

## Code projects: the developer owns the production code

On a `code` project the developer owns the production source code — the implementation logic that makes
the program do what it does. The agent drafts each change in the implementation guide (see
`implementation-guides.md`), and the developer reads it, adjusts anything that's off, applies it, and
commits it. Ownership is in applying and standing behind the code, not in typing it from a blank file:
when `git blame` lands on a production line six months on, it should point at the developer, because
they put it there and they answer for it.

The agent writes everything around the code outright — these are the agent's to produce and land
directly:

- tests,
- in-source comments and API documentation,
- prose documentation,
- everything in `context/`, including the reset file, and
- the implementation guide itself.

The prose in this material — comments and documentation, the files in `context/`, the guide — follows
the voice standard in `writing-voice.md`.

## Context projects: the agent authors the repository

On a `context` project there is no production source code. The repository *is* context — its skills,
prose, and configuration are advanced context, not a program the developer maintains line by line. So
there is nothing to hand off through a guide: the agent authors the repository directly, all of it,
including everything in `context/`. The developer owns the work by directing it and by reviewing and
approving each change; the pull request is where that ownership is exercised, in place of applying a
guide. There are no tests and no implementation guide, because there is no code layer to test or hand
off. The prose the agent authors follows the voice standard in `writing-voice.md`.

## The test

When you're unsure who owns something, first ask what kind of project you're in. On a `code` project,
ask whether the thing is production logic — the code that makes the program behave — or the material
around it: production logic is the developer's to apply and own (the agent drafts it), while tests,
comments, docs, and context are the agent's to write outright. On a `context` project the question
doesn't arise: there is no production logic, so the agent authors everything and the developer reviews.

## Why split it this way

On a `code` project, two reasons. The developer stays fluent in the system: even though the agent
drafts the implementation, the developer has to read, understand, and apply every change, and adjust
anything that's wrong. That active gate is what keeps them in command of a system they'll maintain for
the long haul, rather than passively reviewing code they never engaged with. The guide is detailed
precisely so the developer can take in each change quickly and decide on it — not so they can wave it
through. And the agent takes the work that's valuable but easy to skip — tests, documentation, written
context — which are the first things dropped under time pressure and the dropping of which is what makes
a project hard to sustain.

On a `context` project the deliverable *is* the writing, so the same gate lives in review rather than in
applying a guide: the developer directs and approves, and the agent does the authoring that a context
project exists to produce.

## Keep it language-neutral

State the boundary in plain terms — "production code," "tests," "documentation," "context" — without
tying it to any one language's conventions. The principle holds everywhere; only the specifics differ.
