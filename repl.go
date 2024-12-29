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
	pokedex       map[string]pokeapi.Pokemon
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
		"explore": {
			name:        "explore",
			description: "Displays the inforamation about a location in Pokemon world.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Allows to catch Pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect the caught Pokemon.",
			callback: commandInspect,
		},
	}
}

func commandExit(config *Config, args []string) error {
	if len(args) > 0 {
		fmt.Println("exit command does not expect any arguments, just call 'exit'")
		return nil
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, args []string) error {

	if len(args) > 1 {
		fmt.Println("call 'help' to list all commands or 'help command' to get a help on specific command.")
		return nil
	}

	if len(args) == 1 {
		v, ok := commandRegistry[args[0]]
		if !ok {
			fmt.Printf("command '%s' does not exist, call 'help' to list all commands.\n", args[0])
			return nil
		}

		fmt.Printf("%s: %s\n", v.name, v.description)
		return nil
	}

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
	callback    func(*Config, []string) error
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

func runRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	client := pokeapi.NewClient(5 * time.Second)

	c := &Config{
		pokeapiClient: client,
		pokedex:       make(map[string]pokeapi.Pokemon),
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

		args := make([]string, 0)

		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		command, ok := commandRegistry[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(c, args)
		if err != nil {
			fmt.Println("Error:", err)
		}

	}
}
