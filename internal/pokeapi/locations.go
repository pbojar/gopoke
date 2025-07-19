package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (resource, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resource{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resource{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return resource{}, err
	}

	locationsResp := resource{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return resource{}, err
	}

	return locationsResp, nil
}
