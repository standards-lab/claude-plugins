# Briefs

These are the rules shared by the formats marathon writes for the architect, and by the slice brief
for an implementer; each format is its own file in [`briefs/`](./briefs/). Each one the architect
reads is read in the terminal, fits on about one screen, and leads with what needs a decision.

Every format is markdown, hard-wrapped at 100 characters in the reply, and code fences appear only
around code. As the reply streams, marathon's MessageDisplay hook, `bin/wrap-formats`, rewraps on
screen a format whose lines run over; the reply itself keeps its lines. Each format's template is
shown as a blockquote so it renders as it will in the reply; the reply drops the `> `.
