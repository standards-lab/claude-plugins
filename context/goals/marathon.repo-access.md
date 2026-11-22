# goal · marathon.repo-access

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] ci-awareness
2. [x] read-first-split
3. [ ] review-findings
4. [ ] shared-repos

## Decisions

- goal: one release, at shared-repos: `marathon/v0.23.0` and `marathon-architecture/v0.4.1`;
  ci-awareness and review-findings tag nothing.
- shared-repos: moves claude-plugins' own configuration to `[remote] checks` with the release,
  since the installed 0.22.0 copy and binary don't know the key.
- read-first-split: a read-first budget sits at its set's size rounded up to the next hundred, and
  rises only when the architect approves, with the reason in the CHANGELOG's Unreleased entry.

## Pending edits

- coordinator · roadmap/marathon.repo-access.toml: end the `review-findings` summary's first
  sentence with "and the stale read-first sizes in claude-plugins' `harness-optimization.md`".
