package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start(cfg *config) {
	sc := bufio.NewScanner(os.Stdin)
	const pre string = "Pokedex > "
	commands := makeCommands()
	for {
		fmt.Print(pre)
		sc.Scan()

		words := cleanInput(sc.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		cliComm, ok := commands[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			cliComm.callback(cfg)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
