package main

// NOTICE: caching is not used here, ill use it later tbh once i get a better layout of the project
import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// location
type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please supply location name")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", args[0])

	var final Location

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not get data from resource %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&final); err != nil {
		return fmt.Errorf("Could not decode JSON to GO struct %w", err)
	}
	fmt.Printf("Exploring %s\nFound Pokemon:\n", args[0])
	for _, d := range final.PokemonEncounters {
		fmt.Println(d.Pokemon.Name)
	}

	return nil
}
