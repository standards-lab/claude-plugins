# goal · factory

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] pipeline
2. [x] goals
3. [x] experiments
4. [x] currency
5. [ ] evals

## Decisions

- pipeline, goals: built together as marathon 0.16.0 in one session, which ran under 0.15.
- pipeline: SHIP merges with `[remote] merge`; this repository uses
  `gh pr checks --watch && gh pr merge --merge --delete-branch`.
- goals: a goal's record lives in its `root`, one of the repositories it locks; sync deletes the
  record last.
- experiments: runs before evals, as the first `start factory` under 0.16, and ships as 0.16.1.
- evals: runs under the reinstalled 0.16. Defects its cases expose ship as 0.16.x patches, and the
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

## Pending edits

- architecture · `standards/go-elemental/principles/release-and-ci.md`: release preparation lands through a PR, not a direct commit to main; a tag whose release failed may be deleted and re-pushed at the same version, while a released tag is never re-cut.
