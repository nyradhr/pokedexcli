package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nyradhr/pokedexcli/internal/pokecache"
)

func (c *Client) GetLocationAreas(pageURL *string, cache *pokecache.Cache) (LocationAreaResponse, error) {

	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}
	}
	areasResp := LocationAreaResponse{}
	err := json.Unmarshal(data, &areasResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return areasResp, nil
}
