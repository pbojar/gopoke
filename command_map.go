package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locResp, err := cfg.client.GetLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locResp.Next
	cfg.prevURL = locResp.Previous

	for _, loc := range locResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	if cfg.prevURL == nil {
		return errors.New("you're on the first page")
	}

	locResp, err := cfg.client.GetLocations(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locResp.Next
	cfg.prevURL = locResp.Previous

	for _, loc := range locResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
