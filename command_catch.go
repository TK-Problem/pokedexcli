package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// catchThreshold tunes difficulty. A Pokemon escapes when a roll in
// [0, BaseExperience) lands above it, so higher base experience is harder.
const catchThreshold = 40

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("catch requires exactly one pokemon name")
	}
	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// BaseExperience is nullable in the API and rand.Intn panics on a
	// non-positive argument, so treat a missing value as an automatic catch.
	if pokemon.BaseExperience > 0 && rand.Intn(pokemon.BaseExperience) > catchThreshold {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	return nil
}
