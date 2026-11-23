---
type: regex
target:
  source: file
  path: .git/HEAD
match: contains
flags: m
---
^ref: refs/heads/main$
