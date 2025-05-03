package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type resultLocation struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

type locations struct {
	List        []string
	NextUrl     string
	PreviousUrl string
}

func ListLocations(config *Config) (locations, error) {
	url := baseUrl + fmt.Sprintf("?limit=%d&offset=%d", config.Page.Limit, config.Page.Offset)
	var locationsList locations
	cache, ok := config.PokeapiClient.cache.Get(url)
	if ok {
		err := json.Unmarshal(cache, &locationsList)
		if err == nil {
			return locationsList, nil
		}
	}
	res, err := http.Get(url)
	if err != nil {
		return locations{}, fmt.Errorf("bad request: %s", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var responsesLocations resultLocation
	if err := decoder.Decode(&responsesLocations); err != nil {
		return locations{}, fmt.Errorf("bad parsing of maps: %s", err)
	}
	for _, l := range responsesLocations.Results {
		locationsList.List = append(locationsList.List, l.Name)
	}
	locationsList.NextUrl = responsesLocations.Next
	locationsList.PreviousUrl = responsesLocations.Previous

	jsonData, err := json.Marshal(locationsList)
	if err != nil {
		return locationsList, fmt.Errorf("error serializing locations: %s", err)
	}
	config.NextLocationsURL = locationsList.NextUrl
	config.PrevLocationsURL = locationsList.PreviousUrl
	err = config.PokeapiClient.cache.Add(url, jsonData)
	if err != nil {
		return locationsList, fmt.Errorf("error caching locations: %s", err)
	}
	return locationsList, nil
}
