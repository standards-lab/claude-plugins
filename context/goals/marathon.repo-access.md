# goal · marathon.repo-access

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] ci-awareness
2. [ ] read-first-split
3. [ ] review-findings
4. [ ] shared-repos

## Decisions

- goal: one release, at shared-repos: `marathon/v0.23.0` and `marathon-architecture/v0.4.1`;
  ci-awareness and review-findings tag nothing.
- shared-repos: moves claude-plugins' own configuration to `[remote] checks` with the release,
  since the installed 0.22.0 copy and binary don't know the key.

## Pending edits

None.
