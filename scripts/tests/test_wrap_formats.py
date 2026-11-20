"""Tests for marathon's wrap-formats Stop hook, run on the plan rounds stored from eval runs."""

import json
import re
import subprocess
import unittest
from pathlib import Path

REPOSITORY = Path(__file__).resolve().parents[2]
HOOK = REPOSITORY / "plugins/marathon/bin/wrap-formats"
GRADER = REPOSITORY / "plugins/marathon/evals/plan-round-in-reply/graders/line-width.md"
ROUNDS = Path(__file__).resolve().parent / "fixtures/plan-rounds"
LINE_WIDTH_FAILURES = {
    "2026-11-18T19-29-48-415Z_r1",
    "2026-11-18T19-32-02-939Z_r1",
    "2026-11-18T19-32-02-939Z_r2",
    "2026-11-18T19-32-02-939Z_r3",
    "2026-11-19T21-27-16-897Z_r2",
    "2026-11-20T00-33-31-174Z_r1",
    "2026-11-20T00-54-39-106Z_r3",
}
LONG_WORD = "https://example.com/" + "a" * 100
PROSE = " ".join(["word"] * 15)


def line_width_grader():
    pattern = GRADER.read_text().split("---", 2)[2].strip()
    return re.compile(pattern, re.MULTILINE)


def run_hook(payload):
    result = subprocess.run(
        [str(HOOK)], input=payload, capture_output=True, text=True, check=True
    )
    return json.loads(result.stdout) if result.stdout else None


def stop(message, active=False):
    return run_hook(
        json.dumps({"stop_hook_active": active, "last_assistant_message": message})
    )


def reprint(output):
    return output["reason"].split("\n\n", 1)[1]


def round_with(*lines):
    return "\n".join(["## PLAN ROUND 1 · goal · task t", "", *lines, "", "**reply:** ok"])


class WrapFormatsTest(unittest.TestCase):
    def test_stored_rounds_print_within_the_line_width_grader(self):
        grader = line_width_grader()
        rounds = sorted(ROUNDS.glob("*.txt"))
        self.assertEqual(len(rounds), 27)
        failing = set()
        for path in rounds:
            with self.subTest(path.stem):
                reply = path.read_text()
                if grader.search(reply):
                    failing.add(path.stem)
                output = stop(reply)
                printed = reply if output is None else reprint(output)
                self.assertIsNone(grader.search(printed))
                self.assertEqual(printed.split(), reply.split())
        self.assertEqual(failing, LINE_WIDTH_FAILURES)

    def test_continuation_lines_indent_under_the_text(self):
        output = stop(
            round_with(
                f"1. **Question?** {PROSE} {PROSE}",
                f"   - **rec:** {PROSE} {PROSE}",
                f"> - {PROSE} {PROSE}",
            )
        )
        self.assertEqual(output["decision"], "block")
        lines = reprint(output).split("\n")
        self.assertEqual(
            [line[:6] for line in lines[2:8]],
            ["1. **Q", "   wor", "   - *", "     w", "> - wo", ">   wo"],
        )
        self.assertTrue(all(len(line) <= 100 for line in lines))

    def test_reply_without_a_marathon_format_passes(self):
        self.assertIsNone(stop(f"## Summary\n\n{PROSE} {PROSE}\n"))

    def test_every_architect_format_heading_is_recognized(self):
        headings = [
            "## PLAN ROUND 2 · goal · task t",
            "## INTAKE ROUND 1 · experiment.ai → v1.ai",
            "## ESCALATION · goal · task t",
            "## Task brief · goal · task",
            "## goal · task 3/5 · task",
            "## Status · 2026-10-01",
            "## Retro · goal v1.storage · 2026-09-14..09-24",
        ]
        for heading in headings:
            with self.subTest(heading):
                self.assertEqual(stop(f"{heading}\n\n{PROSE} {PROSE}\n")["decision"], "block")

    def test_wrapped_reply_passes(self):
        self.assertIsNone(stop(round_with(PROSE)))

    def test_reprint_is_not_blocked_again(self):
        self.assertIsNone(stop(round_with(f"{PROSE} {PROSE}"), active=True))

    def test_code_fences_and_tables_are_left_long(self):
        fenced = ["```", f"{PROSE} {PROSE}", "```", f"| {PROSE} | {PROSE} |"]
        self.assertIsNone(stop(round_with(*fenced)))
        output = stop(round_with(*fenced, f"{PROSE} {PROSE}"))
        self.assertEqual(reprint(output).split("\n")[2:6], fenced)

    def test_heading_inside_a_code_fence_is_not_a_format(self):
        self.assertIsNone(stop(f"```\n## PLAN ROUND 1 · g · t\n```\n{PROSE} {PROSE}\n"))

    def test_unwrappable_word_is_reported_not_blocked(self):
        output = stop(round_with(f"- see {LONG_WORD}"))
        self.assertNotIn("decision", output)
        self.assertIn("1 line(s)", output["systemMessage"])

    def test_unwrappable_word_stays_whole_in_a_reprint(self):
        output = stop(round_with(f"- see {LONG_WORD}", f"{PROSE} {PROSE}"))
        self.assertEqual(output["decision"], "block")
        self.assertIn(f"\n- see {LONG_WORD}\n", reprint(output))

    def test_unreadable_input_passes(self):
        for payload in ("not json", "[]", '{"stop_hook_active": false}'):
            with self.subTest(payload):
                self.assertIsNone(run_hook(payload))


if __name__ == "__main__":
    unittest.main()
