package main

import "github.com/pbojar/gopoke/internal/pokeapi"

type config struct {
	client  pokeapi.Client
	pokemon map[string]pokeapi.Pokemon
	nextURL *string
	prevURL *string
}
