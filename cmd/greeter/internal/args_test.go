package internal

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		args []string
		c    Config
		err  error
	}{
		{
			args: []string{"-h"},
			c:    Config{numTimes: 0, printUsage: true},
			err:  nil,
		},
		{
			args: []string{"--help"},
			c:    Config{numTimes: 0, printUsage: true},
			err:  nil,
		},
		{
			args: []string{"-help"},
			c:    Config{numTimes: 0, printUsage: false},
			err:  errors.New(`strconv.Atoi: parsing "-help": invalid syntax`),
		},
		{
			args: []string{},
			c:    Config{numTimes: 0, printUsage: false},
			err:  errors.New("Invalid number of arguments"),
		},
		{
			args: []string{"1"},
			c:    Config{numTimes: 1, printUsage: false},
			err:  nil,
		},
		{
			args: []string{"-1"},
			c:    Config{numTimes: -1, printUsage: false},
			err:  errors.New("The number of greetings must be greater than 0"),
		},
	}

	for _, test := range tests {
		c, err := parseArgs(test.args)
		if !reflect.DeepEqual(c, test.c) {
			t.Errorf("expected %v got %v", test.c, c)
		}
		if test.err == nil && err != nil {
			t.Errorf("expected %v got %v", test.err, err)
		}
		if (test.err != nil && err != nil) && test.err.Error() != err.Error() {
			t.Errorf("expected %v got %v", test.err, err)
		}
	}
}

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		c   Config
		err error
	}{
		{
			c:   Config{},
			err: errors.New("The number of greetings must be greater than 0"),
		},
		{
			c:   Config{numTimes: -1},
			err: errors.New("The number of greetings must be greater than 0"),
		},
		{
			c:   Config{numTimes: 10},
			err: nil,
		},
	}

	for _, test := range tests {
		err := validateArgs(test.c)
		if test.err != nil && err.Error() != test.err.Error() {
			t.Errorf("expected %v but got %v\n", test.err, err)
		}
		if test.err == nil && err != nil {
			t.Errorf("expected nil but got %v\n", err)
		}
	}
}
