package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputText := scanner.Text()
		lowerInput := strings.ToLower(inputText)
		splitInput := strings.Fields(lowerInput)
		fmt.Println("Your command was:", splitInput[0])
	}
}
