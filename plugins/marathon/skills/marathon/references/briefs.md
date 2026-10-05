# Briefs

These are the formats marathon writes for the architect. Each one is read in the terminal, fits on
about one screen, and leads with what needs a decision. Agents look up facts themselves and never
ask the architect for them; the architect makes the decisions.

## Plan round

PLAN asks its questions in rounds. A round holds every question whose prerequisites are already
settled, and the architect answers them all at once, inline.

```
PLAN ROUND 1 · v1.ai · task harness-adapters
facts found: Pi has RPC mode (pi/docs/rpc.md);
             Claude Code: stream-json
currency:    go-core
               go 1.24.2 -> 1.25.1  bug fixes only
               pgx v5.6.0 -> v6.0.0  Query takes options
               chi v5.0.12 -> v5.2.0  held (Decisions)
             go-web-service: current

1. Transport for all harnesses?
   rec: stdin/stdout JSONL — all three
        support it; SSE only in OpenCode
   changes: adapter count 3→1 protocol
2. go-core: pgx v6.0.0, a new major?
   rec: its own task — the Query change
        reaches every caller
   changes: adds task pgx-v6; holds pgx
3. Keep sessions resumable in v1?
   rec: no — out of scope; note in backlog
   changes: drops persistence slice

reply: "1 ok, 2 ok, 3 yes but ..."
```

- Number every question.
- Give every question a `rec:`, and a `changes:` line saying what the answer changes.
- Cite each fact with its source under `facts found:`.
- Round 1 carries a `currency:` block beside `facts found:`, per repository the task touches:
  each trailing item as its currency command reports it, with a one-line summary of its release
  notes, or `current`, or `no currency command`, or the failed command as a fact. A held item is
  listed as a fact (`mechanics/goal-record.md`, Decisions).
- Ask about a trailing item only when it needs a decision: adapt, a breaking change the code must
  absorb; adopt, a new feature worth using; or a new major, to upgrade now, hold, or give its own
  task. A patch or minor bump with nothing to adapt stays a fact, and the upgrade slices take it.
- The rounds end when no question remains open. A BUILD escalation uses the same format, headed
  `ESCALATION` (`references/build.md`).

## Task brief

The task brief is what the architect approves at PLAN, and it is BUILD's contract. It describes
behavior, not procedure, and names no file paths, so it stays valid while other goals change the
code. The goal record keeps it (`mechanics/goal-record.md`).

```
## Task brief · experiment.ai.spike-harness-driver · harness-adapters
Problem       what is missing and why it matters now
Behaviors     numbered, each testable
Test seams    the interfaces the tests exercise (ideally one)
Slices        ordered vertical slices, each demoable alone
Out of scope  what this task will not do
Door          two-way | one-way (why)
Release       the exact tags the task releases (optional)
```

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

## Session brief

The session brief is the architect's view of the finished task, and the body of its pull request.
The diff is already committed when the architect reads it and isn't visible in the terminal, so
Core changes carries the code that matters. The session writes it to `.claude/briefs/<goal>.md` in the
goal's root, which gitignores `.claude/briefs/`.

```
## experiment.ai.spike-harness-driver · task 3/5 · harness-adapters

### Summary
  clutch run --harness=<pi|claude|opencode>
  adapters/ ─┬ pi/       (new)
             ├ claude/   (new)
             └ driver.go Session interface

### Core changes
  1. driver.Session — one interface over every harness
       type Session interface { Send(ctx, Turn) (Reply, error); Close() error }
     why: adapters differ only in transport
  2. pi.New(cfg) (*Session, error) — RPC over stdin/stdout JSONL
     why: matches plan round 2, question 1

### Evidence
  mise run check        ✓ 214 tests, lint, split
  clutch demo adapters  → transcript excerpt

### Decided without you
  1. Kept stdin JSONL over SSE (simpler, spike)
  2. Standards-reviewer: 4 fixes, 2 tautological tests cut

### Merge danger
  two-way door · blast radius: spike repo only

### Needs you
  (none)
```

- **Summary** is the smallest visual that makes the change clear: CLI usage, a tree, or
  pseudocode.
- **Core changes** lists one to five entries. Each gives a short excerpt or a before-and-after
  signature, and one line on why.
- **Evidence** comes from commands the architect can rerun.
- **Decided without you** lists the decisions the brief didn't cover and what the reviewers
  changed.
- **Merge danger** names the door type and the blast radius. A one-way door gets the close read.
  When the task brief has a Release line, Merge danger repeats its tags, so accepting the brief
  authorizes tagging them.
- **Needs you** reads "(none)" when nothing remains.
- **Sync** appears only on the goal's last task, after Needs you. It lists the pending edits the
  sync will apply, per repository, and the goal proposed to stage next: for a spike, the next
  planned spike of its experiment; otherwise the top of `planned`. When a spike's experiment has
  no planned spike left, the line reads "next: none — run `marathon intake experiment.<topic>`".
  Accepting the brief authorizes the sync, and stages that goal when staging's checks pass,
  `exclusive` groups included (`mechanics/goals.md`), counting this goal's lock as released.
  When a check fails, the line reads "proposed, not staged: <reason>", and nothing is staged; a
  later `plan` stages it.

```
### Sync
  coordinator   context/ai.md: answer section for spike-harness-driver
                roadmap.toml: remove experiment.ai.spike-harness-driver
  architecture  standards/harnesses.md: name the driver interface
  next          experiment.ai.spike-local-subagents · staged on accept
```

On a spike's last task, the brief gives the spike's answer: Summary states the question and the
one-line answer, Evidence carries the numbered evidence, and Sync lists the answer section among
the coordinator's edits.

## Answer section

A spike's answer section is the coordinator edit its sync lands in the note its experiment cites
in `context`, one section per spike, under the `## Answers · experiment.<topic>` heading
`experiment` writes there (`commands/experiment.md`). The spike's last task writes the answer
under "The answer" in the spike's README, and adds this section to its goal record's pending
edits.

```markdown
### Answer · experiment.ai.spike-harness-driver

**Question:** Can one driver interface run Pi, Claude Code, and OpenCode sessions?

**Answer:** Yes; stdin/stdout JSONL carries all three.

1. One `Session` interface drives all three harnesses. Proven by a running demo:
   `clutch demo adapters`.
2. Sessions survive a restart. Proven by the validate task `resume-sessions`.
3. OpenCode's tool calls round-trip. Proven only by tests: `adapters/opencode`.

[The answer](https://github.com/example/spike-harness-driver#the-answer) ·
[spike-harness-driver](https://github.com/example/spike-harness-driver)
```

- **The question** is the spike's, as its sub-goal states it.
- **The answer** is one line.
- **The evidence** is numbered, and each item names where it is proven: a test, a validate task,
  or a running demo. An item proven only by tests says so.
- **The links** point to where the spike states its answer, its README's "The answer" for a
  spike `plan` set up, and to the spike's remote.

`marathon intake` later folds each answer section into the served goals' notes and removes it
(`commands/intake.md`).

## Status digest

`status` prints one line per active goal: the goal, its tasks done out of its total, its state,
and the repositories it locks. Goals that need the architect sort first and carry a `!`.

```
── STATUS · 2026-10-01 ──
! experiment.ai.spike-harness-driver    2/3  brief ready  spike-harness-driver
  experiment.ai.spike-local-subagents  1/3  building     spike-local-subagents
  factory                              1/3  building     claude-plugins
  v1.messaging                         0/1  idle         go-messaging, go-core, go-web-service
```

The state is the goal record's State line (`mechanics/goal-record.md`). A goal with no record
yet shows `no record`.

## Retro

`retro` reviews a task, a goal, or a date range. It groups its findings by the layer that should
catch them next time, one line each with a proposed change, and the architect ticks the ones to
apply.

```
RETRO · goal v1.storage · 2026-09-14..09-24
check/eval
  [ ] 1. depguard rule: providers only in internal/app (go-web-service)
standard
  [ ] 2. STANDARDS.md: one contract, one home — doc.go inventory
skill/profile
  [ ] 3. implementer brief: name the check command explicitly
context
  [ ] 4. cull blobfs-composition.md — expressed in blobfs/docs
```

## Intake round

`intake` runs plan rounds led by the spikes' answers (`commands/intake.md`). The header names the
experiment and the goals it serves, then summarizes each spike's answer section: its question,
its one-line answer, and its evidence. The numbered questions follow, in the plan round's form.

```
INTAKE ROUND 2 · experiment.ai → v1.ai, v1.agents
spike-harness-driver
  question  one driver interface over Pi, Claude Code, and OpenCode?
  answer    yes; stdin/stdout JSONL carries all three
  evidence  3 items; item 3 by tests only
spike-local-subagents
  question  do local models hold a subagent's tool loop?
  answer    no; tool calls drift past 20 turns
  evidence  4 items; item 2 waits on a larger model

1. Does v1.ai build on the driver interface?
   rec: yes, as its first task, harness-driver
   changes: adds a task to v1.ai; folds the answer into context/ai.md
2. Does v1.agents keep local subagents?
   rec: no; cloud models only until a larger local model is tested
   changes: drops v1.agents' local-subagents task
3. Archive spike-local-subagents?
   rec: yes; keep spike-harness-driver open as v1.ai's reference
   changes: one remote archived after the merge

reply: "1 ok, 2 ok, 3 yes but ..."
```

- The first round names the served goals, which the header lists from then on, and
  consolidates the spikes' answer sections.
- Each spike's summary comes from its answer section, with any evidence proven only by tests,
  or still waiting, called out.
- Every question has a `rec:` and a `changes:` line, as in a plan round.
