package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMapf(config *config) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}
	config.nextLocationsURL = locationsResp.Next
	config.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *config) error {
	if config.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationResp.Next
	config.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
