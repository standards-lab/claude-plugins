# marathon experiment

Spike an idea before committing to it, when you need to learn whether a design or an approach
works. Experiments are goals under `experiment`, named for the topic they investigate:
`experiment.ai` settles what `v1.ai` is built from. Each spike is a sub-goal,
`experiment.<topic>.<spike>`, named for the spike's repository, which lives outside the project or
workspace it serves (`mechanics/goals.md`).

An experiment decides nothing on its own: a spike that works is evidence, not a decision.

`experiment` runs once per topic. It creates the experiment goal, its `intake` task, and its
planned spikes. When `experiment.<topic>` already exists, stop: a later spike is added with
`plan`. `experiment` runs `plan`'s stages (`commands/plan.md`), with the rounds and edits below.

## Rounds

Settle in plan rounds:

- **The topic**: the experiment's slug, name, and summary, and the note it cites in its
  `context`, where the spikes' answers land.
- **The spikes**, in the order they run. For each spike:
  - **The question** it answers, and the decision its answer changes. If the decision is the same
    either way, don't plan the spike.
  - **Its slug**, which names its repository.
  - **Where it lives**: the local `path` and the `remote`. Propose the coordinator's
    `[workspace.experiments]` convention, with `<slug>` replaced, when it declares one
    (`mechanics/configuration.md`); otherwise assume neither from the served project. A spike
    always has a remote.

## Edits

At the coordinator, in one commit on an `experiment-<topic>` branch:

- Add `experiment.<topic>` with its `intake` task, which `marathon intake` runs once no spike
  remains (`commands/intake.md`). The experiment goal is never listed in the arrays.
- Create the note the experiment cites, or extend it when it exists, with an
  `## Answers · experiment.<topic>` heading under which the spikes' answer sections land
  (`references/briefs.md`, "Answer section").
- Add each spike as a sub-goal with its `root` and `repos` set to its repository, its `remote`
  and `path`, and a summary of its question and the decision it changes
  (`references/manifest.md`). Append the spikes to `planned` in the settled order.

`experiment` creates no repository and stages no spike. Publish and merge the branch with the
approved round outcome as its body (`mechanics/pipeline.md`, "Branches and pull requests").

The architect then sets up each spike with `plan experiment.<topic>.<spike>`, which creates its
repository (`commands/plan.md`, "Spike setup"), and runs it with `start`. The spike's last task's
session brief gives its answer, and its sync lands an answer section in the note the experiment
cites and proposes the next planned spike to stage (`references/briefs.md`, "Answer section";
`mechanics/goals.md`, "Sync").
