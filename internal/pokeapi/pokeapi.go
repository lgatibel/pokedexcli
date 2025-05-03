package pokeapi

type Page struct {
	Limit  int
	Offset int
}

const (
	baseUrl   = "https://pokeapi.co/api/v2/location-area/"
	PageLimit = 20
)
