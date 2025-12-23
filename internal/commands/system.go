package commands

import (
	"fmt"
)

func init() {
	Register(&ExitCommand{})
	Register(&HistoryCommand{})
}

type ExitCommand struct{}

func (c *ExitCommand) Name() string { return "exit" }
func (c *ExitCommand) Execute(ctx *Context) error {
	fmt.Fprintln(ctx.Stdout, "Goodbye!")
	return ErrExit
}

type HistoryCommand struct{}

func (c *HistoryCommand) Name() string { return "history" }
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
