package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex help menu")
	fmt.Println("Here are your available commands")

	avialableCommands := getCommands()
	for _, cmd := range avialableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}

	// fmt.Println(" - help")
	// fmt.Println(" - exit")
	fmt.Println("")
	return nil
}
