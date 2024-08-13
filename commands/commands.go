// This package contains all the commands
package commands

import (
	"github.com/mazzms/pokedex/internal/pokedex"
)

type Config struct {
	PokedexClient    pokedex.Client
	offset           int
	prevCommand      string
	printedLines     int
	PossiblePokemons map[string]pokedex.Pokemon
	CapturedPokemons map[string]pokedex.Pokemon
}

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config, ...string) error
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
		"explore": {
			name:        "explore {location_name} | {location_id}",
			description: "Explore the area. You can use the name of an specific area or the numeration provided by the 'map' commands",
			Callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Try to catch a pokemon.",
			Callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Info about a pokemon",
			Callback:    commandInspect,
		},
	}
}
