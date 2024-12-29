package main

import (
	"fmt"
)

func commandInspect(config *Config, args []string) error {
	if len(args) == 0 || len(args) > 1 {
		fmt.Println("inspect command expects exactly one argument, just call 'inspect name_of_pokemon'")
		return nil
	}

	pokemonName := args[0]
	pokemon, ok := config.pokedex[pokemonName]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println(pokemon)
	return nil
}
