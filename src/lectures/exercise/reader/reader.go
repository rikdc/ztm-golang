//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	CommandHello  = "hello"
	CommandBye    = "bye"
	StatsLines    = "lines"
	StatsCommands = "commands"
)

func main() {
	stats := map[string]int{
		"lines":    0,
		"commands": 0,
	}

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter a command (hello, bye, Q or q): ")
		input, inputErr := r.ReadString('\n') // Alternative would be to use .Scan()
		if inputErr != nil {
			fmt.Printf("Error reading input: %v\n", inputErr)
			continue
		}

		command := strings.TrimSpace(input)

		// Exit conditions
		if strings.ToLower(command) == "q" {
			break
		}
		if inputErr == io.EOF {
			break
		}

		stats[StatsLines]++

		// Command handling
		switch command {
		case CommandHello:
			fmt.Println("Response: Hello!")
			stats[StatsCommands]++
		case CommandBye:
			fmt.Println("Response: Bye!")
			stats[StatsCommands]++
		default:
			fmt.Println("Unknown command, true again!")
		}

	}

	fmt.Printf("Number of commands: %d\n", stats["commands"])
	fmt.Printf("Number of lines: %d\n", stats["lines"])
}
