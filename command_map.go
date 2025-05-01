package main

import (
	"fmt"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL string
	prevLocationsURL string
	page             pokeapi.Page
}

func commandMapB(config *config) error {
	locations, err := pokeapi.ListLocations(config.page)
	if config.page.Offset < 1 {
		return fmt.Errorf("you're on the first page")
	}
	if err != nil {
		return err
	}
	for _, location := range locations.List {
		fmt.Println(location)
	}
	config.nextLocationsURL = locations.NextUrl
	config.prevLocationsURL = locations.PreviousUrl
	config.page.Offset -= len(locations.List)
	return nil
}

func commandMap(config *config) error {
	locations, err := pokeapi.ListLocations(config.page)
	if err != nil {
		return err
	}
	for _, location := range locations.List {
		fmt.Println(location)
	}
	config.nextLocationsURL = locations.NextUrl
	config.prevLocationsURL = locations.PreviousUrl
	config.page.Offset += len(locations.List)
	return nil
}
