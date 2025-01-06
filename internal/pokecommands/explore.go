package pokecommands

func commandExplore(args ...string) error {
	location := args[0]
	println(location)
	return nil
}
