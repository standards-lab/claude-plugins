// Package references checks that the markdown references in a repository's
// comments resolve.
//
// It reads every git-tracked file that isn't markdown; a file that isn't UTF-8,
// a submodule, and a tracked file missing from the working tree are skipped. A
// reference line is one whose first non-blank characters are #, //, --, or /*,
// or one that starts with * inside a block comment. A block comment opens at a
// /* outside a #, //, or -- line comment, and closes at the next */. An opener
// or line comment marker counts at the start of a line or after a character
// that can't continue a path or glob: not a letter, digit, '.', '_', '/', '-',
// '*', '?', '[', ']', '~', ':', or a quote. A reference is a path on such a line made of letters,
// digits, '.', '_', '/', and '-' that ends in '.md' and starts with a letter,
// digit, '_', or '.'. A path that follows one of those characters, or '/', '*',
// '$', '{', '}', '<', or '>', is part of a URL, pattern, or template, and isn't
// a reference.
//
// A reference, normalized, resolves from the repository root when it has only
// one segment, and in the directory beside the repository that its first
// segment names after a leading '../'. Otherwise its first segment resolves in
// the first of these places that has it: the repository root, beside the
// repository, the coordinator's [workspace.paths] entry of that name, and
// beside the coordinator. A reference missing where its first segment resolved
// fails. When its first segment is in none of those places, as in CI, the
// reference is skipped with a note.
//
// The coordinator is read from .claude/marathon.toml at the git top-level: the
// repository itself when that configuration has a [workspace] table, or else
// the directory its [project] coordinator key names, with a leading '~/' read
// from the home directory and a relative path from the top-level. A
// [workspace.paths] location takes '~/' the same way, and a relative one is
// read from the coordinator. A key naming a directory that doesn't exist only
// changes the skip note, which names it; a key naming a directory without a
// readable configuration that has a [workspace] table fails the run.
//
// The package exports:
//
//   - [Run], which checks a repository and returns the command's exit status
package references
