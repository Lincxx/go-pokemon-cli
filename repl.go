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

		if text == "exit" {
			break
		}

		fmt.Println(cleaned)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
