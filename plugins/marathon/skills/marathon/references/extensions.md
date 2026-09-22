# Extensions

An extension layers a convention into marathon's sessions, such as a roadmap manifest the
sessions keep current. It is a skill in its own plugin that declares what it does at the hook
points every session fires (`mechanics/hooks.md`); marathon otherwise never names it.

## Installed and enabled

An extension acts only when both hold:

- **Installed** — its skill resolves in the harness, at user or project scope.
- **Enabled** — the repository lists it in `.claude/marathon.toml`, so enablement travels with a
  clone.

An installed extension nothing enables stays inert. An enabled one whose skill is missing is
reported to the architect, naming the plugin to install.

```toml
[project]
extensions = ["<extension>"]   # this project

[workspace]
extensions = ["<extension>"]   # coordinator only: every project in the workspace
```

Removing the entry disables it; whether its artifact stays is the architect's call.

## What an extension declares

- Its skill-listing description contains "marathon extension".
- Its SKILL.md names the artifact it owns, if any, the hook points it acts at, and the marathon
  version it targets. On an incompatible version, report the mismatch rather than guess.
- It bootstraps its artifact in the first session after it is enabled.
- It never adds a key to `marathon.toml` beyond the `extensions` list. Configuration of its own
  lives in a file it owns, which core never reads.

## Source of truth

The repository is the source of truth. An extension may maintain its artifact inside the
repository; anything it projects outside, onto a tracker or another platform, is a read-only
mirror that never feeds back into core's files or decisions.
