# Marathon optimization pass

Captured at the close of the `marathon-extensions` session (2026-08-25). Unscheduled; the 0.7.0
restructure should see real use first.

A dedicated session over the marathon skill from two standpoints:

- **Data flow** — how a session's reads actually traverse SKILL.md, `mechanics/`, the command
  playbooks, and `references/`, now that the pipeline is the execution spec and the playbooks carry
  only stage content. Testing will show where a session reads a file twice, or reads one it didn't
  need.
- **Context budget** — what each session type costs in loaded context, and where the tiers can
  shed weight without losing a distinction: the always-active `behavior/` load, the pipeline and
  hook specs, and the per-command remainder.
