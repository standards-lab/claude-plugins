# Staged execution

How every working session executes the step it settled. The agent implements the step in
stages, each committed once its check passes, and the architect confirms the result at
checkpoints: behaviors the architect can run or watch run. Execution starts only after the
planning phase: the stage list below is settled and approved at SETTLE, like any other plan,
before anything changes.

## What a stage is

A stage is the smallest change set that leaves one unit of the project consistent on its own. A
stage includes its tests and its in-source comments, so nothing is left invalidated for later.
What the unit is, and what its check is, follows what the stage produces:

- A stage that produces **source** has a compilation unit as its unit, the smallest unit the
  language builds on its own: a package in Go, a crate in Rust, a module in Python. The check is
  scoped to that unit and runs whatever the repository's own tooling declares over the files the
  stage touched: the language's build, vet, and test (in Go: `go build`, `go vet`, and `go test`
  on the package path), and a conventions or lint tool the repository already wires into its own
  CI or task runner. The session runs the tool and resolves its findings; it never restates what
  the tool checks. The module as a whole may be red between stages, because the stage sequence is
  in dependency order and every broken caller is a later stage.
- A stage that produces **prose or configuration** has as its unit the smallest set of files
  that must change together to stay consistent: a command playbook and the reference it cites,
  one note, or one page of the project's documentation. The check is the repository's
  own consistency script where one exists, and a read of the touched files for coherence.

On a **code** project most stages produce source, and a documentation step's stages produce
prose under the second rule. On a **context** project every stage produces prose or
configuration. A `plan` or `review` session runs in stages of context edits under the second
rule.

## The stage list

The stage list is the SETTLE artifact: written during planning, approved by the architect before
execution begins. The session may engage the planner profile to design its first draft
(`behavior/delegation.md`). The list orders the stages by dependency, lowest first, so a stage
that changes an exported interface is followed by the stages that consume it. Each entry names
the stage's unit, the files it touches, and one line on why.

In a workspace, a step that spans member repos has one list: the stages group by repository and
order across repositories by the coordinator's `order` map, lowest layer first.

The list is the plan's artifact and is not committed anywhere. It lives in the conversation and,
on a handoff, in the reset file's Next-focus.

## Checkpoints

The stage list groups its stages under checkpoints. A checkpoint is an observable behavior the
architect can run or watch run, never "the build passed" or "the tests are green". It is the gate
during execution: the session runs every stage up to a checkpoint without stopping, then stops
and shows the architect the behavior.

- Every stage list has at least one checkpoint, the final validation.
- An intermediate checkpoint belongs wherever a behavior first becomes observable that later
  stages build on.
- A library with no runnable surface earns its checkpoint through a consumer exercise written
  for the purpose, such as an example program or a conformance run, and the report shows its
  output.
- A context project has no runnable surface either. Its checkpoint is a walkthrough: the session
  follows one concrete scenario through the changed prose and reports where each changed rule
  applies, and the architect reads the files the walkthrough cites.

The stage list is the dial for how often the architect is consulted. A list that places a
checkpoint after every stage gets a stop at every stage, and no configuration key is needed.

## Executing a stage

Every stage runs in this order:

1. **State the delegation call.** Before executing, say whether the stage goes to the executor
   profile (`behavior/delegation.md`) or stays with the session, and why. When it goes to the
   executor, name the model and the reason for it.
2. **Execute.** Implement the stage and run its check, and fix until the check passes. When the
   executor did the work, the session reads what it produced firsthand before committing.
3. **Commit.** Fire `on-commit`, then commit with the stage's decision lines in the message. No
   stop and no wait: the check is the stage's gate.
4. **Log.** Add the stage's log entry to the conversation.
5. **Move on.** Continue to the next stage. At a checkpoint, stop and report it instead.

## The stage log

Each stage's log entry is conversational, not a file: the `diff --stat`, the check result, and
the delegation call from step 1. Add prose only for a decision the plan did not spell out. Code
carries the what; the log carries the why.

## The checkpoint report

At a checkpoint, stop and report: the stages the checkpoint covers, the behavior and how to see
it (the commands to run, or the walkthrough), what the session observed, and the parts it is
least confident of. Then wait for the outcome.

## Checkpoint outcomes

- **Confirmed.** Continue to the next group of stages.
- **Adjust.** The behavior is wrong within the checkpoint's own stages. Fix it forward in a new
  commit that names the checkpoint it corrects, re-run the checks, and report the checkpoint
  again.
- **Re-plan.** The finding reaches past the checkpoint's stages. The architect enters plan mode
  and sends the findings, and the session re-enters SETTLE for stages k..N. Committed stages stay
  committed unless a finding invalidates them, in which case reverting those commits is the first
  act of the re-plan. A revised stage list from k onward is approved like the original, and
  execution resumes at k.
- **Interrupt.** The architect breaks in mid-run, without waiting for a checkpoint. The session
  finishes the stage in flight if its check can pass quickly, or abandons it and restores the
  working tree otherwise. It then reports as at a checkpoint, and the architect picks one of the
  three outcomes above.

## Validation

After the last stage commits, validate the whole step. The validation is the stage list's final
checkpoint:

- On a **code** project: the whole-module build and full test run (in Go: `./...`), then the
  run-and-verify behavior check, with the concrete commands and what to look for.
- On a **context** project: a read of the whole change for coherence and consistency, and the
  repository's consistency script where one exists.
- In an experiment's project, the final step adds the answer to its question, with the evidence.

Do not close on a failure; fix it and validate again. Once the architect confirms the validation,
`close` opens with the branch review (`commands/close.md`).

## Stay within the step

A session covers its settled step and nothing else. No opportunistic refactors, no unrelated
cleanups; note the temptation as a concept if it is worth keeping, and do not take the detour.
