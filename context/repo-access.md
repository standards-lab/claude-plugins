# Repository access: what a goal holds, and what CI it waits on

Repository access is planned work, tracked as `marathon.repo-access`. It changes how marathon
learns whether a repository runs CI and which repositories a goal holds, and it carries this
repository's open review findings.

## Task `ci-awareness`

A repository declares in `.claude/marathon.toml` whether it runs CI, so SHIP, merge, and the briefs
know whether a pull request will have checks, and no session scripts a wait for CI that doesn't
exist. Today each repository's `[remote] merge` differs by hand: `gh pr checks --watch && gh pr
merge …` where CI runs, and `gh pr merge …` where it doesn't.

- **The declaration**: an optional `[remote] checks` command, such as `gh pr checks --watch`, not a
  boolean. SHIP runs `marathon checks` before `marathon merge` when it is set and merges directly
  when it isn't; `merge` becomes `gh pr merge --merge --delete-branch` everywhere, and `[remote]
  ci` stays the post-merge wait releasing runs. The session brief's Merge danger line says whether
  each repository's merge waits on CI.
- **Migration**: the other repositories move through the CHANGELOG's "Migrating from 0.22.0" step,
  each on its own next task; none joins this goal's `repos`.
- **Running commands**: pipeline.md's "Committing" widens into "Running commands": every command
  that changes git or remote state runs in its own tool call, after the previous one succeeds.
- **Currency**: the task opens by bumping both workflows' `jdx/mise-action` from v5.2.10 to v5.3.0
  and fixing CLP-1: `scripts/currency.sh` prints every `mise outdated` entry unfiltered, so a
  pinned tool that isn't installed reports `-> null` and exits 1. Select only entries whose
  `.value.bump // .value.latest` differs from `.value.requested`.

## Task `review-findings`

The goal also fixes the claude-plugins findings of the `harness.adoption` architecture review,
each with its fix. CLP-1 is in `ci-awareness`; the rest tag nothing, and the marathon-architecture
changes release with the goal:

- CLP-2: `harness-testing.md`'s eval case has five graders and no line-length grader, and one
  numbered question suffices. Say "five graders", drop the 100-character bullet, and require at
  least one numbered question with a `rec:` line and then a `changes:` line.
- CLP-3: `harness-testing.md` says CI installs what `mise.toml` pins, but Claude Code moved to
  `mise.ci.toml`. Name both files, the second read only under `MISE_ENV=ci`, and say
  `scripts/currency.sh` reports a tool or GitHub Action pin that trails its latest release.
- CLP-4: the root README's tree omits `ci.yml`, `mise.ci.toml`, `STANDARDS.md`, and `scripts/`.
  Add them.
- CLP-5: `harness-optimization.md`, "What the goal measures against", is history. Remove it.
- CLP-6: `harness-optimization.md`'s opening and its "One release" entry are change history, and
  task `retro-routes` is in no roadmap file. Open with "marathon's rules are checks, not prose",
  and remove "One release" and the `retro-routes` sentence.
- CLP-7: `harness-optimization.md`'s LOCATE and status CLI entry belongs in
  `tool-based-skills.md`, as a first consumer: a CLI for LOCATE and `status`.
- CLP-8: `marathon-extraction.md`'s last open question is settled by marathon's
  `references/extensions.md`. Remove it.
- CLP-9: `harness-optimization.md`'s judged-suite sentences duplicate `harness-testing.md`.
  Remove them.
- CLP-10: `harness-testing.md` stamps eval with version `2.1.289`. Write "`claude plugin eval`
  runs each case as".
- CLP-11: "the blueprint" and "a consumer of the standard" are undefined. Define both in
  `marathon-extraction.md`'s first paragraph.
- CLP-12: `STANDARDS.md`'s data-file line restates the comments page; end it at "(TOML, YAML,
  `.gitignore`)". "Plugin design lives in the skill" sits in both `CLAUDE.md` and
  `context/README.md`; keep the `CLAUDE.md` one.
- CLP-13: `mise.ci.toml`'s header holds reasoning its `See` target lacks. Make it "The tools CI
  adds to mise.toml's, read only where MISE_ENV=ci is set." with a reference to
  `harness-testing.md`, and move the reason there.
- CLP-14: marathon-architecture's `references/architecture-layer.md` names only the
  `architecture/principles/<page>.md` form, while spikes write `principles/comments.md`. Add that a
  spike outside the workspace directory writes the path relative to the architecture repository.
- CLP-15: what the architecture layer may hold is stated in marathon-architecture's README, its
  `references/architecture-layer.md`, and architecture's README, and the self-containment rule
  only in the last. Shrink the plugin README's section to a link to `architecture-layer.md`, and
  add the self-containment rule there if every layer should hold it.

## Task `shared-repos`

A goal owns its root repository, locked for the goal's life. Its other repositories are not locked
by the goal: each is held only while it is off its default branch, with a task's branch checked
out. Several goals can then change the same non-root repository, queued by whether it is on its
default branch. Today marathon's `mechanics/goals.md`, "The repository lock", locks every
repository in a goal's `repos`.

The task also carries the goal's single release, `marathon/v0.23.0` and
`marathon-architecture/v0.4.1`, and moves claude-plugins' own configuration to `[remote] checks`
with it, since the installed 0.22.0 copy and binary don't know the key.

## Assumptions

- A repository's default branch is the only branch a goal finishes on, so "on its default branch"
  means no goal is mid-task there.

## Open questions

- What a queued goal does when its non-root repository is held.
- How the coordinator's `[workspace] exclusive` groups and a goal's `order` interact with the new
  lock.
- How LOCATE and `status` report a held non-root repository.
- Whether a held non-root repository needs the holding goal named on disk.
