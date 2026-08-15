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

		fmt.Printf("Your command was: %s\n", words[0])
	}
}

// cleanInput lowercases the text and splits it into words on any run of
// whitespace. Leading and trailing whitespace is discarded.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
