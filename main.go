package main

import (
	"time"

	"github.com/pablosukaban/go-pokedex/api"
)

func main() {
	pokeClient := api.NewClient(5 * time.Second)
	cfg := config{
		pokeAPI:  pokeClient,
		commands: getCommands(),
	}

	startRepl(&cfg)
}
