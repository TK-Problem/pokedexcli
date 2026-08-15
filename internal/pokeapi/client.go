package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

// LocationAreasURL is the first page of the location-area listing.
const LocationAreasURL = baseURL + "/location-area"

// Client talks to the PokeAPI.
type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
