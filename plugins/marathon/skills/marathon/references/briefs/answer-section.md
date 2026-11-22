# Answer section

The rules every format shares are in [`../briefs.md`](../briefs.md).

A spike's answer section is the coordinator edit its sync lands in the note its experiment cites
in `context`, one section per spike, under the `## Answers · experiment.<topic>` heading
`experiment` writes there (`commands/experiment.md`). The spike's last task writes the answer
under "The answer" in the spike's `context/README.md`, and adds this section to its goal
record's pending edits. The example is a note's file content, so it stays fenced.

```markdown
### Answer · experiment.ai.spike-harness-driver

**Question:** Can one driver interface run Pi, Claude Code, and OpenCode sessions?

**Answer:** Yes; stdin/stdout JSONL carries all three.

1. One `Session` interface drives all three harnesses. Proven by a running demo:
   `clutch demo adapters`.
2. Sessions survive a restart. Proven by the validate task `resume-sessions`.
3. OpenCode's tool calls round-trip. Proven only by tests: `adapters/opencode`.

[The answer](https://github.com/example/spike-harness-driver/blob/main/context/README.md#the-answer)
· [spike-harness-driver](https://github.com/example/spike-harness-driver)
```

- **The question** is the spike's, as its sub-goal states it.
- **The answer** is one line.
- **The evidence** is numbered, and each item names where it is proven: a test, a validate task,
  or a running demo. An item proven only by tests says so.
- **The links** point to where the spike states its answer, "The answer" in its
  `context/README.md`, and to the spike's remote.

`marathon intake` later folds each answer section into the served goals' notes and removes it
(`commands/intake.md`).
