package main

import (
	"time"

	"github.com/pbojar/gopoke/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		client: pokeClient,
	}

	Start(cfg)
}
