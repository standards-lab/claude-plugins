# marathon experiment

Spike an idea in isolation before committing to it. `experiment` runs the work in the top-level
`experiments/` directory, kept apart from the real tree so throwaway work never mixes into the
project. Use it when you need to try something to learn whether it holds up — a design you're unsure
of, an approach you want to feel out — rather than to build the settled next step (that's `start`).

An experiment earns nothing by default. Its results are concepts, not settled design: a spike that
works is evidence, not a decision. Nothing moves into `design/` — or into the product — without a
deliberate promotion at closeout.

## 1. Plan the spike

Enter plan mode and settle with the developer what the experiment is testing and how you'll know it
worked — the question the spike answers. Keep the scope to that question; an experiment that sprawls
stops being cheap to throw away.

## 2. Work in experiments/

Make the top-level `experiments/<slug>/` directory and do the spike there. Both project kinds spike the
same way — in isolation, treating the result as evidence:

- a **code** project tries an implementation approach in throwaway code;
- a **context** project trials a new skill or agent idea before it becomes real — a draft skill, a
  reworked command shape — kept in `experiments/` until it's proven worth adopting.

Stay inside `experiments/`. Don't reach into the real tree or the product; the point is that this work
can be discarded without a trace.

## 3. Close out: promote deliberately, or don't

At `close`, decide with the developer what the spike earned. A result that proved out is promoted on
purpose — captured as a concept in `concepts/`, or teed up as the next `start` — and the reset file
records that decision. A result that didn't is dropped. Either way, clean up the `experiments/` work
that has served its purpose so it doesn't linger as a third source of truth.
