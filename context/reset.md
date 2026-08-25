# reset · marathon-optimization

- **Status:** closeout
- **Session:** plan
- **Branch:** marathon-optimization

## Disposition

- **Sharpened:** `concepts/marathon-optimization.md` — restructured into two axes. The existing
  mechanical pass (data flow, context budget) is kept as one axis; a new decision-guidance axis
  captures three candidate enhancements from an evaluation of the pathfinding-under-uncertainty
  literature against the skill: a value-of-information test in `experiment` and `plan`, assumption
  annotations on `design/` and `concepts/` notes, and an explicit risk-vs-dependency ordering
  stance naming `experiment` as the risk-first instrument. Each names its target skill file; all
  remain unscheduled until the 0.7.0 restructure sees real use.
- **Retained:** `concepts/marathon-functions.md`, untouched.

## Next-focus

Unchanged from the last closeout: `v1.data.reads` — the coordinated reads slice across
go-database, go-web-sdk, and go-web-service, per
`go-web-service/context/concepts/data-layer.md`. A `coordinate` session from the workspace
root.
