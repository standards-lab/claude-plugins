"""plan-round-in-reply's regex graders, read from their own files, and the stored plan rounds."""

import re
from pathlib import Path

GRADERS = (
    Path(__file__).resolve().parents[2] / "plugins/marathon/evals/plan-round-in-reply/graders"
)
ROUNDS = Path(__file__).resolve().parent / "fixtures/plan-rounds"


def regex_grader(name):
    """The named grader's pattern, multiline only when its front matter declares the m flag."""
    _, front_matter, pattern = (GRADERS / f"{name}.md").read_text().split("---", 2)
    declared = re.search(r"^flags:[ \t]*(\S+)", front_matter, re.MULTILINE)
    multiline = re.MULTILINE if declared and "m" in declared.group(1) else 0
    return re.compile(pattern.strip(), multiline)


def round_with(*lines):
    """A reply holding PLAN ROUND 1 with the given lines, ending at its reply: line."""
    return "\n".join(["## PLAN ROUND 1 · goal · task t", "", *lines, "", "**reply:** ok"])
