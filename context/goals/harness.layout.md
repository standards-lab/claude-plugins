# goal · harness.layout

- **State:** building
- **Task:** marathon-layout
- **Branch:** marathon-layout

## Tasks

1. [x] lint-references
2. [x] standards
3. [ ] marathon-layout
4. [ ] archive
5. [ ] templates
6. [ ] capstones
7. [ ] core
8. [ ] docproc
9. [ ] libraries
10. [ ] go-ai
11. [ ] messaging
12. [ ] docinfer
13. [ ] web
14. [ ] webappgen
15. [ ] service
16. [ ] tools

## Task brief · marathon-layout

**Problem**

The architecture layer now states the layout that separates what people read from what agents
use, but neither marathon nor its own repository follows it. `init` sets up repositories without
the agent area or the `agent` profile. Session briefs are written inside Claude Code's own
directory. A brief's Verify can send the architect to agent harnesses. A session that ships ends
without saying how to start the next one. Every migration task after this one runs marathon and
is reviewed against the layout, so marathon and claude-plugins take it first.

**Behaviors**
1. A session brief's Verify rules out agent harnesses, agent tasks, and internal test commands.
   It still opens with the check's one-line result, and each step is something a person runs
   against the product, with what they should see. The format shows one good step and one bad
   step. The bad step runs an agent's scratch harness or a single package's tests. The good step
   runs the shipped command and names its visible result.
2. The session brief is written to `.agent/state/briefs/<goal>.md` in the goal's root. When the
   root's ignore file doesn't list `.agent/state/`, BRIEF first adds that line on the task's
   branch.
3. `init` writes these files and creates no task directories:
   - ignore lines for `.agent/state/` and `mise.local.toml`, beside `.claude/plans/`
   - a `mise.agent.toml` whose task config includes `mise-tasks` and `.agent/mise-tasks`
   - `MISE_ENV=agent` in the `env` of the Claude Code settings
4. A `start` session that ships ends its last message with the prompt that starts the next
   session, fenced as code:
   - after the goal's next task: `/marathon:marathon start <goal>`
   - after a sync that staged a goal: `/marathon:marathon start <staged goal>`
   - after a sync that only proposed a goal: `/marathon:marathon plan <goal>`
   - after a spike's sync with no planned spike left:
     `/marathon:marathon intake experiment.<topic>`

   A session that stops in `handoff` ends with no prompt, and the session brief carries none.
5. The read-first budgets for `init` and `start` rise to their new sizes, rounded up to the next
   hundred, and the Unreleased entry gives the reason. Every other budget stays.
6. claude-plugins takes the layout:
   - Its tasks run through mise. `check`, `currency`, `gate`, and `build` are one-line tasks over
     the existing scripts, which keep their paths.
   - Its check is three step scripts: plugin consistency, the Go check with its lints, and the
     repository's tests. A person's `check` runs all three in order and still reports every
     failure. Under the `agent` profile, three agent tasks each run one step, and a person's task
     list doesn't show them.
   - Its `mise.agent.toml`, settings `env`, and ignore lines match what `init` writes. It keeps
     ignoring `.claude/briefs/` while the installed marathon still writes briefs there.
   - Each Go package's unit tests sit in that package's `test/` directory, and the Go check passes
     unchanged.
   - Its standards record one exception: marathon's local build stays beside the plugin's shims,
     where the shims find it, rather than in `bin/`.
7. Each change has an entry under the plugin's Unreleased heading. Nothing is tagged.

**Test seams**
- The consistency script, which checks citations, read-first budgets, the Go check, and its own
  bash tests. Those bash tests cover the split check.
- mise's task list with and without `MISE_ENV=agent`.
- A read of the changed formats and playbooks for coherence.

**Slices**
1. **Upgrade claude-plugins.** The CI and release workflows pin mise-action v5.4.0. Done when the
   currency command exits 0 and the check passes.
2. **Go tests in each package's `test/`.** Each marathon package's tests move into its `test/`
   directory and stay external test packages. Done when the check passes.
3. **The check in three steps.** Three step scripts, a check that runs them in order, and the
   check's own tests adapted. Done when the check passes.
4. **claude-plugins' mise tasks and agent profile.** Lands the inline tasks, `mise.agent.toml`,
   the agent tasks that wrap each step, the settings `env`, the ignore lines, the standards line
   for the local-build exception, and current notes and README. Done when mise lists the step
   tasks only under the `agent` profile and the check passes.
5. **Verify and the brief's new path.** The session brief format gains the Verify rule with its
   good and bad step, and the new path. The BRIEF step takes the new path and ignore fallback,
   and the eval fixture's ignore file follows. Done when the check passes and the formats read
   coherently.
6. **init writes the layout.** The setup tree gains the new ignore lines, `mise.agent.toml`, and
   the settings `env`. Done when init's budget is raised with the reason in the Unreleased entry,
   and the check passes.
7. **The next-session prompt.** `start`'s Ship section states the four cases and the handoff
   exception. Done when start's budget is raised with the reason in the Unreleased entry, and the
   check passes.

**Out of scope**
- A check or lint that enforces the layout. It goes to the backlog, for after `tools`.
- An eval case for Verify or the prompt, since no failure has been observed.
- A next-session prompt for `plan`, `experiment`, or `intake` sessions.
- Moving the scripts into file tasks.
- The architecture repository's own layout, which the new `architecture` task covers.
- Any release: marathon releases once, at `archive`.

**Door**

Two-way. Nothing is tagged, the changes land only on claude-plugins' default branch, and
installed consumers keep running 0.24.1 until `archive` releases.

## Progress

slices 0/7 committed · standards — · spec — · editor —

## Decisions

- plan: rejected `_agent/` for the agent area: Go and the web tools treat both alike, and only the
  dot form is hidden, as `.claude/`, `.state/` and `.build/` are.
- plan: an existing `.state/` becomes the ignored `.agent/state/`.
- plan: the session brief's Verify keeps the check's one-line result; its steps never cite agent
  harnesses or internal test commands, with a good and bad example.
- marathon-layout: `archive` is reopened: one central archive reference, in source control in
  one place, sharded so no file outgrows one context, and no per-parent `index.md` scattered across
  the filesystem. `archive`'s PLAN settles its location, sharding, and command shape.
- plan: building the Go `examples/` modules and the web examples is the `examples` backlog goal,
  not this goal's.
- plan: each migration takes its repositories' currency; go-web-service adapts go-core v0.13's
  `EventKind.Raise` at its 12 calls. Only the templates and claude-plugins release; library and
  web moves ride their next feature release.
- plan: `capstones` runs while the architect pauses the capstone goals.
- lint-references: marathon releases once, at `archive`, after the goal's plugin tasks;
  lint-references and marathon-layout leave Unreleased entries.
- standards: this goal's architecture pages run ahead of the code: its criteria commit every
  repository to the layout, and each migration's standards review checks against the pages.
  Rejected writing them after `tools`: every migration would be reviewed against pages that
  forbid it.

- marathon-layout: the session brief lives at `.agent/state/briefs/<goal>.md`; rejected
  `.claude/briefs/`, Claude Code's own directory. Kept `.agent/state/` over collapsing it into
  `.agent/`, which also holds tracked agent tasks.
- marathon-layout: a shipping `start` ends its last message with the fenced
  `/marathon:marathon <command> <goal>` (start, plan, or intake), and with none on handoff;
  rejected bare `marathon`, which is the Go binary on PATH.
- marathon-layout: claude-plugins keeps `scripts/` behind one-line mise tasks, and its local build
  stays at `plugins/marathon/.build/`, where the shims find it.
- marathon-layout: no layout lint during the migrations; review holds the layout.

## Pending edits

- coordinator · roadmap/harness.layout.toml: drop "Releases marathon." from the
  lint-references and marathon-layout summaries; `archive` keeps the goal's one marathon release.
- coordinator · roadmap/harness.layout.toml: end the goal's summary with "The architecture's
  agent-separation principle and the Go and Web Elemental standards state the layout." in place
  of "The goal record's Decisions hold the layout."
- coordinator · roadmap/harness.layout.toml: task `web`'s summary adds that a package's `demo/`
  becomes `examples/` (web-components and web-foundation have a `demo/`).
- coordinator · roadmap/harness.layout.toml: add task `architecture` (the architecture repository
  takes the layout: ignore lines, `mise.agent.toml`, settings `env`; repos = ["architecture"])
  after `archive` in `order`; the record's Tasks gains it at the same position.
- coordinator · roadmap/harness.layout.toml: rewrite `archive`'s summary, and the goal summary's
  "`archive` into an `index.md`", to the central, sharded archive reference.
- coordinator · roadmap: a backlog goal for a layout lint, after `tools`: the ignore lines,
  `mise.agent.toml`, the settings `env`, and Go test placement.
