# Marathon extensions

Queued direction for two sessions, settled in the workspace roadmap re-plan (2026-08-24) and
tracked there as `backlog.marathon-extensions` and `backlog.marathon-roadmap` (standards-lab
`context/roadmap.toml`). The first session redesigns marathon's extension system in the core
and wires the voice standard, with a version bump; the second builds the `marathon-roadmap`
extension plugin on the redesigned contract. This note records the settled direction; each
session's plan mode settles the final wording.

## Why

Two motivating incidents. The voice directive had to be restated by hand in a session's `start`
arguments because the standard's entry-point wiring is a prose citation rather than a loaded
reference. And the workspace adopted a roadmap convention (`context/roadmap.toml` at the
coordinator) that belongs in a skill, but the current extension mechanism cannot host it: it is
scoped to external platform trackers, selected at `init`, and configured per project — none of
which fits a repo-native, self-encapsulated extension.

## The extension system (core, session `backlog.marathon-extensions`)

- **Universal event hooks**, fired in every session type:

  | Hook | Moment |
  |------|--------|
  | `on-start` | session start, orientation phase |
  | `on-execute` | scope settled, branch created, starting work |
  | `on-commit` | work verified and committing |
  | `on-reset` | the reset file being written |
  | `on-close` | session closeout |

  `on-reset` fires wherever `context/reset.md` is authored — both `reset` and `close` — so at
  close the order is `on-reset`, `on-commit`, `on-close`. A handoff session never fires
  `on-close`. `on-init` retires: an `init` session fires the universal hooks like any other,
  and adopting an extension is the explicit act of creating its artifact, in whatever session
  that happens. Session-specific moments follow a naming convention defined in the spec
  (for example `start:guide-written`) and are added only when an extension earns them.
- **Presence-based discovery**, mirroring how marathon detects a project by `context/`. An
  extension is detected at two levels: capability — its skill is installed and its description
  in the skill listing declares it a marathon extension — and adoption — its artifact exists in
  the repository (at the coordinator, in a workspace). No `init`-time selection, no
  `.claude/marathon.toml` key; `.claude/settings.json` carries only ordinary permission grants.
- **Convention points, not executable hooks.** Marathon's SKILL.md instructs: at each hook
  point, apply any installed marathon extension declaring that point. The extension's own
  SKILL.md is the authority on what happens there; marathon never names a specific extension.
- **The source-of-truth rule restated.** The current "reads the repo and projects outward,
  never back into the core" wording was written for external trackers and blocks repo-native
  extensions. The rule it protects: the repository is the source of truth. An extension may own
  and maintain its artifact inside the repository; anything it projects externally is a
  read-only mirror.
- **A contract version.** An extension declares which hook-contract version it targets, so the
  catalog can evolve without silently breaking extensions.
- **Documentation.** The plugin README gains an Extensions section as a peer of Commands;
  SKILL.md's Extension hooks section is rewritten for the generic system;
  `references/extension-hooks.md` becomes the contract spec — hook catalog, declaration format,
  detection rule, source-of-truth rule, contract version.

## Voice-standard wiring (same session)

- In marathon's SKILL.md, replace the "follows the voice standard in
  `references/writing-voice.md`" prose citation with an Always Active reference using the
  `@[filepath]` convention, so the voice standard loads with the skill instead of waiting to be
  consulted.
- Name the framing intent directly in the voice files themselves — the plugin's
  `references/writing-voice.md` and the user-scope `behavior/voice.md`: prose is written in
  natural language rooted in proper American English grammar, avoiding patterns distinctly
  identifiable as AI-generated output. The current habit lists (em-dash cadence, "not X, but
  Y", grandiose wording) gesture at this; the intent should be named.
- The user-scope edit lands in `~/claude-settings/behavior/voice.md` — the repository whose
  files symlink into `~/.claude` — as plain user config alongside the plugin session, not part
  of the plugin release.

## The marathon-roadmap plugin (session `backlog.marathon-roadmap`)

Its own plugin, named `marathon-roadmap` — the name declares the layering in the skill listing,
which is the discovery surface — so anyone can layer a different tracking system on the same
hooks without touching marathon. It codifies the convention the workspace already runs
(standards-lab `context/roadmap.toml` is the lived example):

- **Three primitives**, nested like a filesystem: a goal is a named outcome made of other goals
  and tasks, closing when its criteria hold; a task is one session's requirement — what it is,
  the repositories it touches, its proof; the backlog is the tasks assigned to no goal.
- **One TOML manifest**, `context/roadmap.toml`, at the coordinator in a workspace. Child goals
  nest directly (`[goals.v1.data]`); tasks live under a goal's reserved `tasks` table; backlog
  tasks under the root `backlog` table. Fields: `name`, `summary`, and optional `criteria`,
  `repos`, `proof`, `context` (workspace-relative file links, first segment the repository).
- **Dotted citations.** References cite the dotted slug path, omitting the structural `goals`
  and `tasks` segments: `v1.data.reads`, `backlog.docs-site`. Dots, not slashes, so a task
  reference is never mistaken for a file path; slugs never contain dots.
- **Ephemeral, unordered, proximate.** The manifest captures only what remains — finished
  tasks and closed goals are deleted, the session record keeping the disposition. `next` is the
  only sequence. Resolution matches proximity: the task in front carries detail, everything
  else stays at claim resolution, detail living in linked context files.
- **Hook behavior**: `on-start` — detect the manifest and layer its conventions in; `on-reset`
  — cite task paths in Next-focus; `on-close` — delete the finished task and advance `next`;
  nothing at `on-execute` or `on-commit`.

The skill codifies the manifest contract and the hook behavior, and stays non-prescriptive past
them — how a project shapes its goal tree is its own.
