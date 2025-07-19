package pokeapi

// Structs following the PokeAPI

type resource struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []result `json:"results"`
}

type result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
