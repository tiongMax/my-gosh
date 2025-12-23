package commands

import (
	"errors"
	"io"
)

// ErrExit is returned when the user wants to exit the shell
var ErrExit = errors.New("exit")

// Context holds the execution context for a command
type Context struct {
	Args    []string
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	History *[]string
}

// Command is the interface that all shell commands must implement
type Command interface {
	Name() string
	Execute(ctx *Context) error
}

// Registry maps command names to their implementations
var Registry = map[string]Command{}

// Register adds a command to the registry
func Register(cmd Command) {
	Registry[cmd.Name()] = cmd
}
