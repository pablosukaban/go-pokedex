package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locationAreaURL := "https://pokeapi.co/api/v2/location-area/"
	if cfg.mapURLNext != "" {
		locationAreaURL = cfg.mapURLNext
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
