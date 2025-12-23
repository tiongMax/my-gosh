package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func init() {
	Register(&CatCommand{})
	Register(&GrepCommand{})
	Register(&EchoCommand{})
}

type CatCommand struct{}

func (c *CatCommand) Name() string { return "cat" }
func (c *CatCommand) Execute(ctx *Context) error {
	if len(ctx.Args) < 1 {
		fmt.Fprintln(ctx.Stderr, "cat: missing operand")
		return nil
	}

	// Read and display each file
	for _, targetFile := range ctx.Args {
		content, err := os.ReadFile(targetFile)
		if err != nil {
			fmt.Fprintf(ctx.Stderr, "cat: %s: %s\n", targetFile, err)
			continue // Skip this file, try next one
		}
		fmt.Fprint(ctx.Stdout, string(content))
	}
	return nil
}

type GrepCommand struct{}

func (c *GrepCommand) Name() string { return "grep" }
func (c *GrepCommand) Execute(ctx *Context) error {
	if len(ctx.Args) < 2 {
		fmt.Fprintln(ctx.Stderr, "grep: missing operand")
		return nil
	}
	pattern := ctx.Args[0]
	files := ctx.Args[1:]
	for _, file := range files {
		openFile, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(ctx.Stderr, "grep: %s: %s\n", file, err)
			continue
		}
		scanner := bufio.NewScanner(openFile)
		for scanner.Scan() {
			line := scanner.Text()
			// Inside the loop...
			if strings.Contains(line, pattern) {
				if len(files) > 1 {
					fmt.Fprintf(ctx.Stdout, "%s: %s\n", file, line) // Print filename for multiple files
				} else {
					fmt.Fprintln(ctx.Stdout, line) // Just the line for single file
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(ctx.Stderr, "grep: %s: %s\n", file, err)
		}
		openFile.Close()
	}
	return nil
}

type EchoCommand struct{}

func (c *EchoCommand) Name() string { return "echo" }
func (c *EchoCommand) Execute(ctx *Context) error {
	fmt.Fprintln(ctx.Stdout, strings.Join(ctx.Args, " "))
	return nil
}
