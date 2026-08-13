# marathon functions: a tier for moment-bound operations

Captured 2026-08-13 from a session in the reference-architecture workspace. A `start` session in
go-service outgrew its scope while planning the data composition and CQRS layer, and the session
was converted by hand into a planning capture that produced a seven-rung session ladder
(go-service `concepts/data-cqrs-roadmap.md`). The maneuver worked, but the skill gave it no name:
nothing said when to make the move or what it should produce. This concept proposes the tier that
names it. Everything here is candidate direction for a future session on this repository.

## The gap

The skill's `commands/` directory holds session verbs. Each command defines what a session is
(`plan`, `start`, `experiment`), how it ends (`reset`, `close`), or how it fans out
(`coordinate`); all of them are bound to the session loop. The skill has no vocabulary for an
operation invoked at a moment inside a running session, whatever kind of session it is. The
strain is already visible: `review` and `docs` are documented as on-demand maintenance passes
rather than working sessions, a third category squeezed into the command table.

## Proposal

A `functions/` directory beside `commands/` in the marathon skill
(`plugins/marathon/skills/marathon/`): reusable operations invocable at any point in any session.

Candidate admission criteria, so the tier does not become a junk drawer:

- A function is invocable at any moment; it does not care which command opened the session.
- A function writes durable output into `context/` (a concept, a reset revision), and leaves
  session mechanics to commands: it never opens a branch, publishes, or closes out.
- A function codifies a maneuver that sessions have already needed, not a speculative one.

## First function: scope deconstruction

Invoked when a design under discussion outgrows a single session. The function:

- settles the strategy spine of the oversized design with the developer;
- writes a ladder concept into `concepts/`: a sequence of named, session-sized scopes, each rung
  settling its own API in its own session, all direction candidate until a rung settles it;
- redirects the running session: the reset records it as a planning capture (`Session: plan`)
  whose deliverable is the ladder, and Next-focus points at the first rung.

The guardrail is marathon's own one-step principle. A deconstruction produces named scopes,
never designs; depth stays deferred to each rung's plan mode, because planning far ahead is the
part that usually turns out wrong. The data-cqrs ladder is the worked example: rungs are named,
rung APIs are settled one session at a time.

## Open questions

- The function's name (`deconstruct` and `decompose` are candidates) and the invocation path:
  commands route on the skill's first argument, and functions need either the same routing or a
  distinct form.
- Whether `review` and `docs` migrate into the tier or stay commands; they end through `close`
  and carry their own Session values, which the admission criteria above would exclude.
- What the second function is. The tier is justified by the category, but it ships with one
  member; candidates should be collected before the structure is generalized to other plugins.
