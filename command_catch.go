package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/dweatherstone/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, params []string) error {
	if len(params) == 0 {
		return errors.New("catch command requires a Pokemon name")
	}
	pokemonName := params[0]
	url := cfg.pokemonURL + pokemonName

	var pokemon pokeapi.Pokemon
	if err := cfg.pokeClient.Get(url, &pokemon); err != nil {
		return fmt.Errorf("commandCatch get request: %w", err)
	}
	baseExperience := float64(pokemon.BaseExperience)
	difficultyConstant := float64(200)
	catchProbability := difficultyConstant / (baseExperience + difficultyConstant)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if catchProbability > rand.Float64() {
		// Did manage to catch...
		cfg.pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		// Failed to catch
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
