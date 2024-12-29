package main

import (
	"fmt"
)

func commandPokedex(config *Config, args []string) error {
	if len(args) > 0 {
		fmt.Println("pokedex command does not expect any arguments, just call 'pokedex'")
		return nil
	}

	if len(config.pokedex) == 0 {
		fmt.Println("no Pokemons in your deck, go catch some!")
	}

	fmt.Printf("Your pokedex:")

	for k, _ := range config.pokedex {
		fmt.Printf("\n  - %s", k)
	}

	fmt.Println()

	return nil

}
