package main

import (
	"os"

	"my-gosh/internal/shell"
)

// main initializes and runs the shell with standard I/O streams
func main() {
	s := shell.New(os.Stdin, os.Stdout, os.Stderr)
	s.Run()
}
