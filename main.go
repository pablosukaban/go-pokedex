package main

import (
	"time"

	"github.com/pablosukaban/go-pokedex/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		pokeAPI:  pokeClient,
		commands: getCommands(),
	}

	startRepl(&cfg)
}
