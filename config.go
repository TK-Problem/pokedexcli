package main

import "github.com/TK-Problem/pokedexcli/internal/pokeapi"

// config holds state shared between the REPL and its command callbacks.
type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client

	// Pagination cursors for the location-area listing. A nil nextLocationsURL
	// means there is no page after the current one; likewise prevLocationsURL
	// is nil on the first page.
	nextLocationsURL *string
	prevLocationsURL *string

	// caughtPokemon is the user's pokedex, keyed by pokemon name.
	caughtPokemon map[string]pokeapi.Pokemon
}
