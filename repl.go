package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pablosukaban/go-pokedex/api"
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
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}

type config struct {
	commands   map[string]cliCommand
	pokeAPI    api.Client
	mapURLNext string
	mapURLPrev string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	res := strings.Fields(lower)

	return res
}

func startRepl(cfg *config) {
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
		command, ok := cfg.commands[first]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
