package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"
)

type config struct {
	mapNext    string
	mapPrev    string
}

func startRepl() {
	configState := config {
		mapNext: "",
		mapPrev: "",
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputText := cleanInput(scanner.Text())
		if len(inputText) == 0 {
			continue
		}

		commandName := inputText[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&configState)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(s *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
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
			description: "Displays the next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon world",
			callback:    commandMapb,
		},
	}
}
