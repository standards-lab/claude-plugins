// Package version answers a marathon command's --version flag.
//
// A command built by claude-plugins' scripts/build.sh with a version carries
// it, stamped at link time into this package's version variable with
// -ldflags -X; an unstamped build reports dev. A caller asks for it by
// passing --version as the one argument; any other arguments, including
// --version beside others, go to the command as before.
//
// The package exports:
//
//   - [Report], which prints the command's name and version when its
//     arguments ask for them
package version
