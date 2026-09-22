# Agent profiles: delegation by purpose, with the model chosen per engagement

Captured 2026-09-22 from a harness-realignment planning session at the coordinator
(`standards-lab`). This concept is settled direction for the `v1.harness.workflow-refinement`
session, which implements it alongside `checkpoint-execution.md`.

## The gap

Delegation today binds each purpose to a model. The architect's personal agents in
`~/.claude/agents/` were named `fable` and `opus`, each pinned to its model, and the workspace's
`[workspace.agents]` table in `standards-lab/.claude/marathon.toml` described them by which model
plans and which executes. The pairing had to be inverted once already (standards-lab 0ac75ca),
when the capability ordering between the two models changed, and a model release right after the
workflow-refinement planning session made it stale again. The purpose of a delegation is stable.
The best model for it is not.

## Design

marathon ships three subagent profiles in its plugin, as `plugins/marathon/agents/planner.md`,
`executor.md`, and `reviewer.md`. None sets `model:` in its frontmatter, so each inherits the
session's model unless the session passes one when it engages the profile.

- The planner profile designs the session's initial stage list at SETTLE, from the full context
  the session hands it. It changes nothing and returns a recommendation. The session revises
  the plan itself afterward: a revision during SETTLE or a re-plan does not need a new planner
  engagement.
- The executor profile executes one technical stage of the approved stage list against the
  stage's own check (build, test, lint).
- The reviewer profile reviews the branch's code once the final technical stage completes. This
  is the branch review `checkpoint-execution.md` places ahead of `close`'s context tending, and
  its alignment check defaults to ecosystem idiom as that note describes.

The session chooses the model for each engagement. Before engaging a profile, it states which
profile, which model, and why, so the architect can redirect. A session can also do any of this
work itself when briefing a delegate would cost more than the work.

The `[agents]` and `[workspace.agents]` tables leave `mechanics/configuration.md`. The profiles
are fixed by purpose, so a project has nothing to declare, and a session no longer infers a
reviewer from a project's prose. `behavior/delegation.md` keeps its grain, the limits on what a
delegate does, and what the session owes. Its "Cataloging agents" section is replaced by the
three profiles and the rule for choosing a model.

## Files this touches

- New in `plugins/marathon/agents/`: `planner.md`, `executor.md`, and `reviewer.md`.
- `behavior/delegation.md`: the profiles replace the catalog.
- `mechanics/configuration.md`: the "Agents" section and the `[agents]` example are removed.
- `references/staged-execution.md`: the planner at SETTLE and the executor per stage.
- `commands/close.md`: the reviewer's branch review.
- `.claude-plugin/plugin.json`, only if the plugin needs to declare its agents explicitly.

At `standards-lab`, `.claude/marathon.toml` drops its `[workspace.agents.fable]` and
`[workspace.agents.opus]` tables as a Cross-repo edit in the same session.

## Open questions for the implementing session

- Whether a profile's own instructions should suggest a default model tier (for example, the
  reviewer defaulting to the most capable model available), or leave the choice entirely to the
  session.
