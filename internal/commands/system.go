package commands

import (
	"fmt"
)

func init() {
	Register(&ExitCommand{})
	Register(&HistoryCommand{})
}

// ExitCommand handles the 'exit' command to terminate the shell
type ExitCommand struct{}

// Name returns the name of the command
func (c *ExitCommand) Name() string { return "exit" }

// Execute prints a goodbye message and returns ErrExit
func (c *ExitCommand) Execute(ctx *Context) error {
	fmt.Fprintln(ctx.Stdout, "Goodbye!")
	return ErrExit
}

// HistoryCommand handles the 'history' command to show past commands
type HistoryCommand struct{}

// Name returns the name of the command
func (c *HistoryCommand) Name() string { return "history" }

// Execute prints the command history
func (c *HistoryCommand) Execute(ctx *Context) error {
	if len(ctx.Args) > 0 {
		fmt.Fprintln(ctx.Stderr, "history: too many arguments")
		return nil
	}
	if ctx.History == nil {
		return nil
	}
	for i, entry := range *ctx.History {
		fmt.Fprintf(ctx.Stdout, "%d  %s\n", i+1, entry)
	}
	return nil
}
