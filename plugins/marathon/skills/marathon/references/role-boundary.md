# Role boundary

marathon divides the work between the developer and the agent. The short version: the developer owns the
production code, and the agent does everything around it — including drafting that production code in the
implementation guide for the developer to apply.

## The line

The developer owns the production source code — the implementation logic that makes the program do what
it does. The agent drafts each change in the implementation guide (see `implementation-guides.md`), and
the developer reads it, adjusts anything that's off, applies it, and commits it. Ownership is in applying
and standing behind the code, not in typing it from a blank file: when `git blame` lands on a production
line six months on, it should point at the developer, because they put it there and they answer for it.

The agent writes everything around the code outright — these are the agent's to produce and land
directly:

- tests,
- in-source comments and API documentation,
- prose documentation,
- everything in `context/`, including the reset file, and
- the implementation guide itself.

## The test

When you're unsure who owns something, ask whether it's production logic — the code that makes the
program behave — or the material around it. Production logic is the developer's to apply and own (the
agent drafts it). Tests, comments, docs, and context are the agent's to write outright.

## Why split it this way

Two reasons.

The developer stays fluent in the system. Even though the agent drafts the implementation, the developer
has to read, understand, and apply every change, and adjust anything that's wrong. That active gate is
what keeps them in command of a system they'll maintain for the long haul, rather than passively
reviewing code they never engaged with. The guide is detailed precisely so the developer can take in
each change quickly and decide on it — not so they can wave it through.

The agent takes the work that's valuable but easy to skip. Tests, documentation, and written context are
the first things dropped under time pressure, and dropping them is what makes a project hard to sustain.
Handing them to the agent means they get done consistently instead of being put off.

## Keep it language-neutral

State the boundary in plain terms — "production code," "tests," "documentation" — without tying it to any
one language's conventions. The principle holds everywhere; only the specifics differ.
