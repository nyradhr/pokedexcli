package pokeapi

import (
	"github.com/nyradhr/pokedexcli/internal/pokecache"
)

func (c *Client) GetLocationAreas(pageURL *string, cache *pokecache.Cache) (LocationAreas, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	areasResp, err := GetResource[LocationAreas](cache, c, url)
	return areasResp, err
}
