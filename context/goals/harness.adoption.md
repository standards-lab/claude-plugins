# goal · harness.adoption

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] references
2. [x] go-tools
3. [x] ai-hosting
4. [ ] shim
5. [ ] architecture

## Decisions

- plan: staged over the locks of v1.ai.processing, v1.client.stack, v1.ai.inference, and
  v1.ai.hosting, a one-time exception that the sibling adoption goals share; no session runs on
  those goals until every adoption goal syncs (standards-lab/context/harness-adoption.md).
- plan: the rest of the adoption runs as 14 sibling `harness.adoption-<slug>` goals, in one
  orchestrating session once `shim` releases marathon/v0.21.1; the rules they share live in
  standards-lab/context/harness-adoption.md.
- plan: architecture's `scripts/check.sh` switches to marathon's `lint-references` (the
  `architecture` task).
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
