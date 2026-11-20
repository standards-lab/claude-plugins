---
type: llm
focus: last_message
---
The plan round is the lines from the line that contains `PLAN ROUND 1` through the first line
after it that contains `reply:`. Judge only those lines; ignore all text before and after them.
A question is one numbered item with its `rec:` and `changes:` lines, up to the next numbered
item or the **reply:** line.

FAIL only if one of these holds; otherwise PASS:

1. The round has no numbered question.
2. A question's numbered line does not name what is being decided together with the options it
   chooses between (for example "Format?" alone fails; "Publishing format: static site, single
   PDF, or Markdown?" passes).
3. A question has no `rec:` line, or its `rec:` names a choice but gives no reason for it.
4. A question refers to something it does not itself state: words like "above", "below", "the
   earlier question", "option B", or **Facts found**, a letter or label defined only outside the
   question, or a file, note, or link the reader must open to know what is being decided.
5. A question's options or recommendation are deferred, as in "details to follow".

These are NOT failures: a question that mentions a file or runbook by name while also stating
what it says; a shorthand whose meaning the same question states; a `changes:` line that says a
later round will ask something; facts repeated from **Facts found**.
