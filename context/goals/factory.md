# goal · factory

- **State:** building
- **Task:** evals
- **Branch:** evals

## Tasks

1. [x] pipeline
2. [x] goals
3. [x] experiments
4. [x] currency
5. [ ] evals

## Task brief · evals

```
Problem       A plugin can release with a regression its users already hit:
              nothing runs its evals before a tag, and marathon has no step
              between `[remote] ci` and tagging. The one observed workflow
              failure under 0.16+, a plan round shown as AskUserQuestion
              headlines instead of the round itself (rejected three times), is
              still allowed by 0.17.0. retro routes plugin findings to "an eval
              case" but doesn't say how one is applied.
Behaviors     Core contract (marathon 0.18.0)
              1. `[project] gate` is optional: SHIP runs it at development time
                 on the merge commit, once `[remote] ci` passes (or straight
                 after the merge without ci), before each Release tag, with
                 that tag as its one argument. Exit 0 means the tag may be
                 pushed. Nonzero holds the tag and is handled like a red
                 default branch: fixed forward on `<slug>-fix`, then ci and the
                 gate run again. CI never runs it, and marathon prescribes no
                 tool. Without the key, tagging works as it does today.
              2. A resumed release runs the gate again before each tag it
                 hasn't pushed. A handoff names a gate failure as that
                 repository's position, so the next `start` resumes at the fix.
              3. The configuration reference documents the key beside check
                 and currency, and `init` settles it as an optional founding
                 decision. The 0.18.0 CHANGELOG has a Migrating recipe for each
                 repository kind: claude-plugins declares the key in this
                 release, and repositories that release nothing omit it.
              4. The build reference names the release gate, for a plugin its
                 eval suite (too costly to run per slice), as the release-time
                 part of the automated-checks layer, so the note and the
                 reference agree.
              Plan rounds
              5. PLAN, intake, and escalation print the whole round in the
                 reply, in the briefs format, and the plan file keeps a copy. A
                 round is never presented or collected through
                 AskUserQuestion. The architect answers inline.
              6. Each question states what it is about in its own words, so it
                 can be decided without reading anything else.
              retro
              7. A ticked plugin eval finding is applied as a new case on the
                 retro-<topic> branch, and the case's description cites the
                 failure it came from (session or PR). The case may fail until
                 a task fixes the plugin, and the gate holds the next release
                 until then.
              claude-plugins
              8. The gate command, given `<plugin>/v<version>`, runs the
                 consistency check, then that plugin's eval suite with
                 threshold 1.0, 3 runs, no ablation, scaffolds on, the plugin
                 trusted, and reports kept local. It exits nonzero if either
                 fails or the plugin is unknown.
              9. A plugin with no eval suite is skipped: the gate passes on the
                 check alone and says so on stderr. marathon-architecture
                 releases this way.
              10. marathon's suite holds one case, plan-round-in-reply. An
                  offline scaffold builds a standalone project: a git repo,
                  a marathon config with no remote, a roadmap with one active
                  goal whose single task needs at least two decisions, and the
                  note it cites. The case grants AskUserQuestion so the failure
                  can happen. Graders: AskUserQuestion used 0 times (both arms);
                  the reply contains `PLAN ROUND 1`; it has at least two
                  numbered questions, each with `rec:` and `changes:`; and an
                  llm rubric of concrete PASS/FAIL claims that each question
                  names its subject (behavior 6).
              11. The case runs against 0.17.0 before the fix, and Evidence
                  says whether it failed there. With behaviors 5–6 in place, it
                  scores 1.0.
              12. Eval results stay out of git.
              13. The per-push check also runs `claude plugin validate` on each
                  plugin and the marketplace. CI installs Claude Code at an
                  exact version, and currency reports that pin when it trails.
                  The check's header names the workflow that runs it (ci.yml).
              Record and release
              14. BUILD runs the suite once as a pilot (one run) and reports
                  its costUsd in Evidence.
              15. This release passes its own gate, run by hand before
                  marathon/v0.18.0 is tagged, since the installed 0.17.0 has no
                  gate step. Evidence carries the output.
              16. The harness-testing note describes what exists: a suite for
                  each plugin with a case, cases only from failures via retro,
                  the gate and its flags, the skip rule, and the one seed case
                  in place of the four guessed seeds. The capability map line
                  matches.
Test seams    The gate command (`<plugin>/v<version>` → exit code), and
              `claude plugin eval` on marathon's suite.
Slices        1. Per-push check: validate in check.sh and CI (Claude Code
                 pinned), the currency pin, the header fix (13). Demo: check
                 and currency pass, and CI is green on the branch.
              2. plan-round-in-reply case, scaffold, graders, results ignored
                 (10, 12). Demo: one pilot run against 0.17.0, with its scores
                 and costUsd (11, 14).
              3. Skill fix: the round goes out in full in the reply, and
                 questions carry their subject (5, 6). Demo: the case scores
                 1.0.
              4. Core gate key: configuration, SHIP and Releasing, resume and
                 handoff, init, the build reference (1–4). Demo: check passes,
                 and the prose reads consistently.
              5. claude-plugins gate script and config, skip rule (8, 9). Demo:
                 the gate on marathon-architecture/v0.3.0 skips evals and
                 passes; on marathon/v0.18.0 it runs the suite.
              6. retro Apply bullet (7). Demo: check passes.
              7. Release prep: 0.18.0 version and CHANGELOG with Migrating, the
                 harness-testing note and capability map (3, 16). Demo: check
                 passes, and the gate run by hand passes (15).
Out of scope  Cases for marathon-architecture; the reviewer's module-cache
              write (held until an offline fixture can provoke it); running
              evals per push or in CI; adopting `claude plugin tag`; LLM-only
              cases; a marathon-architecture release; applying the Migrating
              recipe to other repositories.
Door          One-way for the push of marathon/v0.18.0, which reaches plugin
              users. Everything else (the key, the case, the prose, the
              scripts, CI) is two-way and revertable on the branch.
Release       marathon/v0.18.0
```

## Progress

slices 5/7 committed · standards — · spec — · editor —

## Decisions

- pipeline, goals: built together as marathon 0.16.0 in one session, which ran under 0.15.
- pipeline: SHIP merges with `[remote] merge`; this repository uses
  `gh pr checks --watch && gh pr merge --merge --delete-branch`.
- goals: a goal's record lives in its `root`, one of the repositories it locks; sync deletes the
  record last.
- experiments: runs before evals, as the first `start factory` under 0.16, and ships as 0.16.1.
  eval gate starts with the next release.
- The workspace alignment once planned as `factory.alignment` is spread over `quality.checks`,
  `quality.standards`, `quality.architecture-diet`, and `experiment.ai`.
- experiments: an experiment goal is a container never listed in the arrays; its spikes are sub-goals staged on their own; the ancestor rule stays.
- experiments: intake is the experiment's only task, run by `marathon intake`; `start` on an experiment refuses.
- experiments: no manifest field names the served goals; intake's first round names them.
- experiments: a spike's sync lands an answer section in the note the experiment cites; intake folds it into the served goals' notes and removes it.
- experiments: every session lands through one branch and one PR per repository (task slug, plan-, experiment-, intake-, sync-, retro-); only goal-record bookkeeping commits straight to the default branch, a standalone project's included. This replaces 0.16.0's direct coordinator commits.
- experiments: the record is deleted after the sync PR merges.
- experiments: accepting a last-task brief authorizes the sync and stages its proposed goal when staging's checks pass, `exclusive` included.
- experiments: `experiment` creates the goal and planned spikes; `plan <spike>` sets each up through Goal setup.
- experiments: `[experiment] serves` is removed; spike sessions start from the coordinator or workspace.
- experiments: a spike's tasks are its path steps; a finished 0.15 spike gets no lasting record, its closeout becomes its answer section.
- experiments: the migrations run after the 0.16.1 release as standards-lab plan sessions, following the CHANGELOG Migrating recipes.
- experiments: SHIP tags the brief's Release line only after main is green, fixing forward until the planned version releases; a tag whose release failed may be deleted and re-pushed, a released tag never.
- experiments: `[remote] ci` is an optional command that waits for main's CI on the merge commit.
- experiments: this session runs 0.16.0, so marathon/v0.16.1 is tagged by hand following behavior 9.
- experiments: a spike always starts in a new repository; an existing project informs it only as a read-only reference. experiment.ai's personal-agents task becomes `experiment.ai.spike-model-hosting` (models, not agents, until paired with a harness), with personal-agents as a reference; intake decides whether its contents and the spike's findings move into a fresh workspace repository.
- experiments (build, without the architect): a spike sub-goal is named for its repository (`experiment.ai.spike-harness-driver`).
- experiments (build, without the architect): an experiment goal keeps no goal record; each spike keeps its own in its repository.
- experiments (build, without the architect): intake writes only manifest tasks; an active served goal takes them into its record at its next `plan` or `start`.
- experiments (build, without the architect): `experiment` adds an `## Answers · experiment.<topic>` heading to the cited note; answer sections sit under it.
- experiments (build, without the architect): a session resumes on its own coordinator branch, which LOCATE does not treat as held.
- experiments (build, without the architect): a release's "tag <names>" handoff commits on the root's task branch until it merges, then on the default branch as bookkeeping.
- experiments (build, without the architect): the `[remote] ci` example in configuration.md polls for the merge commit's runs.
- experiments (build): spec review closed 10 gaps, among them stuck-sync resume, per-repository release resume, the last spike's Sync line, ordered migration recipes, and existing projects as spike references.
- currency: added before evals by plan, from quality.checks' pending edit; ships as 0.17.0.

- currency: the planner runs `[project] currency` before round 1, including when a stored next brief exists.
- currency: exit 0 with empty stdout means current; non-zero with lines means trailing; non-zero with no lines is a reported, non-blocking failure.
- currency: round 1 asks only adapt, adopt, and new-major questions; patch and minor bumps stay facts.
- currency: upgrade slices come first, one per trailing repo in `order`.
- currency: a held item is a goal-record Decision, asked about again only when its latest moves; there is no config key.
- currency: the key is optional; init settles it beside the check.
- currency: marathon fixes only the contract and prescribes no tool.
- currency: claude-plugins gets its own currency command and exact action pins.
- currency: quality.standards applies the migration recipe's code rows.
- currency (from quality.checks): currency is a development-time step, not CI.
- currency (from quality.checks): it covers direct dependencies and the toolchain; the resolver settles indirect dependencies.
- currency (from quality.checks): rejected as an extension.
- currency (from quality.checks): rejected as a CI gate and as Dependabot version-update PRs.
- currency (from quality.checks): Renovate is rejected in favour of per-repository scripts.
- currency: releases only marathon/v0.17.0.
- currency (build, without the architect): `<where>` in claude-plugins' currency report is the workflow file, so a pin used in two workflows gets one line per file.
- currency (build, without the architect): `scripts/currency.sh` prints its report only after every lookup succeeds, so a failure exits non-zero with empty stdout; it scans `.yml` and `.yaml` workflows.
- currency (build, without the architect): a held item's reason names the latest version it declined, so PLAN can tell when a later one is out; the session logs an answer of "hold" or "its own task" as a held item.
- currency (build, without the architect): `scripts/check.sh` fails when `.claude-plugin/marketplace.json` doesn't parse, a fix from the standards review.
- currency (build, without the architect): PLAN's currency rule lives in step 2 rather than a new step, so `mechanics/hooks.md`'s step references still hold.
- currency (build, without the architect): marathon-architecture is unchanged; it still targets marathon 0.16 and later.

- evals: the release gate is a new optional core key, `[project] gate`, that SHIP runs on the merge commit before each Release tag, with the tag as its argument. Rejected: running evals in release.yml, and reusing `[remote] ci`.
- evals: the gate runs eval with `--threshold 1.0 --runs 3 --ablation none --scaffold --trust-plugin --no-publish` and no cost ceiling. A case that fails in any of its 3 runs counts as a defect.
- evals: the seed suite is one case from observed failures, plan-round-in-reply. Dropped: implementer-no-standards, reviewer-commits, and brief-shape, since none has failed. The reviewer's module-cache write is held until an offline fixture can provoke it.
- evals: runs under 0.17.0. Defects its cases expose are fixed in this task and ship in its release. The gate first applies to that release, run by hand, since 0.17.0 has no gate step.
- evals: retro applies a ticked plugin eval finding as a case on its retro-<topic> branch, citing the failure in the case's description. The gate holds the next release until a fix lands.
- evals: marathon-architecture gets no suite until it has a failure. The gate skips a plugin with no eval suite.
- evals: fixtures are offline scaffold scripts. Replaying a recorded session through `history_file` was rejected because it carries the old skill text.
- evals: releases only marathon/v0.18.0. marathon-architecture is unchanged and doesn't release.
- evals (from round 1): the per-push check also runs `claude plugin validate`, with Claude Code pinned exactly in CI and covered by currency. check.sh's header names ci.yml.
- evals (from round 1): the build reference names the release gate, a plugin's evals, as the release-time part of the automated-checks layer. The harness-testing note is rewritten to match.

## Pending edits

- architecture · `standards/go-elemental/principles/release-and-ci.md`: release preparation lands through a PR, not a direct commit to main; a tag whose release failed may be deleted and re-pushed at the same version, while a released tag is never re-cut.
