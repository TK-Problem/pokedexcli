package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// startRepl reads commands from stdin until the process is interrupted.
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(); err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// commandOrder fixes the display order for help. Go randomizes map iteration,
// so the registry alone can't produce stable output.
var commandOrder = []string{"help", "exit"}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

// cleanInput lowercases the text and splits it into words on any run of
// whitespace. Leading and trailing whitespace is discarded.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
