package internal

import "io"

// runCommand prints the usage and returns if the
// printUsage field in the config struct is set to true.
// If it isn't it will attempt to call getName and
// greetUser.
func runCommand(r io.Reader, w io.Writer, c Config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}
	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}
