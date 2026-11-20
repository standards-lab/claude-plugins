// Tests for plan-round-in-reply's regex graders, run by node --test. Each grader's pattern and
// flags are read from its own file and compiled as a JavaScript RegExp, as the eval compiles them,
// then run on the stored plan rounds in fixtures/plan-rounds and on rounds built here. The
// line-width and numbered-questions graders also run on each stored round as wrap-formats prints
// it, using the binary in $MARATHON_BUILD, or else in plugins/marathon/.build/, as
// scripts/build.sh makes it.
import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { readdirSync, readFileSync } from "node:fs";
import { basename, join } from "node:path";
import { test } from "node:test";

const repo = join(import.meta.dirname, "../..");
const graders = join(repo, "plugins/marathon/evals/plan-round-in-reply/graders");
const rounds = join(import.meta.dirname, "fixtures/plan-rounds");
const build = process.env.MARATHON_BUILD ?? join(repo, "plugins/marathon/.build");

const POINTING_ROUND = "2026-11-18T19-29-48-415Z_r3";
const UNREASONED_REC_ROUND = "2026-11-19T21-33-26-043Z_r2";
const QUESTION = "1. **Format: site or PDF?**";
const REC = "   - **rec:** the site, since it stays current.";
const SECOND_QUESTION = "2. **Audience: on-call only, or on-call and support?**";
const PREAMBLE = "The task has no brief yet, so planning starts.";

function regexGrader(name) {
  const [, frontMatter, ...pattern] = readFileSync(join(graders, `${name}.md`), "utf8").split("---");
  const flags = /^flags:[ \t]*(\S+)/m.exec(frontMatter)?.[1] ?? "";
  return new RegExp(pattern.join("---").trim(), flags);
}

function roundWith(...lines) {
  return ["## PLAN ROUND 1 · goal · task t", "", ...lines, "", "**reply:** ok"].join("\n");
}

function storedRounds() {
  const files = readdirSync(rounds).filter((file) => file.endsWith(".txt")).sort();
  assert.notEqual(files.length, 0, `no stored rounds in ${rounds}`);
  return files.map((file) => ({
    name: basename(file, ".txt"),
    reply: readFileSync(join(rounds, file), "utf8"),
  }));
}

function printed(reply) {
  const hook = spawnSync(join(build, "wrap-formats"), {
    input: JSON.stringify({ stop_hook_active: false, last_assistant_message: reply }),
    encoding: "utf8",
  });
  assert.equal(hook.error, undefined, "wrap-formats did not run; run scripts/build.sh first");
  assert.equal(hook.status, 0, hook.stderr);
  if (hook.stdout === "") return reply;
  const output = JSON.parse(hook.stdout);
  if (output.decision !== "block") return reply;
  const reprintAt = output.reason.indexOf("\n\n");
  assert.notEqual(reprintAt, -1, `block reason without a reprint: ${output.reason}`);
  return output.reason.slice(reprintAt + 2);
}

function words(text) {
  return text.split(/\s+/).filter(Boolean);
}

test("stored rounds print within the line-width grader", () => {
  const lineWidth = regexGrader("line-width");
  let failingUnwrapped = 0;
  for (const { name, reply } of storedRounds()) {
    if (lineWidth.test(reply)) failingUnwrapped++;
    const reprint = printed(reply);
    assert.doesNotMatch(reprint, lineWidth, `${name}: the printed reply fails line-width`);
    assert.deepEqual(words(reprint), words(reply), `${name}: the printed reply changes the words`);
  }
  assert.notEqual(failingUnwrapped, 0, "no stored round fails the line-width grader unwrapped");
});

test("stored rounds keep their numbered questions when printed", () => {
  const numberedQuestions = regexGrader("numbered-questions");
  let kept = 0;
  for (const { name, reply } of storedRounds()) {
    if (!numberedQuestions.test(reply)) continue;
    kept++;
    assert.match(printed(reply), numberedQuestions, `${name}: the printed reply loses them`);
  }
  assert.notEqual(kept, 0, "no stored round has numbered questions");
});

test("no-back-reference flags the round that points at the runbooks above", () => {
  const reply = readFileSync(join(rounds, `${POINTING_ROUND}.txt`), "utf8");
  assert.match(regexGrader("no-back-reference").exec(reply)[0], /the three runbooks above/);
});

test("no-back-reference passes the rounds judged clean", async (t) => {
  const grader = regexGrader("no-back-reference");
  for (const { name, reply } of storedRounds()) {
    if (name === POINTING_ROUND || name === UNREASONED_REC_ROUND) continue;
    await t.test(name, () => assert.doesNotMatch(reply, grader));
  }
});

test("no-back-reference flags each pointer inside a question", async (t) => {
  const grader = regexGrader("no-back-reference");
  for (const pointer of [
    "the files above",
    "the list below",
    "option B",
    "the **Facts found**",
    "the facts found",
    "see the plan file",
    "details to follow",
  ]) {
    await t.test(pointer, () => {
      const first = [PREAMBLE, roundWith(QUESTION, `${REC} It uses ${pointer}.`, SECOND_QUESTION)];
      const second = roundWith(QUESTION, REC, SECOND_QUESTION, `   - **rec:** ${pointer}`);
      assert.match(first.join("\n"), grader);
      assert.match(second, grader);
    });
  }
});

test("no-back-reference passes pointers outside the questions", () => {
  const reply = [
    "As noted below, the round follows.",
    roundWith("**Facts found**", "- the files above", "", QUESTION, REC),
    "Option B and the plan file are described above.",
  ].join("\n");
  assert.doesNotMatch(reply, regexGrader("no-back-reference"));
});

test("no-back-reference passes words that only contain a pointer", () => {
  const reply = roundWith(QUESTION, `${REC} Its optional index stays aboveboard and belowground.`);
  assert.doesNotMatch(reply, regexGrader("no-back-reference"));
});
