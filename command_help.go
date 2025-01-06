package main

import "fmt"

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, c := range commandList {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}
	return nil
}
