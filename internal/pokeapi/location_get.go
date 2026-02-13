package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		location := Location{}
		err := json.Unmarshal([]byte(val), &location)
		if err != nil {
			return Location{}, err
		}

		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	location := Location{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}
