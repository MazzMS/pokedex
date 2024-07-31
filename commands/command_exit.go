package commands

import (
	"os"
	"fmt"
)

func commandExit(cfg *Config) error {
	fmt.Println("Thanks for using this Pokedex!") 
	fmt.Println("Closing it...") 
	defer os.Exit(0)
	// just to keep the same logic, but anyways the program will exit after this
	cfg.prevCommand = "exit"
	return nil
}

