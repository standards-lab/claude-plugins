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
`marathon currency` then ignore its own `mise.toml`, and every adoption goal's validate and
currency steps go through `marathon`. Separately, the plugin's Stop hook blocks a marathon
format with a line over 100 columns and hands the model a rewrapped copy to print again, so the
architect reads each such reply twice.

**Behaviors**
1. A command run through a shim without a local build sees the caller's environment, unchanged:
   no `MISE_OVERRIDE_*` variable the caller didn't set, and the caller's own values as they were.
   A `marathon check` whose configured command runs `mise run` reads the repository's `mise.toml`.
2. The shim still runs the release at the plugin's version, resolved while ignoring the caller's
   mise config and `.tool-versions`. A version not yet installed installs on first use, and a
   local build still wins.
3. claude-plugins' `marathon check` exits 0.
4. The plugin's hook rewraps a marathon format on screen as the reply streams, from the format's
   heading to the end of the message, with today's rewrap rules: lines over 100 columns outside
   code fences and tables, list and blockquote markers kept, continuation lines indented, and no
   continuation starting on a block-start word. The model is never asked to reprint, and a reply
   holding no format displays untouched. A failure or timeout displays the original.
5. The version reads 0.22.0 in the plugin manifest, the skill, each shim, and the pin example in
   the configuration mechanics. The CHANGELOG has a v0.22.0 entry whose Migrating note names the
   hook change and the 0.22.0 pin.

**Test seams**
- `scripts/tests/shims.test.sh`, through the check.
- `wrapformats`' tests, over `MessageDisplay` input across batches and messages.
- `marathon check` in claude-plugins.
- A live session, for the on-screen rewrap (the session brief's Verify).

**Slices**
1. The shims resolve and exec the release, and the tests prove behaviors 1–3.
2. wrap-formats as a `MessageDisplay` hook (behavior 4).
3. Version 0.22.0 and its CHANGELOG entry (behavior 5).

**Out of scope**

Rewrapping files such as the session brief before publishing; replies with no marathon format;
the Go binaries' other commands; adopting repositories' pins.

**Door**

One-way for the tag: once pushed, every installation at 0.22.0 runs these shims and this hook.
Everything else is two-way.

**Release**
- claude-plugins: `marathon/v0.22.0`

## Progress

slices 3/3 committed · standards — · spec — · editor —
- redirect, BUILD: fold the Stop hook's reprint into `shim`; re-planned: wrap-formats moves to a
  `MessageDisplay` hook, and the release becomes v0.22.0.
- slice 2, decided: hook state also opens for a fence still open at a batch's end, so a heading
  inside it isn't a format; one overlong word stays unwrapped, with no report.
- slice 3: the line-width eval grader cannot score displayed text (no grader type runs a command); escalated.

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
  the command sees the caller's environment; released with the hook change as marathon/v0.22.0.
- shim: wrap-formats runs as a `MessageDisplay` hook that rewraps a format on screen; the model
  never reprints.
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
