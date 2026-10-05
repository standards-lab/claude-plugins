---
type: regex
target: last_message
match: contains
# Two numbered questions, each with a rec: line and then a changes: line before the next
# numbered line. The prefix classes allow indentation, list bullets, and bold markup.
flags: m
---
^[ \t>*_]*\d+[.)][ \t]+\S[^\n]*\n(?:(?![ \t>*_]*\d+[.)][ \t])[^\n]*\n)*?[ \t>*_-]*rec:[^\n]*\n(?:(?![ \t>*_]*\d+[.)][ \t])[^\n]*\n)*?[ \t>*_-]*changes:[\s\S]*?^[ \t>*_]*\d+[.)][ \t]+\S[^\n]*\n(?:(?![ \t>*_]*\d+[.)][ \t])[^\n]*\n)*?[ \t>*_-]*rec:[^\n]*\n(?:(?![ \t>*_]*\d+[.)][ \t])[^\n]*\n)*?[ \t>*_-]*changes:
