---
type: regex
target: last_message
match: contains
flags: m
---
^[ \t>*_]*\d+[.)][ \t]+\S[^\n]*\n(?:(?![ \t>*_]*\d+[.)][ \t])[^\n]*\n)*?[ \t>*_-]*rec:[^\n]*\n(?:(?![ \t>*_]*\d+[.)][ \t])[^\n]*\n)*?[ \t>*_-]*changes:[\s\S]*?^[ \t>*_]*\d+[.)][ \t]+\S[^\n]*\n(?:(?![ \t>*_]*\d+[.)][ \t])[^\n]*\n)*?[ \t>*_-]*rec:[^\n]*\n(?:(?![ \t>*_]*\d+[.)][ \t])[^\n]*\n)*?[ \t>*_-]*changes:
