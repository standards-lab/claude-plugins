---
type: regex
target: last_message
match: not_contains
flags: m
---
^(?=[^\n]*PLAN ROUND 1)(?:(?![^\n]*reply:)[^\n]*\n)*?[ \t>*_]*\d+[.)][ \t](?:[^\n]*\n(?![ \t>*_]*\d+[.)][ \t])(?![^\n]*reply:))*?[^\n]*?\b(?:above|below|option [A-Z]|[Ff]acts found|see the plan file|details to follow)\b
