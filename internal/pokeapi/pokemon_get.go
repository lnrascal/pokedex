package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal([]byte(val), &pokemon)
		if err != nil {
			fmt.Println(err)
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	return pokemon, nil
}
