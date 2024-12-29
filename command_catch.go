package main

import "fmt"

func commandCatch(config *Config, args []string) error {

	if len(args) == 0 || len(args) > 1 {
		fmt.Println("catch command expects exactly one argument, just call 'catch name_of_pokemon'")
		return nil
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)


	_, err := config.pokeapiClient.GetPokemon(pokemonName)
	
	if err != nil {
		return err
	}
	
	return nil
}