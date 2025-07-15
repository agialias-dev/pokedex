package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	var err error

	dat, cached := c.cache.Get(url)
	if !cached {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaList{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaList{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaList{}, err
		}
		c.cache.Add(url, dat)
	}

	if cached {
		fmt.Println("Using the Cache!!")
	}

	locationsResp := LocationAreaList{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAreaList{}, err
	}

	return locationsResp, nil
}

func (c Client) GetLocationDetails(location string) (LocationPokemonList, error) {
	url := baseURL + "/location-area/" + location
	var err error

	dat, cached := c.cache.Get(url)
	if !cached {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationPokemonList{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationPokemonList{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationPokemonList{}, err
		}
		c.cache.Add(url, dat)
	}

	if cached {
		fmt.Println("Using the Cache!!")
	}

	pokemonResp := LocationPokemonList{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return LocationPokemonList{}, err
	}

	return pokemonResp, nil
}
