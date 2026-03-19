package main

import "fmt"

func commandHelp(c *config) error {
	fmt.Println("============= HELP =============")
	for key, value := range getCommands() {
		fmt.Printf("\nCommand: %s \n Description: %s \n\n", key, value.description)
	}
	fmt.Println("================================")
	return nil
}
