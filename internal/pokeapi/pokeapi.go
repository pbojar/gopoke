package pokeapi

import (
	"strings"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func makeURL(endpoint string, query map[string]string) string {
	url := baseURL
	if len(endpoint) != 0 {
		url = url + "/" + endpoint
	}
	if query != nil {
		queryStrings := []string{}
		for item, val := range query {
			queryStrings = append(queryStrings, item+"="+val)
		}
		url = url + "?" + strings.Join(queryStrings, "&")
	}
	return url
}
