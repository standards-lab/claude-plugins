# marathon experiment

Spike an idea in isolation before committing to it. `experiment` runs the work in the top-level
`experiments/` directory, kept apart from the real tree so exploratory work never mixes into the
project. Use it when you need to try something to learn whether it holds up — a design you're unsure
of, an approach you want to feel out — rather than to build the settled next step (that's `start`).

An experiment earns nothing by default. Its results are concepts, not settled design: a spike that
works is evidence, not a decision. Nothing moves into `design/` — or into the product — without a
deliberate promotion at closeout.

## 1. Plan the spike

Read `context/reset.md` first. If its Status is `handoff` and its Session is `experiment`, a previous
spike stopped mid-work: check out the open branch, read the Next-focus, and resume it. (A handoff
recorded under another Session belongs to that command.)

Otherwise enter plan mode and settle with the developer what the experiment is testing and how you'll
know it worked — the question the spike answers. Keep the scope to that question; an experiment that
sprawls stops being cheap to set aside. Once the question is agreed, run the `on-session-start` hook
and create a branch named after the spike.

## 2. Work in experiments/

Make the top-level `experiments/<slug>/` directory and do the spike there. Both project kinds spike the
same way — in isolation, treating the result as evidence:

- a **code** project tries an implementation approach in throwaway code;
- a **context** project trials a new skill or agent idea before it becomes real — a draft skill, a
  reworked command shape — kept in `experiments/` until it's proven worth adopting.

Stay inside `experiments/`. Don't reach into the real tree or the product; the isolation is what makes
the spike safe to explore. The directory is tracked like any other: commit the spike's work on the
branch as you go.

## 3. Close out: promote deliberately, or don't

At `close`, decide with the developer what the spike earned. A result that proved out is promoted on
purpose — captured as a concept in `concepts/`, or teed up as the next `start` — and promotion is what
moves proven work into its real home. A result that didn't prove out is simply not promoted.

Either way, the spike itself stays under `experiments/<slug>/` and merges with the branch: the
directory is the durable record of the project's exploratory work, kept isolated from the product
tree. The reset disposition records the outcome — promoted, or retained as record. Stable context
never cites `experiments/`; anything worth referencing has been promoted out of it.
