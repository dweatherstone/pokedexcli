package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	commands map[string]cliCommand
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
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg)
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
	}
}
