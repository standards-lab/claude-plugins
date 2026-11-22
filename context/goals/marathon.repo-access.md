# goal · marathon.repo-access

- **State:** building
- **Task:** read-first-split
- **Branch:** read-first-split

## Tasks

1. [x] ci-awareness
2. [ ] read-first-split
3. [ ] review-findings
4. [ ] shared-repos

## Task brief · marathon.repo-access · read-first-split

**Problem**

Each command's read-first set counts the briefs reference whole: 2,558 words for nine formats,
when each command writes only a few. plan, retro, intake, start, and experiment sit within 50
words of their budgets, so every rule change needs a budget raise, and sentences ci-awareness cut
for budget are still missing.

**Behaviors**
1. Each of the nine formats is its own file in a `briefs/` directory under the skill's
   references; the existing briefs reference keeps only the rules every format shares, and each
   format file points to it. No rule's wording changes.
2. Each command's Read first list names the shared rules and only the formats it writes or cites:
   plan: plan round, task brief · start: plan round, task brief, slice brief, session brief,
   answer section · intake: intake round, plan round, answer section · experiment: plan's formats
   plus answer section · retro: retro · status: status digest, goal status · init and reset: none.
3. Every citation of the single briefs file — in the skill, README, agent profiles, CHANGELOG
   Unreleased, the eval case's description, the hook's package doc, and the context notes —
   points to the format it means, and the check passes.
4. The check fails when a playbook, or a format in its read-first set, cites a briefs file that
   the set doesn't contain; the failure names the file to add to the playbook's Read first list.
5. The cut sentences return word for word: SKILL.md's "The architect approves the task brief and
   accepts the session brief; agents do the rest.", pipeline.md's opening sentence and "Each
   playbook under `commands/` says which stages it runs.", the shared rules' "Agents look up
   facts; the architect makes decisions.", and "never everything" in the commit rule. init.md's
   currency and gate details stay cut.
6. The six briefs-reading commands' budgets drop to their new set sizes rounded up to the next
   hundred; init (3800) and reset (4600) stay. No budget rises.
7. The gate's eval, `plan-round-in-reply`, passes 3/3 against the branch (gate run with
   `marathon/v0.22.0` as label only; nothing tagged).
8. The CHANGELOG's Unreleased section records the split, the check, the restorations, and the
   lowered budgets.

**Test seams**
- The consistency check, with seeded-copy tests for the new rule.
- The release gate's eval run.

**Slices**
1. Split the formats: create the format files, reduce the old file to the shared rules, and
   repoint every Read first list and citation (behaviors 1–3). Budgets unchanged.
2. The check requires listed formats: playbooks cite every format they write, experiment's list
   gains the answer section; add the rule with tests, update the testing note's list of checks,
   and a CHANGELOG entry (behavior 4).
3. Restore the five cut sentences (behavior 5).
4. Lower the budgets, add the CHANGELOG entries for restorations and budgets (behaviors 6, 8),
   then run the gate's eval (behavior 7).

**Out of scope**

Changing any format's rules or headings, the wrap-formats hook's heading list, init's budget, any
other eval case or the judged suite, releasing, and the stale headroom figures in
harness-optimization.md (noted for review-findings).

**Door**

Two-way: no tag. Files move within the plugin, and nothing outside claude-plugins cites them.

## Progress

slices 0/4 committed · standards — · spec — · editor —

## Decisions

- goal: one release, at shared-repos: `marathon/v0.23.0` and `marathon-architecture/v0.4.1`;
  ci-awareness and review-findings tag nothing.
- shared-repos: moves claude-plugins' own configuration to `[remote] checks` with the release,
  since the installed 0.22.0 copy and binary don't know the key.
- read-first-split: the rules every brief format shares stay in `references/briefs.md`; each
  format is its own file under `references/briefs/`.
- read-first-split: init.md's cut currency and gate details stay cut, since
  `mechanics/configuration.md` states them; init's budget stays 3800.

## Pending edits

None.
