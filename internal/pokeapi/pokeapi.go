package pokeapi

type Page struct {
	Limit  int
	Offset int
}

const (
	BaseUrl   = "https://pokeapi.co/api/v2/location-area/"
	PageLimit = 20
)
