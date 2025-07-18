package main

import (
	"fmt"
)

func commandExplore(config *config, args []string) error {
	locationAreaName := args[0]
	fmt.Println("Exploring " + locationAreaName + "...")
	encounterResp, err := config.pokeapiClient.GetAreaEncounters(locationAreaName, config.pokeCache)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounterResp.Encounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}
