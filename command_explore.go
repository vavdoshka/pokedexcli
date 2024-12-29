package main

import "fmt"

func commandExplore(config *Config, args []string) error {

	if len(args) == 0 || len(args) > 1 {
		fmt.Println("explore command expects exactly one argument, just call 'explore name_of_location'")
		return nil
	}

	locationName := args[0]

	locationDetails, err := config.pokeapiClient.GetLocation(locationName)

	if err != nil {
		return err
	}

	if len(locationDetails.PokemonEncounters) == 0 {
		fmt.Println("did not find Pokemons")
	}

	fmt.Printf("location %s, found Pokemens:\n", locationName)
	for _, v := range locationDetails.PokemonEncounters {
		fmt.Printf("  - %s\n", v.Pokemon.Name)
	}
	return nil
}
