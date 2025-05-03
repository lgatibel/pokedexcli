package main

import (
	"fmt"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

func commandExplore(config *pokeapi.Config) error {
	if config.Page.Offset < 1 {
		return fmt.Errorf("you're on the first page")
	}
	if config.Page.Offset >= config.Page.Limit {
		config.Page.Offset -= config.Page.Limit
	}
	locations, err := pokeapi.ListLocationPokemon(config)
	if err != nil {
		return err
	}
	for _, location := range locations.List {
		fmt.Println(location)
	}
	return nil
}
