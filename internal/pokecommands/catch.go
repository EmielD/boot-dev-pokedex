package pokecommands

import (
	"bootdev/emiel/pokedex/internal/pokeapi"
	"fmt"
	"math/rand"
)

func commandCatch(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify a pokemon you'd like to catch")
	}

	pokemonName := args[0]

	if _, caught := pokeapi.CaughtPokemonNames[pokemonName]; caught {
		return fmt.Errorf("you already caught this pokemon")
	}

	pokemonDetails, err := pokeapi.GetPokemonDetails(pokemonName)
	if err != nil {
		return fmt.Errorf("error getting pokemon details: %v", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	chanceOfSuccessfullCatch := 60.0 - (float32(pokemonDetails.BaseExperience)/100.0)*10.0

	catchRoll := rand.Intn(100)
	if catchRoll < int(chanceOfSuccessfullCatch) {
		pokeapi.CaughtPokemonNames[pokemonName] = struct{}{}
		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", pokemonDetails.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonDetails.Name)
	}

	return nil
}
