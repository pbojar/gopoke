package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commands := makeCommands()
	for comm, cliComm := range commands {
		fmt.Printf("%s: %s\n", comm, cliComm.description)
	}
	return nil
}
