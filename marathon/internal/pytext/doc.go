// Package pytext splits and trims text as Python's str methods do, so the
// marathon tools ported from Python keep their line numbers and word counts.
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
package pytext
