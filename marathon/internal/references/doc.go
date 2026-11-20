// Package references checks that the markdown references in a repository's
// comments resolve.
//
// It reads every git-tracked file that isn't markdown; a file that isn't
// UTF-8, a submodule, and a tracked file missing from the working tree are
// skipped. A reference line is one whose first non-blank characters are #,
// //, --, or /*, or one that starts with * inside a block comment a /* line
// opened and no */ has closed yet. A reference is a path on such a line made
// of letters, digits, '.', '_', '/', and '-' that ends in '.md' and starts with
// a letter, digit, '_', or '.'. A path that follows one of those characters,
// or '/', '*', '$', '{', '}', '<', or '>', is part of a URL, pattern, or
// template, and isn't a reference.
//
// A reference, normalized, resolves from the repository root when its first
// segment names something there, or when it has only one segment. Otherwise
// it resolves in the sibling directory of that name beside the repository; a
// reference that starts with '../' names that directory after the '../'. A
// reference resolving in neither place fails. When no directory of that name
// exists in the repository or beside it, as in CI, the reference is skipped
// with a note.
//
// The package exports:
//
//   - [Run], which checks a repository and returns the command's exit status
package references
