package main

import (
	"time"

	"github.com/lnrascal/pokedex/internal/cli"
	"github.com/lnrascal/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &cli.Config{
		PokeapiClient: pokeClient,
	}

	cli.StartREPL(cfg)
}
