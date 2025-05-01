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

type Locations struct {
	List        []string
	NextUrl     string
	PreviousUrl string
}

type Page struct {
	Limit  int
	Offset int
}

func ListLocations(page Page) (Locations, error) {
	url := BaseUrl + fmt.Sprintf("?limit=%d&offset=%d", page.Limit, page.Offset)
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
	var locations Locations
	for _, l := range responsesLocations.Results {
		locations.List = append(locations.List, l.Name)
	}
	locations.NextUrl = responsesLocations.Next
	locations.PreviousUrl = responsesLocations.Next
	return locations, nil
}
