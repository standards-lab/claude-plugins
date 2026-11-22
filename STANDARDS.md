# claude-plugins standards

The judgement calls the standards-reviewer applies to claude-plugins, beyond what `scripts/check.sh` enforces.

- Instructions state the present; a removal lives only in the CHANGELOG.
- A skill programs its own mechanics, never user-scope behavior.
- A skill's configuration and state live in declared files in the repository.
- A command runs through the one pipeline; an extension acts only at hook points.
- SKILL.md indexes its files and never restates them.
- A procedural section in a skill is a candidate tool.
- An instruction that changes no behavior is cut.
- Each change gets an Unreleased entry; a change to the extension contract bumps the floor.
- A read-first budget rises only with the architect's approval, with the reason in the Unreleased entry.
- An eval case comes from an observed failure.
- `architecture/principles/comments.md` governs scripts, tests, and every configuration file (TOML, YAML, `.gitignore`).
- `architecture/principles/rolling-currency.md`
- `architecture/principles/independent-releases.md`
- `architecture/principles/context-architecture.md`
- `architecture/principles/minimal-footprint.md`
- `plugins/marathon/skills/marathon/references/build.md`, "Checks come first", governs the repository's scripts: a mechanical rule goes in the check, and a failure message names its fix.
- `architecture/principles/tool-beside-library.md`: each command under `marathon/cmd/` is a thin standard-library `main` over the `marathon/internal/` package that holds its logic.
- `architecture/standards/go-elemental/principles/topology-and-naming.md`: the one module rooted at `marathon/`, tagged `marathon/vX.Y.Z`, whose `cmd` packages import only `internal` ones.
- `architecture/standards/go-elemental/principles/dependencies.md`: the `marathon/` module takes the standard library plus `github.com/BurntSushi/toml` only.
- `architecture/standards/go-elemental/principles/construction-and-execution.md`: every exported function and type under `marathon/`.
- `architecture/standards/go-elemental/principles/tests-and-docs.md`: the `marathon/` packages' external test packages and their `doc.go` inventories.
- `architecture/standards/go-elemental/principles/release-and-ci.md`: the Go check `scripts/check.sh` runs over `marathon/` with the tools `mise.toml` pins, and CI's one `check` job.
- `architecture/standards/go-elemental/principles/cli-applications.md`, `domain-files.md`, and `lifecycle-and-context.md` don't apply: marathon's tools don't build on go-cli-sdk.
- release-and-ci's Keep-a-Changelog form doesn't apply: `plugins/marathon/CHANGELOG.md`, in its `## vX.Y.Z` and `## Unreleased` form, is the `marathon/` module's record.
