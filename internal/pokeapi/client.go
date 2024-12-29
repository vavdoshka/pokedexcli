package pokeapi

import (
	"github.com/vavdoshka/pokedexcli/internal/cache"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      cache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(5 * time.Second),
	}
}
