package pokeapi

type LocationResponse struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Encounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
