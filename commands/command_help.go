package commands

import (
	"fmt"
)

func commandHelp(cfg *Config) error {
	fmt.Println()
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	cfg.prevCommand = "help"
	return nil
}
