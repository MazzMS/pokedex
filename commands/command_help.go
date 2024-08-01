package commands

import (
	"fmt"
)

func commandHelp(cfg *Config, args ...string) error {
	// TODO: get the description for an specific command if used with an argument
	fmt.Println()
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	cfg.prevCommand = "help"
	return nil
}
