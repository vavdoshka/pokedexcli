package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/vavdoshka/pokedexcli/internal/cache"
	"io"
	"net/http"
	"time"
)

type GenericHttpError struct {
	status string
}

func (e GenericHttpError) Error() string {
	return fmt.Sprintf("something bad happened: %s", e.status)
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "resource was not found"
}

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

func genericRequestWithCache[T any](c *Client, url string) (T, error) {
	var zeroValue T

	dat, ok := c.cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return zeroValue, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return zeroValue, err
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			return zeroValue, NotFoundError{}
		}

		if resp.StatusCode != http.StatusOK {
			return zeroValue, GenericHttpError{resp.Status}
		}

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return zeroValue, fmt.Errorf("can not read the response %v, err %w", resp.Body, err)
		}

		c.cache.Add(url, dat)
	}

	t := zeroValue

	if err := json.Unmarshal(dat, &t); err != nil {
		return t, fmt.Errorf("can not unmarshal response %v, err %w", dat, err)
	}

	return t, nil

}
