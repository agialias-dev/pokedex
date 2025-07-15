package main

import (
	"fmt"
)

func commandExplore(config *config, location string) error {
	pokemonResp, err := config.pokeapiClient.GetLocationDetails(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", pokemonResp.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
