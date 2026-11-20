"""Tests for plan-round-in-reply's no-back-reference grader, run on the stored plan rounds."""

import re
import unittest
from pathlib import Path

REPOSITORY = Path(__file__).resolve().parents[2]
GRADER = (
    REPOSITORY / "plugins/marathon/evals/plan-round-in-reply/graders/no-back-reference.md"
)
ROUNDS = Path(__file__).resolve().parent / "fixtures/plan-rounds"
POINTING_ROUND = "2026-11-18T19-29-48-415Z_r3"
UNREASONED_REC_ROUND = "2026-11-19T21-33-26-043Z_r2"
QUESTION = "1. **Format: site or PDF?**"
REC = "   - **rec:** the site, since it stays current."
SECOND_QUESTION = "2. **Audience: on-call only, or on-call and support?**"


def no_back_reference_grader():
    pattern = GRADER.read_text().split("---", 2)[2].strip()
    return re.compile(pattern, re.MULTILINE)


def round_with(*lines):
    return "\n".join(["## PLAN ROUND 1 · goal · task t", "", *lines, "", "**reply:** ok"])


class NoBackReferenceTest(unittest.TestCase):
    def setUp(self):
        self.grader = no_back_reference_grader()

    def test_flags_the_round_that_points_at_the_runbooks_above(self):
        reply = (ROUNDS / f"{POINTING_ROUND}.txt").read_text()
        self.assertIn("the three runbooks above", self.grader.search(reply).group(0))

    def test_passes_the_rounds_judged_clean(self):
        for path in sorted(ROUNDS.glob("*.txt")):
            if path.stem in (POINTING_ROUND, UNREASONED_REC_ROUND):
                continue
            with self.subTest(path.stem):
                self.assertIsNone(self.grader.search(path.read_text()))

    def test_flags_each_pointer_inside_a_question(self):
        for pointer in [
            "the files above",
            "the list below",
            "option B",
            "the **Facts found**",
            "the facts found",
            "see the plan file",
            "details to follow",
        ]:
            with self.subTest(pointer):
                first = round_with(QUESTION, f"{REC} It uses {pointer}.", SECOND_QUESTION)
                second = round_with(QUESTION, REC, SECOND_QUESTION, f"   - **rec:** {pointer}")
                self.assertIsNotNone(self.grader.search(first))
                self.assertIsNotNone(self.grader.search(second))

    def test_passes_pointers_outside_the_questions(self):
        reply = "\n".join(
            [
                "As noted below, the round follows.",
                round_with("**Facts found**", "- the files above", "", QUESTION, REC),
                "Option B and the plan file are described above.",
            ]
        )
        self.assertIsNone(self.grader.search(reply))

    def test_passes_words_that_only_contain_a_pointer(self):
        reply = round_with(QUESTION, f"{REC} Its optional index stays aboveboard and belowground.")
        self.assertIsNone(self.grader.search(reply))


if __name__ == "__main__":
    unittest.main()
