package main

import (
	"os"

	"github.com/standards-lab/claude-plugins/marathon/internal/wrapformats"
)

func main() {
	os.Exit(wrapformats.Run(os.Stdin, os.Stdout))
}
