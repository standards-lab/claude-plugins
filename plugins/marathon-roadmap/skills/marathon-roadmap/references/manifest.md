# The roadmap manifest

`context/roadmap.toml` — the path to the project's or workspace's target end state, holding only
what remains. One file: at the workspace coordinator when the extension is enabled under
`[workspace]`, at the project itself when enabled under `[project]`.

## Structure

The three primitives nest like a filesystem:

- Root goals sit under `[goals.<slug>]`; child goals nest directly (`[goals.v1.data]`).
- A goal's tasks live under its reserved `tasks` table (`[goals.v1.data.tasks.reads]`).
- Backlog tasks — assigned to no goal — live under the root `backlog` table
  (`[backlog.<slug>]`).
- `next`, a root-level list of dotted paths, is the only sequence the file asserts; the tree
  itself is unordered.

The composed key is the identity. Renaming a slug changes the id and is done deliberately,
updating every reference with it.

## Fields

Goal:

- `name` — the outcome, as a title.
- `summary` — the claim: what holding this outcome means.
- `criteria` (optional) — the claims that, all holding, close the goal.
- `context` (optional) — linked context files carrying the detail.

Task:

- `name` — the requirement, as a title.
- `summary` — one session's worth of work: what it is.
- `repos` (optional) — the repositories it touches.
- `proof` (optional) — the observable result that shows it done.
- `context` (optional) — linked context files carrying the detail.

`context` entries are file paths. In a workspace they are workspace-relative, the first segment
naming the repository (`go-web-service/context/concepts/data-layer.md`); in a standalone project
they are repository-relative.

## Dotted citations

Everything outside the manifest — reset files, concepts, commit and pull-request descriptions —
cites a task or goal by its dotted slug path, omitting the structural `goals` and `tasks`
segments: `v1.data.reads`, `backlog.docs-site`. Dots, not slashes, so a citation is never
mistaken for a file path; slugs never contain dots.

## Lifecycle

The manifest is ephemeral, unordered, and proximate:

- **Ephemeral** — a finished task is deleted, the session record keeping the disposition; a
  goal whose criteria hold is deleted with it; a stale claim is a defect, fixed by the session
  that finds it.
- **Unordered** — `next` is the only sequence; nothing else in the file implies one.
- **Proximate** — the task in front carries detail; everything else stays at claim resolution,
  its depth living in the linked context files, settled by its own session's plan mode.

## Bootstrap

A manifest created at `on-start` opens with this header and an empty `next`; the session
populates it with the developer when it settles scope:

```toml
# Roadmap — the path to the target end state, holding only what remains. Maintained by
# the marathon-roadmap extension; the format is the plugin's references/manifest.md.

next = []
```

## Worked example

```toml
# Roadmap — the path to the target end state, holding only what remains. Maintained by
# the marathon-roadmap extension; the format is the plugin's references/manifest.md.

next = [
  "v1.data.reads",
  "backlog.docs-site",
]

[goals.v1]
name = "The service at v1.0"
summary = "The service reaches v1.0.0 with every capability layer complete."
criteria = [
  "Every capability goal beneath this one is closed.",
]

[goals.v1.data]
name = "Data layer"
summary = "A composed data model over plain SQL with a CQRS-oriented interface."
context = ["service/context/concepts/data-layer.md"]

[goals.v1.data.tasks.reads]
name = "Reads"
repos = ["service"]
summary = "The reads slice: query vocabulary, pagination, and the first domain package."
proof = "Paginated, filtered, sorted reads over HTTP."

[backlog.docs-site]
name = "The docs site"
summary = "The documentation site serving the landing zone's content."
context = ["context/concepts/docs-site.md"]
```
