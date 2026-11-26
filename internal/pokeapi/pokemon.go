package pokeapi

type Pokemon struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	BaseExperience int             `json:"base_experience"`
	Height         int             `json:"height"`
	Weight         int             `json:"weight"`
	Stats          []StatReference `json:"stats"`
	Types          []TypeReference `json:"types"`
}

type StatReference struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type TypeReference struct {
	Type PokemonType `json:"type"`
}

type PokemonType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
