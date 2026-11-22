# Intake round

The rules every format shares are in [`../briefs.md`](../briefs.md).

`intake` runs plan rounds led by the spikes' answers (`commands/intake.md`). The header names the
experiment and the goals it serves, then summarizes each spike's answer section
([`answer-section.md`](./answer-section.md)) under its name: its question, its one-line answer, and
its evidence. The numbered questions follow, in the plan round's form
([`plan-round.md`](./plan-round.md)).

> ## INTAKE ROUND 2 · experiment.ai → v1.ai, v1.agents
>
> **spike-harness-driver**
> - **question:** one driver interface over Pi, Claude Code, and OpenCode?
> - **answer:** yes; stdin/stdout JSONL carries all three
> - **evidence:** 3 items; item 3 by tests only
>
> **spike-local-subagents**
> - **question:** do local models hold a subagent's tool loop?
> - **answer:** no; tool calls drift past 20 turns
> - **evidence:** 4 items; item 2 waits on a larger model
>
> **Questions**
>
> 1. **Does v1.ai build its first task, harness-driver, on the driver interface?**
>    - **rec:** yes, as its first task.
>    - **changes:** adds a task to v1.ai; folds the answer into context/ai.md.
> 2. **Does v1.agents keep its local-subagents task?**
>    - **rec:** no; cloud models only until a larger local model is tested.
>    - **changes:** drops v1.agents' local-subagents task.
> 3. **Archive spike-local-subagents' remote, or keep it open?**
>    - **rec:** archive it; keep spike-harness-driver open as v1.ai's reference.
>    - **changes:** one remote archived after the merge.
>
> **reply:** "1 ok, 2 ok, 3 yes but ..."

- The first round names the served goals, which the header lists from then on, and
  consolidates the spikes' answer sections.
- Each spike's summary comes from its answer section, with any evidence proven only by tests,
  or still waiting, called out.
- Every question has a **rec:** and a **changes:** line, as in a plan round
  ([`plan-round.md`](./plan-round.md)).
