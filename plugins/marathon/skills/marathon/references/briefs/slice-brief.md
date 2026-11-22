# Slice brief

The rules every format shares are in [`../briefs.md`](../briefs.md).

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

- **Behaviors** come from the task brief ([`task-brief.md`](./task-brief.md)), numbered as it
  numbers them.
- **Standards** quotes the sections of the repository's `STANDARDS.md`, and of the architecture
  pages it points to, that bear on the slice, each with its file and section. It always quotes
  the repository's rule on comments, or says the repository has none. A repository with no
  `STANDARDS.md` gets the line "no STANDARDS.md".
