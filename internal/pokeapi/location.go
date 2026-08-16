package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// RespLocationArea is a single location area, including which Pokemon can be
// encountered there.
type RespLocationArea struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// GetLocation fetches a single location area by name.
func (c *Client) GetLocation(locationAreaName string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + locationAreaName

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespLocationArea{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespLocationArea{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return RespLocationArea{}, fmt.Errorf("location area not found: %s", locationAreaName)
	}

	if resp.StatusCode > 299 {
		return RespLocationArea{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	locationResp := RespLocationArea{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil
}
