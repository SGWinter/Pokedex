package pokeapi

import (
	"net/http"
	"time"

	"github.com/SGWinter/Pokedex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	pCache     pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	interval := 5
	nCache := pokecache.NewCache(interval)
	return Client {
		httpClient: http.Client{
			Timeout: timeout,
		},
		pCache:     nCache,
	}
}
