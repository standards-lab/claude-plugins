# Status digest

The rules every format shares are in [`../briefs.md`](../briefs.md).

`status` prints a table with one row per active goal: the goal, its tasks done out of its total,
its state, and the repositories it locks. Goals that need the architect sort first and carry a
`!`.

> ## Status · 2026-10-01
>
> | goal | done | state | locks |
> |---|---|---|---|
> | ! experiment.ai.spike-harness-driver | 2/3 | brief ready | spike-harness-driver |
> | experiment.ai.spike-local-subagents | 1/3 | building | spike-local-subagents |
> | factory | 1/3 | building | claude-plugins |
> | v1.messaging | 0/1 | idle | go-messaging, go-core, go-web-service |

The state is the goal record's State line (`mechanics/goal-record.md`). A goal with no record
yet shows `no record`.
