package pokeapi

import (
	"github.com/nyradhr/pokedexcli/internal/pokecache"
)

func (c *Client) GetAreaPokemon(locationAreaName string, cache *pokecache.Cache) (AreaPokemonResponse, error) {

	url := "https://pokeapi.co/api/v2/location-area/" + locationAreaName

	pokemonResp, err := GetResource[AreaPokemonResponse](cache, c, url)

	return pokemonResp, err
}
