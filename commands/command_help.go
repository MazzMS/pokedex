package commands

import (
	"fmt"
)

func commandHelp(cfg *Config, args ...string) error {
	cfg.prevCommand = "help"
	if len(args) == 0 {
		fmt.Println()
		for _, command := range GetCommands() {
			fmt.Printf("%s: %s\n", command.name, command.description)
		}
		fmt.Println()
		return nil
	} else if len(args) == 1 {
		commands := GetCommands()
		command, ok := commands[args[0]]
		if !ok {
			return fmt.Errorf("%q is not a command", args[0])
		}
		fmt.Println()
		fmt.Printf("%s: %s\n", command.name, command.description)
		fmt.Println()
		return nil
	} else {
		return fmt.Errorf("expected at most 1 argument, found %d", len(args))
	}
}
