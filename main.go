package main

import (
	"time"

	"github.com/mazzms/pokedex/commands"
	"github.com/mazzms/pokedex/internal/pokedex"
)


func main() {
	cfg := &commands.Config{
		PokedexClient: pokedex.NewClient(30 * time.Minute),
	}

	startRepl(cfg)
}
