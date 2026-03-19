package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// creating a way to get user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		// block for user input
		scanner.Scan()
		userInput := scanner.Text()
		cleanedText := cleanInput(userInput)
		fmt.Printf("Your command was: %s\n", cleanedText[0])
	}

}
