package main

import (
	"bootdev/emiel/pokedex/internal/pokecache"
	"bootdev/emiel/pokedex/internal/pokecommands"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	commands := pokecommands.InitializeCommands()
	pokecache.NewCache(5 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

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
			command.Callback(args...)
		}

	}
}

func ParseFromInput(args ...string) string {
	return ""
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}
