package pokecommands

import (
	"bootdev/emiel/pokedex/internal/pokeapi"
	"fmt"
)

func commandExplore(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify a location you'd like to explore")
	}

	location := args[0]
	locationDetails, err := pokeapi.GetLocationDetails(location)
	if err != nil {
		return fmt.Errorf("error getting the location details: %v", err)
	}

	if len(locationDetails.PokemonEncounters) == 0 {
		fmt.Println("No pokemon found in this area...")
	}

	fmt.Println("Found Pokemon:")

	for _, pokemon := range locationDetails.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	return nil
}
