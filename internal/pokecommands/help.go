package pokecommands

import "fmt"

func commandHelp(args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for i := range commands {
		command := commands[i]
		fmt.Printf("%v: %v\n", command.name, command.description)
	}

	return nil
}
