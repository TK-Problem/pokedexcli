package main

import "fmt"

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return showLocationAreas(cfg, *cfg.prevLocationsURL)
}
