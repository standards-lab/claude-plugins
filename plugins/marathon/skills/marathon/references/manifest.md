# The roadmap manifest

The roadmap manifest holds only what remains on the path to the target end state: goals, their
tasks, and the three state arrays (`mechanics/goals.md`). It is an index, `context/roadmap.toml`,
and one goal file per goal in the roadmap directory beside it, `context/roadmap/`. Every file is
plain TOML, with no inline tables and no timestamps.

## The index

The index holds the three arrays, `active`, `planned`, and `backlog`, and a comment of at most two
lines that points to this format. It holds nothing else: no goal tables and no other comments. The
arrays are the only sequences across goals; the goal files have no order among themselves.

## Goal files

- Every goal has its own file, named by the goal's full dotted path:
  `context/roadmap/v1.data.toml`. That includes a container goal that sits in no array, and the
  reserved root goal `experiment`. A backlog goal is an ordinary goal file.
- A sub-goal's file name extends its parent's (`v1.data.toml` under `v1.toml`). The file names are
  the whole tree; no file names another goal as its parent or child.
- The goal's fields sit at the top level of its file. The file doesn't repeat the goal's path; its
  name carries it.
- Each task is a `[tasks.<slug>]` table in its goal's file.
- An experiment is a goal under the reserved root goal `experiment`, named for the topic it
  investigates (`experiment.ai.toml`). It is a container that holds only its `intake` task. Each
  of its spikes is a sub-goal named for the spike's repository
  (`experiment.ai.spike-harness-driver.toml`), and the spike's path steps are its tasks
  (`commands/experiment.md`).

A goal's path, its file name, is its identity, and a task's is its goal's path and its slug.
Renaming a goal's slug means renaming its file and its sub-goals' files, and updating every
reference to it.

## Fields

A goal file has these fields:

- `name`: the outcome, as a title.
- `summary`: what it means for this outcome to hold.
- `repos`: the repositories it locks while active, other than the coordinator. Every goal file
  writes it, and `[]` means the goal touches only the coordinator's `context/`, so only `plan`
  works on it and it has no record. An active goal's `repos` cover every repository its tasks
  touch, other than the coordinator.
- `root`: the repository, one of `repos`, that holds the goal record and receives the context the
  goal syncs. A file writes it once the goal has one, and every active goal has one. A goal file
  with no `root` means `plan` sets the goal up first.
- `order`: every task slug in the file exactly once, in run order. A file that holds tasks has it,
  and the goal record's task list mirrors it.
- `criteria` (optional): the statements that close the goal once they all hold.
- `context` (optional): the files that carry the detail.

A task has these fields:

- `name`: the requirement, as a title.
- `summary`: what the task is, sized to one session's work.
- `repos` (optional): the repositories it touches, among its goal's `repos` and the coordinator. A
  task without it touches its goal's `repos`.
- `proof` (optional): the observable result that shows it is done.
- `context` (optional): the files that carry the detail.

A spike's file adds:

- `remote`: the spike repository's remote.
- `path`: the spike's local project directory.

A spike's `root` and `repos` are its own repository, and nothing else.

No goal field takes its value from an ancestor: each goal file states its own `repos` and `root`,
and a reader learns what a goal locks from its file alone.

`context` entries are file paths: relative to the workspace in a workspace
(`go-web-service/context/data-layer.md`), and relative to the repository otherwise.

## Citing entries

Everything outside the manifest cites a goal by its dotted path, its file name without `.toml`,
and a task by its goal's path and its slug: `v1.data`, `v1.data.reads`,
`experiment.ai.spike-harness-driver`, `docs-site`. Slugs never contain dots.

## Lifecycle

- **Only what remains**: sync deletes a finished goal's file and the files of its sub-goals. It
  then deletes the file of each ancestor whose criteria now hold, or that is left empty: one with
  no tasks of its own and no file whose name begins with its name. A task's table stays until
  then, and the goal record shows its progress. A stale entry is a defect, and the session that
  finds it fixes it, or records it under the goal record's pending edits when the roadmap is in
  another repository.
- **Order as written**: a task added while its goal is active joins `order` when its table is
  written, and joins the goal record's task list at the goal's next `plan` or `start`.
- **Detail only at the front**: an active goal's next task carries detail. Every other entry stays
  a short statement, with its detail in the linked files.

## Starting files

`init` creates the index with the first goal it settles in `active`, and that goal's file; with
none settled, it creates the index alone, empty:

```toml
# Roadmap index: the three arrays. The format is marathon's references/manifest.md.

active = []
planned = []
backlog = []
```

## Example

`context/roadmap.toml`:

```toml
# Roadmap index: the three arrays. The format is marathon's references/manifest.md.

active = ["v1.data", "experiment.cache.spike-cache-redis", "experiment.cache.spike-cache-local"]
planned = ["v1.messaging"]
backlog = ["docs-site"]
```

`context/roadmap/v1.toml`:

```toml
name = "The service at v1.0"
summary = "The service reaches v1.0.0 with every capability layer complete."
repos = []
criteria = ["Every capability goal beneath this one is closed."]
```

`context/roadmap/v1.data.toml`:

```toml
name = "Data layer"
summary = "A composed data model over plain SQL with a CQRS-oriented interface."
repos = ["service"]
root = "service"
order = ["reads", "writes"]
context = ["service/context/data-layer.md"]

[tasks.reads]
name = "Reads"
summary = "The reads slice: query vocabulary, pagination, and the first domain package."
proof = "Paginated, filtered, sorted reads over HTTP."

[tasks.writes]
name = "Writes"
summary = "Commands over the reads slice's first domain package."
```

`context/roadmap/v1.messaging.toml`:

```toml
name = "Messaging"
summary = "A messaging library and the reactor services it enables."
repos = ["messaging-lib", "service"]
```

`context/roadmap/experiment.toml`:

```toml
name = "Experiments"
summary = "The spikes that settle what a goal is built from."
repos = []
```

`context/roadmap/experiment.cache.toml`:

```toml
name = "Does a read-through cache pay for itself?"
summary = "Whether the reads need a cache before v1."
repos = []
order = ["intake"]
context = ["org/context/cache.md"]

[tasks.intake]
name = "The cache intake"
summary = "Decides what v1.data takes from the spikes' answers."
```

`context/roadmap/experiment.cache.spike-cache-redis.toml`:

```toml
name = "A shared Redis cache"
summary = "Whether a shared cache cuts read latency enough to pay for its operation."
repos = ["spike-cache-redis"]
root = "spike-cache-redis"
remote = "https://github.com/example/spike-cache-redis.git"
path = "~/experiments/spike-cache-redis"
order = ["baseline", "cache"]

[tasks.baseline]
name = "Baseline"
summary = "Measures the reads without a cache."

[tasks.cache]
name = "Read-through cache"
summary = "Adds the cache and measures the reads again."
```

`context/roadmap/experiment.cache.spike-cache-local.toml`:

```toml
name = "An in-process cache"
summary = "Whether an in-process cache is enough without shared infrastructure."
repos = ["spike-cache-local"]
root = "spike-cache-local"
remote = "https://github.com/example/spike-cache-local.git"
path = "~/experiments/spike-cache-local"
order = ["cache"]

[tasks.cache]
name = "In-process cache"
summary = "Adds the cache and measures the reads against the same baseline."
```

`context/roadmap/docs-site.toml`:

```toml
name = "The docs site"
summary = "The documentation site that serves the project's published pages."
repos = []
```

`v1.messaging` stays planned while `v1.data` holds `service`, and its file gains `root` when it is
staged. Both spikes of `experiment.cache` are active at once, each locking only its own
repository, while `experiment.cache` itself is never listed. It cites `org/context/cache.md`, a
note in the coordinator, `org`, where each spike's sync lands its answer section.
