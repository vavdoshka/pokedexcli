package main

import "fmt"

func commandExplore(config *Config, args []string) error {

	if len(args) == 0 || len(args) > 1 {
		fmt.Println("explore command expects exactly one argument, just call 'explore name_of_location'")
		return nil
	}

	locationName := args[0]

	fmt.Println(locationName)
	return nil
}
