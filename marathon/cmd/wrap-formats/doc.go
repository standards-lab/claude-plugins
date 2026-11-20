// Command wrap-formats is marathon's MessageDisplay hook: it reads one batch
// of a streaming assistant message as the hook's input on stdin and, when the
// message holds a marathon format, writes the batch rewrapped at 100 columns
// for display. It exits 0, also on input it can't read. The wrapformats
// package holds the logic.
package main
