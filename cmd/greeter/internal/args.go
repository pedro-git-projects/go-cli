package internal

import (
	"errors"
	"strconv"
)

// Config represents the data that defines
// the application behavior at runtime
type Config struct {
	numTimes   int  // number of time the greeting should be printed.
	printUsage bool // wheter the user wants a help msg to be printed.
}

// parseArgs takes a slice of strings that will be gotten by os.Args[1:].
// it checks if the len of the slice is exactly equal to one,
// and whether it corresponds to the correct flag.
// If the value passed is valid, the config struct is populated accordingly.
func parseArgs(args []string) (Config, error) {
	c := Config{}

	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}

	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}

	numTimes, err := strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes

	return c, nil
}

// validateArgs checks if the parsed number of times is greater than 0
func validateArgs(c Config) error {
	if !(c.numTimes > 0) && !c.printUsage {
		return errors.New("The number of greetings must be greater than 0")
	}
	return nil
}
