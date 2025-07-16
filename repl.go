package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/agialias-dev/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.PokemonDetails
}

func startRepl(config *config) {
	fmt.Println("Welcome to the Pokedex")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		for i, commandName := range input {
			command, exists := getCommands()[commandName]
			if exists {
				switch command.name {
				case "help":
					fallthrough
				case "exit":
					fallthrough
				case "map":
					fallthrough
				case "mapb":
					err := command.callback(config, "")
					if err != nil {
						fmt.Printf("Command failed: %v\n", err)
					}
					continue
				case "explore":
					if len(input) <= 1 {
						fmt.Printf("you must provide a location name to explore\n")
						continue
					}
					location := input[i+1]
					err := command.callback(config, location)
					if err != nil {
						fmt.Printf("Command failed: %v\n", err)
					}
					continue
				case "catch":
					if len(input) <= 1 {
						fmt.Printf("you must provide a Pokemon name to catch\n")
						continue
					}
					pokemon := input[i+1]
					err := command.callback(config, pokemon)
					if err != nil {
						fmt.Printf("Command failed: %v\n", err)
					}
					continue
				case "inspect":
					if len(input) <= 1 {
						fmt.Printf("you must provide a Pokemon name to inspect\n")
						continue
					}
					pokemon := input[i+1]
					err := command.callback(config, pokemon)
					if err != nil {
						fmt.Printf("Command failed: %v\n", err)
					}
					continue
				}
				continue
			}

		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	slice := strings.Fields(lower)
	return slice
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores a location, showing you the Pokemon found there. ",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to capture a Pokemon in the current location. The difficulty of catching depends on the Pokemon's experience level.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Returns the stats of a Pokemon contained in your Pokedex.",
			callback:    commandInspect,
		},
	}
}
