package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (resource, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return resource{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return resource{}, err
		}
		defer resp.Body.Close()

		datNew, err := io.ReadAll(resp.Body)
		if err != nil {
			return resource{}, err
		}
		dat = datNew
		c.cache.Add(url, dat)
	}

	locationsResp := resource{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return resource{}, err
	}

	return locationsResp, nil
}

func (c *Client) GetLocationDetails(area *string) (locationDetail, error) {
	if area == nil {
		return locationDetail{}, errors.New("area cannot be nil")
	}
	url := baseURL + "/location-area/" + *area

	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return locationDetail{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return locationDetail{}, err
		}
		defer resp.Body.Close()

		datNew, err := io.ReadAll(resp.Body)
		if err != nil {
			return locationDetail{}, err
		}
		dat = datNew
		c.cache.Add(url, dat)
	}

	locationDetailResp := locationDetail{}
	err := json.Unmarshal(dat, &locationDetailResp)
	if err != nil {
		return locationDetail{}, err
	}

	return locationDetailResp, nil
}
