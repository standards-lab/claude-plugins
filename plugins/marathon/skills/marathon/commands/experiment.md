# marathon experiment

Spike an idea before committing to it, when you need to learn whether a design or an approach
works. A spike is a goal homed in a repository of its own, outside the project or workspace it
serves, so it runs beside the served project's goals. The `experiment` session only sets it up.

An experiment decides nothing on its own: a spike that works is evidence, not a decision.

`experiment` runs `plan`'s stages (`commands/plan.md`), plus the setup below.

## Rounds

Settle in plan rounds:

- **The question** the spike answers, and the decision its answer changes. If the decision is the
  same either way, don't run the experiment.
- **The evidence**: a numbered list of what the spike must show to answer the question.
- **Where it lives**: the local directory, and the account or organization that hosts its
  repository. Assume neither from the served project. Propose the architect's standing convention
  when the served project's context records one. An experiment always has a remote.
- **`init`'s founding decisions** (`commands/init.md`).
- **The catalog**: the file where the served project lists its experiments or repositories, or a
  new one if it has none.
- **The goal**: its slug, its tasks (the spike's path, one task per session), and the task of the
  served goal that will take in the result.

## Setup

1. Run `init` at the settled location with the settled decisions. Its `[experiment]` table names
   the project it serves (`mechanics/configuration.md`), and its `context/README.md` states the
   question, the evidence list, and the goal it serves.
2. Record the repositories the experiment reads, using the served project's own convention. In a
   workspace, that is a committed list of remotes plus a gitignored map to local checkouts, both
   in the coordinator. The experiment reads those repositories and never writes to them. Its code
   dependencies are published versions, never a replace directive.
3. Create the repository on the settled host and push.
4. At the served project, in direct commits: add the spike's goal to the manifest with `home` and
   `repos` set to the spike's repository, stage it into `active`, and add the experiment to the
   catalog.

The architect then runs the spike with `start <goal>`, from either directory. Each task's session
brief narrates the evidence it produced. The last task's validation is the answer to the
question, with the evidence. The goal record carries the answer as a pending coordinator edit
to the served goal's intake task, which sync applies.

## Intake

A `plan` session in the served project takes in the result as a spike intake
(`references/briefs.md`), decides with the architect what the project takes from it, and records
that as notes and roadmap work. A result that deserves its own repository gets a task that runs
`init` in a new sibling directory and rebuilds the result there, instead of copying the spike. The
same session confirms the experiment is pushed, archives its remote as read-only
(`gh repo archive` on GitHub), and keeps the catalog entry as the pointer to it. The experiment
never edits the served project.
