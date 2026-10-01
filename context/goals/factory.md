# goal · factory

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] pipeline
2. [x] goals
3. [ ] experiments
4. [ ] evals

## Decisions

- pipeline, goals: built together as marathon 0.16.0 in one session, which ran under 0.15.
- pipeline: SHIP merges with `[remote] merge`; this repository uses
  `gh pr checks --watch && gh pr merge --merge --delete-branch`.
- goals: a goal's record lives in its `root`, one of the repositories it locks; the coordinator
  changes only through `plan` and sync, and sync deletes last.
- goals: experiments are goals under `experiment`, with spikes as tasks and the intake last.
- experiments: runs before evals, as the first `start factory` under 0.16. Spikes become sub-goals,
  and the new `intake` command takes in a finished experiment.
- evals: runs under the reinstalled 0.16. Defects its cases expose ship as 0.16.x patches, and the
  eval gate starts with the next release.
- The workspace alignment once planned as `factory.alignment` is spread over `quality.checks`,
  `quality.standards`, `quality.architecture-diet`, and `experiment.ai`.

## Pending edits

(none)
