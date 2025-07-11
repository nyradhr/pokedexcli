package main

import (
	"errors"
	"fmt"
)

func commandMapB(config *config) error {

	if config.previous == nil {
		return errors.New("you're already on the first page")
	}

	areasResp, err := config.pokeapiClient.GetLocationAreas(config.previous, config.pokeCache)
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
