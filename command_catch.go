package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.client.GetPokemon(&name)
	if err != nil {
		return err
	}
	catchThreshold := float64(pokemon.BaseExperience) / 255.0
	catchAttempt := normDist(0.7, 0.2)
	caught := catchAttempt > catchThreshold

	fmt.Printf("Throwing a Pokeball at %s", name)
	for range 3 {
		time.Sleep(time.Second / 2)
		fmt.Print(".")
	}
	fmt.Print("\n")
	if caught {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokemon[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func normDist(mean float64, stdDev float64) float64 {
	return rand.NormFloat64()*stdDev + mean
}
