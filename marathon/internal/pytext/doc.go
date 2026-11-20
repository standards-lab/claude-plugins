// Package pytext splits, trims, and reads text as Python does, which is how
// the marathon tools count lines and words.
//
// Python's whitespace is the runes str.isspace accepts: Unicode's White_Space
// runes and the four separators U+001C through U+001F. Its line boundaries
// are the ones str.splitlines breaks on: \n, \r, \r\n, \v, \f, U+001C through
// U+001E, U+0085, U+2028, and U+2029.
//
// The package exports:
//
//   - [SpaceClass], the body of a regular-expression class matching Python's whitespace
//   - [IsSpace], which reports whether a rune is Python's whitespace
//   - [Fields], which splits text as str.split with no arguments does
//   - [TrimSpace], which trims text as str.strip with no arguments does
//   - [Lines], which splits text as str.splitlines does
//   - [UniversalNewlines], which turns text's line endings into \n as Python's
//     universal newlines mode does
package pytext
