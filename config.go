package main

import "github.com/pbojar/gopoke/internal/pokeapi"

type config struct {
	client  pokeapi.Client
	nextURL *string
	prevURL *string
}
