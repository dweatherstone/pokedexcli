package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, params []string) error {
	if len(params) == 0 {
		return errors.New("inspect command requires a Pokemon name")
	}
	pokemon, found := cfg.pokedex[params[0]]
	if !found {
		fmt.Println("you have not caught that Pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, statRef := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", statRef.Stat.Name, statRef.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
