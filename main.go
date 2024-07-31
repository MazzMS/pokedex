package main

import (
	"github.com/mazzms/pokedex/commands"
	"github.com/mazzms/pokedex/internal/pokedex"
)


func main() {
	cfg := &commands.Config{
		PokedexClient: pokedex.NewClient(),
	}

	startRepl(cfg)
}
