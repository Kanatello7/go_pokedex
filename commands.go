package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func commandExit(config *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMapf(config *config, args ...string) error {
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

func commandMapb(config *config, args ...string) error {
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

func commandExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := config.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}

const PowerConstant = 2.0

func catchPokemon(experience int) bool {
	res := rand.Intn(experience)
	return res > 40

}

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide pokemon name")
	}
	name := args[0]
	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catched := catchPokemon(pokemon.Base_experience)
	if catched {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
