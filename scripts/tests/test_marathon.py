"""Tests for marathon's marathon command, run against scratch git repositories."""

import os
import subprocess
import tempfile
import unittest
from pathlib import Path

MARATHON = Path(__file__).resolve().parents[2] / "plugins/marathon/bin/marathon"
KEYS = "keys: check, currency, gate, publish, merge, ci"

CONFIGURATION = """[project]
kind = "context"
check = "printf 'check in %s' \\"$PWD\\""
gate = "printf '[%s]'"

[remote]
publish = "printf '<%s>'"
merge = "printf a && printf b"
ci = '''
printf 'ci %s,' "$#"
for arg in "$@"; do printf '<%s>' "$arg"; done
exit 3
'''
"""


class MarathonTest(unittest.TestCase):
    def setUp(self):
        tmp = tempfile.TemporaryDirectory()
        self.addCleanup(tmp.cleanup)
        self.root = Path(tmp.name).resolve()
        subprocess.run(["git", "init", "-q", str(self.root)], check=True)
        self.configure(CONFIGURATION)

    def configure(self, text):
        path = self.root / ".claude/marathon.toml"
        path.parent.mkdir(exist_ok=True)
        path.write_text(text)

    def marathon(self, *args, cwd=None):
        result = subprocess.run(
            [str(MARATHON), *args],
            cwd=cwd or self.root,
            capture_output=True,
            text=True,
            env={**os.environ, "GIT_CEILING_DIRECTORIES": str(self.root.parent)},
        )
        return result.returncode, result.stdout, result.stderr.strip()

    def assert_fails(self, args, message, cwd=None):
        status, out, err = self.marathon(*args, cwd=cwd)
        self.assertNotEqual(status, 0)
        self.assertEqual((out, err), ("", message))

    def test_chained_merge_runs_as_declared(self):
        self.assertEqual(self.marathon("merge"), (0, "ab", ""))

    def test_chained_merge_stops_at_a_failing_first_command(self):
        self.configure('[remote]\nmerge = "printf a && (exit 4) && printf b"\n')
        self.assertEqual(self.marathon("merge"), (4, "a", ""))
        self.configure('[remote]\nmerge = "(exit 5) && printf b"\n')
        self.assertEqual(self.marathon("merge"), (5, "", ""))

    def test_publish_arguments_arrive_quoted(self):
        tricky = "a b; echo $HOME && `id` 'q' \"d\""
        self.assertEqual(
            self.marathon("publish", "--body-file", tricky),
            (0, f"<--body-file><{tricky}>", ""),
        )

    def test_gate_takes_its_tag(self):
        self.assertEqual(self.marathon("gate", "v1.4.0"), (0, "[v1.4.0]", ""))

    def test_multi_line_ci_runs_as_one_script_with_positional_arguments(self):
        self.assertEqual(
            self.marathon("ci", "x y", "$z;"), (3, "ci 2,<x y><$z;>", "")
        )

    def test_signaled_command_exits_as_a_shell_reports_it(self):
        self.configure('[project]\ncheck = "kill -TERM $$"\n')
        self.assertEqual(self.marathon("check"), (143, "", ""))

    def test_command_runs_from_the_repository_root(self):
        subdir = self.root / "sub/dir"
        subdir.mkdir(parents=True)
        self.assertEqual(
            self.marathon("check", cwd=subdir), (0, f"check in {self.root}", "")
        )

    def test_no_arguments(self):
        self.assertEqual(
            self.marathon(), (2, "", f"usage: marathon <key> [args]; {KEYS}")
        )

    def test_unknown_key(self):
        self.assert_fails(["deploy"], f"marathon: unknown key 'deploy'; {KEYS}")

    def test_unset_key(self):
        self.assert_fails(
            ["currency"], "marathon: set [project] currency in .claude/marathon.toml"
        )

    def test_key_in_the_wrong_table_is_unset(self):
        self.configure('[project]\nmerge = "printf no"\n')
        self.assert_fails(
            ["merge"], "marathon: set [remote] merge in .claude/marathon.toml"
        )

    def test_missing_configuration(self):
        (self.root / ".claude/marathon.toml").unlink()
        self.assert_fails(
            ["merge"],
            f"marathon: no .claude/marathon.toml in {self.root}; write one "
            "(marathon's init does), then set [remote] merge in it",
        )

    def test_configuration_that_does_not_parse(self):
        self.configure("[remote\n")
        status, out, err = self.marathon("merge")
        self.assertEqual((status, out), (2, ""))
        self.assertRegex(err, r"marathon\.toml does not parse: .*; fix it$")

    def test_outside_a_git_repository(self):
        outside = self.root / "outside"
        outside.mkdir()
        (self.root / ".git").rename(self.root / "not-git")
        self.assert_fails(
            ["merge"],
            f"marathon: {outside} is not in a git repository; run marathon inside "
            "the repository whose .claude/marathon.toml declares the command",
            cwd=outside,
        )


if __name__ == "__main__":
    unittest.main()
