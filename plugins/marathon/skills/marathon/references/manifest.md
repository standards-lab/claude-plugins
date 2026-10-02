# The roadmap manifest

`context/roadmap.toml` holds only what remains on the path to the target end state: goals, their
tasks, and the three state arrays (`mechanics/goals.md`). It is plain TOML, with no inline tables
and no timestamps.

## Structure

- The root holds the three arrays, `active`, `planned`, and `backlog`, and nothing else. They are
  the file's only sequences; the rest of the tree has no order.
- Root goals sit under `[goals.<slug>]`, and sub-goals nest directly under them
  (`[goals.v1.data]`). A backlog goal is an ordinary goal table.
- A goal's tasks sit under its reserved `tasks` table (`[goals.v1.data.tasks.reads]`). The goal
  record orders them.
- An experiment is a goal under the reserved root goal `experiment`, named for the topic it
  investigates (`[goals.experiment.ai]`). It is a container that holds only its `intake` task.
  Each of its spikes is a sub-goal named for the spike's repository
  (`[goals.experiment.ai.spike-harness-driver]`), and the spike's path steps are its tasks
  (`commands/experiment.md`).

An entry's full key is its identity. Renaming a slug means updating every reference to it.

## Fields

A goal has these fields:

- `name`: the outcome, as a title.
- `summary`: what it means for this outcome to hold.
- `repos` (optional): the repositories it locks while active. An active goal's `repos` cover
  every repository its tasks touch, other than the coordinator.
- `root` (optional): the repository, one of `repos`, that holds the goal record and receives the
  context the goal syncs. Every active goal has one.
- `criteria` (optional): the statements that close the goal once they all hold.
- `context` (optional): the files that carry the detail.

A task has these fields:

- `name`: the requirement, as a title.
- `summary`: what the task is, sized to one session's work.
- `repos` (optional): the repositories it touches, among its goal's `repos` and the coordinator.
- `proof` (optional): the observable result that shows it is done.
- `context` (optional): the files that carry the detail.

A spike sub-goal adds:

- `remote`: the spike repository's remote.
- `path`: the spike's local project directory.

A spike's `root` and `repos` are its own repository, and nothing else.

`root` and `repos` default to the nearest ancestor's values. A goal with no `repos` touches only
the coordinator's `context/`, so only `plan` works on it and it has no record.

`context` entries are file paths: relative to the workspace in a workspace
(`go-web-service/context/data-layer.md`), and relative to the repository otherwise.

## Citing entries

Everything outside the manifest cites a goal or task by its dotted path, leaving out the `goals`
and `tasks` segments: `v1.data.reads`, `experiment.ai.spike-harness-driver`, `docs-site`. Slugs
never contain dots.

## Lifecycle

- **Only what remains**: sync deletes a finished goal's table with everything under it. A task's
  table stays until then, and the goal record shows its progress. A stale entry is a defect, and
  the session that finds it fixes it, or records it under the goal record's pending edits when
  the manifest is in another repository.
- **Detail only at the front**: an active goal's next task carries detail. Every other entry stays
  a short statement, with its detail in the linked files.

## Starting file

`init` creates this manifest with the first goal it settles in `active`; with none settled, it
starts empty:

```toml
# Roadmap: what remains on the path to the target end state. The format is marathon's
# references/manifest.md.

active = []
planned = []
backlog = []
```

## Example

```toml
# Roadmap: what remains on the path to the target end state. The format is marathon's
# references/manifest.md.

active = ["v1.data", "experiment.cache.spike-cache-redis", "experiment.cache.spike-cache-local"]
planned = ["v1.messaging"]
backlog = ["docs-site"]

[goals.v1]
name = "The service at v1.0"
summary = "The service reaches v1.0.0 with every capability layer complete."
criteria = ["Every capability goal beneath this one is closed."]

[goals.v1.data]
name = "Data layer"
root = "service"
repos = ["service"]
summary = "A composed data model over plain SQL with a CQRS-oriented interface."
context = ["service/context/data-layer.md"]

[goals.v1.data.tasks.reads]
name = "Reads"
summary = "The reads slice: query vocabulary, pagination, and the first domain package."
proof = "Paginated, filtered, sorted reads over HTTP."

[goals.v1.data.tasks.writes]
name = "Writes"
summary = "Commands over the reads slice's first domain package."

[goals.v1.messaging]
name = "Messaging"
repos = ["messaging-lib", "service"]
summary = "A messaging library and the reactor services it enables."

[goals.experiment]
name = "Experiments"
summary = "The spikes that settle what a goal is built from."

[goals.experiment.cache]
name = "Does a read-through cache pay for itself?"
summary = "Whether the reads need a cache before v1."

[goals.experiment.cache.tasks.intake]
name = "The cache intake"
summary = "Decides what v1.data takes from the spikes' answers."

[goals.experiment.cache.spike-cache-redis]
name = "A shared Redis cache"
root = "spike-cache-redis"
repos = ["spike-cache-redis"]
remote = "https://github.com/example/spike-cache-redis.git"
path = "~/experiments/spike-cache-redis"
summary = "Whether a shared cache cuts read latency enough to pay for its operation."

[goals.experiment.cache.spike-cache-redis.tasks.baseline]
name = "Baseline"
summary = "Measures the reads without a cache."

[goals.experiment.cache.spike-cache-redis.tasks.cache]
name = "Read-through cache"
summary = "Adds the cache and measures the reads again."

[goals.experiment.cache.spike-cache-local]
name = "An in-process cache"
root = "spike-cache-local"
repos = ["spike-cache-local"]
remote = "https://github.com/example/spike-cache-local.git"
path = "~/experiments/spike-cache-local"
summary = "Whether an in-process cache is enough without shared infrastructure."

[goals.experiment.cache.spike-cache-local.tasks.cache]
name = "In-process cache"
summary = "Adds the cache and measures the reads against the same baseline."

[goals.docs-site]
name = "The docs site"
summary = "The documentation site that serves the project's published pages."
```

`v1.messaging` stays planned while `v1.data` holds `service`, and it gains its `root` when it is
staged. Both spikes of `experiment.cache` are active at once, each locking only its own
repository, while `experiment.cache` itself is never listed.
