package main

import (
	"fmt"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

func commandPokedex(config *pokeapi.Config, param string) error {
	if len(config.Pokedex) > 0 {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range config.Pokedex {
			fmt.Printf("  - %s\n", pokemon.Name)
		}
	}
	return nil
}
