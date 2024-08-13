
package commands

import (
	"errors"
	"fmt"
)

func commandPokedex(c *Config, args ...string) error {
	c.prevCommand = "pokedex"
	if len(c.CapturedPokemons) == 0 {
		return errors.New("You have not captured any pokemon!")
	}
	fmt.Print("This is your Pokedex! ")
	if len(c.CapturedPokemons) == 1 {
		fmt.Println("You have captured this pokemon!")
	} else {
		fmt.Println("You have captured these pokemons!")
	}
	for _, pokemon := range c.CapturedPokemons {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	fmt.Println()
	return nil
}
