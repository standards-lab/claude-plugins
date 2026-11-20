# goal · harness.adoption

- **State:** brief ready
- **Task:** go-tools
- **Branch:** go-tools

## Tasks

1. [x] references
2. [ ] go-tools
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

## Task brief · go-tools

**Problem**

Every adopting repository is about to wire marathon's tools into its check and CI, and today
those tools are Python scripts. No repository in the organization pins Python. The architect
rejected adding it to about 25 repositories, and porting first spares them a second rewiring.
Go release binaries pinned through mise need no language runtime in the consumer:
- CI gets them from its existing `mise install`;
- `mise outdated` reports a newer tag;
- the web repositories, which have no Go, still get a binary.

**Behaviors**
1. lint-context, lint-references, marathon, and wrap-formats are Go commands in one module at
   claude-plugins' root `marathon/`. Each keeps 0.20.0's arguments, output, and exit status. The
   ported tests compare stdout, stderr, and exit status as the Python tests did.
2. lint-references counts a line that starts with `*` as a comment only inside a `/*` block
   comment, so a `.gitignore` pattern such as `*.md` is not a reference.
3. The module follows the architecture Go pages that claude-plugins' `STANDARDS.md` points to:
   - topology-and-naming, dependencies, construction-and-execution, tests-and-docs, the
     release-and-ci check, comments, and tool-beside-library;
   - logic in internal packages, and a thin standard-library `main` for each tool;
   - dependencies: the standard library plus `BurntSushi/toml`.

   `STANDARDS.md` states two exemptions:
   - from cli-applications, domain-files, and lifecycle-and-context: the tools don't build on
     go-cli-sdk;
   - from release-and-ci's Keep-a-Changelog form: marathon's CHANGELOG is the module's record.
4. Each `marathon/vX.Y.Z` release carries four archives, for linux and darwin on amd64 and
   arm64. Each archive holds all four binaries, beside a `SHA256SUMS`, and each binary reports
   the tag's version. Other plugins' tags release as before, with no files.
5. A repository pins the tools as
   `"github:standards-lab/claude-plugins" = { version = "0.21.0", version_prefix = "marathon/v" }`.
   `mise install` puts all four on PATH, and `mise outdated` reports a newer marathon tag.
6. In a session, each of the plugin's `bin/` commands and the Stop hook runs the binary from the
   gitignored local build when it is present. Otherwise it `mise exec`s the release that matches
   the plugin's own version, ignoring the repository's mise config. claude-plugins' check fails
   when a shim's version differs from `plugin.json`'s. A failed shim in the Stop hook doesn't
   block the reply.
7. The release gate cross-builds all four targets and fails on a broken build. It builds the
   host's binaries into the local build for the eval and removes them when it exits.
8. claude-plugins' check:
   - runs the Go check from release-and-ci: build, vet, gofmt, `go fix -diff`,
     `go mod tidy -diff`, `go test -race`, and golangci-lint;
   - tests the eval graders' patterns against the stored plan rounds with `node --test`;
   - tests itself with a bash script;
   - fails when any Python file is tracked.

   Its `mise.toml` pins Go, Node, and golangci-lint. CI installs through mise, and the currency
   command adds `mise outdated`.
9. marathon's docs:
   - The requirements name mise, not Python.
   - The CI section shows the mise pin in place of the `curl` download.
   - `init`'s check calls the pinned lints.
10. marathon's CHANGELOG `v0.21.0` has `### Migrating from 0.20.0`:
    1. Install 0.21.0 with mise on PATH.
    2. Pin `github:standards-lab/claude-plugins` with `version_prefix = "marathon/v"` at 0.21.0
       or later in each repository's `mise.toml`.
    3. In CI, call `lint-context` and `lint-references` after `mise install`, and drop any
       `curl` download.

    The plugin's version, the skill's Version line, and the CHANGELOG read 0.21.0. The budgets
    hold, and the gate passes.

**Test seams**
- Each Go command, run on a temporary repository; the tests compare stdout, stderr, and exit
  status.
- `node --test` over the graders and the stored plan rounds.
- claude-plugins' check and release gate, including the cross-build and the eval.

**Slices**
1. The module and lint-references (behaviors 1–3 for it, and 8):
   - `mise.toml`, CI through mise, the Go check, and the currency command;
   - `STANDARDS.md`'s Go pointers and exemptions.
2. lint-context in Go, with its tests (behavior 1).
3. marathon and wrap-formats in Go, with their tests (behavior 1).
4. The shims and the local build (behavior 6). The Python tools go, check.sh's tests become
   bash, the grader tests move to `node --test`, and the no-Python check lands (behavior 8).
5. The release workflow and the gate's cross-build (behaviors 4, 7).
6. The docs and `Unreleased` CHANGELOG entries with migration steps 2 and 3 (behaviors 5, 9).
7. The release: `Unreleased` becomes `v0.21.0` with migration step 1, the versions agree, and
   the gate passes (behavior 10).

**Out of scope**

- Any adopting repository's pin, check, or CI: those are the adoption tasks.
- Windows builds.
- Schemas, a `spec/` directory, and new tool commands (`v1.harness.tooling`).
- Any marathon-architecture release.

**Door**

The `marathon/v0.21.0` tag and its release files are one-way: consumers pin them, and a release
is never re-cut. Everything before the tag is two-way.

**Release**

- `marathon/v0.21.0`

## Progress

slices 7/7 committed · standards ✓ · spec ✓ (2 rounds) · editor ✓

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
