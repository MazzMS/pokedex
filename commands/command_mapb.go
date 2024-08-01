package commands

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *Config, args ...string) error {
	// TODO: As we already have arguments, would be interesting to be able to jump more than 1 page
	// if the list was not previously started
	if cfg.prevCommand != "map" && cfg.offset == 0 {
		return errors.New("Use 'map' to start listing.")
	}

	// if list is 20 means that only first page was shown
	if cfg.offset <= 20 {
		cfg.printedLines += 2
		return errors.New("You are already at page 0!")
	}

	// the previous command has to be either map or mapb, so clear it
	for i := 0; i <= cfg.printedLines; i++ {
		fmt.Print("\033[A\033[2K")
	}

	// 1 step is '20', and both map and mapb advances 1 step at the end of the command
	// we need to take 2 steps back to get the previous areas to the last map
	cfg.offset -= 40

	// get locations
	resp, err := cfg.PokedexClient.ListLocationAreas(cfg.offset)
	if err != nil {
		return err
	}

	results := resp.Results

	// set printed lines to 0
	cfg.printedLines = 0
	// print areas
	fmt.Printf("\nThere are %d areas!\n", resp.Count)
	fmt.Printf("Page: %d. Location areas:\n", cfg.offset/20)
	fmt.Println()
	cfg.printedLines += 4
	for i, result := range results {
		fmt.Printf(" %04d. %s\n", cfg.offset+i+1, result.Name)
		cfg.printedLines++
	}
	fmt.Println()
	cfg.printedLines++

	// config for next listing
	cfg.offset += 20
	cfg.prevCommand = "mapb"

	return nil
}
