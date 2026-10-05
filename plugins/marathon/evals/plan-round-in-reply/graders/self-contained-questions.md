---
type: llm
focus: last_message
# Judges only the round text in the final reply; the deterministic graders check its shape.
---
Judge only the plan round in the response: the text from the line that starts `PLAN ROUND 1` to
the end of the round. Ignore everything else in the response. You are reading it as the
architect who must answer it, with nothing open but this text.

The response PASSES only when every claim below holds:

- PASS: the response contains a plan round with at least one numbered question.
- PASS: each numbered question names its subject in its own words, for example "Publishing format
  for the first edition: static site, single PDF, or Markdown in the repository?", so a reader
  knows what is being decided from the question line itself.
- PASS: each question, read with its own `rec:` and `changes:` lines, can be decided without
  reading anything else: the options it chooses between are named, and the recommendation says
  why.

The response FAILS when any claim below holds:

- FAIL: there is no plan round in the response, or it has no numbered questions.
- FAIL: a question names its subject only by reference, such as "this", "the issue above",
  "option B", "the audience question", or a short headline like "Format?" whose options appear
  nowhere in that question, its `rec:`, or its `changes:`.
- FAIL: a question depends on a label, letter, or term defined only outside it, or asks the
  reader to open a file, a note, or a link to understand what is being decided.
- FAIL: a question's options or recommendation are left for another message, such as "details
  to follow" or "see the plan file".
