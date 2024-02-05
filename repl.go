package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRelp() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		avialableCommands := getCommands()

		//lookup key in map
		cmd, ok := avialableCommands[commandName]

		//if not ok return an error
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		cmd.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
