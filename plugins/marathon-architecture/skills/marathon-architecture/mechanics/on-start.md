# on-start

1. Resolve the layer (`mechanics/pipeline.md`), noting whether this session's project is the
   architecture repository itself.
2. Layer in the rules of `references/architecture-layer.md`. In a `review`, add two checks:
   settled notes that have generalized and belong in the layer, and, in the architecture
   repository, pages that restate a repository's implementation.
3. Bootstrap or adopt the layer:
   - **Standalone, no `architecture/` yet** — create it with the bootstrap index; pages arrive by
     promotion only.
   - **Standalone, it exists** — adopt it unchanged.
   - **Workspace, no config or no `repo`** — settle with the architect which member holds the
     layer, write `.claude/marathon-architecture.toml` at the coordinator, and continue without
     the layer until then.
   - **Workspace, `repo` named** — never create that repository; it is its own `context` project
     from `marathon init`. If it has no checkout, report it and continue without the layer.
