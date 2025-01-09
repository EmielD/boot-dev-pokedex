package pokecommands

type cliCommand struct {
	name        string
	description string
	Callback    func(...string) error
}

var commands map[string]cliCommand

func InitializeCommands() map[string]cliCommand {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			Callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			Callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Returns more information about the specified location",
			Callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempts to catch the specified pokemon",
			Callback:    commandCatch,
		},
		"clear": {
			name:        "clear",
			description: "Clears the terminal",
			Callback:    commandClear,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Gives detailed information of the caught pokemon",
			Callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Returns all the pokemon you have caught",
			Callback:    commandPokedex,
		},
	}

	return commands
}
