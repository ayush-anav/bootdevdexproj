package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *config) {
	fmt.Println("Welcome to the Pokedex!")
	// creating a way to get user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		// block for user input
		scanner.Scan()
		cleanedText := cleanInput(scanner.Text())

		// if nothing was entered, skip the iteration
		if len(cleanedText) == 0 {
			continue
		}

		command, exists := getCommands()[cleanedText[0]]
		if !exists {
			fmt.Println("Unknown Command")
		} else {
			if err := command.callback(c); err != nil {
				fmt.Println(err)
			}
		}
	}

}

func cleanInput(text string) []string {
	output := strings.Fields(strings.ToLower(text))
	return output
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists 20 maps, and keeps listing 20 maps the more you call",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back a page on your map",
			callback:    commandMapb,
		},
	}
}
