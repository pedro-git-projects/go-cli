package internal

import (
	"fmt"
	"io"
	"os"
)

// printUsage prints the instructions for using the program
// to the specified io writer.
func printUsage(w io.Writer) {
	usageString := fmt.Sprintf("Usage: %s <integer> [-h|--help]\nA greeter application which prints the name you entered <integer> number of times.\n", os.Args[0])
	fmt.Fprintf(w, usageString)
}
