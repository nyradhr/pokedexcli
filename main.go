package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	text = strings.Trim(text, " ")
	slice := strings.Fields(text)
	for i, word := range slice {
		slice[i] = strings.ToLower(word)
	}
	return slice
}
