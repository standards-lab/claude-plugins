# Session brief

The rules every format shares are in [`../briefs.md`](../briefs.md).

The session brief is the architect's view of the finished task, and the body of its pull request. It
opens with what changed, then how to see it behave: the architect judges outcomes, not code. The
session writes it to `.claude/briefs/<goal>.md` in the goal's root, which gitignores
`.claude/briefs/`.

> ## experiment.ai.spike-harness-driver · task 3/5 · harness-adapters
>
> ### Summary
>
> - `clutch run --harness=<pi|claude|opencode>` picks a session's harness.
> - New `pi/` and `claude/` adapters sit behind one `driver.Session` interface.
>
> ### Core changes
>
> 1. **driver.Session**: one interface over every harness. Why: the adapters differ only in
>    transport.
>    ```go
>    type Session interface { Send(ctx, Turn) (Reply, error); Close() error }
>    ```
>
> ### Verify
>
> `mise run check`: ✓ 214 tests, lint, split
>
> 1. `clutch run --harness=pi "hi"` → the reply streams, exit 0.
> 2. `clutch run --harness=claude "hi"` → the same, through Claude Code.
>
> ### Decided without you
>
> 1. Kept stdin JSONL over SSE: simpler, and a spike needs one transport.
> 2. Dropped decision "harness: a Session closes on EOF": the `driver` package docs state it.
>
> ### Redirects and escalations
>
> 1. Escalation, slice 2: raise pi's floor to 0.9 for RPC mode? Answer: yes.
>
> **Merge danger:** two-way door · blast radius: the spike repository only ·
> spike-harness-driver: merge waits on CI
>
> **Needs you:** none

- **Summary** gives the details of what changed that the architect orients by, as markdown
  bullets or a table, never a fenced block.
- **Core changes** lists one to five entries, each with one line on why and, where the code is
  the change, a short excerpt or signature fenced as code.
- **Verify** opens with the check's one-line result. Then come the steps the architect can run,
  outside the check and tests, to see the behavior, each with what they should see. The
  spec-reviewer followed each. A task with no such step, such as a prose change, reads "none".
- **Decided without you** lists the decisions the brief didn't cover and each Decision the
  editor dropped (`mechanics/goal-record.md`). A drop the architect rejects is a redirect, and
  the session restores the line.
- **Redirects and escalations** lists each BUILD escalation and each BRIEF redirect in the task,
  one line each with the architect's answer, or reads "none". Progress holds them
  (`mechanics/goal-record.md`).
- **Merge danger** names the door type, the blast radius, and whether each repository's merge
  waits on CI. A one-way door gets the close read.
  When the task brief ([`task-brief.md`](./task-brief.md)) has a Release line, Merge danger
  repeats its tags, which accepting the brief authorizes (`mechanics/releasing.md`).
- **Needs you** reads "none" when nothing remains.
- **Sync** appears only on the goal's last task, after Needs you. It lists the pending edits the
  sync will apply, per repository, and the goal proposed to stage next: for a spike, the next
  planned spike of its experiment; otherwise the top of `planned`. When a spike's experiment has
  no planned spike left, the line reads "next: none — run `marathon intake experiment.<topic>`".
  Accepting the brief authorizes the sync, and stages that goal when staging's checks pass,
  `exclusive` groups included (`mechanics/goals.md`), counting this goal's lock as released.
  When a check fails, the line reads "proposed, not staged: <reason>", and nothing is staged; a
  later `plan` stages it.

> ### Sync
>
> - **coordinator**
>   - context/ai.md: answer section for spike-harness-driver
>   - roadmap/index.toml: remove experiment.ai.spike-harness-driver from `active`
>   - roadmap/experiment.ai.spike-harness-driver.toml: delete
> - **architecture**
>   - standards/harnesses.md: name the driver interface
> - **next:** experiment.ai.spike-local-subagents · staged on accept

On a spike's last task, the brief gives the spike's answer: Summary states the question, the
one-line answer, and the numbered evidence, and Sync lists the answer section
([`answer-section.md`](./answer-section.md)) among the coordinator's edits.
