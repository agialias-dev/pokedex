package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, pokemon string) error {
	pokemonResp, err := config.pokeapiClient.GetPokemonDetails(pokemon)
	if err != nil {
		return err
	}

	difficulty := float64(pokemonResp.BaseExperience) * 0.01
	attempt := float64(rand.Intn(7)) + rand.Float64()

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	if attempt > difficulty {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		_, exists := config.pokedex[pokemonResp.Name]
		if exists {
			fmt.Printf("you already have %s in your collection, you let %s go free.\n", pokemonResp.Name, pokemonResp.Name)
		} else {
			config.pokedex[pokemonResp.Name] = pokemonResp
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}

	return nil
}
