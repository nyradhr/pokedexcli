package main

import (
	"fmt"
	"os"
)

func commandExit(config *config, args []string) error {
	fmt.Println("Closing the Pokédex... Goodbye!")
	os.Exit(0)
	return nil
}
