package main

import (
	"bootdev/emiel/pokedex/internal/pokecommands"
	"fmt"
	"strings"

	"github.com/chzyer/readline"
)

func main() {

	commands := pokecommands.InitializeCommands()

	rl, err := readline.NewEx(&readline.Config{
		Prompt: "Pokedex > ",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		input, err := rl.Readline()
		if err != nil {
			fmt.Printf("an error occurred parsing your input text: %v", err)
		}

		if len(input) == 0 {
			continue
		}

		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if command, ok := commands[words[0]]; ok {
			err := command.Callback(args...)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}
