package main

import (
	"bufio"
	"fmt"
	"github.com/vavdoshka/pokedexcli/internal/pokeapi"
	"os"
	"strings"
	"time"
)

type Config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}

var commandRegistry map[string]cliCommand

func init() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world, goes forward and shows 20 locations at once.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 location areas in the Pokemon world, goes backward and shows 20 locations at once.",
			callback:    commandMapb,
		},
	}
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	helpInstruction := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, v := range commandRegistry {
		helpInstruction += fmt.Sprintf("%s: %s\n", v.name, v.description)
	}
	fmt.Println(helpInstruction)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

func runRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	client := pokeapi.NewClient(5 * time.Second)

	c := &Config{
		pokeapiClient: client,
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		if scanner.Err() != nil {
			fmt.Println("Error: %", scanner.Err())
		}

		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			fmt.Println("Error: please enter the command")
			continue
		}

		commandName := cleanedInput[0]

		command, ok := commandRegistry[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(c)
		if err != nil {
			fmt.Println("Error: ", err)
		}

	}
}
