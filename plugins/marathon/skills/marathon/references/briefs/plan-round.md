# Plan round

The rules every format shares are in [`../briefs.md`](../briefs.md).

A PLAN round holds every question whose prerequisites are settled, and the architect answers
them all at once, inline. A question that depends on another open question's answer waits for the
next round; it is never asked conditionally, as in "only if 1 is yes". The round is presented as
`behavior/planning.md` says.

> ## PLAN ROUND 1 · v1.ai · task harness-adapters
>
> **Facts found**
> - Pi has an RPC mode (pi/docs/rpc.md).
> - Claude Code streams JSON (stream-json).
>
> **Currency**
> - go-core
>   - go 1.24.2 → 1.25.1: bug fixes only
>   - pgx v5.6.0 → v6.0.0: Query takes options
>   - chi v5.0.12 → v5.2.0: held (Decisions)
> - go-web-service: current
>
> **Questions**
>
> 1. **Harness transport: stdin/stdout JSONL or SSE?**
>    - **rec:** stdin/stdout JSONL. Pi, Claude Code, and OpenCode all support it; only OpenCode
>      supports SSE.
>    - **changes:** the adapters share one protocol instead of three.
> 2. **go-core's pgx v6.0.0, a new major: upgrade now, hold, or give it its own task?**
>    - **rec:** its own task. v6's Query takes options, which reaches every caller.
>    - **changes:** adds task pgx-v6; holds pgx at v5.6.0.
> 3. **Session resume after a driver restart: build it in v1, or leave it out?**
>    - **rec:** leave it out of v1 and note it in the backlog.
>    - **changes:** drops the persistence slice.
>
> **reply:** "1 ok, 2 ok, 3 yes but ..."

- Number every question.
- Name each question's subject in its own words, with the options it chooses between, so the
  architect can decide it from the question, its **rec:**, and its **changes:** alone. All three
  name the things they refer to and never point elsewhere for them: not "this", "option B",
  "the three files above", **Facts found**, another question, a term defined outside the
  question, or a file to open. A fact the question leans on is restated in it. Before printing
  the round, reread each question for "above", "below", "this", and "these", and replace each
  with the thing it points at.
- Decide one thing per question. When two choices are coupled, ask first the one that settles
  the other. Never bundle two decisions into lettered combinations.
- Give every question a **rec:** with its reason, and a **changes:** line saying what the answer
  changes.
- Cite each fact with its source under **Facts found**.
- Round 1 carries a **Currency** list after **Facts found**, per repository the task touches: each
  trailing item as its currency command reports it, with a one-line summary of its release notes,
  or `current`, or `no currency command`, or the failed command as a fact. A held item is listed
  as a fact (`mechanics/goal-record.md`, Decisions).
- Ask about a trailing item only when it needs a decision: adapt, a breaking change the code must
  absorb; adopt, a new feature worth using; or a new major, to upgrade now, hold, or give its own
  task. A patch or minor bump with nothing to adapt stays a fact, and the upgrade slices take it.
- The rounds end when no question remains open. A BUILD escalation uses the same format, headed
  `ESCALATION` in place of `PLAN ROUND <n>` ([`references/build.md`](../build.md)).
