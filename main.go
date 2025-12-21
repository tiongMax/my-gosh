package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("gosh> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		switch command {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "echo":
			fmt.Println(strings.Join(parts[1:], " "))
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error getting current directory: %v\n", err)
				continue
			}
			fmt.Println(dir)
		case "cd":
			var targetDir string

			// Handle no arguments: go to home directory
			if len(parts) < 2 {
				home, err := os.UserHomeDir()
				if err != nil {
					fmt.Printf("cd: %s\n", err)
					continue
				}
				targetDir = home
			} else {
				targetDir = parts[1]

				// Handle tilde expansion for ~ and ~/path or ~\path
				if targetDir == "~" {
					// Just tilde: expand to home directory
					home, err := os.UserHomeDir()
					if err != nil {
						fmt.Printf("cd: %s\n", err)
						continue
					}
					targetDir = home
				} else if strings.HasPrefix(targetDir, "~/") || strings.HasPrefix(targetDir, "~\\") {
					// Tilde with separator: replace ~ with home, keeping the rest
					home, err := os.UserHomeDir()
					if err != nil {
						fmt.Printf("cd: %s\n", err)
						continue
					}
					// Remove the ~ and join with home directory
					targetDir = home + targetDir[1:]
				}
				// Note: paths like ~foo are left unchanged (not expanded)
			}

			// Change directory
			if err := os.Chdir(targetDir); err != nil {
				fmt.Printf("cd: %s\n", err)
			}
		case "ls":
			// If no arguments, list current directory
			targets := []string{"."}
			if len(parts) >= 2 {
				targets = parts[1:]
			}

			// List each target
			for i, targetDir := range targets {
				// Print header if multiple targets
				if len(targets) > 1 {
					if i > 0 {
						fmt.Println() // Blank line between listings
					}
					fmt.Printf("%s:\n", targetDir)
				}

				files, err := os.ReadDir(targetDir)
				if err != nil {
					fmt.Printf("ls: cannot access %s: No such file or directory\n", targetDir)
					continue // Skip this target, move to next
				}

				for _, file := range files {
					if file.IsDir() {
						fmt.Printf("%s/\n", file.Name())
					} else {
						fmt.Println(file.Name())
					}
				}
			}
		case "cat":
			if len(parts) < 2 {
				fmt.Println("cat: missing operand")
				continue
			}

			// Read and display each file
			for _, targetFile := range parts[1:] {
				content, err := os.ReadFile(targetFile)
				if err != nil {
					fmt.Printf("cat: %s: %s\n", targetFile, err)
					continue // Skip this file, try next one
				}
				fmt.Print(string(content))
			}
		case "mkdir":
			if len(parts) < 2 {
				fmt.Println("mkdir: missing operand")
				continue
			}
			for _, targetDir := range parts[1:] {
				err := os.Mkdir(targetDir, 0755)
				if err != nil {
					fmt.Printf("mkdir: %s: %s\n", targetDir, err)
				}
			}
		case "touch":
			if len(parts) < 2 {
				fmt.Println("touch: missing operand")
				continue
			}
			for _, targetFile := range parts[1:] {
				// Check if file exists
				_, err := os.Stat(targetFile)
				if err == nil {
					// File exists: update timestamp to current time
					now := time.Now()
					err = os.Chtimes(targetFile, now, now)
					if err != nil {
						fmt.Printf("touch: %s: %s\n", targetFile, err)
					}
				} else if os.IsNotExist(err) {
					// File doesn't exist: create it
					file, err := os.Create(targetFile)
					if err != nil {
						fmt.Printf("touch: %s: %s\n", targetFile, err)
						continue
					}
					file.Close()
				} else {
					// Other error (permission denied, etc.)
					fmt.Printf("touch: %s: %s\n", targetFile, err)
				}
			}
		case "cp":
			if len(parts) < 3 {
				fmt.Println("cp: missing operand")
				continue
			}
			source := parts[1]
			destination := parts[2]
			sourceFile, err := os.Open(source)
			if err != nil {
				fmt.Printf("cp: %s: %s\n", source, err)
				continue
			}
			destinationFile, err := os.Create(destination)
			if err != nil {
				fmt.Printf("cp: %s: %s\n", destination, err)
				sourceFile.Close()
				continue
			}
			_, err = io.Copy(destinationFile, sourceFile)
			if err != nil {
				fmt.Printf("cp: %s: %s\n", destination, err)
			}
			sourceFile.Close()
			destinationFile.Close()
		case "mv":
			if len(parts) < 3 {
				fmt.Println("mv: missing operand")
				continue
			}
			source := parts[1]
			destination := parts[2]
			err := os.Rename(source, destination)
			if err != nil {
				fmt.Printf("mv: %s: %s\n", destination, err)
			}
		default:
			fmt.Printf("Command not found: %s\n", command)
		}
	}
}
