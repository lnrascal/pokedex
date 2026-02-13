package pokeapi

import (
	"net/http"
	"time"

	"github.com/lnrascal/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
	Pokemons   map[string]Pokemon
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache:    pokecache.NewCache(cacheInterval),
		Pokemons: map[string]Pokemon{},
	}
}
