# Marathon optimization pass

Candidate direction for a dedicated future session (or sessions) over the marathon skill. The
pass is scheduled as `backlog.marathon-optimization` in the workspace roadmap; the 0.7.0
restructure should see real use before the mechanical and decision-guidance axes run. Two
conventions within it — the workspace reset and the comment-free guide — are already settled
in practice and wait only on codification.

## Workspace axis: a single reset file at the coordinator

Settled in practice at the go-database `query-vocabulary` closeout (2026-08-25). A workspace
maintains one reset file, at the
coordinator; member projects carry none. The current mechanics route continuity through member
reset files — the coordinator's reset names the member, and the member's own `reset.md` carries
the record — which splits one session's story across repositories and is exactly where a session
already failed to keep both current. Under the convention, LOCATE at any workspace directory
resolves the coordinator's reset directly, Next-focus names the member project a session
continues in, and the member-reset routing (including the resting-point deletion rule) drops out
of `mechanics/pipeline.md` and `SKILL.md`. A standalone project keeps its own reset unchanged.

## Guide axis: no comments in implementation guides

Settled in the go-web-sdk `read-contract` session (2026-08-26), after guides had carried godoc
comments for several sessions running. The rule was already implied but never stated in one
place: `commands/start.md` excludes "the tests or documentation the agent adds at closeout,"
`references/role-boundary.md` puts in-source comments and API documentation on the agent's
side of the line, and `references/implementation-guides.md` bans "explanatory comments" in
code blocks. What defeats the implication each session is the same reference's "the same as
the production code will be" — the finished code *does* carry godoc, so the idiomatic-Go prior
resolves the ambiguity toward documented code every time a fresh context reads it.

Codification: one explicit sentence in `references/implementation-guides.md` — the guide's
code blocks carry no comments of any kind; doc comments and API documentation are the agent's
closeout work, not part of the handoff. Under the git-blame test, a comment the developer
types from the guide would blame to the wrong party.

## Mechanical axis: how the skill files load

Captured at the close of the `marathon-extensions` session (2026-08-25). Two standpoints:

- **Data flow** — how a session's reads actually traverse SKILL.md, `mechanics/`, the command
  playbooks, and `references/`, now that the pipeline is the execution spec and the playbooks carry
  only stage content. Testing will show where a session reads a file twice, or reads one it didn't
  need.
- **Context budget** — what each session type costs in loaded context, and where the tiers can
  shed weight without losing a distinction: the always-active `behavior/` load, the pipeline and
  hook specs, and the per-command remainder.

## Decision-guidance axis: how the workflow chooses under uncertainty

Captured 2026-08-25 from an evaluation of the pathfinding-under-uncertainty literature — the
Canadian Traveller Problem, D\* incremental replanning, receding-horizon control, value of
information (Howard), Boehm's spiral model — against the skill.

The evaluation mostly confirmed the design, and the future session shouldn't relitigate what it
confirmed: one-step planning with shallow far notes is receding-horizon control; deliberate
promotion out of `experiments/` and `concepts/` is staged commitment; the decay/promote/cull cycle
is incremental repair rather than replanning from scratch. Three gaps surfaced, each a candidate
enhancement with a named target:

- **Value-of-information test.** An experiment's worth is not how much uncertainty it removes but
  whether its answer changes what you'd do next. Target: `commands/experiment.md`, Settle — alongside
  "what the spike is testing," name the decision the answer changes; if no decision changes either
  way, the spike isn't worth running. The same test applies in `commands/plan.md` when choosing
  Next-focus: probe where uncertainty × consequence is highest, not where uncertainty alone is.
- **Assumption annotations.** The skill tracks settled intent (`design/`) and unsettled ideas
  (`concepts/`) but not what a note silently rests on, so when a build falsifies an assumption,
  finding the invalidated notes is a judgment sweep at `review`. Candidate convention: a `design/`
  or `concepts/` note names the unverified assumptions it depends on, and the reset Disposition
  records when one is falsified — a surprise then invalidates identified notes, not the whole tree.
  Open question: a core convention in `references/context-engineering.md`, or a separate extension.
- **Ordering stance.** "Start from the lowest-level requirement" is dependency-driven ordering; the
  spiral model orders by risk, attacking the highest-consequence unknown first. The two conflict
  when the scariest unknown sits high in the stack. Candidate reconciliation: name `experiment` as
  the risk-first instrument — spike the high-consequence unknown cheaply while the builds proceed in
  dependency order — as a sentence in SKILL.md "Iterative development" or `commands/plan.md`.
