package main

import (
	"fmt"
	"strings"
)

func commandHelp(cfg *config, _params []string) error {
	var helpString strings.Builder
	helpString.WriteString("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range cfg.commands {
		commandString := fmt.Sprintf("%s: %s\n", command.name, command.description)
		helpString.WriteString(commandString)
	}
	fmt.Print(helpString.String())
	return nil
}
