package main

import (
	"errors"
	"fmt"
)

func commandMapb(config *Config, args []string) error {

	if len(args) > 0 {
		fmt.Println("mapb command does not expect any arguments, just call 'mapb'")
		return nil
	}

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

func commandMap(config *Config, args []string) error {

	if len(args) > 0 {
		fmt.Println("map command does not expect any arguments, just call 'map'")
		return nil
	}

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
