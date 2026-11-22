# goal · marathon.repo-access

- **State:** building
- **Task:** ci-awareness
- **Branch:** ci-awareness

## Tasks

1. [ ] ci-awareness
2. [ ] review-findings
3. [ ] shared-repos

## Task brief · ci-awareness

**Problem**

marathon can't tell whether a repository runs CI. Each repository writes the wait for checks
into its `merge` command by hand, and in a repository without CI that form fails, because
`gh pr checks` exits 1 when no checks are reported. Separately, sessions that put several git or
remote operations into one tool call fail more often than the architect accepts. Until
`scripts/currency.sh` is fixed, claude-plugins' currency check falsely reports trailing tools.

**Behaviors**
1. `marathon currency` in claude-plugins exits 0 when nothing trails. An entry whose bump is null,
   or whose latest version equals the requested one, is not reported. Both workflows use
   jdx/mise-action v5.3.0.
2. `[remote] checks` is an optional configuration key: the command that waits for a published
   pull request's checks. `marathon checks` runs it the way it runs the other keys, and reports it
   as unset when it isn't configured.
3. At SHIP, and at every merge in "Branches and pull requests" and releasing, a repository with
   `checks` set runs `marathon checks` before `marathon merge`. A failing check goes to `handoff`
   as it does today. A repository without `checks` merges directly.
4. The configuration reference's example splits the wait out: `checks = "gh pr checks --watch"`
   and `merge = "gh pr merge --merge --delete-branch"`. `[remote] ci` stays the wait after the
   merge, which only releasing runs.
5. init's Remote decision asks whether the repository runs CI on pull requests, and writes
   `checks` when it does.
6. The session brief's Merge danger line says, for each repository, whether its merge waits on CI.
   The format and its example show this.
7. A "Running commands" rule replaces "Committing". Every command that changes git or remote
   state runs in its own tool call, and the next one runs only after the previous one succeeds.
   That covers commit, branch, checkout or switch, merge, rebase, tag, push, `marathon publish`,
   `marathon merge`, `marathon checks`, and `gh` writes. Read-only commands may still be chained.
   The existing commit rules (confirm the branch first, stage explicit paths) stay. Every agent
   profile and delegation.md's Limits point to the rule.
8. The CHANGELOG's Unreleased section records the change, with a "Migrating from 0.22.0" step for
   repositories with CI: move `gh pr checks --watch` out of `merge` and into `checks`.

**Test seams**
- The marathon CLI's key resolution and command run, through its existing tests.
- `scripts/currency.sh`, run against the repository.
- `scripts/check.sh`, which resolves the skill's pointers.

**Slices**
1. Upgrade: bump jdx/mise-action to v5.3.0 in both workflows and fix `scripts/currency.sh`'s
   filter (CLP-1). Done when `marathon currency` exits 0 and the check passes.
2. CLI: the `checks` key and `marathon checks` (behavior 2).
3. Skill: configuration, SHIP, "Branches and pull requests", releasing, init, the session brief's
   Merge danger line, and the CHANGELOG's Unreleased section (behaviors 3–6 and 8).
4. The "Running commands" rule and the pointers to it (behavior 7).

**Out of scope**

- claude-plugins' own `.claude/marathon.toml` and its `mise.toml` pin: shared-repos migrates them
  with the release.
- The other repositories' configurations: each applies the Migrating step.
- Any version bump or tag: the goal releases once, at shared-repos.
- CLP-2 through CLP-15 (task review-findings), and the repository lock (task shared-repos).

**Door**

Two-way. The task tags nothing, and installed copies don't change until the goal's release.

## Progress

slices 0/4 committed · standards — · spec — · editor —

## Decisions

- ci-awareness: CI is declared by an optional `[remote] checks` command, not a boolean; a boolean
  would make marathon know each platform's wait.
- ci-awareness: the other repositories migrate through the CHANGELOG's "Migrating from 0.22.0"
  step, not as goal repos, so the goal locks only claude-plugins.
- goal: one release, at shared-repos: `marathon/v0.23.0` and `marathon-architecture/v0.4.1`;
  ci-awareness and review-findings tag nothing.
- shared-repos: moves claude-plugins' own configuration to `[remote] checks` with the release,
  since the installed 0.22.0 copy and binary don't know the key.

## Pending edits

None.
