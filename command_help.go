package main

import "fmt"

func commandHelp(config *config, s string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s:\t\t%s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
