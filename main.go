package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

				// Handle tilde expansion for ~ and ~/path
				if strings.HasPrefix(targetDir, "~") {
					home, err := os.UserHomeDir()
					if err != nil {
						fmt.Printf("cd: %s\n", err)
						continue
					}
					// Replace ~ with home directory
					targetDir = strings.Replace(targetDir, "~", home, 1)
				}
			}

			// Change directory
			if err := os.Chdir(targetDir); err != nil {
				fmt.Printf("cd: %s\n", err)
			}
		default:
			fmt.Printf("Command not found: %s\n", command)
		}
	}
}
