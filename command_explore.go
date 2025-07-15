package main

import (
	"fmt"
)

func commandExplore(config *config, location string) error {
	locationResp, err := config.pokeapiClient.GetLocationDetails(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationResp.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
