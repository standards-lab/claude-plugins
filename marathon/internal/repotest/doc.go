// Package repotest runs a marathon tool's Run function on a temporary
// repository in a test and checks its exit status and output.
//
// The package exports:
//
//   - [Command], the Run signature of each tool's package
//   - [Repo], a temporary repository directory alone in its parent, whose
//     methods write files in it or beside it, run git in it, and run a Command
//     on it and check the result
//   - [New], which creates a Repo for a test
package repotest
