package main

import (
	"fmt"
)

func commandInspect(config *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("you must provide a Pokemon name to inspect")
	}

	pokemonData, exists := config.pokedex[pokemon]
	if !exists {
		return fmt.Errorf("%s was not found in your Pokedex", pokemon)
	} else {
		var hp int
		var attack int
		var defense int
		var spAttack int
		var spDefense int
		var speed int

		for _, stat := range pokemonData.Stats {
			switch stat.Stat.Name {
			case "hp":
				hp = stat.BaseStat
				continue
			case "attack":
				attack = stat.BaseStat
				continue
			case "defense":
				defense = stat.BaseStat
				continue
			case "special-attack":
				spAttack = stat.BaseStat
				continue
			case "special-defense":
				spDefense = stat.BaseStat
				continue
			case "speed":
				speed = stat.BaseStat
				continue
			}
		}
		fmt.Printf("Name: %s\n", pokemonData.Name)
		fmt.Printf("Height: %dm\n", pokemonData.Height)
		fmt.Printf("Weight: %dkg\n", pokemonData.Weight)
		fmt.Println("Stats:")
		fmt.Printf(" -hp: %d\n", hp)
		fmt.Printf(" -attack: %d\n", attack)
		fmt.Printf(" -defense: %d\n", defense)
		fmt.Printf(" -special-attack: %d\n", spAttack)
		fmt.Printf(" -special-defense: %d\n", spDefense)
		fmt.Printf(" -speed: %d\n", speed)
		fmt.Println("Types:")
		for _, pokemonType := range pokemonData.Types {
			fmt.Printf(" -%s\n", pokemonType.Type.Name)
		}
		return nil
	}
}
