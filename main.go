package main

import (
	"time"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config, param string) error
}

// "Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex",
func main() {
	config := &pokeapi.Config{
		PokeapiClient:    *pokeapi.NewClient(3*time.Second, 5*time.Second),
		PrevLocationsURL: "",
		NextLocationsURL: "",
		Page: pokeapi.Page{
			Limit:  20,
			Offset: 0,
		},
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(config)
}
