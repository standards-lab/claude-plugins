package main

import (
	"os"

	"github.com/standards-lab/claude-plugins/marathon/internal/marathon"
)

func main() {
	os.Exit(marathon.Run(".", os.Args[1:], os.Stdout, os.Stderr))
}
