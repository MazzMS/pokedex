package commands

import (
	"errors"
	"fmt"

	"github.com/mazzms/pokedex/internal/pokedex"
)

func commandInspect(c *Config, args ...string) error {
	c.prevCommand = "inspect"
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
	_, ok := c.CapturedPokemons[pokemon.Name]
	if !ok {
		return fmt.Errorf("%s has not been captured!", pokemon.Name)
	}

	// print info
	fmt.Println()
	fmt.Printf("inspecting %s!\n", pokemon.Name)
	fmt.Printf("height: %d\n", pokemon.Height)
	fmt.Printf("weight: %d\n", pokemon.Weight)
	fmt.Println("stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("types: ")
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
