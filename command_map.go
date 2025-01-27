package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.mapNext)
	if err != nil {
		return err
	}

	cfg.mapNext = locationsResp.Next
	cfg.mapPrev = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.mapPrev == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.mapPrev)
	if err != nil {
		return err
	}

	cfg.mapNext = locationResp.Next
	cfg.mapPrev = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
