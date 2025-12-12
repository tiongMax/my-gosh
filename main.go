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
		default:
			fmt.Printf("Command not found: %s\n", command)
		}
	}
}
