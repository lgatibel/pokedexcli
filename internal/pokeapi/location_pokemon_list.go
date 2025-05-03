package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type pokemonList []string

type ressultLocationPokemon struct {
	Pokemonlist []struct {
		Pokemon struct {
			Name string
			Url  string
		}
	} `json:"pokemon_encounters"`
}

func ListLocationPokemon(config *Config, param string) (pokemonList, error) {
	url := baseUrl + param
	var pokemons pokemonList
	res, err := http.Get(url)
	if err != nil {
		return pokemonList{}, fmt.Errorf("bad request: %s", err)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var responsesLocations ressultLocationPokemon
	if err := decoder.Decode(&responsesLocations); err != nil {
		return pokemonList{}, fmt.Errorf("bad parsing of maps: %s", err)
	}

	for _, p := range responsesLocations.Pokemonlist {
		pokemons = append(pokemons, p.Pokemon.Name)
	}
	return pokemons, nil
}
