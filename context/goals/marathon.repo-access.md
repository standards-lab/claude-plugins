# goal · marathon.repo-access

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] ci-awareness
2. [x] read-first-split
3. [x] review-findings
4. [x] shared-repos

## Decisions

- goal: one release, at shared-repos: `marathon/v0.23.0` and `marathon-architecture/v0.4.1`;
  ci-awareness and review-findings tag nothing.
- shared-repos: SHIP waits on CI by hand before merging, since the installed 0.22.0 copy and
  binary don't run the `[remote] checks` key this task adds to claude-plugins' configuration.

## Pending edits

- coordinator · roadmap/v1.harness.toml: `door-scaled-review`'s `context` becomes
  `claude-plugins/context/harness-findings.md`.
- coordinator · context/waves.md: wave B's sentence on v1.ai.hosting sharing "the architecture
  lock" with v1.ai.processing, freed by the architect, becomes: both goals hold the architecture
  repository only while a task's branch is checked out there.
- coordinator · docs/ai-hosting.md: "The goal's coordination"'s sentence on the `architecture`
  task sharing "the architecture lock" with v1.ai.processing, freed by the architect, becomes:
  both goals hold the architecture repository only while a task's branch is checked out there.
