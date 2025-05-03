package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

func commandCatch(config *pokeapi.Config, param string) error {
	if param == "" {
		return fmt.Errorf("no param has been pased")
	}
	pokemon, err := pokeapi.GetPokemon(config, param)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if isCatch(pokemon.BaseExperience) {
		config.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func isCatch(xp int) bool {
	random := rand.IntN(xp)
	return xp/3 < random
}
