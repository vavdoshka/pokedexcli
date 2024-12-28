package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	helpInstruction := "Welcome to the Pokedex!\nUsage:\n\n"
	for _,v := range commandRegistry {
		helpInstruction += fmt.Sprintf("%s: %s\n", v.name, v.description)
	}
	fmt.Println(helpInstruction)
	return nil
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}





func runRepl() {


	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		
		if scanner.Err() != nil {
			fmt.Println("Error: %", scanner.Err)
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

		err := command.callback()
		if err != nil {
			fmt.Println("Error: %v", err)
		}
		
	}
}
