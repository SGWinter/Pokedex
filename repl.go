package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"
)

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

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
