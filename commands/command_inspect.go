package commands

import (
	"errors"
	"fmt"

	"github.com/mazzms/pokedex/internal/pokedex"
)

func commandInspect(c *Config, args ...string) error {
	// no argument, show all the captured pokemons
	if len(args) == 0 {
		if len(c.capturedPokemons) == 0 {
			return errors.New("You have not captured any pokemon!")
		}
		if len(c.capturedPokemons) == 1 {
			fmt.Println("You have captured this pokemon!")
		} else {
			fmt.Println("You have captured these pokemons!")
		}
		for _, pokemon := range c.capturedPokemons {
			fmt.Printf("- %s\n", pokemon.Name)
		}
		fmt.Println()
		return nil
	}

	// else show the info about that pokemon
	if len(args) != 1 {
		return errors.New("Not a valid amount of arguments")
	}

	// check if valid pokemon
	pokemon, err :=  c.PokedexClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	// if pokemon is not captured
	_, ok := c.capturedPokemons[pokemon.Name]
	if !ok {
		return fmt.Errorf("%s has not been captured!", pokemon.Name)
	}

	// print info
	fmt.Println()
	fmt.Printf("inspecting %s!\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	/*
	fmt.Println("\t-hp: %s")
	fmt.Println("\t-attack: %s")
	fmt.Println("\t-defense: %s")
	fmt.Println("\t-special-attack: %s")
	fmt.Println("\t-special-defense: %s")
	fmt.Println("\t-speed: %s")
	fmt.Println("Types: ")
	*/
	for _, t := range pokemon.Types {
		color, err := pokedex.GetColor(t.Type.Name)
		if err != nil {
			return err
		}
		fmt.Printf("\t- \033%s%s\033[0m\n", color, t.Type.Name)
	}
	fmt.Println()
	return nil
}
