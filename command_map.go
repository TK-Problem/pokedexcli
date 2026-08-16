package main

import "fmt"

func commandMap(cfg *config) error {
	if cfg.nextLocationsURL == nil {
		fmt.Println("you're on the last page")
		return nil
	}
	return showLocationAreas(cfg, *cfg.nextLocationsURL)
}

// showLocationAreas fetches one page of location areas, advances both
// pagination cursors, and prints the area names.
func showLocationAreas(cfg *config, pageURL string) error {
	resp, err := cfg.pokeapiClient.ListLocations(&pageURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}
