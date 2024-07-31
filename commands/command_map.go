package commands

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config) error {
	// if previous command was not either 'map' or 'mapb' restart the offset
	if cfg.prevCommand != "map" && cfg.prevCommand != "mapb" {
		cfg.offset = 0
	}

	// get locations
	resp, err := cfg.PokedexClient.ListLocationAreas(cfg.offset)
	if err != nil {
		return err
	}

	// if beyond count it's an error
	if cfg.offset > resp.Count {
		return errors.New("You are already at the last page!")
	}

	results := resp.Results

	// clear previous list if the last command was either map or mapb
	if cfg.prevCommand == "map" || cfg.prevCommand == "mapb" {
		for i := 0; i <= (24 + 1); i++ {
			fmt.Print("\033[A\033[2K")
		}
	}

	// print areas
	fmt.Printf("\nThere are %d areas!\n", resp.Count)
	fmt.Printf("Page: %d. Location areas:\n", cfg.offset/20)
	fmt.Println()
	for i, result := range results {
		fmt.Printf(" %04d. %s\n", cfg.offset+i+1, result.Name)
	}
	fmt.Println()

	// config for next listing
	cfg.offset += 20
	cfg.prevCommand = "map"

	return nil
}
