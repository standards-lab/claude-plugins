# Task brief

The rules every format shares are in [`../briefs.md`](../briefs.md).

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
  the repositories. The line says which parts of the task are one-way.
- **Release** appears only when the task releases. It lists each tag in its repository's own
  naming, such as `marathon/v0.16.1`, or, in a Go repository, `v1.4.0` at the root and
  `<submodule>/v1.4.0` for a sub-module, one tag per artifact. The plan rounds
  ([`plan-round.md`](./plan-round.md)) settle it; it is never inferred from a version bump, and no
  manifest field holds it. SHIP tags it (`mechanics/pipeline.md`, 6 · SHIP).
