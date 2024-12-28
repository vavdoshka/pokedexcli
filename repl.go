package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		fmt.Println("Your command was: " + cleanedInput[0])
	}
}
