---
type: regex
target: last_message
match: not_contains
# The plan round is markdown, never inside a code fence (references/briefs.md). Matches when the
# line that contains PLAN ROUND 1 follows an odd number of fence lines: the fences before it
# pair off, and the last one opens a block the round sits in. A fence line may be indented or
# quoted, and opens with ``` or ~~~.
---
^(?:(?:(?![ \t>]*(?:```|~~~))[^\n]*\n)*[ \t>]*(?:```|~~~)[^\n]*\n(?:(?![ \t>]*(?:```|~~~))[^\n]*\n)*[ \t>]*(?:```|~~~)[^\n]*\n)*(?:(?![ \t>]*(?:```|~~~))[^\n]*\n)*[ \t>]*(?:```|~~~)[^\n]*\n(?:(?![ \t>]*(?:```|~~~))[^\n]*\n)*[^\n]*PLAN ROUND 1
