package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

type Locations struct {
	List        []string
	NextUrl     string
	PreviousUrl string
}

type Page struct {
	Limit  int
	Offset int
}

func ListLocations(config *Config) (Locations, error) {
	url := BaseUrl + fmt.Sprintf("?limit=%d&offset=%d", config.Page.Limit, config.Page.Offset)
	cache, err := config.Cache.Get(url)
	var locations Locations
	if err == nil {
		err = json.Unmarshal(cache, &locations)
		if err == nil {
			return locations, nil
		}
	}
	waitTime := time.Second * 2
	time.Sleep(waitTime)
	res, err := http.Get(url)
	if err != nil {
		return Locations{}, fmt.Errorf("bad request: %s", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var responsesLocations resultLocation
	if err := decoder.Decode(&responsesLocations); err != nil {
		return Locations{}, fmt.Errorf("bad parsing of maps: %s", err)
	}
	for _, l := range responsesLocations.Results {
		locations.List = append(locations.List, l.Name)
	}
	locations.NextUrl = responsesLocations.Next
	locations.PreviousUrl = responsesLocations.Previous

	jsonData, err := json.Marshal(locations)
	if err != nil {
		return locations, fmt.Errorf("error serializing locations: %s", err)
	}
	config.NextLocationsURL = locations.NextUrl
	config.PrevLocationsURL = locations.PreviousUrl
	err = config.Cache.Add(url, jsonData)
	if err != nil {
		return locations, fmt.Errorf("error caching locations: %s", err)
	}
	return locations, nil
}
