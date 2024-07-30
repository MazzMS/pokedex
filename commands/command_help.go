package commands

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println()
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
