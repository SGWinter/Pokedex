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

		fmt.Println("Your command was:", inputText[0])
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
