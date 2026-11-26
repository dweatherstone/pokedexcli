package main

import "github.com/dweatherstone/pokedexcli/internal/pokeapi"

func main() {
	cfg := &config{
		commands:        getCommands(),
		pokeClient:      pokeapi.NewClient(),
		mapNextURL:      nil,
		mapPrevURL:      nil,
		locationAreaURL: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}
	startRepl(cfg)
}
