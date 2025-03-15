package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	user := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokédex > ")
		user.Scan()
		input := user.Text()
		cleanText := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanText[0])
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
