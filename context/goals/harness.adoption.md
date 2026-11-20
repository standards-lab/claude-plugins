# goal · harness.adoption

- **State:** building
- **Task:** shim
- **Branch:** shim

## Tasks

1. [x] references
2. [x] go-tools
3. [x] ai-hosting
4. [ ] shim
5. [ ] architecture

## Task brief · shim

**Problem**

marathon 0.21.0's shims run the release through `mise exec` with two `MISE_OVERRIDE_*` variables
set, and both reach the command marathon runs. A repository's `marathon check` and
`marathon currency` then ignore its own `mise.toml`. Every adoption goal's validate and currency
steps go through `marathon`, so the 14 sibling goals wait on this fix.

**Behaviors**
1. A command run through a shim without a local build sees the caller's environment, unchanged:
   no `MISE_OVERRIDE_*` variable the caller didn't set, and the caller's own values as they were.
   A `marathon check` whose configured command runs `mise run` reads the repository's `mise.toml`.
2. The shim still runs the release at the plugin's version, resolved while ignoring the caller's
   mise config and `.tool-versions`. A version not yet installed installs on first use. A local
   build still wins, and the Stop hook still passes when the release fails.
3. claude-plugins' `marathon check` exits 0.
4. The version reads 0.21.1 in the plugin manifest, the skill, each shim, and the pin example in
   the configuration mechanics. The CHANGELOG has a v0.21.1 entry whose Migrating note says pins
   may move to 0.21.1.

**Test seams**
- `scripts/tests/shims.test.sh`, through the check.
- `marathon check` in claude-plugins.

**Slices**
1. The shims resolve and exec the release, and the tests prove behaviors 1–3.
2. Version 0.21.1 and its CHANGELOG entry (behavior 4).

**Out of scope**

The Go binaries; adopting repositories' pins (ai-hosting's currency reports 0.21.1, and the
sibling goals pin it).

**Door**

One-way for the tag: once pushed, every installation of the plugin at 0.21.1 runs these shims.
Everything else is two-way.

**Release**
- claude-plugins: `marathon/v0.21.1`

## Progress

slices 0/2 committed · standards — · spec — · editor —

## Decisions

- plan: staged over the locks of v1.ai.processing, v1.client.stack, v1.ai.inference, and
  v1.ai.hosting, a one-time exception that the sibling adoption goals share; no session runs on
  those goals until every adoption goal syncs (standards-lab/context/harness-adoption.md).
- plan: the rest of the adoption runs as 14 sibling `harness.adoption-<slug>` goals, in one
  orchestrating session once `shim` releases marathon/v0.21.1; the rules they share live in
  standards-lab/context/harness-adoption.md.
- plan: architecture's `scripts/check.sh` switches to marathon's `lint-references` (the
  `architecture` task).
- shim: the shims resolve the release with `mise where` and exec its binary, never `mise exec`, so
  the command sees the caller's environment; released as the patch marathon/v0.21.1.
- references: rejected a pinned architecture checkout in CI, so `architecture/...` references
  skip there: architecture has no tags, so it would pin a bare SHA no currency command tracks.

## Pending edits

- coordinator · standards-lab: repoint each mention of ai-hosting's removed `context/design.md`
  at the ai-hosting note that now holds its subject:
  - `context/ai-hosting.md`, "What `ai-hosting` holds", the suite's design:
    `ai-hosting/context/hosting-suite.md`.
  - `docs/ai-hosting.md`, the opening paragraph's design note (each task's spike parts and open
    points): ai-hosting's capability notes, mapped in `ai-hosting/context/README.md`; "The spikes'
    answers", e5's scope and price on Foundry: `ai-hosting/context/foundry-overflow.md`; "The
    `gateway` task", the six open points: `ai-hosting/context/gateway.md`, "Open at the gateway
    plan".
  - `roadmap/v1.ai.hosting.toml`: `tasks.secrets`' "design's note": `ai-hosting/context/secrets.md`;
    `tasks.gateway`'s "what design's note defers to it": `ai-hosting/context/gateway.md`, "Open
    at the gateway plan"; `tasks.cleanup`'s "the design note's inventory":
    `ai-hosting/context/cleanup.md`. `tasks.design`'s summary describes the note that task wrote,
    and stays.
