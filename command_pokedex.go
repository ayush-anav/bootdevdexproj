package main

import (
	"fmt"
)

func commandPokedex(c *config, args ...string) error {
	for _, v := range caughtPokemon {
		fmt.Printf("- %s\n", v.Name)
	}
	return nil
}
