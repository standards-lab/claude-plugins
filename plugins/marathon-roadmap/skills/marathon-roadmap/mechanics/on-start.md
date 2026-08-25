# on-start

Fires as the session begins, before orientation.

1. Resolve the manifest per `mechanics/pipeline.md` and read it. Its `next` list, and the
   entries for whatever the session's focus cites, join `context/reset.md` as orientation
   input.
2. Layer the conventions into the session:
   - Cite tasks by dotted slug path — `v1.data.reads`, `backlog.docs-site` — everywhere the
     session writes: reset files, concepts, commit and pull-request descriptions.
   - Hold resolution to proximity. Only the task the session is advancing gains detail in the
     manifest; everything else stays at claim resolution.
   - Treat a stale claim as a defect: the session that finds it fixes it.
3. If the extension is enabled and the manifest does not exist, bootstrap it: create
   `context/roadmap.toml` at the resolved location with the header and empty `next` from
   `references/manifest.md`. Populate it with the developer when the session settles its scope —
   the roadmap is the developer's path, not the agent's guess.
