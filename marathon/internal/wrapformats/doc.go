// Package wrapformats rewraps a marathon format at 100 columns on screen: the
// plugin's MessageDisplay hook.
//
// Claude Code runs the hook once per batch of lines while an assistant message
// streams to the screen, and shows the hook's displayContent in place of the
// batch; the transcript and the model keep the original. The hook finds a
// format from plugins/marathon/skills/marathon/references/briefs.md that the
// architect reads in the terminal by its heading at the start of a line, bare
// or under heading, quote, or emphasis markup, outside code fences and tables.
// From that heading to the end of the message it rewraps each line over 100
// characters outside code fences and tables; lines before the heading, and
// every message holding no format, display untouched, and a batch the rewrap
// doesn't change gets no output. The rewrap breaks only at spaces, keeps each
// list and blockquote marker, and indents continuation lines under the text.
// It never starts a continuation line on a word Markdown would read as a block
// start, such as an ordered marker (also under emphasis, which the
// numbered-questions grader reads as one) or a bullet marker, a heading,
// quote, fence, or table opener, or a rule or setext underline: it carries the
// previous word down with it, so the line keeps the paragraph it was in. A
// line that one word keeps over 100 columns stays as it is. The final batch
// may end mid-line, and its last line is rewrapped like the others.
//
// A batch can't see the batches before it, so the hook keeps each message's
// state, whether its format heading has passed and whether a code fence is
// open, in a file in the OS temp directory named for the session_id and
// message_id. The file exists only while either holds, and the message's
// final batch removes it; a missing or unreadable file means neither holds.
//
// A width counts code points, words split at Python's whitespace, digits and
// word characters are Unicode's, and the output is JSON as Python's
// json.dumps writes it, with every character outside printable ASCII escaped.
// Input that isn't UTF-8, isn't a JSON object, lacks a string message_id or
// delta, or gives a field the wrong type gets no output, so the batch displays
// as it is.
//
// The package exports:
//
//   - [Run], which reads one batch's hook input and writes its output
package wrapformats
