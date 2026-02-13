package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lnrascal/pokedex/internal/pokeapi"
)

func StartREPL(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if command, exists := getCommands()[commandName]; exists {
			command.callback(cfg, args...)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	res := strings.Fields(text)

	return res
}

type Config struct {
	PokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "Explore",
			description: "Get the list of all the Pokémon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "Catch",
			description: "Catch the Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "Inspect",
			description: "Inspect the Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "Pokedex",
			description: "Displays all caught pokemons",
			callback:    commandPokedex,
		},
	}
}
