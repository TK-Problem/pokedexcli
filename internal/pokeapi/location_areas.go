package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LocationAreasResp is one page of the location-area listing. Next and Previous
// are null at the last and first page respectively, hence *string.
type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// ListLocationAreas fetches a single page of location areas.
func (c *Client) ListLocationAreas(pageURL string) (LocationAreasResp, error) {
	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreas := LocationAreasResp{}
	if err := json.Unmarshal(dat, &locationAreas); err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil
}
