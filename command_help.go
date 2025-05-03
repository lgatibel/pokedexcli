package main

import (
	"fmt"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

func commandHelp(config *pokeapi.Config, param string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	return nil
}
