package main

import (
	"fmt"
)

func commandExplore(config *config, args []string) error {

	locationAreaName := args[0]

	fmt.Println("Exploring " + locationAreaName + "...")

	pokemonResp, err := config.pokeapiClient.GetAreaPokemon(locationAreaName, config.pokeCache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemon := range pokemonResp.Encounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}

	return nil
}
