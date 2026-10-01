# on-start

1. Find the layer (`mechanics/pipeline.md`), and note whether this session's project is the
   architecture repository itself.
2. Create or adopt the layer:
   - **Standalone, with no `architecture/` directory**: create it with the starting index
     (`references/architecture-layer.md`).
   - **Standalone, with an `architecture/` directory**: adopt it unchanged.
   - **Workspace, with no configuration file or no `repo` key**: agree with the architect which
     member holds the layer, write `.claude/marathon-architecture.toml` at the coordinator, and
     continue without the layer until then.
   - **Workspace, with `repo` set**: never create that repository. It is a `context` project of
     its own, set up with `marathon init`. If it isn't checked out, report that and continue
     without the layer.
3. Give the session the layer's location. The session passes it to the standards-reviewer, which
   resolves the pointers in the repository's `STANDARDS.md` against it, and to `retro`, which
   checks that each pointer resolves.
