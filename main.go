package main

import (
	"time"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

// "Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex",
func main() {
	config := config{
		pokeapiClient:    pokeapi.NewClient(3 * time.Second),
		prevLocationsURL: "",
		nextLocationsURL: "",
		page: pokeapi.Page{
			Limit:  20,
			Offset: 0,
		},
	}
	startRepl(config)
}
