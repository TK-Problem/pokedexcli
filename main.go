package main

import (
	"time"

	"github.com/TK-Problem/pokedexcli/internal/pokeapi"
)

func main() {
	// Seeded so that a nil nextLocationsURL unambiguously means "last page".
	firstPage := pokeapi.LocationAreasURL

	cfg := &config{
		commands:         getCommands(),
		pokeapiClient:    pokeapi.NewClient(5*time.Second, 5*time.Minute),
		nextLocationsURL: &firstPage,
	}

	startRepl(cfg)
}
