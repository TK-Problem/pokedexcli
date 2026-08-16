package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("explore requires exactly one location area name")
	}
	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)

	resp, err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
