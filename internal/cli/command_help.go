package cli

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	keys := make([]string, 0, len(getCommands()))
	for k := range getCommands() {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		cmd := getCommands()[k]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
