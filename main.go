package main

import (
	"os"

	"my-gosh/internal/shell"
)

func main() {
	s := shell.New(os.Stdin, os.Stdout, os.Stderr)
	s.Run()
}

