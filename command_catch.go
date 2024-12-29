package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args []string) error {

	if len(args) == 0 || len(args) > 1 {
		fmt.Println("catch command expects exactly one argument, just call 'catch name_of_pokemon'")
		return nil
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := config.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	caught := rand.Intn(2)

	if caught == 1 {
		config.pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
