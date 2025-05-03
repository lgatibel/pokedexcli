package main

import (
	"fmt"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

func commandInspect(config *pokeapi.Config, param string) error {
	if param == "" {
		return fmt.Errorf("no param has been pased")
	}

	pokemon, ok := config.Pokedex[param]
	if !ok {
		return fmt.Errorf("%s: was not found in your Pokedex", param)
	}
	stats := []string{"hp", "attack", "defense", "special-attack", "special-defense", "speed"}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	if err := printStats(stats, pokemon.Stats); err != nil {
		return err
	}
	if err := printType(pokemon.Types); err != nil {
		return err
	}
	return nil
}

func printType(types []string) error {
	fmt.Printf("Types:\n")
	for _, t := range types {
		fmt.Printf("  - %s\n", t)
	}
	return nil
}

func printStats(statRef []string, stats map[string]int) error {
	fmt.Printf("Stats:\n")
	for _, sr := range statRef {
		stat, ok := stats[sr]
		if !ok {
			return fmt.Errorf("stat: %s Not Found", sr)
		}
		fmt.Printf("  -%s: %d\n", sr, stat)
	}
	return nil
}
