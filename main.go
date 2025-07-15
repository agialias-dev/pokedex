package main

import (
	"time"

	"github.com/agialias-dev/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.PokemonDetails),
	}

	startRepl(cfg)
}
