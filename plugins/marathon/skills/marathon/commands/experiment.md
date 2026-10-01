# marathon experiment

Spike an idea before committing to it, when you need to learn whether a design or an approach
works. Experiments are goals under `experiment`, named for the topic they investigate:
`experiment.ai` settles what `v1.ai` is built from. Each spike is a task of its experiment goal,
named for the spike's repository, which lives outside the project or workspace it serves. The
experiment's last task is the intake. The `experiment` session sets up one spike.

An experiment decides nothing on its own: a spike that works is evidence, not a decision.

`experiment` runs `plan`'s stages (`commands/plan.md`), plus the setup below.

## Rounds

Settle in plan rounds:

- **The question** the spike answers, and the decision its answer changes. If the decision is the
  same either way, don't run the experiment.
- **The evidence**: a numbered list of what the spike must show to answer the question.
- **The experiment goal**: an existing `experiment.<topic>` the spike joins, or a new one with an
  `intake` task last.
- **Where it lives**: the local directory and the remote. Propose the workspace's convention when
  the coordinator's `[workspace.experiments]` declares one (`mechanics/configuration.md`);
  otherwise assume neither from the served project. An experiment always has a remote.
- **`init`'s founding decisions** (`commands/init.md`), with the spike's path as its first steps.

## Setup

1. Run `init` at the settled location with the settled decisions. Its `[experiment] serves`
   holds the coordinator's path (`mechanics/configuration.md`), and its `context/README.md`
   states the question, the evidence list, and the experiment goal.
2. Record the repositories the spike reads, using the served project's own convention. In a
   workspace, that is a committed list of remotes plus a gitignored map to local checkouts, both
   in the coordinator. The spike reads those repositories and never writes to them. Its code
   dependencies are published versions, never a replace directive.
3. Create the repository on the settled host and push.
4. At the coordinator, in one direct commit: add the spike as a task of its experiment goal, with
   its `remote` and `path` (`references/manifest.md`), add its repository to the goal's `repos`,
   and stage the goal into `active` if it isn't already. A new goal's `root` is its first spike.

The architect then runs the spike with `start experiment.<topic>`. Each task's session brief
narrates the evidence it produced, and the spike task's last brief gives the answer to its
question, with the evidence.

## Intake

The experiment's last task is its intake, a `start` session led by the spike intake round
(`references/briefs.md`, "Spike intake"). It decides with the architect what the served goal takes
from each answer, and records the decisions in the goal record: notes and roadmap work for the
coordinator, such as the served goal's tasks, as pending edits. A result that deserves its own
repository becomes a task of the served goal that runs `init` in a new sibling directory and
rebuilds the result there, instead of copying the spike.

The intake's sync applies those edits, moves each spike into the workspace's repository catalog
as archived, and archives the spikes' remotes (`mechanics/goals.md`, "Sync"). The spikes never
edit the served project.
