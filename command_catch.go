package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args []string) error {
	pokemonName := args[0]
	pokemonResp, err := config.pokeapiClient.GetPokemon(pokemonName, config.pokeCache)
	if err != nil {
		return err
	}
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	base_exp := pokemonResp.BaseExperience
	if rand.Float64() < float64(100)/float64(base_exp) {
		fmt.Println(pokemonName + " was caught!")
		config.pokedex[pokemonName] = pokemonResp
		fmt.Println("Added " + pokemonName + "'s data to the Pokédex.")
	} else {
		fmt.Println(pokemonName + " escaped!")
	}
	return nil
}
