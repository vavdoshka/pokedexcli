package main

import (
	"fmt"
	"errors"
)

func commandMapb(config *Config) error {

	if config.previous == nil {
		return errors.New("there is no backward url, you are on the first page")

	}

	locationAreaResponse, err := config.pokeapiClient.ListLocations(config.previous)
	if err != nil {
		return fmt.Errorf("can not map, internal error %w", err)
	}

	config.next = locationAreaResponse.Next
	config.previous = locationAreaResponse.Previous


	for _, v := range locationAreaResponse.Results {
		fmt.Println(v.Name)
	}

	return nil

}

func commandMap(config *Config) error {
	locationAreaResponse, err := config.pokeapiClient.ListLocations(config.next)
	if err != nil {
		return fmt.Errorf("can not map, internal error %w", err)
	}


	config.next = locationAreaResponse.Next
	config.previous = locationAreaResponse.Previous


	for _, v := range locationAreaResponse.Results {
		fmt.Println(v.Name)
	}

	return nil
}