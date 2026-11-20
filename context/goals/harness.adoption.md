# goal · harness.adoption

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [ ] references
2. [ ] ai-hosting
3. [ ] core
4. [ ] docproc
5. [ ] go-ai
6. [ ] go-libraries
7. [ ] templates
8. [ ] web
9. [ ] go-web-service
10. [ ] architecture
11. [ ] spikes

## Decisions

- plan: staged over the locks of v1.ai.processing, v1.client.stack, v1.ai.inference, and
  v1.ai.hosting, a one-time exception; no session runs on them until harness.adoption and then
  v1.ai.hosting's `secrets` task ship.
- plan: the coordinator's roadmap moves to `roadmap/index.toml` in a `plan harness.adoption`
  session right after `references` ships and is installed, not at sync.
- plan: one reference check, shipped by marathon beside `lint-context` and fetched pinned in CI;
  architecture's `scripts/check.sh` switches to it. How a code repository resolves an
  `architecture/...` reference without an architecture checkout in CI is `references`' question.
- plan: the 33 archived spikes stay as they are; only the 4 unarchived spikes adopt.
- plan: webappgen sits on its local, unpublished `template` branch: its adoption branches from
  main, and once merged, main merges into `template`. go-storage's modified `go.work.sum` is
  committed with its adoption.

## Pending edits

- none
