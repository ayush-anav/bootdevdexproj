package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func commandMap(c *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.nextLocationsURL != nil {
		url = *c.nextLocationsURL
	}

	var serverData mapData

	if cachedData, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(cachedData, &serverData); err != nil {
			return fmt.Errorf("could not decode cached data: %w", err)
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("could not get resource at endpoint: %w", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("could not read response body: %w", err)
		}

		c.cache.Add(url, body)

		if err := json.Unmarshal(body, &serverData); err != nil {
			return fmt.Errorf("could not decode JSON data: %w", err)
		}
	}

	c.nextLocationsURL = serverData.Next
	c.prevLocationsURL = serverData.Previous

	for _, loc := range serverData.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(c *config, args ...string) error {
	if c.prevLocationsURL == nil {
		fmt.Println("You're on the first page!")
		return nil
	}

	url := *c.prevLocationsURL
	var serverData mapData

	if cachedData, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(cachedData, &serverData); err != nil {
			return fmt.Errorf("could not decode cached data: %w", err)
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("could not get resource at endpoint: %w", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("could not read response body: %w", err)
		}

		c.cache.Add(url, body)

		if err := json.Unmarshal(body, &serverData); err != nil {
			return fmt.Errorf("could not decode JSON data: %w", err)
		}
	}

	c.nextLocationsURL = serverData.Next
	c.prevLocationsURL = serverData.Previous

	for _, loc := range serverData.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
