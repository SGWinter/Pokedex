package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandList = map[string]cliCommand{}

func init() {
	commandList = map[string]cliCommand {
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

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputText := cleanInput(scanner.Text())
		if len(inputText) == 0 {
			continue
		}

		commandName := inputText[0]

		if commandName == commandList["exit"].name {
			commandList["exit"].callback()
		} else if commandName == commandList["help"].name {
			commandList["help"].callback()
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	for _, c := range commandList {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}

	return nil
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
