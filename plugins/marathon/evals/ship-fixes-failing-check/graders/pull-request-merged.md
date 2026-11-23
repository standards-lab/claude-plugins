---
type: regex
target:
  source: file
  path: .forge/pulls/summary
match: contains
flags: m
---
^state: merged$
