# Writing implementation guides

The implementation guide is how a `start` on a **code** project hands work to the developer. The agent
writes it at the end of planning; the developer applies it. It lives at `context/guide.md` and is
deleted at closeout — a working document, not a lasting record. A context project has no guide (there is
no code to hand off), so this reference applies only to code projects.

A good guide is exhaustive about *what* to change and sparing with prose. The developer reads code
fluently and often prefers it; prose is for what code can't show — the reasoning and the concepts. Aim
for the point where someone can apply the whole change from the guide alone, without filling in gaps and
without wading through explanation they didn't need.

## Shape

Open with a short preamble: the shape of the change and why, in a few sentences, so the developer has
the model before the mutations. Then the mutations, in the order they should be applied — dependencies
before the code that uses them, so the developer can work top to bottom and keep the tree coherent.
Close with how to run and verify the work.

## Spell out every mutation

Every change to the codebase appears in the guide, in full. If the same change applies to ten call
sites, write all ten — don't say "and the same across the other call sites." That shortcut is exactly
where mistakes hide, and it pushes the work of finding every site onto the developer. Each mutation
names its exact location: the file, and a precise anchor within it (the function, the symbol, the line
it follows), so there's no question where it goes.

## Code carries the what; prose carries the why

Show the code, not a description of it. For new code, give the full block. For a change to existing
code, show the region as it should look afterward, with enough surrounding lines to place it
unambiguously — not the whole file.

Keep the code blocks clean, the same as the production code will be: no explanatory comments inside
them. To point out what changed, use a line of prose before the block ("in `Load`, the final return now
validates the parsed config") rather than marking it inside the code — that way the block matches
exactly what lands. When a change needs reasoning — why this approach, what a concept means, a
constraint that isn't obvious — put a line or two of prose above the block. When the code is
self-evident, let it stand on its own.

The balance to hold: enough that the change is unambiguous and easy to follow, and no more. Too terse
and the developer is reconstructing your intent; too thorough and the guide is exhausting to read.
Include exactly what's needed to apply the change with confidence. The prose itself follows the voice
standard in `references/writing-voice.md`.

## Stay within the step

A guide covers the session's single step and nothing else. No opportunistic refactors, no unrelated
cleanups riding along. If something else needs doing, it's a concept or a future step, not a
detour in this guide.

## End with how to run and verify

Finish with a short, concrete sequence for running the work and confirming it behaves as intended — the
commands to run and what to look for. Enough that the developer can see the change is working before
handing back for tests at closeout. Keep it to the actual steps of a run-and-see; the tests come at
closeout.
