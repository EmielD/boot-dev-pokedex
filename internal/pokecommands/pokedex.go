package pokecommands

import (
	"bootdev/emiel/pokedex/internal/pokeapi"
	"fmt"
)

func commandPokedex(args ...string) error {

	if len(pokeapi.CaughtPokemonNames) == 0 {
		return fmt.Errorf("you didn't catch any Pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for pokemonName := range pokeapi.CaughtPokemonNames {
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}
