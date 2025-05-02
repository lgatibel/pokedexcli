package main

import (
	"fmt"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

func commandMapB(config *pokeapi.Config) error {
	if config.Page.Offset < 1 {
		return fmt.Errorf("you're on the first page")
	}
	if config.Page.Offset >= config.Page.Limit {
		config.Page.Offset -= config.Page.Limit
	}
	locations, err := pokeapi.ListLocations(config)
	if err != nil {
		return err
	}
	for _, location := range locations.List {
		fmt.Println(location)
	}
	return nil
}

func commandMap(config *pokeapi.Config) error {
	locations, err := pokeapi.ListLocations(config)
	if err != nil {
		return err
	}
	for _, location := range locations.List {
		fmt.Println(location)
	}
	config.Page.Offset += len(locations.List)
	return nil
}
