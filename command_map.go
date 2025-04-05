package main

import (
	"fmt"
)

func commandMap(config *config) error {

	areasResp, err := config.pokeapiClient.GetLocationAreas(config.next)
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
