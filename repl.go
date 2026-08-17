package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// startRepl reads commands from stdin until the process is interrupted.
func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, ok := cfg.commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(cfg, words[1:]); err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

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
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon and add it to your Pokedex (usage: catch <name>)",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore",
			description: "Lists the Pokemon in a location area (usage: explore <area-name>)",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect",
			description: "Prints details about a Pokemon you have caught (usage: inspect <name>)",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			callback:    commandMapb,
		},
	}
}

// cleanInput lowercases the text and splits it into words on any run of
// whitespace. Leading and trailing whitespace is discarded.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
