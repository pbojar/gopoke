package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func makeCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location>",
			description: "Lists Pokemon found at <location>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempts to catch <pokemon>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Lists stats of <pokemon>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists names of pokemon in pokedex",
			callback:    commandPokedex,
		},
	}
	return commands
}
