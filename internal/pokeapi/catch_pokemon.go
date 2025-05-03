package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responsePokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int
		Type struct {
			Name string
		}
	} `json:"types"`
	BaseExperience int `json:"base_experience"`
}

func GetPokemon(config *Config, param string) (Pokemon, error) {
	url := baseUrl + pokemonEnpoint + param
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("bad request: %s", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var pokemonResult responsePokemon
	if err := decoder.Decode(&pokemonResult); err != nil {
		return Pokemon{}, fmt.Errorf("bad parsing of maps: %s", err)
	}
	pokemon := Pokemon{
		Name:           pokemonResult.Name,
		Height:         pokemonResult.Height,
		Weight:         pokemonResult.Weight,
		Stats:          make(map[string]int),
		Types:          []string{},
		BaseExperience: pokemonResult.BaseExperience,
	}

	for _, s := range pokemonResult.Stats {
		pokemon.Stats[s.Stat.Name] = s.BaseStat
	}
	for _, t := range pokemonResult.Types {
		pokemon.Types = append(pokemon.Types, t.Type.Name)
	}
	return pokemon, nil
}
