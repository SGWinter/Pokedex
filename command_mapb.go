package main

import (
	"fmt"
	"io"
	"net/http"
	"log"
	"encoding/json"
)

func commandMapb(s *config) error {
	type resultsBody struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	type bodyResults struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []resultsBody  `json:"results"`
	}

	getURL := ""
	if s.mapPrev == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		getURL = s.mapPrev
	}
	res, err := http.Get(getURL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return err
	}
	if err != nil {
		log.Fatal(err)
		return err
	}

	b := bodyResults{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Fatal(err)
		return err
	}
	for i := range b.Results {
		fmt.Println(b.Results[i].Name)
	}
	s.mapNext = b.Next
	s.mapPrev = b.Previous
	return nil
}
