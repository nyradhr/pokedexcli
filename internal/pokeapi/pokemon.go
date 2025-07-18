package pokeapi

import (
	"github.com/nyradhr/pokedexcli/internal/pokecache"
)

func (c *Client) GetAreaEncounters(locationAreaName string, cache *pokecache.Cache) (AreaEncounters, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationAreaName
	return GetResource[AreaEncounters](cache, c, url)
}

func (c *Client) GetPokemon(pokemonName string, cache *pokecache.Cache) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	return GetResource[Pokemon](cache, c, url)
}
