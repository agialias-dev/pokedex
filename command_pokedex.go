package main

import (
	"fmt"
)

func commandPokedex(config *config, s string) error {
	if len(config.pokedex) == 0 {
		return fmt.Errorf("you have no Pokemon in your Pokedex")
	}
	fmt.Println("Your Pokedex:")
	for _, entry := range config.pokedex {
		fmt.Printf(" -%s\n", entry.Name)
	}
	return nil
}
