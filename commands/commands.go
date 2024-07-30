// This package contains all the commands
package commands

type cliCommand struct {
	name        string
	description string
	Callback    func() error
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
	}
}
