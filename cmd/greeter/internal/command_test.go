package internal

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRunCommand(t *testing.T) {
	tests := []struct {
		c      Config
		input  string
		output string
		err    error
	}{
		{
			c:      Config{printUsage: true},
			output: fmt.Sprintf("Usage: %s <integer> [-h|--help]\nA greeter application which prints the name you entered <integer> number of times.\n", os.Args[0]),
		},
		{
			c:      Config{numTimes: 5},
			input:  "",
			output: "Please enter your name and press Return when done\n",
			err:    errors.New("A name should be entered"),
		},
		{
			c:      Config{numTimes: 5},
			input:  "Pedro",
			output: "Please enter your name and press Return when done\n" + strings.Repeat("Nice to meet you, Pedro\n", 5),
			err:    errors.New("A name should be entered"),
		}}

	byteBuff := new(bytes.Buffer)
	for _, test := range tests {
		rd := strings.NewReader(test.input)
		err := runCommand(rd, byteBuff, test.c)
		if err != nil && test.err == nil {
			t.Fatalf("expected nil got %v\n", err.Error())
		}
		if (test.err != nil && err != nil) && err.Error() != test.err.Error() {
			t.Fatalf("expected %v but got %v\n", test.err.Error(), err.Error())
		}
		gotMsg := byteBuff.String()
		if gotMsg != test.output {
			t.Errorf("expected %v but got %v\n", test.output, gotMsg)
		}
		byteBuff.Reset()
	}
}
