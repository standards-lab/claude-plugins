# Extensions

A marathon extension layers a convention into marathon's sessions: a roadmap manifest the sessions
keep current, or a tracker that mirrors the repository's state. The extension is a skill, installed
as its own plugin. It declares what happens at the hook points every session fires; marathon applies
it at those points and otherwise never names it. This reference documents the system — what an
extension is, how a repository enables one, and what the hook points mean. The firing spec the
pipeline executes is `mechanics/hooks.md`.

## Installed and enabled

An extension acts in a session when two conditions hold:

- **Installed** — its skill resolves in the harness. Claude Code resolves skills at user scope and
  project scope alike; either satisfies the condition. Installation only makes the extension's
  declarations available.
- **Enabled** — the repository declares it in `.claude/marathon.toml`. Enablement is repository
  data: it travels with a clone, and the repository remains the record of its own process.

An installed extension that no repository enables stays inert. An enabled extension whose skill is
missing gets surfaced to the architect, naming the plugin to install, rather than silently skipped.

## Enabling in marathon.toml

The `extensions` key lists enabled extensions by skill name, and the table it sits in sets the
scope:

```toml
[project]
extensions = ["<extension>"]   # enabled for this project

[workspace]
extensions = ["<extension>"]   # coordinator only: enabled for every project in the workspace
```

A project enables an extension for itself under `[project]`. A workspace coordinator may enable one
for every member project under `[workspace]`; a member's session reads its own file and the
coordinator's. Removing the entry disables the extension. Its artifact is the architect's call:
delete it when the convention is retired, or keep it as a plain file the extension no longer
maintains.

## What an extension declares

The extension's own files carry its whole declaration:

- Its description in the skill listing contains the phrase "marathon extension", so the listing
  identifies it without loading the skill.
- Its SKILL.md names the artifact it owns, if it owns one; the hook points it acts at; and the
  marathon version it targets. A session that finds the targeted version incompatible with the
  marathon it runs surfaces the mismatch to the architect instead of guessing.

An extension that owns an artifact also bootstraps it: when the extension is enabled and the
artifact does not exist yet, the next session creates it as the extension's SKILL.md directs.

## The hook points

Five universal hooks fire in every session type, each just before the moment it names, so an
extension can shape what that moment produces. The hooks themselves, their firing table, the
resolution order, and the per-command constraints live in `mechanics/hooks.md`; session-specific
moments beyond the five follow the naming convention `<command>:<moment>`, added to that spec only
when an extension earns them.

## Source of truth

The repository is the source of truth. An extension may own and maintain its artifact inside the
repository; that is the normal shape of a repo-native extension. Anything it projects outside the
repository, onto a tracker or any other platform surface, is a read-only mirror of repository
state, and the projection never feeds back into the core's files or decisions.
