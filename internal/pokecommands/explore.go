package pokecommands

import (
	"bootdev/emiel/pokedex/internal/pokeapi"
	"fmt"
)

func commandExplore(args ...string) error {
	location := args[0]
	locationDetails, err := pokeapi.GetLocationDetails(location)
	if err != nil {
		return fmt.Errorf("error getting the location details: %v", err)
	}

	if len(locationDetails.PokemonEncounters) == 0 {
		fmt.Println("No pokemon found in this area...")
	}

	for _, pokemon := range locationDetails.PokemonEncounters {
		println(pokemon.Pokemon.Name)
	}

	return nil
}
