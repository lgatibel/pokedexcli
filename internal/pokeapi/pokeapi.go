package pokeapi

type Page struct {
	Limit  int
	Offset int
}

type Pokemon struct {
	Name           string
	Height         int
	Weight         int
	Stats          map[string]int
	Types          []string
	BaseExperience int
}

const (
	baseUrl          = "https://pokeapi.co/api/v2/"
	locationEndpoint = "location-area/"
	pokemonEnpoint   = "pokemon/"
	PageLimit        = 20
)
