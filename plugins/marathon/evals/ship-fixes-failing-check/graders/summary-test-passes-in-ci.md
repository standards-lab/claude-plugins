---
type: regex
target:
  source: file
  path: .forge/ci/summary.log
match: contains
flags: m
---
^ok tests/summary\.test\.sh$
