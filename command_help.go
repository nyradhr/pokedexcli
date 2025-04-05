package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Printf("\nWelcome to the Pokédex!\nUsage:\n\n")
	for _, command := range commands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
