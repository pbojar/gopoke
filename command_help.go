package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commands := makeCommands()
	for _, cliComm := range commands {
		fmt.Printf("%s: %s\n", cliComm.name, cliComm.description)
	}
	return nil
}
