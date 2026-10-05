# goal · factory

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] pipeline
2. [x] goals
3. [x] experiments
4. [x] currency
5. [x] evals

## Decisions

- pipeline, goals: built together as marathon 0.16.0 in one session, which ran under 0.15.
- pipeline: SHIP merges with `[remote] merge`; this repository uses
  `gh pr checks --watch && gh pr merge --merge --delete-branch`.
- goals: a goal's record lives in its `root`, one of the repositories it locks; sync deletes the
  record last.
- experiments: runs before evals, as the first `start factory` under 0.16, and ships as 0.16.1.
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
- evals (escalation): `claude plugin eval` 2.1.289 runs cases as `claude -p --permission-mode dontAsk` with no permission-prompt tool, which drops AskUserQuestion and plan mode, so plan-round-in-reply can't reproduce its failure. It stays as a check of the round's shape in the reply, keeping the AskUserQuestion grader for when eval offers the tool; the skill fix is unproven by eval.
- evals (redirect): CI pins Claude Code to its latest release (2.1.289), and currency compares the pin against npm's `latest` dist-tag, matching the dev machine and rolling currency.
- evals (build, without the architect): `scripts/check.sh` runs `claude plugin validate --strict` and fails when `claude` isn't on PATH.
- evals (build, without the architect): `scripts/gate.sh` takes `<plugin>/v<major>.<minor>.<patch>` and doesn't compare versions, since SHIP already does. It runs the check before the evals, and it finds cases where eval does: under the manifest's `experimental.evals`, or else `evals/`.
- evals (build, without the architect): the gate passes no `--allow-tools`, and the case doesn't list EnterPlanMode.
- evals (build, without the architect): the fixture's two decisions, format and audience, are independent, so a compliant round 1 asks both.
- evals (build, without the architect): the first hand-run gate failed at 0.92 on the self-contained-questions grader, which exposed rounds with conditional questions, bundled decisions, and pointers to "above". `references/briefs.md` now rules these out: each question decides one thing, and a dependent question waits for the next round.
- evals (build, without the architect): the duplicated pointer clauses to the plan-round rule stay, because trimming them changes the guarded path, and eval can't detect that.
- evals (build, without the architect): the numbered-questions grader checks two questions; the `llm` grader judges every question.
- evals (build, without the architect): go-web-sdk-template gets its own Migrating recipe for the gate, because it tags releases.

## Pending edits

- architecture · `standards/go-elemental/principles/release-and-ci.md`: release preparation lands through a PR, not a direct commit to main; a tag whose release failed may be deleted and re-pushed at the same version, while a released tag is never re-cut.
