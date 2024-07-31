// This package contains all the commands
package commands

import (
	"github.com/mazzms/pokedex/internal/pokedex"
)

type Config struct {
	PokedexClient pokedex.Client
	offset        int
	prevCommand   string
}

type cliCommand struct {
	name        string
	description string
	Callback    func(cfg *Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show all possible commands",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the first, or next, 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back 20 locations. You should start with map",
			Callback:    commandMapb,
		},
	}
}
