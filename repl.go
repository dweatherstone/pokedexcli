package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dweatherstone/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, params []string) error
}

type config struct {
	commands        map[string]cliCommand
	pokeClient      *pokeapi.Client
	mapNextURL      *string
	mapPrevURL      *string
	locationAreaURL string
}

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	return strings.Fields(lower_text)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			}
			return
		}

		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		parameters := words[1:]
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, parameters)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error running callback: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", commandName)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explores a location to find which Pokemon are there",
			callback:    commandExplore,
		},
	}
}
