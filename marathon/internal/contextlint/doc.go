// Package contextlint checks a marathon repository's context: the files under
// context/ and roadmap/, the root README, and the markdown under docs/.
//
// Every file under context/ is of a known kind: the context README, a note
// (context/*.md), or a goal record (context/goals/*.md); context/roadmap.toml
// and context/roadmap/*.toml name the roadmap's old home. Every file under
// roadmap/ is the roadmap index, index.toml, or a goal file, <goal>.toml. A
// file or directory whose name starts with '.' is skipped.
//
// Notes, READMEs, goal-record sections, and roadmap entries stay within fixed
// word budgets, where a word is a run of characters between Python's
// whitespace, as wc -w counts it. The index and goal files parse as TOML; the
// index holds only the active, planned, and backlog arrays and at most two
// comment lines, and a goal file carries no comment, writes repos, and lists
// each of its tasks once in order, which its goal record's ## Tasks list
// mirrors. The context README's capability map lists each feature as built,
// linking its home, or planned. Every relative markdown link outside code in
// the root README, context/, and docs/ resolves once percent-decoded.
//
// Files are read as UTF-8 with Python's universal newlines, and lines break
// where Python's str.splitlines breaks them.
//
// The package exports:
//
//   - [Run], which checks a repository and returns the command's exit status
package contextlint
