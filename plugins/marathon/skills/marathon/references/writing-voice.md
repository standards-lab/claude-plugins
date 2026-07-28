# Writing voice

How the agent writes prose. The standard covers every piece of prose the agent is responsible for, not
`context/` alone: design notes and concepts, reset files, implementation guides, godoc and `doc.go`,
in-source comments, prose documentation, project and organization profiles, and the skill files
themselves.

## The voice

Write in a plain, conventional technical-documentation voice.

- Use concrete nouns. Name the file, the directory, the function, the note. Avoid abstractions like
  "surface" or "space" where a concrete noun exists.
- State what the thing does, in present tense, as fact. Skill and infrastructure documentation records
  objective implementation detail: what happens, where, in what order. It does not speculate, and it
  does not argue for the design it describes; a reader who wants the rationale finds it in a design
  note, where rationale is the subject.
- Keep sentences declarative and even. The subject matter carries the interest.

## Habits to avoid

These habits read as machine-generated prose. Avoid them:

- Bold or italics scattered through running text for emphasis. Emphasis belongs to structure: a defined
  term at first use, the key of a list entry.
- Em-dashes used for dramatic pauses or as a repeated cadence. An em-dash is an ordinary punctuation
  mark; use it occasionally, where a parenthetical genuinely helps.
- The "not X, but Y" frame and its variants as a recurring rhetorical device.
- Grandiose and promotional wording: "powerful", "robust", "seamless", "crucial", "comprehensive".
- Comparisons to alternatives the document has no need to mention. Describe the thing itself.

## godoc

godoc keeps its idiomatic form: the comment opens with the identifier it documents ("Load reads the
configuration from…"). Within that form, the prose follows the same voice as everything else.

## When in doubt

Plainer and terser. Cut the sentence that can go; simplify the word that can be simpler.
