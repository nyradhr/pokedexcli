package main

import (
	"fmt"
)

func commandMap(config *config, args []string) error {
	areasResp, err := config.pokeapiClient.GetLocationAreas(config.next, config.pokeCache)
	if err != nil {
		return err
	}
	config.next = areasResp.Next
	config.previous = areasResp.Previous
	for _, result := range areasResp.Results {
		fmt.Println(result.Name)
	}
	return nil
}
