# marathon briefs: what the architect reads

This note defines every output marathon produces for the architect. Each one is read in the
terminal, fits on about one screen, and leads with what needs a decision. Agents look up facts
themselves and never ask the architect for them. Decisions belong to the architect. The
pipeline that produces these formats is `marathon-factory.md`.

## Plan round

A plan session asks its questions in rounds. Each round holds every question whose prerequisites
are already settled. The architect answers inline, all at once.

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
- Give every question a `rec:`, and a `changes:` saying what the answer changes.
- Cite facts with their source under `facts found:`.
- The rounds end when no open question remains. An escalation during BUILD uses the same format.

## Task brief

The brief is what the architect approves, and it becomes the contract for BUILD. It describes
behavior, not procedure. It names no file paths, so it stays valid while other goals change the
code.

```
## Task brief · v1.ai.experiment · harness-adapters
Problem       what is missing and why it matters now
Behaviors     numbered, each testable
Test seams    the interfaces the tests exercise (ideally one)
Slices        ordered vertical slices, each demoable alone
Out of scope  what this task will not do
Door          two-way | one-way (why)
```

## Session brief

The session brief is the architect's view of the finished task, and the PR body. The diff is
already committed by the time the architect reads it, and isn't visible in the terminal, so Core
changes carries the code that matters.

```
## v1.ai.experiment · task 3/5 · harness-adapters

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
  2. Reviewer: 4 fixes, 2 tautological tests cut

### Merge danger
  two-way door · blast radius: spike repo only

### Needs you
  (none)
```

- **Summary** is the smallest visual that makes the change clear: CLI usage, a tree, or pseudocode.
- **Core changes** lists 1–5 entries. Each gives a short excerpt or a before-and-after
  signature, and one line on why.
- **Evidence** comes from commands the architect can rerun.
- **Merge danger** names the door type and the blast radius. A one-way door gets the close read.
- **Needs you** reads "(none)" when nothing remains.

## Status digest

`status` prints one line per active goal. Items that need the architect sort first and carry a
`!`. The digest reads only the roadmap's `active` array and the goal records.

```
── STATUS · 2026-10-01 ──
! v1.ai.experiment  1/4  brief ready  spike-harness-driver
  factory           1/3  building     claude-plugins
  v1.messaging      1/1  intake next  standards-lab
```

## Retro

`retro` reviews a session, a goal, or a date range. Its findings are grouped by the layer that
should catch them next time (`marathon-factory.md`, "The retro"), one line each with a proposed
change. The architect ticks the findings to apply.

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

The retro includes the context-drift pass that `review` ran in 0.15. It covers notes the code
now expresses or contradicts.

## Spike intake

Spike intake takes in a finished spike. It uses the plan-round form, led by the spike's result:

```
INTAKE · spike-messaging → v1.messaging
question  one broker-agnostic event + reactor contract on JetStream and in-memory?
answer    yes — 8/8 evidence (README "The answer")
evidence  2, 3 by tests only; 7 waits on go-storage

1. Does go-core take reactor and event?
   rec: yes, as lifecycle components
   changes: go-core minor release before go-messaging
```
