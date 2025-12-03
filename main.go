package main

import (
	"time"

	"github.com/Kanatello7/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	config := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}
