package commands

import (
	"os"
	"fmt"
)

func commandExit() error {
	fmt.Println("Thanks for using this Pokedex!") 
	fmt.Println("Closing it...") 
	defer os.Exit(0)
	return nil
}

