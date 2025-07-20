package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokemon) == 0 {
		return errors.New("no pokemon in pokedex")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
