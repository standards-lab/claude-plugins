# Extensions

An extension adds a convention to marathon's sessions, such as a roadmap manifest the sessions
keep current. An extension is a skill in its own plugin. It declares what it does at the hooks
every session fires (`mechanics/hooks.md`), and marathon never refers to it otherwise.

## Installed and enabled

An extension acts only when it is both:

- **Installed**: its skill is available in the harness, at user or project scope.
- **Enabled**: the repository lists it in `.claude/marathon.toml`, so every clone has the same
  setting.

An installed extension that no repository enables does nothing. When an enabled extension's skill
is missing, tell the architect, and name the plugin to install.

```toml
[project]
extensions = ["<extension>"]   # this project

[workspace]
extensions = ["<extension>"]   # coordinator only: every project in the workspace
```

Removing the entry disables the extension. The architect decides whether its files stay.

## What an extension declares

- Its skill description contains "marathon extension".
- Its SKILL.md names the artifact it owns, if any, the hooks it acts at, and the marathon version
  it targets. When the version is incompatible, report the mismatch instead of guessing.
- It creates its artifact in the first session after it is enabled.
- It adds no key to `marathon.toml` other than its entry in the `extensions` list. Its own
  configuration lives in a file it owns, which marathon never reads.

## Source of truth

The repository is the source of truth. An extension may keep its artifact in the repository.
Anything it copies outside the repository, such as to an issue tracker or another platform, is a
read-only mirror that never feeds back into marathon's files or decisions.
