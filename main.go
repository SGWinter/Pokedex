package main

import (
	"time"

	"github.com/SGWinter/Pokedex/internal/pokeapi"
	"github.com/SGWinter/Pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache()
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
	}

	startRepl(cfg)
}
