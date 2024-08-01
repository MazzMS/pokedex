package commands

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Thanks for using this Pokedex!")
	fmt.Println("Closing it...")
	defer os.Exit(0)
	// just to keep the same logic, but anyways the program will exit after this
	cfg.prevCommand = "exit"
	return nil
}
