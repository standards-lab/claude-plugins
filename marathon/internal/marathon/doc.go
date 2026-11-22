// Package marathon runs a command declared in a repository's
// .claude/marathon.toml.
//
// The configuration is read at the git top-level of the directory marathon
// runs in, and a key's value runs with bash from that top-level, as declared,
// so a chained value keeps its && and a multi-line value runs as one script.
// The keys are check, currency, and gate in the [project] table, and publish,
// merge, ci, and checks in the [remote] table. Arguments are quoted as Python's
// shlex.quote quotes them: a single-line value gets them appended, and a
// multi-line value gets them as its positional parameters. The configuration
// is read as UTF-8 with Python's universal newlines.
//
// marathon exits with the command's status, or 128 plus the signal's number
// when a signal ends it. It exits 2 with a message naming the fix when no key
// is given, the key is unknown or unset, the configuration is missing or
// doesn't parse, or the directory is outside a git repository, and 1 when git
// or bash can't start.
//
// The package exports:
//
//   - [Run], which runs a key's command and returns marathon's exit status
package marathon
