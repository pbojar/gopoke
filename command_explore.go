package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	locDetResp, err := cfg.client.GetLocationDetails(&name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, pkEncount := range locDetResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pkEncount.Pokemon.Name)
	}
	return nil
}
