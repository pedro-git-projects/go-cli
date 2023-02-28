package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

// getName prints a message to an io.Writer
// and scans for the user input, if it erros out
// or is an empty string, an error and zero value string is returned.
// Else the input text is returned.
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Please enter your name and press Return when done\n"
	fmt.Fprintf(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("A name should be entered")
	}
	return name, nil
}

// greetUser extracts the name from the Config struct
// and prints the greeting message as many times as
// specified in the numTimes field to the specified writer.
func greetUser(c Config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you, %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}
