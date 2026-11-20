// Command wrap-formats is marathon's Stop hook: it reads the hook's input on
// stdin and blocks a reply holding a marathon format with lines over 100
// columns, handing back the reply rewrapped. It exits 0, also on input it
// can't read. The wrapformats package holds the logic.
package main
