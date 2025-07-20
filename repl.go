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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cliComm, exists := commands[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		} else {
			err := cliComm.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
