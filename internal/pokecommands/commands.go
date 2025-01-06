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
		"clear": {
			name:        "clear",
			description: "Clears the terminal",
			Callback:    commandClear,
		},
	}

	return commands
}
