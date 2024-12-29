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

	maxCatchProbability := 100
	maxBaseExp := 600.0
	probability := maxCatchProbability - int((float64(pokemon.BaseExperience)*0.99)/maxBaseExp*100)

	rand := rand.Intn(100)

	if rand < probability {
		config.pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
