package commands

import (
	"fmt"

	"github.com/mazzms/pokedex/internal/pokedex"
)

func commandMap() error {
	pokedexClient := pokedex.NewClient()

	resp, err := pokedexClient.ListLocationAreas()

	if err != nil {
		return err
	}

	results := resp.Results

	for _, result := range results {
		fmt.Println(result.Name)
	}
	fmt.Println()
	return nil
}
