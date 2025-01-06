package pokecommands

import (
	"bootdev/emiel/pokedex/internal/pokeapi"
	"fmt"
)

func commandMapB(args ...string) error {
	locations, err := pokeapi.GetLocations(true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
