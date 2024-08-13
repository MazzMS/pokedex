package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Not a valid amount of arguments")
	}

	// if no zone explored
	if c.PossiblePokemons == nil {
		return errors.New("You have not visited any zone!")
	}

	// if name is not a pokemons
	pokemon, err := c.PokedexClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	// if pokemon is not in list
	_, ok := c.PossiblePokemons[pokemon.Name]
	if !ok {
		return fmt.Errorf("%s is not in the area!", pokemon.Name)
	}

	fmt.Printf("Throwing a Pokeball to \033[38;2;255;87;51m%s\033[0m.\r", pokemon.Name)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Throwing a Pokeball to \033[38;2;255;87;51m%s\033[0m..\r", pokemon.Name)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Throwing a Pokeball to \033[38;2;255;87;51m%s\033[0m...\n", pokemon.Name)

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if threshold < randNum {
		return fmt.Errorf("%s escaped!", pokemon.Name)
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Printf("\033[38;2;50;205;50m%s was caught!\033[0m\n", pokemon.Name)
	c.CapturedPokemons[pokemon.Name] = pokemon
	fmt.Println()
	return nil
}
