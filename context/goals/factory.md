# goal · factory

- **State:** brief ready
- **Task:** experiments
- **Branch:** experiments

## Tasks

1. [x] pipeline
2. [x] goals
3. [ ] experiments
4. [ ] evals

## Task brief · experiments

```
Problem   In 0.16.0 a spike is a task, but a spike spans many sessions, so spikes can't run
          on their own or in parallel, and nothing consolidates their answers. Coordinator
          edits skip review through direct commits, and releases are tagged by hand outside
          the pipeline. marathon 0.16.1 fixes all three before evals and the experiment
          migrations depend on them.

Behaviors
  1. A spike is a sub-goal experiment.<topic>.<spike> carrying `remote` and `path`, with its
     own repository as `root` and `repos`, and its path steps as tasks. The experiment goal
     is never listed in the arrays and holds only its `intake` task.
  2. Spikes of one experiment can be active at once; each locks only its own repository.
  3. `marathon experiment` runs once per topic: it creates the experiment goal, its intake
     task, and planned spike sub-goals (remote/path from [workspace.experiments], plus a
     summary of each spike's question and the decision it changes).
  4. `plan experiment.<topic>.<spike>` sets the spike up through plan's Goal setup: evidence
     list, founding decisions, path as tasks, repository and remote, read-only references,
     record. `[experiment] serves` is removed: spike sessions start from the coordinator or
     workspace, and the manifest's `path` locates the spike.
  5. `start` on a spike runs its tasks; the last task's brief gives the answer, and its sync
     lands an answer section in the experiment's cited note (question; one-line answer;
     numbered evidence, each with where it is proven; links to the spike README's "The
     answer" and its remote) and stages the next planned spike. `start` on an experiment
     goal refuses and points to `intake`.
  6. `marathon intake <experiment>` runs only when no spike sub-goal remains: it consolidates
     the answer sections, decides with the architect in rounds what each served goal builds,
     writes those tasks and removes the experiment in one intake-<topic> coordinator PR, then
     archives the named spike remotes, skipping any already archived. The plan file holds
     its state; there is no record.
  7. A last task's session brief has a Sync section: the pending edits and the goal proposed
     to stage next. Accepting the brief authorizes the sync and stages that goal when
     staging's checks pass, `exclusive` included; otherwise it reads "proposed, not staged:
     <reason>".
  8. Every session lands through a branch and a PR: plan-<goal>, experiment-<topic>,
     intake-<topic>, sync-<goal>, retro-<topic>, and task branches; a project with no remote
     merges locally. Only goal-record bookkeeping (header lines, a next brief, deleting a
     record after its sync PR merges) commits straight to the default branch. LOCATE treats
     a coordinator on a session branch as held.
  9. When the task brief has a `Release` line (exact tags, in each repository's own naming),
     SHIP runs per repository, lowest layer first: merge → `[remote] ci` passes on main's
     merge commit (when set) → each tag, annotated "<artifact> <version>" on that commit,
     pushed alone after checking its version matches the artifact's. A tag is never pushed
     while main is red. Failures are fixed forward in the same session until the planned
     version releases: a red main is fixed on <slug>-fix through PR → merge → ci; a failed
     release is fixed the same way, then its tag is deleted and re-pushed at the same
     version. A successfully released tag is never re-cut. ESCALATE only for decisions the
     brief doesn't cover; `handoff` with next move "tag <names>" only when the context fills.
  10. A planned release makes the task brief's Door one-way, and the session brief's Merge
      danger names the tags.
  11. The plugin ships as 0.16.1: version consistent in all three places; `intake` in the
      skill, README, and plugin description; a CHANGELOG Migrating section with one recipe
      per state, run as standards-lab `plan` sessions through PRs:
      - experiment.messaging (spike done, intake next): keeps only its intake task;
        spike-messaging's closeout becomes its answer section; drop `serves` and reset.md;
        leaves the arrays → ready for `intake`.
      - experiment.ai, spike-harness-driver (done, awaiting sync): sync under the new
        rules with no lasting sub-goal: closeout → answer section, drop `serves`, stage
        the next planned spike.
      - experiment.ai, planned spikes (spike-local-subagents, personal-agents): become
        planned sub-goals with remote/path kept; each set up later by `plan`; intake stays
        the experiment's only task.

Test seams  scripts/check.sh (versions, links, marketplace), plus a spec-reviewer
            walk-through of the skill against five scenarios: new experiment with two
            parallel spikes; a spike's sync; an intake; a released last task across two
            repositories, including a red main and a failed release; the three migration
            recipes against the current standards-lab roadmap.

Slices
  1. Spike sub-goal shape: manifest, goals, goal record, configuration (serves removed),
     LOCATE, status (behaviors 1, 2, part of 4).
  2. PR landing for every session: commit rules, admin exception, lock (8).
  3. `experiment` creates, `plan <spike>` sets up (3, 4).
  4. Spike sync, answer section, Sync section with stage-on-accept (5, 7).
  5. `intake` command: playbook, command table, hooks, intake round (6).
  6. Release at SHIP: Release line, `[remote] ci`, tag step, fix-forward loop, Door (9, 10).
  7. Release 0.16.1: versions, CHANGELOG with Migrating, README, descriptions (11).

Out of scope  Running the migrations (follow-on standards-lab plan sessions). Evals (task
              4). marathon-architecture changes. Setting `[remote] ci` in other
              repositories. The architecture repository's release-and-ci.md.

Door     one-way: marathon/v0.16.1 is published; every other change is two-way.
Release  marathon/v0.16.1
```

## Progress

slices 7/7 committed · standards ✓ · spec ✓ · editor ✓

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

## Pending edits

- architecture · `standards/go-elemental/principles/release-and-ci.md`: release preparation lands through a PR, not a direct commit to main; a tag whose release failed may be deleted and re-pushed at the same version, while a released tag is never re-cut.
