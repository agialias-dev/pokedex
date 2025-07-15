package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocationDetails(location string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + location
	var err error

	dat, cached := c.cache.Get(url)
	if !cached {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationDetails{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationDetails{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationDetails{}, err
		}
		c.cache.Add(url, dat)
	}

	if cached {
		fmt.Println("Using the Cache!!")
	}

	locationResp := LocationDetails{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationDetails{}, err
	}

	return locationResp, nil
}
