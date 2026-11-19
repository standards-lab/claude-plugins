# Briefs

These are the formats marathon writes for the architect, plus the slice brief it writes for an
implementer. Each one the architect reads is read in the terminal, fits on about one screen, and
leads with what needs a decision. Agents look up facts themselves and never ask the architect for
them; the architect makes the decisions.

Every format is markdown, hard-wrapped at 100 characters in the reply, and code fences appear only
around code. Each template below is shown as a blockquote so it renders as it will in the reply;
the reply drops the `> `.

## Plan round

PLAN asks its questions in rounds. A round holds every question whose prerequisites are already
settled, and the architect answers them all at once, inline. A question that depends on the answer
to another open question waits for the next round; it is never asked conditionally, as in "only if
1 is yes". The round is presented as `behavior/planning.md`, "Ask for decisions, never for
facts", says.

> ## PLAN ROUND 1 · v1.ai · task harness-adapters
>
> **Facts found**
> - Pi has an RPC mode (pi/docs/rpc.md).
> - Claude Code streams JSON (stream-json).
>
> **Currency**
> - go-core
>   - go 1.24.2 → 1.25.1: bug fixes only
>   - pgx v5.6.0 → v6.0.0: Query takes options
>   - chi v5.0.12 → v5.2.0: held (Decisions)
> - go-web-service: current
>
> **Questions**
>
> 1. **Harness transport: stdin/stdout JSONL or SSE?**
>    - **rec:** stdin/stdout JSONL. Pi, Claude Code, and OpenCode all support it; only OpenCode
>      supports SSE.
>    - **changes:** the adapters share one protocol instead of three.
> 2. **go-core's pgx v6.0.0, a new major: upgrade now, hold, or give it its own task?**
>    - **rec:** its own task. v6's Query takes options, which reaches every caller.
>    - **changes:** adds task pgx-v6; holds pgx at v5.6.0.
> 3. **Session resume after a driver restart: build it in v1, or leave it out?**
>    - **rec:** leave it out of v1 and note it in the backlog.
>    - **changes:** drops the persistence slice.
>
> **reply:** "1 ok, 2 ok, 3 yes but ..."

- Number every question.
- Name each question's subject in its own words, with the options it chooses between, so the
  architect can decide it from the question, its **rec:**, and its **changes:** alone. All three
  name the things they refer to and never point elsewhere for them: not "this", "option B",
  "above", **Facts found**, another question, a term defined outside the question, or a file to
  open.
- Decide one thing per question. When two choices are coupled, ask the one that settles the
  other; the dependent choice waits for the next round. Never bundle two decisions into lettered
  combinations.
- Give every question a **rec:**, and a **changes:** line saying what the answer changes.
- Cite each fact with its source under **Facts found**.
- Round 1 carries a **Currency** list after **Facts found**, per repository the task touches: each
  trailing item as its currency command reports it, with a one-line summary of its release notes,
  or `current`, or `no currency command`, or the failed command as a fact. A held item is listed
  as a fact (`mechanics/goal-record.md`, Decisions).
- Ask about a trailing item only when it needs a decision: adapt, a breaking change the code must
  absorb; adopt, a new feature worth using; or a new major, to upgrade now, hold, or give its own
  task. A patch or minor bump with nothing to adapt stays a fact, and the upgrade slices take it.
- The rounds end when no question remains open. A BUILD escalation uses the same format, headed
  `ESCALATION` in place of `PLAN ROUND <n>` (`references/build.md`).

## Task brief

The task brief is what the architect approves at PLAN, and it is BUILD's contract. It describes
behavior, not procedure, and names no file paths, so it stays valid while other goals change the
code. The goal record keeps it (`mechanics/goal-record.md`).

> ## Task brief · experiment.ai.spike-harness-driver · harness-adapters
>
> **Problem**
>
> What is missing and why it matters now.
>
> **Behaviors**
> 1. Numbered, each testable.
>
> **Test seams**
> - The interfaces the tests exercise, ideally one.
>
> **Slices**
> 1. Ordered vertical slices, each demoable alone.
>
> **Out of scope**
>
> What this task will not do.
>
> **Door**
>
> Two-way, or one-way, and why.
>
> **Release**
>
> The exact tags the task releases (optional).

- **Slices** begin with upgrade slices when a repository the task touches trails: one per
  trailing repository, in the coordinator's `order`, each done when that repository's currency
  command exits 0, or reports only held items, with its check passing. The task's own slices
  build on them.
- **Door** is one-way when the task plans a release, since a pushed tag reaches people outside
  the repositories. The rest of the task may still be two-way, and the line says which part is
  which.
- **Release** appears only when the task releases. It lists each tag in its repository's own
  naming, such as `marathon/v0.16.1`, or, in a Go repository, `v1.4.0` at the root and
  `<submodule>/v1.4.0` for a sub-module, one tag per artifact. The plan rounds settle it; it is
  never inferred from a version bump, and no manifest field holds it. SHIP tags it
  (`mechanics/pipeline.md`, 6 · SHIP).

## Slice brief

The slice brief is what the session gives a fresh implementer for one slice at BUILD
(`references/build.md`, "The loop"). It is all the implementer reads before the code, so it
carries the standards that bear on the slice, quoted, in place of the whole `STANDARDS.md`
(`behavior/delegation.md`).

> ## Slice brief · experiment.ai.spike-harness-driver · harness-adapters · slice 2/3
>
> **Behaviors**
> 1. Numbered, each testable: what this slice adds.
>
> **Test seam**
>
> The interface the tests exercise.
>
> **Branches**
> - repository → branch, for each repository the slice touches
>
> **Check**
> - the check command, per repository
>
> **Standards**
> - the sections that bear on the slice, quoted, each with its source

- **Behaviors** come from the task brief, numbered as it numbers them.
- **Standards** quotes the sections of the repository's `STANDARDS.md`, and of the architecture
  pages it points to, that bear on the slice, each with its file and section. It always quotes
  the repository's rule on comments, or says the repository has none. A repository with no
  `STANDARDS.md` gets the line "no STANDARDS.md".

## Session brief

The session brief is the architect's view of the finished task, and the body of its pull request.
It opens with what changed, then how to verify each behavior: the architect judges outcomes, not
code. The session writes it to `.claude/briefs/<goal>.md` in the goal's root, which gitignores
`.claude/briefs/`.

> ## experiment.ai.spike-harness-driver · task 3/5 · harness-adapters
>
> ### Summary
>
> ```
> clutch run --harness=<pi|claude|opencode>
> adapters/ ─┬ pi/       (new)
>            ├ claude/   (new)
>            └ driver.go Session interface
> ```
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
> 1. **Every harness answers.** `clutch run --harness=pi "hi"` → reply streams, exit 0.
> 2. **A closed session refuses turns.** Tests only: `TestSendAfterClose`.
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
> **Merge danger:** two-way door · blast radius: the spike repository only
>
> **Needs you:** none

- **Summary** is the smallest visual that makes the change clear: CLI usage, a tree, or
  pseudocode, fenced as code.
- **Core changes** lists one to five entries, each with one line on why and, where the code is
  the change, a short excerpt or signature fenced as code.
- **Verify** opens with the check's one-line result. Then one item per task-brief behavior,
  numbered alike: the steps the architect can follow, outside the check and tests, and what
  they should see. The spec-reviewer followed each. "Tests only", naming the test, is the
  exception, when no such step exists.
- **Decided without you** lists the decisions the brief didn't cover and each Decision the
  editor dropped (`mechanics/goal-record.md`). A drop the architect rejects is a redirect, and
  the session restores the line.
- **Redirects and escalations** lists each BUILD escalation and each BRIEF redirect in the task,
  one line each with the architect's answer, or reads "none". Progress holds them
  (`mechanics/goal-record.md`).
- **Merge danger** names the door type and the blast radius. A one-way door gets the close read.
  When the task brief has a Release line, Merge danger repeats its tags, which accepting the
  brief authorizes (`mechanics/releasing.md`).
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
>   - roadmap.toml: remove experiment.ai.spike-harness-driver from `active`
>   - roadmap/experiment.ai.spike-harness-driver.toml: delete
> - **architecture**
>   - standards/harnesses.md: name the driver interface
> - **next:** experiment.ai.spike-local-subagents · staged on accept

On a spike's last task, the brief gives the spike's answer: Summary states the question and the
one-line answer, Verify carries the numbered evidence, and Sync lists the answer section among
the coordinator's edits.

## Answer section

A spike's answer section is the coordinator edit its sync lands in the note its experiment cites
in `context`, one section per spike, under the `## Answers · experiment.<topic>` heading
`experiment` writes there (`commands/experiment.md`). The spike's last task writes the answer
under "The answer" in the spike's `context/README.md`, and adds this section to its goal
record's pending edits. The example is a note's file content, so it stays fenced.

```markdown
### Answer · experiment.ai.spike-harness-driver

**Question:** Can one driver interface run Pi, Claude Code, and OpenCode sessions?

**Answer:** Yes; stdin/stdout JSONL carries all three.

1. One `Session` interface drives all three harnesses. Proven by a running demo:
   `clutch demo adapters`.
2. Sessions survive a restart. Proven by the validate task `resume-sessions`.
3. OpenCode's tool calls round-trip. Proven only by tests: `adapters/opencode`.

[The answer](https://github.com/example/spike-harness-driver/blob/main/context/README.md#the-answer)
· [spike-harness-driver](https://github.com/example/spike-harness-driver)
```

- **The question** is the spike's, as its sub-goal states it.
- **The answer** is one line.
- **The evidence** is numbered, and each item names where it is proven: a test, a validate task,
  or a running demo. An item proven only by tests says so.
- **The links** point to where the spike states its answer, "The answer" in its
  `context/README.md`, and to the spike's remote.

`marathon intake` later folds each answer section into the served goals' notes and removes it
(`commands/intake.md`).

## Status digest

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

## Goal status

`status <goal>` prints the goal's tasks in `order`, then one line per sub-goal.

> ## Status · v1.data · 2026-10-01
>
> | task | done |
> |---|---|
> | reads | x |
> | writes | |
>
> - v1.data.cache: 0/2

## Retro

`retro` groups its findings under one heading per layer that should catch them next time, one
checkbox each with a proposed change, and the architect ticks the ones to apply.

> ## Retro · goal v1.storage · 2026-09-14..09-24
>
> transcripts missing: go-core
>
> ### design
> - [ ] 1. new goal v1.config: one typed config loader (go-web-service)
>
> ### check/eval
> - [ ] 2. depguard rule: providers only in internal/app (go-web-service)
>
> ### standard
> - [ ] 3. STANDARDS.md: one contract, one home — doc.go inventory
>
> ### skill/profile
> - [ ] 4. implementer brief: name the check command explicitly
>
> ### context
> - [ ] 5. cull blobfs-composition.md — expressed in blobfs/docs

## Intake round

`intake` runs plan rounds led by the spikes' answers (`commands/intake.md`). The header names the
experiment and the goals it serves, then summarizes each spike's answer section under its name:
its question, its one-line answer, and its evidence. The numbered questions follow, in the plan
round's form.

> ## INTAKE ROUND 2 · experiment.ai → v1.ai, v1.agents
>
> **spike-harness-driver**
> - **question:** one driver interface over Pi, Claude Code, and OpenCode?
> - **answer:** yes; stdin/stdout JSONL carries all three
> - **evidence:** 3 items; item 3 by tests only
>
> **spike-local-subagents**
> - **question:** do local models hold a subagent's tool loop?
> - **answer:** no; tool calls drift past 20 turns
> - **evidence:** 4 items; item 2 waits on a larger model
>
> **Questions**
>
> 1. **Does v1.ai build its first task, harness-driver, on the driver interface?**
>    - **rec:** yes, as its first task.
>    - **changes:** adds a task to v1.ai; folds the answer into context/ai.md.
> 2. **Does v1.agents keep its local-subagents task?**
>    - **rec:** no; cloud models only until a larger local model is tested.
>    - **changes:** drops v1.agents' local-subagents task.
> 3. **Archive spike-local-subagents' remote, or keep it open?**
>    - **rec:** archive it; keep spike-harness-driver open as v1.ai's reference.
>    - **changes:** one remote archived after the merge.
>
> **reply:** "1 ok, 2 ok, 3 yes but ..."

- The first round names the served goals, which the header lists from then on, and
  consolidates the spikes' answer sections.
- Each spike's summary comes from its answer section, with any evidence proven only by tests,
  or still waiting, called out.
- Every question has a **rec:** and a **changes:** line, as in a plan round.
