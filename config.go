package main

// config holds state shared between the REPL and its command callbacks.
type config struct {
	commands map[string]cliCommand
}
