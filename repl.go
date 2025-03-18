package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	user := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokédex > ")
		user.Scan()
		input := user.Text()
		cleanText := cleanInput(input)
		commandName := cleanText[0]
		command, flag := commands()[commandName]
		if flag {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.Trim(text, " ")
	slice := strings.Fields(text)
	for i, word := range slice {
		slice[i] = strings.ToLower(word)
	}
	return slice
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}
