# The remote platform and extension hooks

## Git workflow

Work happens on a branch and lands as commits. Publishing a finished branch as a change proposal is
platform-specific and is determined by the remote platform declared at `init`.

## Declaring the remote platform

`init` records the project's remote platform and the command that proposes a change there:

- GitHub — `gh pr create`
- GitLab — `glab mr create`
- another platform — the equivalent command the project supplies
- none — local-only; closeout commits and stops

marathon stores this in the `[remote]` table (`platform`, `publish`) of `.claude/marathon.toml`; the
canonical example of the whole file is in `SKILL.md`. Plugin and skill permissions live in
`.claude/settings.json`.

At closeout the core commits, then runs the declared publish command with the change description taken
from `context/reset.md`.

## Hook points

The core calls these points, where an optional extension can act:

| Hook | Trigger | Fired by | Frequency |
|------|---------|----------|-----------|
| `on-init` | a project is set up | `init` | once per project |
| `on-session-start` | a session begins: scope approved, branch created | `plan`, `start`, `experiment`, `review`, `docs`; per project in a `coordinate` fan-out | once per session |
| `on-commit` | the core makes a commit | `init` (setup commit), `reset` (WIP commit, when made), `close` (closeout commit) | once per commit |
| `on-closeout` | a session closes | `close` | once per session |

With no extension configured, the points do nothing. The core never loads a platform skill on its own.

## Extensions

A project-management extension projects the repo's state onto a platform's tracker — for example,
copying a session's Next-focus into an issue, or linking a board item to the published change. It is
chosen at `init` and recorded in `.claude/settings.json`, with its plugin enabled and its skill and
tools allowed. The file is per-project, so each repository opts in on its own.

Information flows one way: an extension reads the repo and projects outward; it does not write back into
the core. The repository is authoritative; the platform mirrors it.

## Hooks under coordination

A `coordinate` run fans out to per-project sessions (see `commands/coordinate.md`), and each project's
hooks fire as they normally would within its own session — `on-session-start` when its part begins,
`on-commit` and `on-closeout` at its close. What is *not* yet defined is how a single coordinated change
should present across several trackers at once — one linked item spanning the projects, versus an
independent item per project. Until that is settled, the hooks stay per project; cross-tracker
coordination is left to a later step.

## Core neutrality

Nothing in the core's skill or commands names a specific platform or assumes one is present. The platform
appears only in the project's configuration, set at `init`.
