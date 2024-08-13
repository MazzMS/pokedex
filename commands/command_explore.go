package commands

import (
	"errors"
	"fmt"

	"github.com/mazzms/pokedex/internal/pokedex"
)

func commandExplore(c *Config, args ...string) error {
	c.prevCommand = "explore"
	if len(args) != 1 {
		return errors.New("Not a valid amount of arguments")
	}

	// get location
	resp, err :=  c.PokedexClient.LocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("You are in \033[38;2;255;87;51m%s\033[0m\n", resp.Name)
	fmt.Println("You see these pokemons!")
	c.PossiblePokemons = make(map[string]pokedex.Pokemon)
	fmt.Println()
	for _, pokemon := range resp.PokemonEncounters {
		pokemon, err := c.PokedexClient.GetPokemon(pokemon.Pokemon.Name) 
		if err != nil {
			return err
		}
		fmt.Printf(" - %s\n", pokemon.Name)
		c.PossiblePokemons[pokemon.Name] = pokemon
	}
	fmt.Println()
	return nil
}
