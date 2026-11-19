---
type: regex
target: last_message
match: not_contains
flags: mu
---
^(?=[^\n]*PLAN ROUND 1)(?:(?![^\n]*reply:)[^\n]*\n)*?[^\n]{101}
