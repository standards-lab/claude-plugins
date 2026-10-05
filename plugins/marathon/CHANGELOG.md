# Changelog

All notable changes to the marathon plugin are documented here. Versions follow
[Semantic Versioning](https://semver.org/spec/v2.0.0.html); dates and release links live on the
GitHub releases the tags cut.

## v0.17.0

A task begins by bringing what it touches up to date. A repository may declare a currency
command, which PLAN runs before round 1; what trails becomes facts and questions in the round and
upgrade slices at the head of the task. The key is optional, so 0.16.1 repositories keep working;
see Migrating below to declare it.

### Added

- **`[project] currency`** (`mechanics/configuration.md`): an optional, read-only command that
  reports which direct dependencies, toolchain versions, and pins, such as CI actions and images,
  trail their latest release. Indirect dependencies are left to the ecosystem's resolver. It exits
  0 with empty stdout when everything is current, and nonzero with one stdout line per trailing
  item, suggested as `<where>: <item> <pin> -> <latest>`, when something trails. Nonzero with no
  stdout lines is a failed command. Diagnostics go to stderr. PLAN runs it at development time and
  CI never does. marathon prescribes no tool.
- **The currency command as a founding decision** of `init`, optional, beside the check.
- **The `currency:` block** of plan round 1 (`references/briefs.md`): per repository the task
  touches, each trailing item with a one-line summary of its release notes, or `current`, or
  `no currency command`, or the failed command as a fact. The round asks only about an item to
  adapt (a breaking change the code must absorb), to adopt (a new feature worth using), or at a
  new major (upgrade now, hold, or give it its own task); a patch or minor bump with nothing to
  adapt stays a fact.
- **Upgrade slices** (`references/briefs.md`): a task's slices begin with one upgrade slice per
  trailing repository it touches, in the coordinator's `order`, each done when that repository's
  currency command exits 0, or reports only held items, with its check passing.
- **Held items** (`mechanics/goal-record.md`): a goal-record Decision
  `<task>: held <item> at <pin>: <reason>`, whose reason names the latest version it declined,
  so the planner asks again only once a later one is out.

### Changed

- **PLAN step 2** (`mechanics/pipeline.md`): before round 1, the planner runs each touched
  repository's currency command on its default branch and reads the release notes of what trails.
  A repository without the key is noted, and a failed command is reported as a fact and doesn't
  block. A stored next brief no longer skips currency: when something trails that the brief
  doesn't cover, one round opens on it before approval.
- **This repository** declares `currency = "scripts/currency.sh"`, which checks each workflow's
  `uses:` pin against the action's latest release tag, and pins its actions to exact release tags.

### Migrating from 0.16.1

Install 0.17.0 first. Declaring the key is optional; a repository without it is noted in round 1
and otherwise plans as before. Each repository adds the key under `[project]` in its
`.claude/marathon.toml`, through its own pull request, with a currency command that meets the
contract (a `mise` task, a script, or the ecosystem's outdated report):

1. **A code repository with its tasks at the root** (go-core, sqlate, go-database, go-web-sdk,
   go-observability, go-storage, blobfs, go-web-service): `currency = "mise run currency"`.
2. **A code repository whose tasks live in a subdirectory** (go-web-sdk-template):
   `currency = "mise -C template run currency"`.
3. **A context repository that pins actions** (claude-plugins): `currency = "scripts/currency.sh"`,
   done in this release.
4. **A repository that pins nothing** (architecture, standards-lab, and the two `.github` profile
   repositories): no key.
5. **A spike**: optional, when the spike has a currency task.

The `quality.standards` task applies recipes 1 and 2.

## v0.16.1

Spikes become goals of their own, every session lands through a pull request, and releases are
tagged inside the pipeline. An experiment's spikes run in parallel and their answers are taken in
by the new `intake` command. 0.16.0 experiments need migrating; see Migrating below.

### Added

- **`intake`** (`commands/intake.md`): `marathon intake experiment.<topic>` takes in a finished
  experiment once no spike sub-goal remains. In plan rounds led by the spikes' answers it decides
  what each served goal builds, then one `intake-<topic>` coordinator pull request writes those
  tasks, folds the answer sections into the served goals' notes, removes the experiment, and
  updates the repository catalog; the named spike remotes are archived after it merges. The plan
  file holds its state, and it writes no goal record.
- **Release tagging at SHIP** (`mechanics/pipeline.md`, "Releasing"). A task brief may carry a
  `Release` line naming the exact tags it releases, which makes its Door one-way; the session
  brief's Merge danger repeats them, so accepting it authorizes tagging. SHIP merges and tags one
  repository at a time, lowest layer first, never while the default branch is red, and fixes
  forward on `<slug>-fix` until the planned version releases. A released tag is never re-cut. A
  release whose context fills hands off with "tag <names>", committed on the root's task branch
  until it merges and on the default branch after; the next `start` resumes each repository from
  the position the handoff records.
- **`[remote] ci`**: an optional command that waits for the default branch's CI run on the merge
  commit before any tag is pushed.
- **The answer section** (`references/briefs.md`): a spike's last task gives its answer, and its
  sync lands an answer section (question, one-line answer, numbered evidence with where each item
  is proven, links to where the spike states its answer and to its remote) in the note the
  experiment cites, under the heading `experiment` writes there.
- **The Sync section** of a last task's session brief: the pending edits per repository and the
  goal proposed to stage next. Accepting the brief authorizes the sync and stages that goal when
  staging's checks pass; otherwise it reads "proposed, not staged: <reason>".

### Changed

- **A spike is a sub-goal**, `experiment.<topic>.<spike>`, carrying `remote` and `path`, with its
  own repository as `root` and `repos` and its path steps as tasks. Spikes of one experiment can
  be active at once, each locking only its own repository. The experiment goal is never listed in
  `active`, `planned`, or `backlog`, keeps no goal record, and holds only its `intake` task;
  `start` on it refuses and points to `intake`.
- **`experiment`** runs once per topic: in one `experiment-<topic>` pull request it creates the
  experiment goal, its `intake` task, and its planned spike sub-goals, and the note the
  experiment cites. It creates no repository and stages no spike.
- **`plan experiment.<topic>.<spike>`** stages a spike and sets it up ("Spike setup" in
  `commands/plan.md`): evidence list, founding decisions, path as tasks, repository and remote,
  read-only references, and record. A spike always starts in a new repository; an existing
  project it builds on is only a read-only reference. A spike session starts from the
  coordinator or the workspace, and LOCATE finds the spike's repository through its `path`.
- **Every session lands through a branch and a pull request** in each repository it changes: the
  task's slug, `plan-<goal>`, `experiment-<topic>`, `intake-<topic>`, `sync-<goal>`, and
  `retro-<topic>` (`mechanics/pipeline.md`, "Branches and pull requests"). A project with no
  `[remote] publish` merges locally. Only goal-record bookkeeping commits straight to the default
  branch. The coordinator changes only through `plan`, `experiment`, `intake`, `retro`, and sync,
  and LOCATE treats a coordinator on another session's branch as held; a session resumes on its
  own.
- **Sync** of a spike stages the next planned spike of its experiment when staging's checks pass,
  and no longer archives spike remotes; `intake` does.
- **`on-build`** no longer fires when a session resumes only a sync or a release; a release fires
  it whenever it opens or resumes a `<slug>-fix` branch (`mechanics/hooks.md`).

### Removed

- `[experiment] serves`, from configuration, `init`, and LOCATE.
- Spikes as tasks of an experiment goal, and the spike intake brief, replaced by the intake round.

### Migrating from 0.16.0

Install 0.16.1 first. Migrate each experiment in one `plan` session at the coordinator, landing
as one `plan-<experiment>` coordinator pull request whose body is the approved round outcome. No
session brief or goal record is involved: a 0.15 spike has none. Its rounds settle what the
recipes below leave open.

In every case:

- The experiment goal keeps only its `intake` task. Remove its spike tasks, its `root` and
  `repos` (an experiment goal is a container and has neither; `experiment.ai` has
  `spike-harness-driver` as `root` and all three spikes' repositories as `repos`,
  `experiment.messaging` has `spike-messaging` as both), and the experiment itself from
  `active`, `planned`, and `backlog`. Its spike sub-goals are listed instead.
- A finished spike's 0.15 closeout, its `context/reset.md` and where the spike states its answer
  (`spike-harness-driver`'s `context/findings.md`, `spike-messaging`'s `context/README.md`),
  becomes its answer section (`references/briefs.md`, "Answer section"). The section lands in
  the note the experiment cites, under an `## Answers · experiment.<topic>` heading added when
  the note has none. When the experiment cites several notes, as `experiment.ai` cites
  `ai-strategy.md` and `ai-hosting.md`, the plan round chooses the one that receives it.
- A finished spike's repository changes through a pull request on a `plan-<experiment>` branch
  there: it deletes `[experiment] serves` and `context/reset.md`. Neither spike's
  `marathon.toml` has `[remote] merge`, so the same pull request adds
  `merge = "gh pr checks --watch && gh pr merge --merge --delete-branch"` under `[remote]`;
  without it, the session stops for a manual merge.
- Update any repository catalog line that describes a spike as a task of an experiment goal.

The recipes, by the experiment's state:

1. **A finished experiment whose spike is done and whose intake is next** (for example,
   `experiment.messaging` with `spike-messaging`). Apply the steps above: the spike's answer
   section lands, its repository's pull request merges, and no sub-goal remains. Then run
   `marathon intake experiment.<topic>`.
2. **A spike that is done and awaiting sync** (for example, `experiment.ai`'s
   `spike-harness-driver`). Sync it under the new rules, leaving no sub-goal behind: its answer
   section lands, its repository's pull request merges, and the next planned spike is staged
   into `active`.
3. **Spikes still planned** (for example, `experiment.ai`'s `spike-local-subagents` and
   `personal-agents`). Turn each spike task into a planned sub-goal
   `experiment.<topic>.<spike>` that keeps its `remote` and `path`, with the spike's repository
   as `root` and `repos` and a summary of its question; list it in `planned`. `intake` stays the
   experiment's only task. Each spike is set up later by `plan experiment.<topic>.<spike>`.
   A spike always starts in a new repository (`commands/plan.md`, "Spike setup"), so a spike
   task that names an existing project becomes a sub-goal for a new spike repository instead.
   `experiment.ai`'s `personal-agents` becomes `experiment.ai.spike-model-hosting`, at the
   `[workspace.experiments]` path `~/experiments/spike-model-hosting` with the remote
   `https://github.com/JaimeStill/spike-model-hosting.git`. Its summary names
   `~/code/personal-agents` as a read-only reference, recorded when `plan` sets the spike up.
   Nothing is written to or archived from `personal-agents`. The migration also adds an open
   question to `experiment.ai`'s `intake` task entry, for the intake to decide: whether
   `personal-agents`' contents and the spike's findings move into a fresh workspace repository.

`experiment.ai` takes recipes 3 and 2 in its one pull request, in that order: convert the planned
spikes to sub-goals first, then land `spike-harness-driver`'s answer section and stage the next
planned spike.

Each repository may add `[remote] ci` to its `marathon.toml`. `.claude/briefs/` stays in
`.gitignore`.

## v0.16.0

marathon becomes a software factory. The architect engages twice per task, approving the task brief
and accepting the session brief, and subagents make the work good in between. Goals and tasks
replace waves, lanes, and folds, and the roadmap moves into core. This release breaks 0.15
repositories; see Migrating below.

### Added

- **The PLAN, BUILD, BRIEF, SHIP pipeline** (`mechanics/pipeline.md`). PLAN runs plan rounds to an
  approved task brief and slices. BUILD runs without stopping, escalating only for a one-way door,
  a decision the brief doesn't cover, or scope beyond the task. BRIEF shows the session brief.
  SHIP publishes it as the pull request's body and merges.
- **Goals in core** (`mechanics/goals.md`, `references/manifest.md`). The manifest's root holds
  `active`, `planned`, and `backlog`; goals carry `root` and `repos`, inherited from ancestors.
  An active goal locks its repositories, so goals run side by side on main checkouts. The
  coordinator is never locked and changes only through `plan` and sync. Sync carries a finished
  goal's pending edits into the coordinator and other repositories, and deletes last.
- **Experiments as goals.** An experiment is a goal under `experiment` (`experiment.ai`), its
  spikes are tasks with `remote` and `path`, and its last task is the intake. Sync archives the
  spikes into the workspace's repository catalog.
- **Goal records** (`mechanics/goal-record.md`): `context/goals/<goal>.md` in the goal's root,
  one of the repositories it locks, holding State, tasks, the current brief, progress, handoff,
  decisions, and pending edits. An open plan round lives in the plan file and is logged into the
  record on approval.
- **Briefs** (`references/briefs.md`): plan round, task brief, session brief, status digest, retro,
  and spike intake.
- **The build loop** (`references/build.md`): checks first, slices, implementer per slice, then the
  standards-reviewer, the spec-reviewer, and the editor.
- **Profiles**: `implementer` (formerly `executor`), loading only the slice brief and the check;
  `standards-reviewer` and `spec-reviewer` (formerly `reviewer`), in separate contexts, the first
  committing fixes and sweeping for tests that lie.
- **Commands**: `status` prints the digest; `retro` routes findings to a check, a standard, a skill,
  or a note, and absorbs the drift pass.
- **Configuration**: `[project] check`, `[remote] merge`, `[workspace] exclusive`, and
  `[workspace.experiments]`; `[experiment] serves` takes the coordinator's path.
- **Session briefs** at `.claude/briefs/<goal>.md`, one per goal.
- **`init`** writes `roadmap.toml`, a `STANDARDS.md` stub, and a pointers-only `CLAUDE.md`.

### Changed

- **`start`** runs one task end to end, including what `close` did, and syncs a goal's last task.
- **`plan`** defines goals and makes them ready to run: it creates goals and tasks, sets up their
  repositories and records, stages and orders them, and writes notes and next briefs.
- **`reset`** writes the handoff into the goal record and runs on its own when the context fills.
- **`experiment`** adds a spike as a task of an `experiment.<topic>` goal.
- **Hooks**: `on-start`, `on-build` (formerly `on-execute`), `on-ship` (formerly `on-close`), and
  `on-record` (formerly `on-reset`).
- **Delegation**: profiles inherit the session's model, and the session announces none.

### Removed

- `close` and `review`, the reset file (`mechanics/reset-file.md`), waves, lanes, folds, and
  worktrees (`mechanics/waves.md`), stages and checkpoints (`references/staged-execution.md`),
  `.claude/report.md`, delegation statements, and the `on-commit` hook.
- The marathon-roadmap extension, now part of core.

### Migrating from 0.15

1. Let any session already running under 0.15 finish there; 0.16 applies only to new sessions.
2. Disable and uninstall marathon-roadmap; remove it from every `extensions` list. A third-party
   extension renames `on-execute` to `on-build`, `on-close` to `on-ship`, and `on-reset` to
   `on-record`, and drops `on-commit`.
3. In `roadmap.toml`, replace `next` with `active` and `planned`: a task path or a wave member maps
   to its goal's path. Turn each `[backlog.x]` table into `[goals.x]` listed in `backlog`, and
   rewrite `backlog.x` citations in notes as `x`. Give each goal you stage a `root` and `repos`
   that cover its tasks.
4. Move experiments into `experiment.<topic>` goals, each spike a task with `remote` and `path`
   and an `intake` task last. Move archived spikes into the repository catalog, and point each
   spike's `[experiment] serves` at the coordinator's path.
5. Turn `context/reset.md`, each lane record, and each spike's own reset file into the progress of
   the goal records in each goal's root, then delete them.
6. Add `check` and `merge` to each repository's `marathon.toml`. In `.gitignore`, replace
   `.claude/report.md` with `.claude/briefs/` and drop `.claude/worktrees/`; run
   `git worktree remove` on any worktree left over.
7. Optionally add a `STANDARDS.md` and reduce `CLAUDE.md` to navigation pointers.
8. Update marathon-architecture to 0.3.0.

## v0.15.0

### Added

- **Waves in a standalone project.** A standalone project can run a wave: its main lane plus
  experiment lanes, with lane records and folding as in a workspace. `mechanics/waves.md` is the
  one home for waves in both project kinds, and folding moves there from
  `mechanics/reset-file.md`.
- **A worktree at the shared repository.** The shared repository holds the reset file: the
  project itself, or the coordinator. The lane that owns it works in its main checkout. Every
  other lane changes it only in a worktree of its own, at `.claude/worktrees/<branch>`, created
  with `git worktree add` from the fetched `origin/<default-branch>` once the shared repository
  gitignores `.claude/worktrees/`. The worktree holds only the lane's `context/`
  changes, meaning its record and its own notes. `reset` keeps the worktree and records its
  path, and `close` removes it after publishing. With no owning lane, the main checkout stays on
  the default branch. LOCATE reads a lane's record from the checkout that holds its open branch.
- **The branch check before every commit.** Every commit a session makes first confirms that the
  checkout is on the session's branch, and the session stops and reports if it isn't.

### Changed

- **A lane owns exactly one thing**: the standalone project, a workspace member, the coordinator,
  or an experiment. A step that spans several members never runs as a lane.
- **Extension artifacts wait for the fold.** An extension's artifact is shared wherever it lives,
  so a lane records its changes to one, and its experiment catalog entry, in its Disposition. The
  fold applies them.
- **A lane's branch report** is written in its worktree, so lanes that close together don't
  collide.
- **Branches and steps in a wave.** Every branch a lane creates starts from the fetched default
  branch, and a lane's next step waits until its previous branch has merged.
- **Folding.** The fold can change files outside `context/` and in other repositories, with one
  commit in each. The session that finishes the last lane first brings its branch up to date with
  the default branch. A fold found at LOCATE runs as the session's first stage, after approval.
- **Shared files during a wave.** Tending the context at close, intake, and extension hooks,
  including an artifact's creation at `on-start`, record changes outside the lane's own files in
  its Disposition. A gitignored local map is changed in the main checkout. `context/reset/` is
  the one subdirectory `context/` holds.

### Migrating

Add `.claude/worktrees/` to `.gitignore` in each repository that holds a reset file. Split any
planned lane that spans several workspace members into lanes of one member each, or run the step
outside a wave. Update marathon-roadmap and marathon-architecture to 0.2.2, which target marathon
0.15.

## v0.14.0

### Changed

- **Notes and documentation state what is true now.** A note describes what exists, in the
  present tense, or what is planned, marked as planned. It has no settledness line, history,
  dates, or other detail that only holds for a while. The roadmap and the reset file track status.
  The rule covers every document outside project tracking, including `docs/` and READMEs, and
  CHANGELOGs are exempt. It replaces 0.13's rule that a note states its settledness in its
  opening lines.
- **The prose is rewritten for clarity.** Every file in the skill, the agent profiles, and the
  README is rewritten so a reader new to the project can follow it on first read. The rules are
  unchanged.

### Added

- **The editor profile.** `agents/editor.md` edits the prose a branch changed, once, after the
  last stage commits and before the final checkpoint. It works to a six-rule standard for concise
  technical documentation, defers to the ecosystem's API documentation form and to a voice
  standard the project declares, keeps meaning unchanged, and never commits. The session commits
  its edits as one stage.

### Migrating

Remove the settledness line from each note in `context/`, and rewrite the note to describe what
exists or what is planned. Move status to the roadmap or the reset file. Add `.claude/report.md`
to `.gitignore` in any repository that doesn't list it.

## v0.13.0

### Changed

- **Checkpoints replace per-stage review.** A stage commits once its check passes and logs its
  diff stat, check result, and delegation call. The stage list groups stages under checkpoints,
  observable behaviors the architect runs or watches, with Confirmed, Adjust, Re-plan, and
  Interrupt as outcomes. A context project's checkpoint is a scenario walkthrough.
- **`close` opens with a branch review.** The reviewer profile reviews the whole branch against
  the ecosystem's idiom, or a practice the project declares, and writes a gitignored
  `.claude/report.md` (overview, per-change changelog, verification) that closeout deletes. The
  reset's Disposition gains **Validated**, and a handoff records its checkpoint position.
- **Agent profiles replace the agent catalog.** marathon ships `planner`, `executor`, and
  `reviewer` profiles in `agents/`, none pinning a model; the session states the profile, model,
  and reason before each engagement. `[agents]` and `[workspace.agents]` are removed.
- **`context/` is flat.** `design/` and `concepts/` are gone: every note sits under `context/` and
  states its settledness in its opening lines. A note is Integrated once the built work or its
  documentation expresses it, with lasting reasoning moved into the owning repository's docs.
  The ledger is Integrated, Culled, Retained, Add or sharpen, and Cross-repo; Promote leaves core.
- **Experiments are standalone projects.** `experiment` settles the question and where the
  experiment lives (directory and hosting account), runs `init` there with `[experiment] serves`,
  and catalogs it in the served project. The experiment runs its own sessions in parallel; a
  `plan` session takes in its result and archives its remote. Top-level `experiments/` is gone.
- **Waves.** A Next-focus may name a wave of lanes that run concurrently, each a single step or a
  sequence run in order. Each lane keeps its own record at `context/reset/<lane>.md`, and the
  session that finishes the last lane folds the wave into one commit.
- **The skill is about a third smaller**, with restated rules and cross-file duplication removed.

### Migrating

Move each `context/design/` and `context/concepts/` note to `context/` with a settledness line, or
integrate, cull, or relocate it; move each `experiments/<slug>/` spike to its own repository and
catalog it; drop `[agents]` and `[workspace.agents]` from `marathon.toml`; and add
`.claude/report.md` to `.gitignore`.

## v0.12.0

### Changed

- **The architecture layer is no longer mandatory** — extracted out of core entirely into the new
  `marathon-architecture` extension. `init` no longer scaffolds an `architecture/` directory or
  asks it as a founding decision; `context-engineering.md` no longer describes the layer;
  `close` and `review` lose their architecture-specific clauses; `configuration.md` carries no
  `architecture` key. A project that never generalizes past itself now enables nothing and
  scaffolds nothing, instead of carrying a directory that never gains a second sentence.
- **Extensions never touch core's config schema** — a new rule in `references/extensions.md`: an
  extension's only foothold in `marathon.toml` is the shared `extensions` list; configuration
  beyond enablement lives in a file the extension owns, that only it reads. Named explicitly
  because `marathon-architecture` is the first extension that needed it.
- **Two new context disciplines** in `context-engineering.md`: a pre-write check before anything
  lands in `context/` (does this fact already have exactly one home outside it? cite it, never
  restate it), and a rule against promoting a design note or authoring a skill in the same
  session that designed the shape it documents — promotion waits for a session where something
  real exercised the shape and it held.
- **Context states current truth only** — `design/` and `concepts/` notes, and an architecture
  page, hold what is true now and never a changelog of how it changed; that accounting belongs
  solely to the reset file's Disposition.
- **The stage-loop delegation call is stated, not silently skippable** — every stage now opens by
  stating whether it goes to a declared technical agent or stays with the session, and why,
  mirroring how SETTLE already requires stating an escalation before engaging it.

## v0.11.0

### Added

- **Delegation** — a session may hand one unit of its own work (a stage's implementation, a
  design decision, a call too consequential to settle alone) to another agent and stays its
  owner; marathon states the contract a delegation runs under and names no agent of its own. A
  project catalogs which agents its sessions may reach for under `[agents]` in `marathon.toml`,
  one sub-table per agent with a `delegation` field describing what it's for and why, in the
  project's own words; a session judges fit against that text, never a fixed role lookup. A
  workspace coordinator's `[workspace.agents]` is the baseline, a project's own table overrides
  it key by key. `behavior/delegation.md` holds the contract; `mechanics/configuration.md` holds
  the schema. SETTLE names where a design decision may go to a declared escalation agent, and the
  stage loop of `references/staged-execution.md` names where a stage's implementation may go to
  a declared technical agent; either way the session reports and commits exactly as it would
  have for work it did itself.

- **Sufficiency before building it directly** — before the stage list is settled, a step that
  plans to develop a solution directly asks whether the problem is already resolved by the
  established, idiomatic approach for the language — the standard library, or a dependency the
  ecosystem already treats as standard. If it is, the step needs an adequate reason to build its
  own instead of adopting that approach. Asked at SETTLE, while the answer still changes the
  plan, not discovered at review. A step that proceeds with its own implementation anyway
  carries its reason into the design note it touches, as a rejected alternative; a step with no
  adequate reason adopts the existing approach instead. Stated in `behavior/planning.md`, named
  at the pipeline's SETTLE step. Found by the SQL strategy's own sufficiency rule
  (`standards-lab/context/design/dsl-driven-services.md` §2.4), which names the marathon harness
  as where the question belongs.

- **A stage's check runs the repository's own tooling** — on a code project, a stage that
  produces source now runs whatever conventions or lint tool the repository already wires into
  its own CI or task runner, over the files the stage touched, alongside the language's build,
  vet, and test. The session runs the tool and resolves its findings; it never restates what the
  tool checks. Stated once in `references/staged-execution.md`'s "What a stage is." First
  consumer: sqlate's `sqlint`, which a consuming repository already runs through its own `mise`
  task or CI step.

### Changed

- **The architecture layer, and documentation as the project's own** — every marathon project
  has an architecture layer: the principles, definitions, and conventions that have generalized
  past one repository, the top of the context lifecycle and the target a design note promotes
  to when the built work cannot express it. A standalone project keeps it as a top-level
  `architecture/` directory, scaffolded by `init`; a workspace keeps it as one repository, named
  by the coordinator's `[workspace] architecture` key, whose tree is the architecture with a
  README as every directory's index. It holds nothing a reader could infer from a repository's
  source, and a page that restates a repository is a defect. `docs/` is now only the project's
  own documentation, the guide over its README, API documentation, and source, with the code as
  its source of truth, written by a documentation step of a `start` session like any other
  built work. Under 0.10.0 the docs tier doubled as the workspace's landing zone, named by
  `[workspace] docs`, a member never grew a `docs/` of its own, and the landing zone documented
  each member's implementation, which drifted with every release. `review` gains the
  promotion-candidates check and `close` lands a generalized design note in the architecture
  repository under **Cross-repo**. Found by the `v1.alignment.docs` session, whose module pages
  restated package documentation and were reverted.

### Removed

- **The `docs` command** — project documentation is built work, so a `start` step writes and
  curates it and `review` flags its drift; the architecture repository is a `context` project
  that `start` and `plan` author. What the command's playbook stated about establishing a
  `docs/` directory is one paragraph of `references/context-engineering.md`. The stage rule in
  `references/staged-execution.md` generalizes to carry it: a stage's unit and check follow what
  the stage produces, source or prose, rather than the project kind alone, so a code project's
  `start` runs a documentation step under the prose rule. The Session values of the reset file
  lose `docs`.

## v0.10.0

### Changed

- **The review gate, for every working session** — a stage is reported with the working tree
  uncommitted, so the diff reads cleanly in the architect's tools; it commits only on the
  architect's approval, who then states whether a `reset` follows. Under 0.9.0 the commit
  preceded the report, and the loop applied to a code project's `start` alone. Now every command
  that executes runs it: `start` on either kind, `experiment`, and the context edits of `plan`,
  `review`, and `docs`. A stage's unit follows the project kind: a compilation unit on a code
  project; on a context project, the smallest set of files that must change together, checked by
  the repository's own consistency script and a coherent read. Validation is per kind, and an
  experiment's is its answer to the question it was settled to answer. Stated once in
  `references/staged-execution.md`; the pipeline gains the invariant, and the reset-file schema,
  `start`, `experiment`, `reset`, `close`, `plan`, `review`, `docs`, and the cross-repo stage
  list follow. Found by the `v1.data.sql.prototype` experiment, whose first three stages were
  committed before review.
- **Experiments live at the workspace coordinator** — in a workspace, every experiment sits
  under the coordinator's top-level `experiments/`, on the coordinator's branch, with the reset
  file's Project line naming the coordinator; a standalone project keeps its own `experiments/`.
  A spike depends on member modules as published versions, never through a replace directive,
  and a change it implies for a member's code reaches that repository only by promotion at
  `close`. The reasons are a new section of `references/workspace-coordination.md`;
  `commands/experiment.md`, the tier list of `references/context-engineering.md`, and the
  reset-file schema notes state the rule. Settled by the same experiment, which moved from a
  member repository to the coordinator for exactly this reason.

## v0.9.0

### Changed

- **Staged execution replaces the implementation guide** — on a `code` project the agent
  implements the settled step directly, in stages, and the architect reviews each stage before
  the next begins. A stage is the smallest change set that leaves one compilation unit green —
  its tests and in-source comments included; the check is scoped to that unit, and the module may
  be red between stages because the sequence is in dependency order. The stage list — unit,
  files, one-line why per entry — is the SETTLE artifact, produced in the planning phase and
  approved before any code changes. Each stage lands as its own commit (`on-commit` fires first)
  with a conversational report: `diff --stat`, the check result, prose only on decisions the plan
  didn't spell out. Review outcomes are approve, adjust, and re-plan (the architect enters plan
  mode with findings; invalidated commits revert first; a revised list from stage k is approved
  like the original). Validation — the whole-module build, full test run, and run-and-verify
  check — runs once after the last stage. The new `references/staged-execution.md` holds the
  loop; `start`, `close`, `reset`, the pipeline, and the reset-file schema (stage position in a
  handoff's Next-focus) are rewritten around it.
- **The human is the architect** — vision, direction, and quality control, exercised where the
  processes name them: the SETTLE approval, the per-stage review, the tending confirmation, and
  the pull request. Each playbook states what it expects of the architect at the moment it
  occurs; the standalone role-boundary reference and the git-blame test are retired, and
  "developer" becomes "architect" throughout the skill. `kind` survives on narrower grounds —
  `code` means a build-and-test loop defines stage boundaries and a validation phase; `context`
  has neither.
- **Cross-repo steps replace `coordinate`** — the fan-out model is retired; an orchestrated
  change across workspace projects runs as one working session whose settled step names the
  repos it touches: one spanning stage list in the coordinator's dependency order, a branch per
  touched repo under the step's shared slug, one close that commits and publishes each.
  `references/workspace-coordination.md` is rewritten around what remains: the coordinator and
  its order map, the workspace-holds-no-context rule, continuity at the single reset, downward
  awareness, coordinator conventions.
- **The reset file is purely ephemeral** — its sole purpose is contextual bootstrapping between
  sessions; durable detail belongs in the context layers. Stated outright in
  `mechanics/reset-file.md`, with two fixes: the file is rewritten each session and git is the
  archive (the "accumulates" wording is gone), and an interrupted cross-repo step is expressible —
  the Project line lists the touched repos, the branches share the step's slug, and Next-focus
  carries whatever bootstrap state resuming cold needs. **Cross-repo** joins the declared
  Disposition ledger vocabulary.
- **Close tends the full context scope** — the tending pass establishes the written context the
  step touched before operating on it: standalone, the project's own `context/`; in a workspace,
  each touched repo's `context/`, the coordinator's notes on the changed capability, the docs
  landing zone pages the change moved out from under, and claims about the changed behavior in
  other member repos' context — a stale claim found anywhere in that scope is a defect the pass
  fixes, recorded under **Cross-repo**.
- **Single-source pass** — the decay rule is stated once, in
  `references/context-engineering.md` with its protective qualifier; `close`, `reset`, and
  `review` cite it instead of restating the lossy short form. The `[workspace]` TOML block is
  printed once (`mechanics/configuration.md`), the five-hook list once (`mechanics/hooks.md`),
  and `init`'s hook constraints live only there.

### Added

- **The skill states its version** — a `Version:` line under the SKILL.md title makes the
  extension compatibility check of `mechanics/hooks.md` executable and installed-vs-source skew
  detectable at runtime.
- **Plugin CI** — a push-triggered workflow at the host checks, per plugin, that the manifest
  version, the top CHANGELOG heading, and the SKILL.md version line agree, that
  `marketplace.json` sources resolve, and that every `@`-pointer and `./`-link inside each skill
  resolves to a real file.

### Removed

- `references/implementation-guides.md`, `references/role-boundary.md`, `context/guide.md`,
  `commands/coordinate.md`, and the coordinated fan-out's consolidated guide. The guide-era
  workflow remains available at the `marathon/v0.8.0` tag.

## v0.8.0

### Changed

- **A single workspace reset at the coordinator** — a workspace maintains one reset file, at the
  coordinator; member projects carry none, and a standalone project keeps its own unchanged. LOCATE
  resolves the coordinator's reset from anywhere in the workspace, `reset` and `close` write it
  there (a commit in the coordinator's repository), and the member-reset routing — including the
  resting-point deletion rule — is retired. The schema gains a **Project** line naming the member a
  workspace record concerns. The new `mechanics/reset-file.md` holds the schema, locations, and
  Status semantics; `coordinate` and `references/workspace-coordination.md` re-anchor continuity on
  the coordinator's record plus each project's open branch.
- **Comment-free implementation guides** — stated outright in
  `references/implementation-guides.md`: the guide's code blocks carry no comments of any kind; doc
  comments and API documentation are the agent's closeout work. The "same as the production code"
  clause that pulled sessions toward documented blocks is reworded.
- **Context tending follows validation** — on a code project, the context edits noted at SETTLE
  wait for CONCLUDE, after the developer has applied and validated the guide; EXECUTE produces the
  guide and nothing else. A context project is unchanged. Stated in `mechanics/pipeline.md`
  (SETTLE, CONCLUDE, and a new invariant), `commands/start.md`, and `commands/close.md`.
- **SKILL.md is an index** — the entry file reduces to front matter, the description, and an index
  into the sub-layers: `@` pointers for always-loaded conventions (`behavior/planning.md`,
  `mechanics/pipeline.md`), `./` links for material consulted on demand. The content it restated
  moves to one home each: the reset schema to `mechanics/reset-file.md`, the `marathon.toml`
  canonical layout to the new `mechanics/configuration.md`, the planning philosophy (one step at a
  time, planning is half the work) to the new `behavior/planning.md`, and the hook table lives only
  in `mechanics/hooks.md`, summarized in prose by `references/extensions.md`.
- **Voice is out of scope** — the skill no longer defines or references a writing voice.
  Communication style is identity-level behavior and belongs to user-scoped configuration;
  `behavior/voice.md` is removed and the skill's `behavior/` tier holds workflow conduct only.

### Added

- **Value-of-information test** — `experiment` settles, alongside the question a spike answers, the
  decision the answer changes; a spike that changes no decision isn't run. `plan` weighs
  uncertainty against consequence when choosing the next focus.
- **Assumption annotations** — a `design/` or `concepts/` note names the unverified assumptions it
  rests on, and the reset Disposition records a falsification, so a surprise invalidates identified
  notes instead of forcing a judgment sweep. Stated in `references/context-engineering.md`.
- **Risk-first spiking** — builds proceed in dependency order while `experiment` probes the
  highest-consequence unknown out of band. Stated in `behavior/planning.md`.

## v0.7.0

### Added

- **Mechanics tier** — a new `mechanics/` directory holds the skill's execution specs, written for
  agent action. `mechanics/pipeline.md` defines the five-stage session pipeline every command runs —
  locate, start, settle, execute, conclude — and how each command layers into it; the nine command
  playbooks are slimmed to the content of their stages. `mechanics/hooks.md` is the hook firing spec:
  resolution of the enabled extension set, the firing table, and per-command ordering constraints.
- **Generic extension system** — the platform-tracker hook points (`on-init`, `on-session-start`,
  `on-commit`, `on-closeout`) become five universal hooks bound to the pipeline stages: `on-start`,
  `on-execute`, `on-commit`, `on-reset`, `on-close`, each firing before the moment it names. An
  extension is a separately installed skill; a repository enables it with an `extensions` key in
  `.claude/marathon.toml` — `[project]` for itself, `[workspace]` at a coordinator for every member —
  replacing `init`-time selection. `references/extensions.md` documents the system: installed vs.
  enabled, the extension's SKILL.md declaration, artifact bootstrapping, and the source-of-truth rule
  restated for repo-native extensions (an extension may own its artifact inside the repository;
  anything projected outward is a read-only mirror). `references/extension-hooks.md` is retired.
- **Behavior tier** — a new `behavior/` directory holds always-active conduct, mirroring the
  user-scope `behavior/` convention. The voice standard moves there
  (`references/writing-voice.md` → `behavior/voice.md`) and loads with the skill through an `@`
  reference in SKILL.md's new Behavior section, instead of waiting to be consulted; the standard now
  names its intent — natural language rooted in proper American English grammar, free of the patterns
  identifiable as machine-generated prose.

### Changed

- **SKILL.md reorganized around the sub-structures** — new Behavior and Mechanics sections initialize
  those tiers directly below the direct skill context; the Extensions section replaces Extension
  hooks.
- **`docs/` articulated as the standardized tier of context** — the lifecycle reads
  `concepts/` (volatile) → `design/` (settled) → `docs/` (standardized convention), with the
  established rule unchanged: a `docs/` page itself never decays. Stated in
  `references/context-engineering.md` and `commands/docs.md`.

## v0.6.0

### Added

- **Workspace docs centralization** — a project in a workspace no longer bootstraps its own `docs/`;
  documentation centralizes in one landing-zone project, named by a new optional `[workspace] docs`
  field in the coordinator's `marathon.toml` (an order key, alongside `role` and `order`). `docs` checks
  which case applies — standalone, this project is the landing zone, or another project is — before it
  bootstraps or curates anything. A standalone project is unaffected. `context-engineering.md`'s Decay
  rule gains a second target: a `design/` note also decays once a landing-zone page expresses it, not
  only when the code does. A repository whose own convention narrows or adds to a linked landing-zone
  page records that addition beside the link instead of duplicating the page.

### Fixed

- **Workspace entry routing** — `plan`, `start`, and `experiment` located their project by reading
  `context/reset.md` first, which left a workspace root (no `context/` of its own) an unhandled case
  discovered only by the absence turning up empty. A new `SKILL.md` section, "Finding your project,"
  states the two directory kinds — project and workspace root — as the first check every session-opening
  command makes, and the three commands now point to it ahead of their reset-file logic.

## v0.5.1

### Changed

- **Specificity in the writing voice** — `references/writing-voice.md` adds the discipline of naming
  what is actually happening in a detail rather than reaching for a stock noun or verb, a second
  test (does each term belong to the vocabulary of the discipline it describes?), the
  colon-versus-semicolon rule, and the stock-vocabulary habit in the illustrative aside.
- **Skill prose swept for stock vocabulary** — the commands and references now use the terms their
  subjects already have: the session works on a *step*, the capability map lists *capabilities*,
  the guide preamble states *what the change does*, `marathon.toml` has a canonical *layout*, and
  the role boundary is drawn per kind rather than coming in "shapes".

## v0.5.0

### Added

- **Resting point and workspace-entry routing** — `close` deletes the reset file when the project's
  deliverable is released, it has no next step of its own, and a workspace coordinator carries
  continuity; a session that finds no reset file defers to the coordinator, and a marathon command
  run from the workspace root resolves the coordinator through `[workspace] role` and routes into
  the member repository named by its reset file's Branch line.
- **`[workspace.paths]`** — an optional table in the coordinator's `marathon.toml` that resolves an
  `order` key living outside the workspace root to a directory on this machine.
- **LICENSE** — the repository is licensed Apache-2.0, recorded in the plugin manifest.

### Changed

- **Writing voice restructured** — `references/writing-voice.md` is principle-led: the goal (what a
  capable technical colleague would write), the discipline, and one applicable test, with the former
  habit list demoted to an illustrative aside. Adds the built/planned tense rule and generalizes the
  godoc section to API documentation.
- **Secondary sessions specified** — `plan` and `experiment` open with the same reset-file Status
  check as `start` and resume their own handoffs; `review` and `docs` close through `close`; the
  Session enum grows to `init | plan | start | experiment | review | docs`; branch creation and
  `on-session-start` are sequenced for every session; the session-loop diagram is redrawn so `reset`
  branches off the step.
- **Experiments are tracked** — spikes commit under `experiments/<slug>/` and merge with the branch
  as the durable record of exploratory work; promotion moves proven work out, and stable context
  never cites the directory.
- **Hook contract table** — `references/extension-hooks.md` states each hook's trigger, firing
  commands, and frequency; `init`'s setup commit and `reset`'s WIP commit fire `on-commit`.
- **Canonical configuration** — `SKILL.md` documents the whole `marathon.toml` in one example:
  `[project]`, `[remote]`, and the optional `[workspace]` and `[workspace.paths]`.
- **Coordinator conventions** — an organization-level coordinator's conventions bind member
  repositories through sessions: a member's `review` consults them, a coordinated fan-out applies
  them, and the awareness rule keeps member repos from citing them.
- **Decay rule refined** — a `design/` note decays only when the built work expresses it and the
  note holds no conceptual or pattern detail beyond it; the reset ledger vocabulary (Integrated,
  Retained) is mapped to the operations.
- **Prose normalized** — the skill corpus is rewritten against the restructured voice standard, with
  root-relative cross-references throughout.

## v0.4.0

### Added

- **Writing-voice standard (`references/writing-voice.md`)** — the voice for every piece of prose the
  agent is responsible for: design notes and concepts, reset files, implementation guides, godoc and
  `doc.go`, in-source comments, prose documentation, profiles, and the skill files themselves. Plain
  technical-documentation voice, concrete nouns, objective implementation detail; a list of habits to
  avoid. godoc keeps its idiomatic form. Cited from `SKILL.md`, the role-boundary and
  implementation-guides references, and the commands that author prose.

## v0.3.0

### Added

- **Project kinds (`code` / `context`)** — a project declares its kind at `init` in
  `.claude/marathon.toml`. A `code` project holds production source the developer owns: `start` drafts
  an implementation guide, the developer applies it, and closeout adds tests and documentation. A
  `context` project *is* context — skills, prose, configuration — which the agent authors directly under
  the developer's review, with no guide and no tests. The role boundary now has two shapes, one per kind.
- **`plan` sub-command** — a planning-and-curation session that touches only `context/`: create and
  refine concepts, settle a design, and decide what the next `start` should focus on. Forward-looking,
  where `review` is the backward-looking drift audit. Lands on a branch like `review`.
- **`experiment` sub-command** — a spike in the isolated top-level `experiments/` directory; results are
  concepts, promoted deliberately at closeout, if at all. (Previously `start experiment`.)
- **Workspace coordination (`coordinate`)** — run one change across several marathon projects that live
  as siblings in a workspace. `coordinate` detects the workspace, reads a coordinator project's declared
  dependency `order` (layered, with adjacent peers as sub-arrays), and fans a session out to each project
  in order, honoring its kind. The workspace holds no context; continuity stays per repository.
- **Plugin README** — a quick-reference `README.md` for the plugin directory.

### Changed

- **Sub-command restructure (breaking).** `start <development|context|experiment>` is replaced by
  top-level commands: `plan`, `start`, and `experiment`. `start` no longer takes a type argument — it
  advances the product one step, resolving by project kind. Older `reset.md` files carrying a
  `Session type:` line are still read.
- **Vocabulary** — the notes in `concepts/` are consistently called *concepts*, not *candidate notes*.
- The role-boundary, session-loop, and context-engineering references are generalized so nothing assumes
  a code-only, single-repository model.

## v0.2.0

### Added

- **Human-oriented `docs/` tier** — an optional top-level `docs/` directory for reference documentation
  written for people, a peer to the agent-oriented `context/`. It shares context engineering's
  maintenance discipline (curate it, keep it in sync) but not its lifecycle: a `context/` note decays
  once the code expresses it, while a `docs/` page exists to explain code that already does. A repository
  opts in at any time; most stay `context/`-only.
- **`docs` sub-command** — the deliberate authoring and curation pass for `docs/`. On a repository with
  no `docs/`, the first run bootstraps the tier; later runs extend and restructure it. Plan-mode-driven,
  no code handoff, landed on a branch like `review`.
- **Docs drift in `review`** — `review` now also flags `docs/` pages that have drifted from the code,
  once a `docs/` tier exists. The core build loop is unchanged; documentation is never cram-written at
  closeout and never silently rots.

## v0.1.0

Initial release — the standalone core workflow.

### Added

- **Concept-driven `init`** — evaluate a project-planning concept in plan mode and align with the
  developer, then scaffold a project's top-level `context/` (orientation, `design/`, `concepts/`, and
  single `guide.md`/`reset.md` files) in one pass.
- **Typed sessions** — `start` a `development`, `context`, or `experiment` session. A fresh session
  plans in plan mode and settles scope with the developer before any implementation guide is written;
  a session resuming from a handoff picks up the prior plan in place.
- **Mid-session handoff** — `reset` captures in-flight state and a resume pointer without closing
  the work, so a fresh context window picks up where the last left off.
- **Closeout** — `close` runs the reset transaction (integrate / promote / cull / retain), decays
  design that code now expresses, deletes the spent implementation guide, commits, and publishes the
  branch (pull request, merge request, or the project's equivalent) with its description from the reset
  file.
- **Drift review** — `review` enters plan mode to audit design-vs-code drift and aligns with the
  developer before culling or promoting context.
- **Context engineering** — a volatile-vs-stable context model with deliberate promotion and decay,
  keeping the repository the single source of truth. Any operation that culls or promotes context
  shows its dispositions for developer alignment first.
- **Role boundary** — the developer owns production source; the agent owns tests, documentation,
  and context artifacts.
- **Remote platform** — the remote and its publish command (`gh pr create`, `glab mr create`, or
  another) are declared at init and stored in `.claude/marathon.toml`; the core's git workflow stays
  platform-neutral.
- **Extension hooks** — named, no-op-by-default hook points (`on-init`, `on-session-start`, `on-commit`,
  `on-closeout`) for opt-in platform project-management extensions selected at init.
