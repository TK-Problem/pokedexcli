package main

import (
	"fmt"
	"slices"
)

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")

	// Sorted so the listing is stable; ranging the map directly would
	// reshuffle it on every run.
	names := make([]string, 0, len(cfg.caughtPokemon))
	for name := range cfg.caughtPokemon {
		names = append(names, name)
	}
	slices.Sort(names)

	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
