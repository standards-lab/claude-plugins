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
- An eval case comes from an observed failure.
- `architecture/principles/comments.md`: scripts, tests, and every configuration file (TOML, YAML, `.gitignore`), not just source; a data file's header is its package documentation.
- `architecture/principles/rolling-currency.md`
- `architecture/principles/independent-releases.md`
- `architecture/principles/context-architecture.md`
- `architecture/principles/minimal-footprint.md`
- `plugins/marathon/skills/marathon/references/build.md` "Checks come first": the repository's scripts; a mechanical rule goes in the check, and a failure message names its fix.
