package main

import (
	"time"

	"github.com/nyradhr/pokedexcli/internal/pokeapi"
	"github.com/nyradhr/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
	}

	repl(cfg)
}
