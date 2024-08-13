package commands

import (
	"errors"
	"fmt"
)

func commandExplore(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Not a valid amount of arguments")
	}

	// get location
	resp, err :=  c.PokedexClient.LocationArea(args[0])
	if err != nil {
		return err
	}

	// print areas
	// fmt.Printf("You are in %q!\n", resp.Name)
	fmt.Printf("You are in \033[38;2;255;87;51m%s\033[0m\n", resp.Name)
	fmt.Println("You see these pokemons!")
	fmt.Println()
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
