# marathon experiment

Spike an idea before committing to it, when you need to learn whether a design or an approach holds
up rather than build the settled next step (that's `start`). An experiment is a standalone marathon
project of its own, outside the project or workspace it serves, with its own reset file, so its
sessions run in parallel with the served project's. The `experiment` session only sets it up.

An experiment earns nothing by default: a spike that works is evidence, not a decision.

`experiment` runs the session pipeline (`mechanics/pipeline.md`). A handoff recorded under Session
`experiment` resumes here.

## Settle

- **The question** the spike answers, and the decision its answer changes. If no decision changes
  either way, don't run it.
- **Where it lives**: the local directory, and the account or organization hosting its repository.
  Neither is assumed from the served project; propose the architect's standing convention where
  the served project's context records one. The experiment always has a remote.
- **`init`'s founding decisions** (`commands/init.md`), with the first spike step as its first
  step, so `init` runs without a second approval.
- **The catalog**: the file the served project keeps its experiments or repositories in, or a new
  one if it keeps none.

Branch slug: the spike.

## Execute

1. Run `init` at the settled location with the settled decisions. `[experiment]` in its
   `.claude/marathon.toml` names the project it serves (`mechanics/configuration.md`), and its
   `context/README.md` names the question and the goal it serves.
2. Record the repositories it reads in the served project's own convention for that (in the
   workspace, a committed list with remotes and a gitignored map to local checkouts). It reads
   them and never writes them; code dependencies are published versions, never a replace
   directive.
3. Create the repository on the settled host and push.
4. Add the experiment and its remote to the settled catalog in the served project (the
   coordinator, in a workspace).

The architect then works the experiment in its own directory with `start`. Its final step's
validation is the answer to its question.

## Conclude

The experiment ends with an ordinary `close` in its own project, whose Disposition states the
question, the answer, and the evidence. A `plan` session in the served project then reads it,
decides with the architect what it earned, and captures that as notes and roadmap work. A result
that earns a repository gets a step that runs `init` in a new sibling directory and rebuilds the
result there rather than copying the spike. The same `plan` session confirms the experiment is
pushed, archives its remote read-only (`gh repo archive` on GitHub), and keeps the catalog entry
as the pointer. The experiment never edits the served project.
