package pokeapi

import (
	"github.com/nuuttisir/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
