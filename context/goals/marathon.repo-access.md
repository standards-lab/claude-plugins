# goal · marathon.repo-access

- **State:** building
- **Task:** shared-repos
- **Branch:** shared-repos

## Tasks

1. [x] ci-awareness
2. [x] read-first-split
3. [x] review-findings
4. [ ] shared-repos

## Task brief · shared-repos

**Problem**

Today an active goal locks every repository in its `repos` for its whole life. So two goals can't
share a library even when their tasks never overlap in time, and the coordinator already runs two
active goals that both list the architecture repository, kept apart by hand. Git's checked-out
branch already shows when a session is mid-work in a repository. The lock should follow that branch
state rather than the goal's lifetime, and the goal ends by releasing the plugins it changed.

**Behaviors**
1. **Staging.** Staging fails only when another active goal declares the same root, or when
   `[workspace] exclusive` would be broken (at most one active goal whose `repos` touch a group). A
   goal's `repos` may include another active goal's root or non-root repositories.
2. **When a repository is held.** Any repository, root included, is held while it is off its
   default branch: on a task branch, `<slug>-fix`, `sync-<goal>`, or a session branch. It is also
   held while an active goal's Handoff reads "tag <names>" and names it short of released.
   Otherwise it is free.
3. **LOCATE.** LOCATE accepts a checked-out branch as the goal's own only when the goal's record is
   mid-task on it: State building, brief ready, or handoff, with Branch equal to the task slug, or
   the release's `<slug>-fix`, or the goal's `sync-<goal>`. Any other branch means the repository
   is held: LOCATE stops and names each holder (each active goal whose record is mid-task on that
   branch name, or the goal or topic a session branch names). A waiting goal never skips ahead to
   a later task and never polls.
4. **PLAN, non-root held.** Before creating any branch, PLAN rechecks every repository the task
   touches. If a non-root repository is held, PLAN commits the approved brief to the root's record
   as bookkeeping with State idle, reports the holder, and stops; the next `start` takes the stored
   brief to approval without new rounds.
5. **PLAN, root held.** If the root itself is held, PLAN keeps the approved brief in the plan file,
   reports the holder, and stops; the next `start` continues from the plan file to approval.
6. **Changes from sync, retro, and plan.** A change to an active goal's root becomes a task of that
   goal. A change to any other repository goes on the session's own branch while that repository
   is on its default branch. When it is held, retro and plan stop and report; sync applies its
   other edits and goes to handoff with the next move "sync <repo>, then delete the record".
7. **Status digest.** The `locks` column lists the goal's root plus each other repository currently
   checked out on the goal's branch.
8. **The last task's Sync line** stages the proposed goal under behavior 1, counting the finishing
   goal's root as released.
9. **Architecture layer.** A page task from the architecture layer belongs to a goal rooted in the
   architecture repository. marathon-architecture still requires marathon ≥ 0.19.0.
10. **claude-plugins' own configuration** declares its pull-request CI wait as `[remote] checks`, and
    its merge command no longer waits. This task's own merge waits on CI by hand, since the
    installed 0.22.0 copy can't run that key.
11. **CHANGELOGs and versions.** Each plugin's CHANGELOG records the changes under its release
    heading, and every version marker agrees: marathon 0.23.0, marathon-architecture 0.4.1. The
    planned-work note for repository access is deleted, and the capability map lists repository
    access as built, linking where the lock rule now lives.
12. **Word budgets.** Every read-first set stays within budget; a set the rewrite grows is raised by
    at most 100 words, and the session brief lists each raise.

**Test seams**
- The repository's consistency check: versions, CHANGELOG headings, links and citations, word
  budgets, the formats each playbook must list, and the context lint.
- A coherence read across the playbooks: no statement of the lifetime lock remains, and each rule
  is stated once and cited elsewhere.
- The release gate per tag: the check, plus for marathon the dist build and the eval suite.

**Slices**
1. **The lock rule.** The goals mechanics' lock and staging (behaviors 1, 2); the manifest's `repos`
   and `root`, `exclusive` in the configuration reference, workspace coordination, the goal
   record's "where it is committed"; plan's staging check, the session brief's Sync staging check
   (behavior 8), the plugin README and description. CHANGELOG entry. Check passes.
2. **LOCATE and PLAN.** LOCATE's lock step points to the rule and says how it accepts a branch and
   names holders (behavior 3); PLAN's recheck and brief storage (behaviors 4, 5); resume's sync
   wording. CHANGELOG entry. Budgets within behavior 12. Check passes.
3. **Routing.** Sync's step for other repositories with its new handoff move; retro's Apply rules;
   plan's rule for member-repository changes; the marathon-architecture page-task sentence. Both
   CHANGELOG entries. Check passes.
4. **Status.** The `locks` column in the status playbook and the digest format, including its
   example. CHANGELOG entry. Check passes.
5. **Context.** Delete the repository-access note; mark its capability map entry built. Check passes.
6. **Release preparation.** claude-plugins' configuration moves the CI wait to `checks`; marathon's
   Unreleased becomes `v0.23.0`, marathon-architecture's becomes `v0.4.1`; every version marker
   agrees. Check passes.

**Out of scope**
- A binary subcommand or tooling-layer CLI for finding holders.
- A marker naming the holder on disk, goal-qualified branch names, unique slugs across goals.
- Any change to what `exclusive` means; raising marathon-architecture's floor.
- Editing the coordinator directly (pending edits, applied at sync).
- Reinstalling the installed marathon copy (the architect, after the release).

**Door**

One-way for the two pushed tags and marathon's released binaries, which reach people who install
the plugins. Everything before tagging is two-way.

**Release**

`marathon/v0.23.0`, `marathon-architecture/v0.4.1`

## Progress

slices 6/6 committed · standards ✓ · spec ✓ (1 gap closed) · editor ✓

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
