// Package wrapformats rewraps a marathon reply at 100 columns: the plugin's
// Stop hook.
//
// The hook blocks a reply that holds both a format from
// plugins/marathon/skills/marathon/references/briefs.md that the architect
// reads in the terminal and a line over 100 characters outside code fences and tables. It finds a format by
// its heading at the start of a line, bare or under heading, quote, or
// emphasis markup. The block reason hands the session the whole reply
// rewrapped at 100 columns to print again as is. The rewrap breaks only at
// spaces, keeps each list and blockquote marker, and indents continuation
// lines under the text. It never starts a continuation line on a word
// Markdown would read as a block start, such as an ordered marker (also under
// emphasis, which the numbered-questions grader reads as one) or a bullet
// marker, a heading, quote, fence, or table opener, or a rule or setext
// underline: it carries the previous word down with it, so the line keeps the
// paragraph it was in. Every other reply passes untouched. The hook blocks at
// most once in a row, since stop_hook_active marks the reprint. A line that
// one word keeps over 100 columns stays as it is; when only such lines run
// over, the hook reports them to the architect instead of blocking.
//
// A width counts code points, words split at Python's whitespace, digits and word characters are
// Unicode's, and the output is JSON as Python's json.dumps writes it, with
// every character outside printable ASCII escaped. Input that isn't UTF-8 or
// isn't a JSON object passes.
//
// The package exports:
//
//   - [Run], which reads the hook's input and writes its output
package wrapformats
