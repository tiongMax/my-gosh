package commands

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func init() {
	Register(&CdCommand{})
	Register(&PwdCommand{})
	Register(&LsCommand{})
	Register(&MkdirCommand{})
	Register(&TouchCommand{})
	Register(&CpCommand{})
	Register(&MvCommand{})
	Register(&RmCommand{})
	Register(&RmdirCommand{})
}

type CdCommand struct{}

func (c *CdCommand) Name() string { return "cd" }
func (c *CdCommand) Execute(ctx *Context) error {
	var targetDir string
	parts := ctx.Args

	// Handle no arguments: go to home directory
	if len(parts) < 1 {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(ctx.Stderr, "cd: %s\n", err)
			return nil
		}
		targetDir = home
	} else {
		targetDir = parts[0]

		// Handle tilde expansion for ~ and ~/path or ~\path
		if targetDir == "~" {
			// Just tilde: expand to home directory
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Fprintf(ctx.Stderr, "cd: %s\n", err)
				return nil
			}
			targetDir = home
		} else if strings.HasPrefix(targetDir, "~/") || strings.HasPrefix(targetDir, "~\\") {
			// Tilde with separator: replace ~ with home, keeping the rest
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Fprintf(ctx.Stderr, "cd: %s\n", err)
				return nil
			}
			// Remove the ~ and join with home directory
			targetDir = home + targetDir[1:]
		}
		// Note: paths like ~foo are left unchanged (not expanded)
	}

	// Change directory
	if err := os.Chdir(targetDir); err != nil {
		fmt.Fprintf(ctx.Stderr, "cd: %s\n", err)
	}
	return nil
}

type PwdCommand struct{}

func (c *PwdCommand) Name() string { return "pwd" }
func (c *PwdCommand) Execute(ctx *Context) error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(ctx.Stderr, "Error getting current directory: %v\n", err)
		return nil
	}
	fmt.Fprintln(ctx.Stdout, dir)
	return nil
}

type LsCommand struct{}

func (c *LsCommand) Name() string { return "ls" }
func (c *LsCommand) Execute(ctx *Context) error {
	// If no arguments, list current directory
	targets := []string{"."}
	if len(ctx.Args) >= 1 {
		targets = ctx.Args
	}

	// List each target
	for i, targetDir := range targets {
		// Print header if multiple targets
		if len(targets) > 1 {
			if i > 0 {
				fmt.Fprintln(ctx.Stdout) // Blank line between listings
			}
			fmt.Fprintf(ctx.Stdout, "%s:\n", targetDir)
		}

		files, err := os.ReadDir(targetDir)
		if err != nil {
			fmt.Fprintf(ctx.Stderr, "ls: cannot access %s: No such file or directory\n", targetDir)
			continue // Skip this target, move to next
		}

		for _, file := range files {
			if file.IsDir() {
				fmt.Fprintf(ctx.Stdout, "%s/\n", file.Name())
			} else {
				fmt.Fprintln(ctx.Stdout, file.Name())
			}
		}
	}
	return nil
}

type MkdirCommand struct{}

func (c *MkdirCommand) Name() string { return "mkdir" }
func (c *MkdirCommand) Execute(ctx *Context) error {
	if len(ctx.Args) < 1 {
		fmt.Fprintln(ctx.Stderr, "mkdir: missing operand")
		return nil
	}
	for _, targetDir := range ctx.Args {
		err := os.Mkdir(targetDir, 0755)
		if err != nil {
			fmt.Fprintf(ctx.Stderr, "mkdir: %s: %s\n", targetDir, err)
		}
	}
	return nil
}

type TouchCommand struct{}

func (c *TouchCommand) Name() string { return "touch" }
func (c *TouchCommand) Execute(ctx *Context) error {
	if len(ctx.Args) < 1 {
		fmt.Fprintln(ctx.Stderr, "touch: missing operand")
		return nil
	}
	for _, targetFile := range ctx.Args {
		// Check if file exists
		_, err := os.Stat(targetFile)
		if err == nil {
			// File exists: update timestamp to current time
			now := time.Now()
			err = os.Chtimes(targetFile, now, now)
			if err != nil {
				fmt.Fprintf(ctx.Stderr, "touch: %s: %s\n", targetFile, err)
			}
		} else if os.IsNotExist(err) {
			// File doesn't exist: create it
			file, err := os.Create(targetFile)
			if err != nil {
				fmt.Fprintf(ctx.Stderr, "touch: %s: %s\n", targetFile, err)
				continue
			}
			file.Close()
		} else {
			// Other error (permission denied, etc.)
			fmt.Fprintf(ctx.Stderr, "touch: %s: %s\n", targetFile, err)
		}
	}
	return nil
}

type CpCommand struct{}

func (c *CpCommand) Name() string { return "cp" }
func (c *CpCommand) Execute(ctx *Context) error {
	if len(ctx.Args) < 2 {
		fmt.Fprintln(ctx.Stderr, "cp: missing operand")
		return nil
	}
	source := ctx.Args[0]
	destination := ctx.Args[1]
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Fprintf(ctx.Stderr, "cp: %s: %s\n", source, err)
		return nil
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		fmt.Fprintf(ctx.Stderr, "cp: %s: %s\n", destination, err)
		return nil
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Fprintf(ctx.Stderr, "cp: %s: %s\n", destination, err)
	}
	return nil
}

type MvCommand struct{}

func (c *MvCommand) Name() string { return "mv" }
func (c *MvCommand) Execute(ctx *Context) error {
	if len(ctx.Args) < 2 {
		fmt.Fprintln(ctx.Stderr, "mv: missing operand")
		return nil
	}
	source := ctx.Args[0]
	destination := ctx.Args[1]
	err := os.Rename(source, destination)
	if err != nil {
		fmt.Fprintf(ctx.Stderr, "mv: %s: %s\n", destination, err)
	}
	return nil
}

type RmCommand struct{}

func (c *RmCommand) Name() string { return "rm" }
func (c *RmCommand) Execute(ctx *Context) error {
	if len(ctx.Args) < 1 {
		fmt.Fprintln(ctx.Stderr, "rm: missing operand")
		return nil
	}

	recursive := false
	args := ctx.Args

	if args[0] == "-r" {
		recursive = true
		if len(args) < 2 {
			fmt.Fprintln(ctx.Stderr, "rm: missing operand after -r")
			return nil
		}
		args = args[1:]
	}

	for _, target := range args {
		// Common check for existence
		info, err := os.Stat(target)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(ctx.Stderr, "rm: %s: No such file or directory\n", target)
			} else {
				fmt.Fprintf(ctx.Stderr, "rm: %s: %s\n", target, err)
			}
			continue
		}

		if recursive {
			if err := os.RemoveAll(target); err != nil {
				fmt.Fprintf(ctx.Stderr, "rm: %s: %s\n", target, err)
			}
		} else {
			if info.IsDir() {
				fmt.Fprintf(ctx.Stderr, "rm: %s: is a directory\n", target)
				continue
			}
			if err := os.Remove(target); err != nil {
				fmt.Fprintf(ctx.Stderr, "rm: %s: %s\n", target, err)
			}
		}
	}
	return nil
}

type RmdirCommand struct{}

func (c *RmdirCommand) Name() string { return "rmdir" }
func (c *RmdirCommand) Execute(ctx *Context) error {
	if len(ctx.Args) < 1 {
		fmt.Fprintln(ctx.Stderr, "rmdir: missing operand")
		return nil
	}

	for _, targetDir := range ctx.Args {
		info, err := os.Stat(targetDir)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(ctx.Stderr, "rmdir: %s: No such directory\n", targetDir)
			} else {
				fmt.Fprintf(ctx.Stderr, "rmdir: %s: %s\n", targetDir, err)
			}
			continue
		}

		if info.IsDir() {
			if err := os.Remove(targetDir); err != nil {
				fmt.Fprintf(ctx.Stderr, "rmdir: %s: %s\n", targetDir, err)
			}
		} else {
			fmt.Fprintf(ctx.Stderr, "rmdir: %s: not a directory\n", targetDir)
		}
	}
	return nil
}
