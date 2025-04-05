package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nyradhr/pokedexcli/internal/pokeapi"
)

func repl(config *config) {
	user := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokédex > ")
		user.Scan()

		input := user.Text()
		cleanText := cleanInput(input)

		if len(cleanText) == 0 {
			continue
		}

		commandName := cleanText[0]
		command, exists := commands()[commandName]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokédex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 location areas from PokéAPI",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the 20 previous location areas from PokéAPI",
			callback:    commandMapB,
		},
	}
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}
