# Build

BUILD turns an approved task brief into committed, reviewed work without stopping for the
architect. The session orchestrates: it engages one profile per unit (`behavior/delegation.md`)
and holds only pointers to the goal record, the brief, and the commits.

## Checks come first

Quality comes in three layers, cheapest first: automated checks, automated review, then the
architect's review. Each layer exists to make the next one lighter, so the architect's attention
goes to one-way doors.

- **The check.** Each repository exposes one deterministic check command, named by `[project]
  check` in `.claude/marathon.toml` (`mechanics/configuration.md`), such as `mise run check` or
  `scripts/check.sh`. It runs the build, vet, tests, lint, and the repository's own tools. A
  mechanical rule becomes part of the check rather than prose in `STANDARDS.md`.
- **Standards.** A repository's `STANDARDS.md` holds judgement calls only, with pointers to the
  architecture pages that apply. Only the standards-reviewer reads it. `CLAUDE.md` holds
  navigation pointers only.
- **Tests can lie.** A green check proves nothing when a test is tautological,
  structure-sensitive, or unable to fail. The standards-reviewer sweeps for all three. The design
  cure is deep modules: small interfaces over large implementations, tested only at their seams.

## Slices

A slice is an ordered vertical piece of the task that can be demonstrated on its own. The task
brief lists them, lowest dependency first. A slice leaves each repository it touches with its
check passing.

- On a **code** project, a slice adds behavior through the test seam, with its tests.
- On a **context** project, a slice is the smallest set of files that must change together, such
  as a playbook and the reference it cites, and its check is the consistency script, where one
  exists, plus a read for coherence.

A task that spans repositories orders its slices by the coordinator's `order`, lowest layer
first, so each higher layer builds against the real change below it.

## The loop

1. **Implement.** For each slice in order, engage a fresh implementer with the slice brief, the
   branches, and the check command. It commits the slice with the check passing.
2. **Review standards.** Once every slice is committed, engage the standards-reviewer on the whole
   diff. It commits its fixes with the check passing.
3. **Review the spec.** Engage the spec-reviewer with the task brief. For each gap it returns,
   engage an implementer to close it, then engage the spec-reviewer again. After two rounds that
   leave the same gap open, escalate.
4. **Edit.** Engage the editor once. It updates the goal record, tends the notes, and edits the
   changed prose. Read its edits and commit them.
5. **Validate.** Run the check in every touched repository, plus whatever the spec-reviewer ran to
   show each behavior. The results are the session brief's Evidence. Don't go on to BRIEF with a
   failing check.

The session records progress in the goal record as slices commit, so a handoff loses nothing.

## Escalation

BUILD stops for the architect only for:

- **a one-way door**: an irreversible migration, data loss, or an action with effects outside the
  repositories
- **a decision the task brief doesn't cover**, which a delegate couldn't settle without changing
  the brief
- **scope beyond the task**

An escalation uses the plan-round format, headed `ESCALATION` (`references/briefs.md`). BUILD
resumes once it is answered. Anything else a delegate decides goes under Decided without you.

## Handoff

When the session's context fills before the task reaches BRIEF, it runs `reset` on its own: it
writes the handoff state into the goal record and stops (`commands/reset.md`). The next `start`
on the goal resumes BUILD from there without asking the architect.

## Stay within the task

Make no opportunistic refactors or unrelated cleanups. When a change outside the task looks
worthwhile, record it as a note, or as a pending roadmap edit in the goal record, and leave it.
