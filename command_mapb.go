package main

import (
	"fmt"
)

func commandMapb(cfg *config) error {
	locationAreaURL := cfg.mapURLPrev
	if locationAreaURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locAreaResp, err := cfg.pokeAPI.GetPokemons(locationAreaURL)
	if err != nil {
		return err
	}

	if len(locAreaResp.Results) == 0 || locAreaResp.Count == 0 {
		return fmt.Errorf("locations not found")
	}

	for _, area := range locAreaResp.Results {
		fmt.Println(area.Name)
	}

	cfg.mapURLNext = locAreaResp.Next
	cfg.mapURLPrev = locAreaResp.Previous

	return nil
}
