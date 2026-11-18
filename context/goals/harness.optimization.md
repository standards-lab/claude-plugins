# goal · harness.optimization

- **State:** building
- **Task:** version-floor
- **Branch:** version-floor

## Tasks

1. [ ] version-floor
2. [ ] skill-footprint
3. [ ] content-budgets
4. [ ] retro-routes
5. [ ] evidence-artifacts
6. [ ] comment-standard
7. [ ] release

## Task brief · version-floor

```
Problem       An extension names the marathon version it "targets", and the contract says to
              report any mismatch, so agents running marathon 0.18 flag marathon-architecture
              (targets 0.16) as incompatible although nothing between them broke.
Behaviors     1. The extension contract says an extension declares the minimum marathon version
                 it requires; any later marathon is compatible unless its CHANGELOG marks an
                 extension-contract change after that floor.
              2. An agent reports a mismatch only when the installed marathon is below the floor,
                 or a later release marked a contract change the extension hasn't absorbed.
              3. The contract states that a marathon release that changes the extension contract
                 updates the affected extensions and bumps their floor in the same goal.
              4. marathon's CHANGELOG convention gives contract changes a visible marker.
              5. marathon-architecture declares "Requires: marathon ≥ 0.19.0" (the release that
                 carries this contract), and its README and CHANGELOG agree.
              6. The repository's check passes.
              7. CI pins are current: create-gh-release-action v1.12.0, claude-code 2.1.327.
Test seams    the repository check (pointer resolution, plugin validate); scripts/currency.sh
              reports nothing; a reading of the contract against marathon-architecture's
              declaration
Slices        1. bump create-gh-release-action to v1.12.0 and @anthropic-ai/claude-code to
                 2.1.327 in .github/workflows/
              2. the contract and CHANGELOG convention in marathon
              3. marathon-architecture's declaration, README, and CHANGELOG entry (unreleased)
Out of scope  releasing or tagging either plugin (task `release`); other extensions' notes
Door          two-way (no release)
```

## Progress

slices 3/3 committed · standards — · spec — · editor —

## Decisions

- Extensions declare a minimum marathon version, not a target; a contract-breaking marathon
  release updates its extensions and bumps their floor in the same goal (plan).
- Nothing releases until the last task, `release`; sessions run the installed 0.18.0 until then
  (plan).
- Content budgets are a `lint-context` check the plugin ships, not prose rules (plan).
- The architecture comment standard becomes a near-ban; `comment-standard` adds architecture to
  the goal's repos through a plan once v1.ai.processing frees it (plan).
- version-floor: applied marathon-architecture 0.3.0 (targets 0.16) under marathon 0.18.0; its
  on-start only locates the layer, and the question itself is the problem this task removes.
- version-floor: bumped CI pins create-gh-release-action v1.12.0 and claude-code 2.1.327 as
  slice 1; neither needed adapting.
