package pokeapi

import (
	"net/http"
	"time"

	"github.com/lgatibel/pokedexcli/internal/pokecache"
)

type Config struct {
	PokeapiClient    Client
	NextLocationsURL string
	PrevLocationsURL string
	Page             Page
	Cache            pokecache.Cache
}

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, inertval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(inertval),
	}
}
