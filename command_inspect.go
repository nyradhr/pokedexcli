package main

import (
	"fmt"
)

func commandInspect(config *config, args []string) error {
	pokemon, ok := config.pokedex[args[0]]
	if !ok {
		fmt.Println("You have not caught that Pokémon yet.")
	} else {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Println(" - ", stat.Stat.Name+":", stat.BaseStat)
		}
		fmt.Println("Types: ")
		for _, t := range pokemon.Types {
			fmt.Println(" - ", t.Type.Name)
		}
	}
	return nil
}
