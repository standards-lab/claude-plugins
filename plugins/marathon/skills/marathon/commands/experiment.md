# marathon experiment

Spike an idea before committing to it, when you need to learn whether a design or an approach
works instead of building the planned next step (that is `start`). An experiment is a standalone
marathon project, outside the project or workspace it serves, with its own reset file, so its
sessions run in parallel with the served project's sessions. The `experiment` session only sets
it up.

An experiment decides nothing on its own: a spike that works is evidence, not a decision.

`experiment` runs the session pipeline (`mechanics/pipeline.md`). A handoff under Session
`experiment` resumes here.

## Settle

- **The question** the spike answers, and the decision its answer changes. If the decision is the
  same either way, don't run the experiment.
- **Where it lives**: the local directory, and the account or organization that hosts its
  repository. Assume neither from the served project. Propose the architect's standing convention
  when the served project's context records one. An experiment always has a remote.
- **`init`'s founding decisions** (`commands/init.md`), with the first spike step as the first
  step, so `init` needs no second approval.
- **The catalog**: the file where the served project lists its experiments or repositories, or a
  new one if it has none.

Branch slug: the spike.

## Execute

1. Run `init` at the settled location with the settled decisions. The `[experiment]` table in its
   `.claude/marathon.toml` names the project it serves (`mechanics/configuration.md`), and its
   `context/README.md` states the question and the goal it serves.
2. Record the repositories the experiment reads, using the served project's own convention. In a
   workspace, that is a committed list of remotes plus a gitignored map to local checkouts. The
   experiment reads those repositories and never writes to them. Its code dependencies are
   published versions, never a replace directive.
3. Create the repository on the settled host and push.
4. Add the experiment and its remote to the catalog in the served project (the coordinator, in a
   workspace).

The architect then works on the experiment in its own directory with `start`. The validation of
its final step is the answer to its question.

## Conclude

The experiment ends with an ordinary `close` in its own project. That close's Disposition states
the question, the answer, and the evidence. A `plan` session in the served project then reads the
result, decides with the architect what the project takes from it, and records that as notes and
roadmap work. A result that deserves its own repository gets a step that runs `init` in a new
sibling directory and rebuilds the result there, instead of copying the spike. The same `plan`
session confirms the experiment is pushed, archives its remote as read-only (`gh repo archive` on
GitHub), and keeps the catalog entry as the pointer to it. The experiment never edits the served
project.
