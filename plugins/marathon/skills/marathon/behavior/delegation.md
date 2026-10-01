# Delegation

A session orchestrates a task: it hands each unit of work to a subagent and holds only pointers to
the results, meaning the goal record, the task brief, and the commits. It still owns the work. It
reads each delegate's account, decides what follows, and reports the task as its own.

## The profiles

marathon ships five subagent profiles in the plugin's `agents/` directory. Each engagement starts
a fresh context.

| Profile | Stage | Loads | Does |
|---|---|---|---|
| **planner** | PLAN | the goal record, notes, roadmap, code | Finds the facts, drafts the round's questions, then the task brief and slices; changes nothing |
| **implementer** | BUILD | the slice brief and the check command only | Makes one slice work with the check passing, and commits it |
| **standards-reviewer** | BUILD | `STANDARDS.md`, the architecture pages it points to, the diff | Applies the standards, sweeps for tests that lie, and commits its fixes |
| **spec-reviewer** | BUILD | the task brief, the diff, the running behavior | Returns where the work is missing, wrong, or out of scope; changes nothing |
| **editor** | BUILD | the goal record, notes, READMEs and docs, the diff | Keeps the context current and edits the changed prose; never commits |

- **The implementer never loads standards or architecture.** Implementing fills a context window
  with exploring, editing, and debugging. Review is a lighter load, so the conventions go there.
- **The reviewers stay apart.** Standards and spec run in separate contexts, so work that follows
  every convention but builds the wrong thing can't pass on its conventions.
- **No profile pins a model.** Every profile inherits the session's model, and the session doesn't
  announce one.

## Limits

- A delegate takes one unit: one plan round or brief, one slice, one review, or one edit pass. A
  delegate that reaches beyond its unit reports back, and the session decides whether to escalate
  (`references/build.md`).
- Only the implementer and the standards-reviewer commit, and each first confirms the checkout is
  on the task's branch. No delegate publishes, merges, or writes the goal record's state; the
  editor drafts the record's content and the session commits it.
- The session does a unit itself when briefing a delegate would cost more than the work, such as a
  one-line fix. It never implements and reviews the same work itself.
