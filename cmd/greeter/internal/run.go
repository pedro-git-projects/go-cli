package internal

import (
	"fmt"
	"os"
)

// Run parses the command line args properly populating the config struct,
// it then validates the required fields
// and, if they're valid, runs the command.
// If at any point a function returns an error, the program will exit with
// status code 1.
func Run() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	if err := validateArgs(c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	if err := runCommand(os.Stdin, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
