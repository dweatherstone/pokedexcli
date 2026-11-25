package main

import (
	"fmt"

	"github.com/dweatherstone/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := cfg.locationAreaURL
	if cfg.mapNextURL != nil {
		url = *cfg.mapNextURL
	}

	return getLocationAreas(cfg, url, "commandMap")
}

func commandMapB(cfg *config) error {
	if cfg.mapPrevURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	return getLocationAreas(cfg, *cfg.mapPrevURL, "commandMapB")
}

func getLocationAreas(cfg *config, url string, fnName string) error {
	var response pokeapi.LocationAreaResponse
	if err := cfg.pokeClient.Get(url, &response); err != nil {
		return fmt.Errorf("%s get request: %w", fnName, err)
	}
	// Set the next and previous map URLs into config
	cfg.mapNextURL = response.Next
	cfg.mapPrevURL = response.Previous
	// Print names to stdout
	for _, r := range response.Results {
		fmt.Println(r.Name)
	}

	return nil
}
