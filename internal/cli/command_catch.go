package cli

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a single pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.PokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	isCatched := tryCatch(pokemon.BaseExperience)

	if isCatched {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.PokeapiClient.Pokemons[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func tryCatch(baseExperience int) bool {
	chance := float64(1 / baseExperience)
	if chance < 0.25 {
		chance = 0.25
	}

	if rand.Float64() < chance {
		return true
	}

	return false
}
