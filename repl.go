package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	res := strings.Fields(lower)

	return res
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if err := scanner.Err(); err != nil {
			fmt.Printf("error: %v", err)
			continue
		}
		ok := scanner.Scan()
		if !ok {
			fmt.Printf("not ok")
			continue
		}
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			fmt.Println("no text")
			continue
		}
		first := cleaned[0]
		command, ok := getCommands()[first]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
