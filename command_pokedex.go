package main

import (
	"fmt"
)

func commandPokedex(config *config, args []string) error {
	fmt.Println("Your Pokédex: ")
	for _, pokemon := range config.pokedex {
		fmt.Println("-", pokemon.Name)
	}
	return nil
}
