# goal · harness.optimization

- **State:** building
- **Task:** skill-footprint
- **Branch:** skill-footprint

## Tasks

1. [x] version-floor
2. [ ] skill-footprint
3. [ ] roadmap-layout
4. [ ] content-budgets
5. [ ] retro-routes
6. [ ] evidence-artifacts
7. [ ] comment-standard
8. [ ] release

## Task brief · skill-footprint

**Problem**

marathon's SKILL.md `@`-includes three files that Claude Code never expands. Agents therefore read
about 4,300 words by hand, including release and resume rules they rarely need. Rules repeat across
the skill and its agent profiles, and no check bounds the size of either. The formats the architect
reads are code-fenced templates, which lose their formatting in the terminal and on GitHub.

**Behaviors**

1. No plugin's markdown contains an `@`-include line. marathon's SKILL.md tells the session to read
   its command's playbook first. Each playbook opens with links to the files it needs before acting.
2. Releasing and RESUME each live in their own file. A session reads the Releasing file only when a
   task brief has a Release line, and the RESUME file only when LOCATE routes to RESUME. Stage
   names and the moments hooks fire are unchanged.
3. Each rule that repeated across the skill's files has one home, and every other file cites it.
   This covers rounds never going through AskUserQuestion, currency before round 1, nothing
   changing before approval, the branch check before a commit, and the order of on-record.
4. Each agent profile holds only its delegate's own behavior. It cites the skill's sections, the
   delegate limits included, by `${CLAUDE_PLUGIN_ROOT}` paths that Claude Code fills in when it
   loads the agent. The skill keeps one line per profile and cites the profile for what that
   delegate does.
5. The slice brief has a written format: behaviors, test seam, branches, check command, and the
   standards sections that bear on the slice, quoted with their source and always including the
   comment rule. The delegation rule says the implementer never loads `STANDARDS.md` whole.
6. The implementer writes only the comments its quoted standard requires, and never one that
   excuses a workaround. The standards-reviewer sweeps for comments that excuse a workaround and
   edits doc-comment wording. The editor never edits source files.
7. These formats are markdown with code fences only around code: plan rounds, intake rounds,
   escalations, task briefs, session briefs, the status digest, and retros.
   - Rounds and escalations use a heading, bold labels over lists, and numbered questions with
     bold rec and changes lines.
   - Task briefs use bold labels over lists.
   - The status digest is a table, and retros are checkbox lists under one heading per layer.
   - Every format is hard-wrapped at 100 characters in the reply, under one width rule in the
     brief-formats reference.
   - A goal record stores its task brief in the same markdown shape.
8. The repository check fails, naming the count and the limit, when:
   - a command's read-first set exceeds its word budget;
   - an agent profile exceeds its word budget;
   - a skill description exceeds 1,024 characters;
   - a backtick file citation, or one built on `${CLAUDE_PLUGIN_ROOT}`, doesn't resolve.

   Each word budget is the size at the end of this task, rounded up to the next 100 words.
9. The plan-round eval fails a round printed inside a code fence, and a round with a line over 100
   characters. Its self-contained grader finds the round by the line that contains the round
   header.
10. Both plugins' CHANGELOGs record the changes under Unreleased with no extension-contract marker.
    Their READMEs link only to files that exist.
11. The repository check passes, and the currency command reports nothing.

**Test seams**

- The repository check, including its failure on a planted violation of each new rule
- The plan-round eval's graders, run on a passing round and on a fenced, over-wide round
- One live delegate engagement showing that a `${CLAUDE_PLUGIN_ROOT}` citation in a profile resolves

**Slices**

1. Explicit loading in marathon and marathon-architecture, with the check rejecting `@`-include
   lines
2. The check resolves backtick and `${CLAUDE_PLUGIN_ROOT}` file citations
3. Releasing and RESUME move into their own files, and the citations to them are repointed
4. Repeated rules get one home within the skill
5. Profiles cite the skill by `${CLAUDE_PLUGIN_ROOT}` paths and keep only their own behavior. The
   slice brief gets its format, with quoted standards. The comment rules land in the implementer,
   standards-reviewer, and editor profiles.
6. The architect-facing formats and the goal-record brief become markdown under the 100-character
   width rule. The plan-round eval gains its code-fence and line-width graders.
7. Word budgets for the read-first sets and the profiles, and the description cap. marathon's
   description is trimmed to fit.

**Out of scope**

- `lint-context` and budgets on produced content (task content-budgets)
- The roadmap layout (task roadmap-layout)
- The architecture layer's comment standard (task comment-standard)
- Evidence artifacts in the session brief (task evidence-artifacts)
- Rewriting briefs already stored fenced in goal records
- Releasing or tagging either plugin (task release)

**Door**

Two-way. Nothing is released, and sessions keep running the installed 0.18.0.

## Progress

slices 1/7 committed · standards — · spec — · editor —

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
- version-floor: a contract change is a CHANGELOG entry starting `**Extension contract:**`;
  marathon's CHANGELOG preamble states the convention and `references/extensions.md`,
  Compatibility, cites it.
- version-floor: unreleased changes sit under `## Unreleased` in each CHANGELOG; `release` renames
  them to marathon v0.19.0 and marathon-architecture's next version and releases both together,
  which satisfies the 0.19.0 floor.
- version-floor: marathon-architecture's CHANGELOG cites marathon's Compatibility section instead
  of restating the rule (standards review).
- roadmap-layout: the roadmap stays in repository files, one `context/roadmap/<goal>.toml` per
  goal, with `context/roadmap.toml` holding only the three arrays; GitHub Issues were rejected for
  losing PR-reviewed planning, offline sessions, and checks (plan).
- roadmap-layout: a goal's task order is an `order` array in its goal file, which the record's
  task list mirrors; every goal file states its own `root` and `repos` (plan).
- roadmap-layout: `status` prints the cross-goal view on demand; a GitHub Project mirror is
  deferred to `v1.harness.tooling` (plan).
- roadmap-layout: the coordinator's roadmap is split during this goal's final sync; the task
  precedes content-budgets so its manifest rules are written once (plan).
- skill-footprint: `@`-includes become links with an explicit read instruction; `@` expands only in
  CLAUDE.md (plan).
- skill-footprint: Each playbook opens with its own read-first list (plan).
- skill-footprint: Word budgets are the task-end size, rounded up to the next 100 (plan).
- skill-footprint: Skill descriptions are capped at 1,024 characters (plan).
- skill-footprint: Releasing and RESUME move to two on-demand files; the "tag <names>" resume moves
  with Releasing (plan).
- skill-footprint: The check resolves backtick file citations, and rejects `@`-include lines in
  every plugin (plan).
- skill-footprint: No Extension contract marker; marathon-architecture keeps its 0.19.0 floor
  (plan).
- skill-footprint: Profiles cite the skill through `${CLAUDE_PLUGIN_ROOT}` paths, not a `skills:`
  preload. They only cite shared rules, with delegation Limits as the home for delegate limits. A
  profile is the home of its own delegate's behavior (plan).
- skill-footprint: Profiles have word budgets (plan).
- skill-footprint: The session quotes relevant standards, the comment rule always included, into the
  slice brief; the implementer never loads `STANDARDS.md` whole. The slice brief gets a written
  format (plan).
- skill-footprint: The standards-reviewer owns doc-comment wording and the sweep for comments that
  excuse; the editor leaves source files (plan).
- skill-footprint: Architect-facing formats are markdown, fenced only around code: rounds as heading
  plus lists; the task brief as bold labels over lists; the session brief as sections; status as a
  table; retro as checkboxes. Goal records store the brief unfenced. All of it is hard-wrapped at
  100 characters (plan).
- skill-footprint: The plan-round eval gains a code-fence grader and a line-width grader (plan).
- skill-footprint: The format work joins skill-footprint; the roadmap redesign is its own task,
  roadmap-layout (plan).
