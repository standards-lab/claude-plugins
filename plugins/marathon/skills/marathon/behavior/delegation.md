# Delegation

A session orchestrates a task: it hands each unit of work to a subagent and holds only pointers to
the results, meaning the goal record, the task brief, and the commits. It still owns the work. It
reads each delegate's account, decides what follows, and reports the task as its own.

## The profiles

marathon ships five subagent profiles in the plugin's `agents/` directory. Each engagement starts
a fresh context.

| Profile | Stage | Loads | Does |
|---|---|---|---|
| **planner** | PLAN | the goal record, notes, roadmap, code, the currency output and the release notes of what trails | Drafts the rounds, then the task brief ([`agents/planner.md`](../../../agents/planner.md)) |
| **implementer** | BUILD | the slice brief, with the standards it quotes | Builds and commits one slice ([`agents/implementer.md`](../../../agents/implementer.md)) |
| **standards-reviewer** | BUILD | `STANDARDS.md`, the architecture pages it points to, the diff | Fixes the work to the standards ([`agents/standards-reviewer.md`](../../../agents/standards-reviewer.md)) |
| **spec-reviewer** | BUILD | the task brief, the diff, the running behavior | Returns the gaps against the brief ([`agents/spec-reviewer.md`](../../../agents/spec-reviewer.md)) |
| **editor** | BUILD | the goal record, notes, READMEs and docs, the diff | Keeps the context current and edits the changed prose ([`agents/editor.md`](../../../agents/editor.md)) |

- **The implementer never loads `STANDARDS.md` whole.** The session quotes the sections that bear
  on the slice, and the architecture pages they point to, into the slice brief
  (`references/briefs.md`, "Slice brief"). Whole standards cost the implementer context it needs
  for exploring, editing, and debugging; quoted lines cost less than a fix at standards review.
- **The reviewers stay apart.** Standards and spec run in separate contexts, so work that follows
  every convention but builds the wrong thing can't pass on its conventions.
- **No profile pins a model.** Every profile inherits the session's model, and the session doesn't
  announce one.

## Limits

These limits bind every delegate, and each profile cites them.

- A delegate takes one unit: one plan round or brief, one slice, one review, or one edit pass. A
  delegate that reaches beyond its unit, or can't finish it without something its brief doesn't
  cover, stops and reports back. The session decides whether to escalate (`references/build.md`).
  The planner returns a finding beyond its task as a note, not as part of the brief.
- A delegate recommends what the architect decides; it never decides it. A brief that is itself
  wrong goes back to the session as a question. The spec-reviewer judges against the brief, not
  against what it would have built.
- Only the implementer and the standards-reviewer commit, each as `mechanics/pipeline.md`,
  "Committing", says. No delegate publishes, merges, or writes the goal record's State or Tasks.
  Only the editor writes in the goal record: it drafts the Decisions and Pending edits, and the
  session commits them.
- The planner and the spec-reviewer change nothing. They run only commands that write nothing to
  the working tree, such as `git log`, `git diff`, the currency command, or a build or test.
- The standards-reviewer and the editor don't change what the task does. The standards-reviewer
  changes how the work is built; the editor changes how the prose is written, never what it says,
  and returns a passage whose meaning is unclear to the session with the question it raises.
- The editor edits prose only. It never edits a source file: code, tests, configuration, or their
  comments.
- The session does a unit itself when briefing a delegate would cost more than the work, such as a
  one-line fix. It never implements and reviews the same work itself.
