package pokecommands

import (
	"bootdev/emiel/pokedex/internal/pokeapi"
	"fmt"
)

func commandInspect(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify a pokemon you caught to inspect")
	}

	pokemonName := args[0]

	if _, caught := pokeapi.CaughtPokemonNames[pokemonName]; !caught {
		return fmt.Errorf("you have not caught that pokemon")
	}

	pokemonDetails, err := pokeapi.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}

	// Basic fields
	fmt.Printf("Name: %s\n", pokemonDetails.Name)
	fmt.Printf("Height: %d\n", pokemonDetails.Height)
	fmt.Printf("Weight: %d\n", pokemonDetails.Weight)

	// Stats section
	fmt.Println("Stats:")
	for _, stat := range pokemonDetails.Stats {
		fmt.Printf("    - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	// Types section
	fmt.Println("Types:")
	for _, t := range pokemonDetails.Types {
		fmt.Printf("    - %s\n", t.Type.Name)
	}

	return nil
}
