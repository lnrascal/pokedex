package cli

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a single pokemon name")
	}

	name := args[0]
	if pokemon, exists := cfg.PokeapiClient.Pokemons[name]; exists {
		fmt.Printf("Name: %s \n", pokemon.Name)
		fmt.Printf("Height: %v \n", pokemon.Height)
		fmt.Printf("Weight: %v \n", pokemon.Weight)
		fmt.Println("Stats:")

		for i := 0; i < len(pokemon.Stats); i++ {
			stat := pokemon.Stats[i]
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for j := 0; j < len(pokemon.Types); j++ {
			fmt.Printf("  -%s\n", pokemon.Types[j].Type.Name)
		}

		return nil
	}

	fmt.Println("you have not caught that pokemon")

	return nil
}
