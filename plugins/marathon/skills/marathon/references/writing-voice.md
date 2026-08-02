# Writing voice

How the agent writes prose. The standard covers all prose the agent is responsible for:
design notes and concepts, reset files, implementation guides, API documentation and in-source
comments, prose documentation, repository and organization landing pages (profiles), and the skill
files themselves.

## The principle

Write what a capable technical colleague would write. The prose is natural, clear, and
professional, and it uses only the complexity the idea requires. The reader notices the subject,
never the writing.

## The discipline

- Say it once, concretely. Name the file, the directory, the function, the note.
- Name what is actually happening in the detail at hand, never the nearest stock noun or verb. A
  generic term that would fit a dozen unrelated subjects unchanged — a "shape", a "piece",
  something the code "handles" — describes none of them; the subject has its own term (the
  signature, the step, the readiness probe), and it is worth another word to use it. Surrounding
  prose that leans on a generic term is not license to repeat it.
- State what exists in present tense, as fact; mark planned work as planned.
- Keep sentences declarative and even, in ordinary words, with technical terms where they are the
  precise ones. The subject matter supplies the interest.
- Let structure convey emphasis: a defined term at first use, the key of a list entry.
- Record objective detail where documentation is the subject; keep rationale in design notes,
  where rationale is the subject.
- A colon introduces a list; a semicolon joins the clauses of one sentence.

## The tests

Does the sentence draw attention to its own craft? Rewrite it plainer.

Does each term belong to the vocabulary of the discipline it describes? Every discipline names its
own mechanisms; rewrite the word an outsider would reach for into the one its practitioners use.

Habits that fail the tests: emphasis styling scattered through running text, em-dashes as a
recurring cadence, the "not X, but Y" frame as a repeated device, grandiose wording ("powerful",
"robust", "seamless"), stock nouns and verbs standing in where the subject has its own term, and
comparisons to alternatives the document has no need to mention.

## API documentation

API documentation keeps its idiomatic form; godoc, for example, opens with the identifier it
documents ("Load reads the configuration from…"). Within that form, the prose follows the same
voice as everything else.

## When in doubt

Plainer and terser. Cut the sentence that can go; simplify the word that can be simpler.
