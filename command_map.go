package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type mapData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.nextLocationsURL != nil {
		url = *c.nextLocationsURL
	}

	// 1. make the get request
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Could not get resource at endpoint %w", err)
	}
	defer res.Body.Close()

	// 2. setup decoder
	var serverData mapData
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&serverData); err != nil {
		return fmt.Errorf("Could not decode JSON data to GO struct %w", err)
	}

	// set our config now
	c.nextLocationsURL = serverData.Next
	c.prevLocationsURL = serverData.Previous

	for _, loc := range serverData.Results {
		fmt.Println(loc.Name)
	}

	return nil

}

func commandMapb(c *config) error {
	if c.prevLocationsURL != nil {
		res, err := http.Get(*c.prevLocationsURL)
		if err != nil {
			return fmt.Errorf("Could not get resource at endpoint %w", err)
		}
		defer res.Body.Close()

		var serverData mapData
		decoder := json.NewDecoder(res.Body)

		if err := decoder.Decode(&serverData); err != nil {
			return fmt.Errorf("Could not decode JSON data to GO struct %w", err)
		}

		// set our config now
		c.nextLocationsURL = serverData.Next
		c.prevLocationsURL = serverData.Previous

		for _, loc := range serverData.Results {
			fmt.Println(loc.Name)
		}
	} else {
		fmt.Println("You're on the first page!")
	}

	return nil
}
