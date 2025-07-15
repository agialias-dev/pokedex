package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetPokemonDetails(pokemon string) (PokemonDetails, error) {
	url := baseURL + "/pokemon/" + pokemon
	var err error

	dat, cached := c.cache.Get(url)
	if !cached {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonDetails{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonDetails{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return PokemonDetails{}, err
		}
		c.cache.Add(url, dat)
	}

	if cached {
		fmt.Println("Using the Cache!!")
	}

	pokemonResp := PokemonDetails{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonDetails{}, err
	}

	return pokemonResp, nil
}
