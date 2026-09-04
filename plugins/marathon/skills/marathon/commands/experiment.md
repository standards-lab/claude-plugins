# marathon experiment

Spike an idea in isolation before committing to it. `experiment` runs the work in the top-level
`experiments/` directory, kept apart from the real tree so exploratory work never mixes into the
project. Use it when you need to try something to learn whether it holds up — a design you're unsure
of, an approach you want to feel out — rather than to build the settled next step (that's `start`).

An experiment earns nothing by default. Its results are concepts, not settled design: a spike that
works is evidence, not a decision. Nothing moves into `design/` — or into the product — without a
deliberate promotion at closeout.

`experiment` runs the session pipeline (`mechanics/pipeline.md`). A handoff recorded under Session
`experiment` resumes here.

## Settle

The scope to settle is what the experiment is testing and how you'll know it worked — the question
the spike answers. Name, alongside the question, the decision its answer changes: a spike is worth
running only if some choice comes out differently depending on the result. If no decision changes
either way, don't run it. Keep the scope to that question; an experiment that sprawls stops being
cheap to set aside.

Branch slug: the spike.

## Execute: work in experiments/

Make the `experiments/<slug>/` directory and do the spike there. Where that directory is depends
on how the project sits:

- A **standalone** project keeps its own top-level `experiments/`.
- In a **workspace**, every experiment lives at the coordinator, under the coordinator's top-level
  `experiments/`, whatever the spike's scope and whichever member repository's question it
  answers. The branch is created at the coordinator, and the reset file's Project line names the
  coordinator, plus any member repository the session also edits. The reasons are in
  `references/workspace-coordination.md`.

Both project kinds spike the same way, in isolation, treating the result as evidence:

- a **code** project tries an implementation approach in throwaway code;
- a **context** project trials a new skill or agent idea before it becomes real — a draft skill, a
  reworked command playbook — kept in `experiments/` until it's proven worth adopting.

Stay inside `experiments/`. Don't reach into the real tree or the product; the isolation is what makes
the spike safe to explore. The directory is tracked like any other, and the spike runs in stages
under the review gate of `references/staged-execution.md`: each stage is reported with the working
tree uncommitted and commits on the architect's approval.

## Conclude: promote deliberately, or don't

At `close`, decide with the architect what the spike earned. A result that proved out is promoted on
purpose — captured as a concept in `concepts/`, or teed up as the next `start` — and promotion is what
moves proven work into its real home. A result that didn't prove out is simply not promoted.

Either way, the spike itself stays under `experiments/<slug>/` and merges with the branch: the
directory is the durable record of the project's exploratory work, kept isolated from the product
tree. The reset disposition records the outcome — promoted, or retained as record. Stable context
never cites `experiments/`; anything worth referencing has been promoted out of it.
