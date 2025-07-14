package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommands map[string]cliCommand

func init() {
	cliCommands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a README",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
	fmt.Println("Welcome to the Pokedex")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		for _, word := range input {
			_, exists := cliCommands[word]
			if exists {
				cliCommands[word].callback()
			}
			continue
		}

	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	slice := strings.Fields(lower)
	return slice
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n", "Usage:\n")
	for _, comms := range cliCommands {
		fmt.Printf("%s: %s\n", comms.name, comms.description)
	}
	return nil
}
