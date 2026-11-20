// Command lint-references checks that the markdown references in the comments
// of the git repository at the working directory resolve. It takes no
// arguments, prints one line per failure or skip, and exits 1 if any reference
// fails or the repository can't be read. The references package holds the
// rules.
package main
