package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon *string) (Pokemon, error) {
	if pokemon == nil {
		return Pokemon{}, errors.New("area cannot be nil")
	}
	url := baseURL + "/pokemon/" + *pokemon

	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		datNew, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}
		dat = datNew
		c.cache.Add(url, dat)
	}

	pokemonResp := Pokemon{}
	err := json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}
