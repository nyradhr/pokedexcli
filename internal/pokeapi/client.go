package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/nyradhr/pokedexcli/internal/pokecache"
)

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

func GetResource[T any](cache *pokecache.Cache, client *Client, url string) (T, error) {
	var resource T
	data, ok := cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return resource, err
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return resource, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return resource, err
		}

		cache.Add(url, data)
	}

	result := resource
	err := json.Unmarshal(data, &result)
	if err != nil {
		return resource, err
	}
	return result, nil
}
