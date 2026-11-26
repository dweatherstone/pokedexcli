package main

import (
	"errors"
	"fmt"

	"github.com/dweatherstone/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, params []string) error {
	if len(params) == 0 {
		return errors.New("explore command requires a location area name")
	}
	url := cfg.locationAreaURL + params[0]
	fmt.Printf("Exploring %s...\n", params[0])

	var response pokeapi.LocationResponse
	if err := cfg.pokeClient.Get(url, &response); err != nil {
		return fmt.Errorf("commandExplore get request: %w", err)
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range response.Encounters {
		fmt.Printf("\t- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
