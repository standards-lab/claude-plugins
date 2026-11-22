# goal · marathon.repo-access

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] ci-awareness
2. [ ] review-findings
3. [ ] shared-repos

## Decisions

- goal: one release, at shared-repos: `marathon/v0.23.0` and `marathon-architecture/v0.4.1`;
  ci-awareness and review-findings tag nothing.
- shared-repos: moves claude-plugins' own configuration to `[remote] checks` with the release,
  since the installed 0.22.0 copy and binary don't know the key.

## Pending edits

- coordinator · roadmap/marathon.repo-access.toml: add task `read-first-split` before
  `review-findings` in `order`: split `references/briefs.md` into one file per format, each
  command's Read first list naming only the formats it writes; a check that every command lists
  each format it writes; re-baselined budgets; the gate's eval run; restore the sentences
  ci-awareness cut for budget where the room allows.
