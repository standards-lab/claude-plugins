# on-start

1. Find the layer (`mechanics/pipeline.md`), and note whether this session's project is the
   architecture repository itself.
2. Apply the rules in `references/architecture-layer.md`. In a `review`, add two checks: notes
   that validated work proved and that apply beyond their repository, which belong in the layer;
   and, in the architecture repository, pages that restate a repository's implementation.
3. Create or adopt the layer:
   - **Standalone, with no `architecture/` directory**: create it with the starting index. Pages
     arrive only by promotion.
   - **Standalone, with an `architecture/` directory**: adopt it unchanged.
   - **Workspace, with no configuration file or no `repo` key**: agree with the architect which
     member holds the layer, write `.claude/marathon-architecture.toml` at the coordinator, and
     continue without the layer until then.
   - **Workspace, with `repo` set**: never create that repository. It is a `context` project of
     its own, set up with `marathon init`. If it isn't checked out, report that and continue
     without the layer.
