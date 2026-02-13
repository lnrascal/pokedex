package cli

import "fmt"

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.PokeapiClient.Pokemons {
		fmt.Println(cfg.PokeapiClient.Pokemons[name].Name)
	}

	return nil
}
