# on-start

Fires as the session begins, before orientation.

1. Resolve the layer per `mechanics/pipeline.md`. In a workspace, note whether this session's
   project is the architecture repository itself; the review check in step 2 runs the other way
   there.

2. Layer the conventions into the session:
   - The layer holds only what has generalized past one repository. Nothing a reader could infer
     from a repository's source belongs in it (`references/architecture-layer.md`).
   - Knowledge reaches the layer by promotion and only by promotion: a concept note in the
     repository that owns it, the same note once it states itself settled, and a page once the
     knowledge has outgrown that repository.
   - A repository links the layer's principles from its README and states beside the link any
     convention of its own that narrows a principle; it never restates the page.
   - An architecture page, like every note in `context/`, states what is true now — never
     the session or repository that promoted it, or how the principle has changed
     (`references/architecture-layer.md`). That history belongs to the promoting session's own
     reset file, never the page.
   - A settled note also decays once an architecture page fully expresses it, recorded as
     **Integrated** under marathon's decay rule. The removal points at the page.
   - In a `review` session, two checks join the drift scan: settled notes that have generalized
     past this repository and belong in the layer, and, in the architecture repository itself,
     pages that restate a repository's implementation, which are defects reduced to the principle
     they state or removed.

3. Bootstrap the layer, or adopt the one that exists:
   - **Standalone, no `architecture/` directory yet** — create it with a README as its index,
     seeded from the bootstrap index in `references/architecture-layer.md`. It rides the
     session's own commits. Pages arrive by promotion afterward; the session writes none here.
   - **Standalone, the directory exists** — adopt it as it stands. A layer scaffolded before the
     extension was enabled is the artifact; nothing is rewritten, regenerated, or marked.
   - **In a workspace, `.claude/marathon-architecture.toml` missing or names no `repo`** — this
     extension's own config, never core's. Settle it with the architect (which member holds the
     layer, or whether one is still to be created) and write the file at the coordinator with the
     resolved `repo` key. Continue without the layer until it's settled.
   - **In a workspace, the file names a `repo`** — never create the repository it names. The
     architecture repository is a marathon project of its own, initialized by `marathon init` as
     a `context` project whose tree is the architecture. If the key resolves to no checkout,
     report it to the architect and continue without the layer.
