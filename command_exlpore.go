package main

import (
	"fmt"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

func commandExplore(config *pokeapi.Config, param string) error {
	if param == "" {
		return fmt.Errorf("no param has been pased")
	}
	pokemons, err := pokeapi.ListLocationPokemon(config, param)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
	return nil
}
