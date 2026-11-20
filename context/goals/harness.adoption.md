# goal · harness.adoption

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] references
2. [x] go-tools
3. [ ] ai-hosting
4. [ ] core
5. [ ] docproc
6. [ ] go-ai
7. [ ] go-libraries
8. [ ] templates
9. [ ] web
10. [ ] go-web-service
11. [ ] architecture
12. [ ] spikes

## Decisions

- plan: staged over the locks of v1.ai.processing, v1.client.stack, v1.ai.inference, and
  v1.ai.hosting, a one-time exception; no session runs on them until harness.adoption and then
  v1.ai.hosting's `secrets` task ship.
- plan: architecture's `scripts/check.sh` switches to marathon's `lint-references` (the
  `architecture` task).
- plan: the 33 archived spikes stay as they are; only the 4 unarchived spikes adopt.
- plan: webappgen sits on its local, unpublished `template` branch: its adoption branches from
  main, and once merged, main merges into `template`. go-storage's modified `go.work.sum` is
  committed with its adoption.
- references: rejected a pinned architecture checkout in CI, so `architecture/...` references
  skip there: architecture has no tags, so it would pin a bare SHA no currency command tracks.
- ai-hosting: `context/design.md` splits into one note per planned capability, each under 1,500
  words and linked from a `(planned)` map entry: engines, the `hosting` suite, host-class
  profiles, gateway, secrets, Foundry embedding overflow, and cleanup. The H/G/P/Q findings table
  goes (git history keeps it); settlements only it states move to their tasks' notes;
  v1.ai.hosting's `design:` decision and P14 pending edit are reworded, its `secrets` brief
  untouched; ai-hosting's `STANDARDS.md` source line changes.
- ai-hosting: the Foundry embedding-overflow prices lose their dates and keep their figures and
  source URLs; the `embedding` task re-reads them before relying on them.
- adoption: comments outside the five kinds go where the code already says them (a key name, a
  mise task description) or where they describe planned work; a script's section comments become
  named functions; a reference comment points where the reasoning has a home, and a new `docs/`
  page holds only reasoning with none. In ai-hosting this removes 15 lines and creates no `docs/`.
- adoption: each repository's `STANDARDS.md` comment rule covers every configuration file (TOML,
  YAML, `.gitignore`), as claude-plugins' does.
- ai-hosting: `scripts/check.sh` keeps its own link check over every tracked `.md` beside
  `lint-context`, which doesn't check `CLAUDE.md` or `STANDARDS.md`.

## Pending edits

- coordinator, once ai-hosting splits `context/design.md`: standards-lab's
  `context/ai-hosting.md`, `docs/ai-hosting.md`, and `roadmap/v1.ai.hosting.toml` name it; point
  each at the note that now holds its subject.
