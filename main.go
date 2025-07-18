package main

import (
	"time"

	"github.com/nyradhr/pokedexcli/internal/pokeapi"
	"github.com/nyradhr/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(15 * time.Second)
	pokedex := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
		pokedex:       pokedex,
	}
	repl(cfg)
}
