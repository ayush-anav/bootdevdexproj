// all pokemon stored at caughtPokemon
package main

import (
	"fmt"
)

func commandInspect(c *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please enter the name of the Pokemon you wish to inspect")
	}

	pokemon, ok := caughtPokemon[args[0]]
	if !ok {
		return fmt.Errorf("Pokemon does not exist in Pokeball!")
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		// types
		fmt.Println("Types:")
		for _, v := range pokemon.Types {
			fmt.Printf("- %s\n", v.Type.Name)
		}
	}

	return nil
}
