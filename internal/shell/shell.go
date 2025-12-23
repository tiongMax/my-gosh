package shell

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"my-gosh/internal/commands"
)

type Shell struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	History []string
}

func New(stdin io.Reader, stdout, stderr io.Writer) *Shell {
	return &Shell{
		Stdin:   stdin,
		Stdout:  stdout,
		Stderr:  stderr,
		History: []string{},
	}
}

func (s *Shell) Run() {
	scanner := bufio.NewScanner(s.Stdin)
	for {
		fmt.Fprint(s.Stdout, "gosh> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		s.History = append(s.History, input)

		commandName := parts[0]
		args := parts[1:]

		cmd, ok := commands.Registry[commandName]
		if !ok {
			fmt.Fprintf(s.Stdout, "Command not found: %s\n", commandName)
			continue
		}

		ctx := &commands.Context{
			Args:    args,
			Stdin:   s.Stdin,
			Stdout:  s.Stdout,
			Stderr:  s.Stderr,
			History: &s.History,
		}

		err := cmd.Execute(ctx)
		if err == commands.ErrExit {
			return
		}
		if err != nil {
			fmt.Fprintf(s.Stderr, "Error executing %s: %v\n", commandName, err)
		}
	}
}
