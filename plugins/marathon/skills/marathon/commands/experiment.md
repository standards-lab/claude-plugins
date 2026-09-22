# marathon experiment

Spike an idea in isolation before committing to it. Use `experiment` when you need to try something
to learn whether it holds up — a design you're unsure of, an approach you want to feel out — rather
than to build the settled next step (that's `start`).

An experiment earns nothing by default. Its results are concepts, not settled design: a spike that
works is evidence, not a decision. Nothing is recorded as settled — or moves into the product —
without a deliberate promotion at closeout.

`experiment` runs the session pipeline (`mechanics/pipeline.md`). A handoff recorded under Session
`experiment` resumes here.

## Settle

The scope to settle is what the experiment is testing and how you'll know it worked — the question
the spike answers. Name, alongside the question, the decision its answer changes: a spike is worth
running only if some choice comes out differently depending on the result. If no decision changes
either way, don't run it. Keep the scope to that question; an experiment that sprawls stops being
cheap to set aside.

Then settle the experiment's grain with the architect:

- **In-tree** — one session, producing evidence for the step directly in front. The spike runs in
  the top-level `experiments/` directory of this project, or of the coordinator in a workspace.
- **Isolated** — a standalone marathon project of its own, outside the workspace tree. Choose it
  when the spike needs several sessions, may become a new repository, or should run in parallel
  with the workspace's own sessions. It is not a workspace member, so it keeps its own reset
  file and branches, and the coordinator's reset never sees its sessions.

Branch slug: the spike.

## Execute: in-tree

Make the `experiments/<slug>/` directory and do the spike there. Where that directory is depends
on how the project sits:

- A **standalone** project keeps its own top-level `experiments/`.
- In a **workspace**, every in-tree experiment lives at the coordinator, under the coordinator's
  top-level `experiments/`, whatever the spike's scope and whichever member repository's question
  it answers. The branch is created at the coordinator, and the reset file's Project line names
  the coordinator, plus any member repository the session also edits. The reasons are in
  `references/workspace-coordination.md`.

Both project kinds spike the same way, in isolation, treating the result as evidence:

- a **code** project tries an implementation approach in throwaway code;
- a **context** project trials a new skill or agent idea before it becomes real — a draft skill, a
  reworked command playbook — kept in `experiments/` until it's proven worth adopting.

Stay inside `experiments/`. Don't reach into the real tree or the product; the isolation is what
makes the spike safe to explore. The directory is tracked like any other, and the spike runs in
stages under `references/staged-execution.md`.

## Execute: isolated

This session sets the experiment up and ends; the spike itself runs in the experiment's own
sessions. At the location the architect chooses:

1. Run `init` there as a new project. When it serves a workspace, `init`'s founding decisions
   record that in `.claude/marathon.toml` (`[experiment]`, `mechanics/configuration.md`), and its
   `context/README.md` names the question and the workspace goal the experiment serves.
2. Record the workspace repositories the experiment draws on: a committed list of the
   repositories and their remotes, and a gitignored map to their local checkouts. Where the
   coordinator keeps its own convention for this, the experiment follows it. The experiment's
   sessions read those checkouts and never write them. Code dependencies on member modules are
   published versions, through `go.mod` or the equivalent, never a replace directive.
3. Write the experiment's first reset file, with the first spike step as its Next-focus.
4. In the coordinator, record the experiment where the workspace catalogs its repositories.

The session then closes as usual. The architect opens a session in the experiment's directory and
works it with `start`.

## Conclude: promote deliberately, or don't

For an in-tree experiment, decide with the architect at `close` what the spike earned. A result
that proved out is promoted on purpose — captured as a concept note in `context/`, or teed up as
the next `start` — and promotion is what moves proven work into its real home. A result that
didn't prove out is simply not promoted. Either way, the spike itself stays under
`experiments/<slug>/` and merges with the branch: the directory is the durable record of the
project's exploratory work, kept isolated from the product tree. The reset disposition records the
outcome — promoted, or retained as record. Stable context never cites `experiments/`; anything
worth referencing has been promoted out of it.

An isolated experiment ends with an ordinary `close` in its own project, whose Disposition makes
the record self-contained. Graduation is a read, not a write: a `plan` session at the coordinator
reads the closed experiment and does the intake, capturing its findings as notes and adding the
work the goal now needs. When the result earns a repository of its own, that work includes a
step that runs `init` in a new sibling directory, rebuilds the result there rather than copying
the spike, and adds the repository to the coordinator's `order`. The experiment never edits the
workspace; the workspace reads the experiment.
