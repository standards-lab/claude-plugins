# Briefs

These are the formats marathon writes for the architect. Each one is read in the terminal, fits on
about one screen, and leads with what needs a decision. Agents look up facts themselves and never
ask the architect for them; the architect makes the decisions.

## Plan round

PLAN asks its questions in rounds. A round holds every question whose prerequisites are already
settled, and the architect answers them all at once, inline.

```
PLAN ROUND 2 · v1.ai · task harness-adapters
facts found: Pi has RPC mode (pi/docs/rpc.md);
             Claude Code: stream-json

1. Transport for all harnesses?
   rec: stdin/stdout JSONL — all three
        support it; SSE only in OpenCode
   changes: adapter count 3→1 protocol
2. Keep sessions resumable in v1?
   rec: no — out of scope; note in backlog
   changes: drops persistence slice

reply: "1 ok, 2 yes but ..."
```

- Number every question.
- Give every question a `rec:`, and a `changes:` line saying what the answer changes.
- Cite each fact with its source under `facts found:`.
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
```

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
- **Needs you** reads "(none)" when nothing remains.

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

## Spike intake

A spike intake is a plan round led by the spike's result:

```
INTAKE · spike-messaging → v1.messaging
question  one broker-agnostic event + reactor contract on JetStream and in-memory?
answer    yes — 8/8 evidence (README "The answer")
evidence  2, 3 by tests only; 7 waits on go-storage

1. Does go-core take reactor and event?
   rec: yes, as lifecycle components
   changes: go-core minor release before go-messaging
```
