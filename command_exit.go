package main

import (
	"fmt"
	"os"
)

func commandExit(s *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
