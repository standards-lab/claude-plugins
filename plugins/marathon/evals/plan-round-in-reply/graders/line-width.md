---
type: regex
target: last_message
match: not_contains
# Every line of the plan round is at most 100 characters (references/briefs.md). Matches a line
# of 101 or more characters from the line that contains PLAN ROUND 1 through the first line
# that contains reply:, or to the end of the reply when the round has no reply: line. The u flag
# counts characters, not UTF-16 units.
flags: mu
---
^(?=[^\n]*PLAN ROUND 1)(?:(?![^\n]*reply:)[^\n]*\n)*?[^\n]{101}
