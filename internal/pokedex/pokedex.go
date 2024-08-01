package pokedex

import (
	"net/http"
	"time"

	"github.com/mazzms/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(interval),
	}
}
