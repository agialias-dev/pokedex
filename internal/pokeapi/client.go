package pokeapi

import (
	"net/http"
	"time"

	"github.com/agialias-dev/pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cInterval time.Duration) Client {

	return Client{
		cache: pokecache.NewCache(cInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
