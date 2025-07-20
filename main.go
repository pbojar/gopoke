package main

import (
	"time"

	"github.com/pbojar/gopoke/internal/pokeapi"
)

func main() {
	const timeOut = 5 * time.Second
	const cacheTime = 5 * time.Minute
	pokeClient := pokeapi.NewClient(timeOut, cacheTime)
	cfg := &config{
		client: pokeClient,
	}

	Start(cfg)
}
