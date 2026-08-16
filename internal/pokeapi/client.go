package pokeapi

import (
	"net/http"
	"time"

	"github.com/TK-Problem/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// LocationAreasURL is the first page of the location-area listing. The offset
// and limit are spelled out so this matches the "previous" URL the API returns
// for page two -- the cache is keyed by URL, so an equivalent-but-different
// spelling would cache the same page twice and make mapb miss.
const LocationAreasURL = baseURL + "/location-area?offset=0&limit=20"

// Client talks to the PokeAPI. Responses are cached by request URL, so the
// cache is shared by every request the client makes.
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
